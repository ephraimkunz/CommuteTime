package main

import (
	"fmt"
	"time"

	gomail "gopkg.in/gomail.v2"
)

func sendEmail(toWork bool) {
	m := gomail.NewMessage()
	email := parameters().emailAddress
	password := parameters().emailPassword
	username := parameters().name

	now := time.Now()
	noon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())

	m.SetHeader("From", email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Commute report for "+now.Format("Mon Jan 2"))

	m.SetBody("text/html", fmt.Sprintf("Hi %s,<br>Here's your commute information for %s:<br><br>", username, timeOfDay(now, noon)))

	if toWork {
		m.Attach(toWorkFilename)
	} else {
		m.Attach(fromWorkFilename)
	}

	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
