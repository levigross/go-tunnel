package main

import "log"

func checkError(err error, fatal bool) {
	if err != nil {
		if fatal {
			log.Fatalln("Fatal error:", err.Error())
		} else {
			log.Println("Error:", err.Error())
		}

	}
}
