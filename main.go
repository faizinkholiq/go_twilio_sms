package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set account keys & information
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
    	Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE"))
	params.SetFrom(os.Getenv("FROM_PHONE"))
	params.SetBody("Hello bocah kuntul")

	resp, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println("SMS sent successfully!")
		response, _ := json.Marshal(*resp)
        fmt.Println("Response: " + string(response))
	}

}
