package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
}

type Calendar struct {
	events map[int]Event
	mu     sync.Mutex
	currID int
}

var calendar = Calendar{
	events: make(map[int]Event),
	currID: 1,
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", createEvent)
	mux.HandleFunc("/update_event", updateEvent)
	mux.HandleFunc("/delete_event", deleteEvent)
	mux.HandleFunc("/events_for_day", eventsForDay)
	mux.HandleFunc("/events_for_week", eventsForWeek)
	mux.HandleFunc("/events_for_month", eventsForMonth)

	loggedMux := loggingMiddleware(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	log.Println("Started on :8080")
	log.Fatal(server.ListenAndServe())
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	title := r.FormValue("title")
	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid date format")
		return
	}

	calendar.mu.Lock()
	defer calendar.mu.Unlock()

	event := Event{
		ID:     calendar.currID,
		UserID: userID,
		Name:   title,
		Date:   date,
	}

	calendar.events[event.ID] = event
	calendar.currID++

	jsonResult(w, fmt.Sprintf("Event %d created", event.ID))
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	eventID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid event ID")
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	title := r.FormValue("title")
	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid date format")
		return
	}

	calendar.mu.Lock()
	defer calendar.mu.Unlock()

	event, exists := calendar.events[eventID]
	if !exists {
		jsonError(w, http.StatusServiceUnavailable, "Event not found")
		return
	}

	event.UserID = userID
	event.Name = title
	event.Date = date
	calendar.events[eventID] = event

	jsonResult(w, fmt.Sprintf("Event %d updated", event.ID))
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	eventID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid event ID")
		return
	}

	calendar.mu.Lock()
	defer calendar.mu.Unlock()

	_, exists := calendar.events[eventID]
	if !exists {
		jsonError(w, http.StatusServiceUnavailable, "Event not found")
		return
	}

	delete(calendar.events, eventID)

	jsonResult(w, fmt.Sprintf("Event %d deleted", eventID))
}

func eventsForDay(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r.URL.Query().Get("date"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid date format")
		return
	}

	calendar.mu.Lock()
	defer calendar.mu.Unlock()

	var events []Event
	for _, event := range calendar.events {
		if event.Date.Equal(date) {
			events = append(events, event)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r.URL.Query().Get("date"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid date format")
		return
	}

	calendar.mu.Lock()
	defer calendar.mu.Unlock()

	var events []Event
	start := date
	end := date.AddDate(0, 0, 7)
	for _, event := range calendar.events {
		if event.Date.After(start) && event.Date.Before(end) {
			events = append(events, event)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r.URL.Query().Get("date"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid date format")
		return
	}

	calendar.mu.Lock()
	defer calendar.mu.Unlock()

	var events []Event
	start := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	end := start.AddDate(0, 1, 0)
	for _, event := range calendar.events {
		if event.Date.After(start) && event.Date.Before(end) {
			events = append(events, event)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func parseDate(dateStr string) (time.Time, error) {
	layout := "0000-00-00"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, errors.New("Invalid date format")
	}
	return date, nil
}

func jsonError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func jsonResult(w http.ResponseWriter, result string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
		},
	)
}
