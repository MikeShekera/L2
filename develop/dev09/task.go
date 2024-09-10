package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var siteLink string
	_, err := fmt.Scan(&siteLink)
	linksMap := make(map[string]string)
	if err != nil {
		log.Fatal(err)
	}
	err = downloadSite(siteLink, linksMap)
	if err != nil {
		fmt.Println(err)
	}
	select {}
}

func downloadSite(link string, linksMap map[string]string) error {
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	linksMap[getPageName(link)] = link
	body, err := html.Parse(resp.Body)
	extractLinks(body, linksMap)

	for k, v := range linksMap {
		fmt.Println(k + ":" + v)
		saveToFile(v)
	}
	return nil
}

func extractLinks(n *html.Node, linksMap map[string]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" && strings.Contains(attr.Val, "https") {
				pageName := getPageName(attr.Val)
				if _, ok := linksMap[pageName]; !ok {
					linksMap[pageName] = attr.Val
				}
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c, linksMap)
	}
}

/*func downloadSite2(link string, depthIndex int, linksMap map[string]string) error {
	if _, exists := linksMap[link]; exists {
		return nil // предотвращает повторное скачивание той же страницы
	}

	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Сохраняем текущую страницу
	pageName := getPageName(link)
	linksMap[link] = pageName
	err = saveToFile(pageName, resp.Body)
	if err != nil {
		return err
	}

	// Парсим HTML и извлекаем ссылки
	body, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	extractLinks(body, link, linksMap)

	// Скачиваем все найденные ссылки
	for _, v := range linksMap {
		if v != link { // избегаем скачивания уже обработанной страницы
			err := downloadSite(v, depthIndex+1, linksMap)
			if err != nil {
				fmt.Println("Ошибка при скачивании страницы:", v, err)
			}
		}
	}
	return nil
}

func extractLinks2(n *html.Node, baseURL string, linksMap map[string]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				// Обрабатываем относительные и абсолютные ссылки
				link := resolveURL(attr.Val, baseURL)
				if link != "" {
					pageName := getPageName(link)
					if _, exists := linksMap[link]; !exists {
						linksMap[link] = pageName
					}
				}
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c, baseURL, linksMap)
	}
}

func resolveURL(href, base string) string {
	// Обрабатываем ссылку относительно базового URL
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	parsedURL, err := url.Parse(href)
	if err != nil {
		return ""
	}
	// Объединяем базовый и относительный URL
	resolvedURL := baseURL.ResolveReference(parsedURL)
	return resolvedURL.String()
}*/

func saveToFile(link string) {
	resp, err := http.Get(link)
	if err != nil {
		return
	}
	newPage, err := os.Create(getPageName(link))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newPage.Close()
	_, err = io.Copy(newPage, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getPageName(link string) string {
	splittedLink := strings.Split(link, "/")
	return splittedLink[len(splittedLink)-1] + ".html"
}
