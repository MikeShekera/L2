package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

const ntpServer = "0.ru.pool.ntp.org"
const ntpServerDamaged = "0.ru.pool.ntp.og"

func main() {
	time, err := ntp.Time(ntpServer)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(time)

	time, err = ntp.Time(ntpServerDamaged)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(time)
}
