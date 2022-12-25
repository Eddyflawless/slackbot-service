package helpers

import (
	"log"
	"time"
)

func StartCron(worker func(), duration int) {

	ticker := time.NewTicker(time.Second * time.Duration(duration)) // for x seconds

	for t := range ticker.C {
		//
		worker()

		log.Print(t)
	}
}
