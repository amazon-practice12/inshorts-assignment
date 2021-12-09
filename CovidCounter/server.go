package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {

	// gocron.NewScheduler(time.UTC).Every(5).Second().Do(fmt.Println("Hello World"))
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(30).Minute().Do(refreshDetails)
	// scheduler starts running jobs and current thread continues to execute
	s1.StartAsync()

	e := echo.New()
	e.GET("/users/:id", getUser)
	e.GET("/details", getDetails)
	// e.POST("/users", saveUser)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	// refreshDetails()
	return c.String(http.StatusOK, id)
}

// e.GET("/details/:id", getUser)
func getDetails(c echo.Context) error {
	longitude := c.QueryParam("longitude")
	latitude := c.QueryParam("latitude")
	//get Distict and State by running reverse geocoding
	fmt.Println(longitude + "  " + latitude)
	// district := "Araria"
	// state := "BR"

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://eu1.locationiq.com/v1/reverse.php?key=pk.7b267195cbcb3707a6c59479bd302906&lat=26.397701&lon= 87.256081&format=json&statecode=true", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	// fmt.Println(string(ioutil.ReadAll(resp.Body)))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	location := LocationDetails{}

	json.Unmarshal(body, &location) // Convert JSON data into interface{} type
	state := location.Address.State
	district := location.Address.StateDistrict

	//get details from database using the district and state info
	details := getInfoFromDatabase(state, district)
	return c.JSONPretty(http.StatusOK, details, " ")
}

func getInfoFromDatabase(state string, district string) CovidDetails {
	fmt.Println("State: " + state + " District: " + district)
	getDatabaseConnection()
	covidInfo := CovidDetails{
		State:           state,
		LastUpdatedDate: time.Now().GoString(),
	}
	return covidInfo
}

func getDatabaseConnection() {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://rootuser:testuser1234@cluster0.n409n.mongodb.net/inshorts?retryWrites=true&w=majority"))
	// clientOptions := options.Client().
	// ApplyURI("mongodb+srv://rootuser:testuser1234@cluster0.n409n.mongodb.net/sample_airbnb?retryWrites=true&w=majority")

	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://rootuser:testuser1234@cluster0.n409n.mongodb.net/sample_airbnb?retryWrites=true&w=majority"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	// clientOptions := options.Client().ApplyURI("mongodb+srv://rootuser:testuser1234@cluster0.n409n.mongodb.net/")
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// collection = client.Database("sample_airbnb").Collection("listingsAndReviews")

	client, ctx, cancel, err := connect("mongodb+srv://rootuser:testuser1234@cluster0.n409n.mongodb.net/")
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	ping(client, ctx)
}
