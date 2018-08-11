package main

import (
	"log"
	"os"
	"strconv"
)

const (
	toWorkFilename   = "to-work.png"
	fromWorkFilename = "from-work.png"
)

type environmentParameters struct {
	toWork        bool
	emailAddress  string
	emailPassword string
	name          string
	homeLocation  string
	workLocation  string
	apiKey        string
}

func parameters() environmentParameters {

	toWork, err := strconv.ParseBool(os.Getenv("TOWORK"))
	if err != nil {
		log.Fatal(err)
	}

	email := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	username := os.Getenv("NAME")
	homeLocation := os.Getenv("HOME_LOCATION")
	workLocation := os.Getenv("WORK_LOCATION")
	apiKey := os.Getenv("APIKEY")

	return environmentParameters{
		toWork:        toWork,
		emailAddress:  email,
		emailPassword: password,
		name:          username,
		homeLocation:  homeLocation,
		workLocation:  workLocation,
		apiKey:        apiKey,
	}
}
