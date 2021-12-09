package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func refreshDetails() {
	fmt.Println("Fetching Staticstics")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://data.covid19india.org/v4/min/data.min.json", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(responseBody))
	ioutil.WriteFile("test1.json", responseBody, 777)
	fmt.Println("Fetching complete, updating database")
	// Update database
	fmt.Println("Database update complete, Infomation up-to-date as on: " + time.Now().Format("2006-01-02 03:04:05 PM"))
}
