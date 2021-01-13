package main

import "log"

func main() {
	if err := startApp(); err != nil {
		log.Fatal(err)
	}
}
