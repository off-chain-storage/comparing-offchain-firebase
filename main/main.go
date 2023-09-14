package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	F "github.com/off-chain-storage/comparing-offchain-firebase/firebase"
	O "github.com/off-chain-storage/comparing-offchain-firebase/offchain"
	U "github.com/off-chain-storage/comparing-offchain-firebase/utils"
	// service "github.com/off-chain-storage/go-off-chain-storage/service"
)

func main() {
	U.Log_init()

	err := godotenv.Load("/root/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Firebase FireStore - CREATE
	startTime := time.Now()
	F.CreateDoc()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("FIREBASE CREATE TIME: %v", elapsedTime)

	// Firebase FireStore - READ
	startTime1 := time.Now()
	F.ReadDoc()
	endTime1 := time.Now()
	elapsedTime1 := endTime1.Sub(startTime1)
	fmt.Println("FIREBASE READ TIME: %v", elapsedTime1)

	// Firebase FireStore - UPDATE
	startTime2 := time.Now()
	F.UpdateDoc()
	endTime2 := time.Now()
	elapsedTime2 := endTime2.Sub(startTime2)
	fmt.Println("FIREBASE UPDATE TIME: %v", elapsedTime2)

	// Firebase FireStore - UPDATE
	startTime4 := time.Now()
	F.DeleteDoc()
	endTime4 := time.Now()
	elapsedTime4 := endTime4.Sub(startTime4)
	fmt.Println("FIREBASE DELETE TIME: %v", elapsedTime4)

	// Off-chain - CREATE
	startTime3 := time.Now()
	O.CreateDoc_off_chain()
	endTime3 := time.Now()
	elapsedTime3 := endTime3.Sub(startTime3)
	fmt.Printf("OFF_CHAIN STORAGE CREATE TIME: %v\n", elapsedTime3)

	// Off-chain - UPDATE
	startTime5 := time.Now()
	O.ReadDoc_off_chain()
	endTime5 := time.Now()
	elapsedTime5 := endTime5.Sub(startTime5)
	fmt.Printf("OFF_CHAIN STORAGE UPDATE TIME: %v\n", elapsedTime5)

	// Off-chain - READ
	startTime6 := time.Now()
	O.UpdateDoc_off_chain()
	endTime6 := time.Now()
	elapsedTime6 := endTime6.Sub(startTime6)
	fmt.Printf("OFF_CHAIN STORAGE UPDATE TIME: %v\n", elapsedTime6)

	// Off-chain - CREATE
	startTime7 := time.Now()
	O.DeleteDoc_off_chain()
	endTime7 := time.Now()
	elapsedTime7 := endTime7.Sub(startTime7)
	fmt.Printf("OFF_CHAIN STORAGE DELETE TIME: %v\n", elapsedTime7)
}
