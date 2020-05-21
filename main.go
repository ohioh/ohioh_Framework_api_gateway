package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

type Locationlatitude struct {
	LocationID string  `json:"location_id"`
	Latitude   float64 `json:"latitude"`
	Departure  bool    `json:"departure"`
}

type BluetoothEncounter struct {
	UserID          string `json:"user_id"`
	EncounterUserId string `json:"encounter_user_id"`
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

func insertBluethoothEncounter(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	bluetoothEncounter := BluetoothEncounter{}

	erp := json.NewDecoder(r.Body).Decode(&bluetoothEncounter)
	fmt.Println(erp)

	if erp != nil {
		log.Fatalf("Error parsing request body1")
	}

	pl, errpl := json.Marshal(bluetoothEncounter)

	if errpl != nil {
		log.Fatalf("Error parsing request body2")
	}

	postURL := os.Getenv("BULETOOTH_ENCOUNTER_URI") + "/" + os.Getenv("BULETOOTH_ENCOUNTER_SERVICE")
	fmt.Println(postURL)
	response, err := http.Post(postURL, "application/json", bytes.NewBuffer(pl))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		json.NewEncoder(w).Encode(response.Body)
	}
}

func insertLocationLatitude(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	locationLat := Locationlatitude{}

	erp := json.NewDecoder(r.Body).Decode(&locationLat)

	if erp != nil {
		log.Fatalf("Error parsing request body1")
	}

	pl, errpl := json.Marshal(locationLat)

	if errpl != nil {
		log.Fatalf("Error parsing request body2")
	}

	postURL := os.Getenv("LOCATION_LAT_URI") + "/" + os.Getenv("LOCATION_LAT_SERVICE")
	fmt.Println(postURL)
	response, err := http.Post(postURL, "application/json", bytes.NewBuffer(pl))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		json.NewEncoder(w).Encode(response.Body)
	}
}

// jwt authenticaton

var mySigningKey = []byte(os.Getenv("SIGNING_KEY"))

func getToken(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Ellit Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something wernt wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil

}

//token verification

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	gatewayPrefix := "/gateway/ohioh"
	gatewayRouter := mux.NewRouter().StrictSlash(true)
	gatewayRouter.HandleFunc(gatewayPrefix+"/", defaultRoute)
	gatewayRouter.HandleFunc(gatewayPrefix+"/users", isAuthorized(allUsers)).Methods(http.MethodGet)
	//http.Handle(gatewayPrefix+"/users", isAuthorized(allUsers))
	gatewayRouter.HandleFunc(gatewayPrefix+"/users", insertUserRecord).Methods(http.MethodPost)
	gatewayRouter.HandleFunc(gatewayPrefix+"/user_locations", insertUserLocationRecord).Methods(http.MethodPost)
	gatewayRouter.HandleFunc(gatewayPrefix+"/location_latitude", insertLocationLatitude).Methods(http.MethodPost)
	gatewayRouter.HandleFunc(gatewayPrefix+"/bluetooth_encounter", insertBluethoothEncounter).Methods(http.MethodPost)

	gatewayRouter.HandleFunc(gatewayPrefix+"/token", getToken)
	log.Fatal(http.ListenAndServe(":8000", gatewayRouter))
}

func main() {
	handleRequests()
}
