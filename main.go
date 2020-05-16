package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//User this struct is used for mapping ohioh user microservice api calls
type User struct {
	UserID                string `json:"user_id"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Age                   int    `json:"age"`
	Country               string `json:"country"`
	IsInfected            bool   `json:"is_infected"`
	TrackingSaveDuration  int    `json:"tracking_save_duration"`
	BluetoothSaveDuration int    `json:"bluetooth_save_duration"`
	Phone                 string `json:"phone"`
	ZipCode               string `json:"zip_code"`
}

//UserLocation this struct is used for mapping ohion user location microservice api calls
type UserLocation struct {
	UserID       string  `json:"user_id"`
	LocationID   string  `json:"location_id"`
	LocationType int     `json:"location_type"`
	Longitude    float64 `json:"longitude"`
	Arrival      bool    `json:"arrival"`
	Speed        int     `json:"speed"`
	Splitted     bool    `json:"splitted"`
}

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "default route")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	users := []User{}
	getURL := os.Getenv("USER_URI") + "/" + os.Getenv("USER_SERVICE")
	response, err := http.Get(getURL)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		fmt.Printf("whats wrong? ")
		json.NewDecoder(response.Body).Decode(&users)
		json.NewEncoder(w).Encode(users)
	}
}

func insertUserRecord(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	user := User{}

	errParse := json.NewDecoder(r.Body).Decode(&user)

	if errParse != nil {
		log.Fatalf("Error parsing request body")
	}

	pl, errpl := json.Marshal(user)

	if errpl != nil {
		log.Fatalf("Error parsing request body")
	}

	postURL := os.Getenv("USER_URI") + "/" + os.Getenv("USER_SERVICE")
	response, err := http.Post(postURL, "application/json", bytes.NewBuffer(pl))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		json.NewEncoder(w).Encode(response.Body)
	}
}

func insertUserLocationRecord(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	userLocation := UserLocation{}

	errParse := json.NewDecoder(r.Body).Decode(&userLocation)

	if errParse != nil {
		log.Fatalf("Error parsing request body")
	}

	pl, errpl := json.Marshal(userLocation)

	if errpl != nil {
		log.Fatalf("Error parsing request body")
	}

	postURL := os.Getenv("USER_LOCATION_URI") + "/" + os.Getenv("USER_LOCATION_SERVICE")
	response, err := http.Post(postURL, "application/json", bytes.NewBuffer(pl))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		json.NewEncoder(w).Encode(response.Body)
	}
}

func handleRequests() {
	gatewayPrefix := "/gateway/ohioh"
	gatewayRouter := mux.NewRouter().StrictSlash(true)
	gatewayRouter.HandleFunc(gatewayPrefix+"/", defaultRoute)
	gatewayRouter.HandleFunc(gatewayPrefix+"/users", allUsers).Methods(http.MethodGet)
	gatewayRouter.HandleFunc(gatewayPrefix+"/users", insertUserRecord).Methods(http.MethodPost)
	gatewayRouter.HandleFunc(gatewayPrefix+"/user_locations", insertUserLocationRecord).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8000", gatewayRouter))
}

func main() {
	handleRequests()
}
