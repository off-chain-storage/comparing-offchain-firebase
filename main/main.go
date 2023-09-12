package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	F "github.com/off-chain-storage/comparing-offchain-firebase/firebase"
	U "github.com/off-chain-storage/comparing-offchain-firebase/utils"
)

func main() {
	U.Log_init()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	startTime := time.Now()

	ctx := context.Background()
	client := F.DBClient()

	wr, err := client.Doc("User/User1").Create(ctx, map[string]interface{}{
		"capital": "Denver",
		"pop":     5.5,
	})
	if err != nil {
		log.Fatalf("firestore Doc Create error:%s\n", err)
	}
	endTime := time.Now()

	// 경과 시간 계산
	elapsedTime := endTime.Sub(startTime)

	// 경과 시간 출력
	fmt.Printf("작업이 완료되는 데 걸린 시간: %v\n", elapsedTime)

	fmt.Println(wr.UpdateTime)
}
