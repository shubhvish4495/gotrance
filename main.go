package main

import (
	"gotrance/pkg/service"
	"log"
	"time"
)

func main() {

	//get db instance
	db := service.GetDbInstance()

	//get current epoch timestamp
	currEpochTs := time.Now().UTC().Unix()

	//get next db runtime
	waitTs := db.GetNextRuntime()

	//waitforit interval
	waitForIt := waitTs - currEpochTs

	log.Println("Going to sleep before next execution wake up call")

	//sleep for the same
	time.Sleep(time.Duration(waitForIt) * time.Second)

	log.Println("Program back up")

	//go for the transactional step
	db.TransactionalStep()

	//print final execution
	log.Println("Program served it's purpose exiting now")
}
