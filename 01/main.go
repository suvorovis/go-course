package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, e := fmt.Fprintln(os.Stderr, err)
		if e != nil {
			fmt.Println(err)
		}
		os.Exit(1)
	}
	fmt.Println(time)
}
