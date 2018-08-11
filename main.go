package main

import "log"

func main() {
	toWork := parameters().toWork

	log.Printf("Generating route, toWork: %t", toWork)
	generateGraph(toWork)
	sendEmail(toWork)
}
