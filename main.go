package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//port to run on
var port = 8080

//Robot is a json object
type Robot struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

//Barrier is a json object
type Barrier struct {
	ID     int `json:"id"`
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

//Robots is an array of the type Robot
type Robots []Robot

//Barriers is an array of the type Barrier
type Barriers []Barrier

//Json object containing the robots
var robots = Robots{
	Robot{
		ID:     0,
		Name:   "Curiosity",
		X:      1,
		Y:      2,
		Width:  8,
		Height: 7,
	},
	Robot{
		ID:     1,
		Name:   "Tess",
		X:      3,
		Y:      4,
		Width:  10,
		Height: 15,
	},
	Robot{
		ID:     2,
		Name:   "Oportunity",
		X:      9,
		Y:      5,
		Width:  13,
		Height: 9,
	},
	Robot{
		ID:     3,
		Name:   "Mars III",
		X:      22,
		Y:      18,
		Width:  6,
		Height: 12,
	},
	Robot{
		ID:     4,
		Name:   "Beagle 2",
		X:      32,
		Y:      45,
		Width:  10,
		Height: 15,
	},
}

//Json object containing the barriers
var barriers = Barriers{
	Barrier{
		ID:     0,
		X:      16,
		Y:      34,
		Width:  5,
		Height: 8,
	},
	Barrier{
		ID:     1,
		X:      12,
		Y:      34,
		Width:  22,
		Height: 15,
	},
	Barrier{
		ID:     2,
		X:      34,
		Y:      22,
		Width:  12,
		Height: 13,
	},
	Barrier{
		ID:     3,
		X:      26,
		Y:      12,
		Width:  14,
		Height: 17,
	},
}

//index route
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w,
		`<h1>GESTRA API</h1>
<p>
    Welcome to a modified instance of the <b>Gestra api</b>.<br>
    the following routes are available:
    <pre>
    /robots

    /robot/{id}

    /barriers
    </pre>
</p>`)
}

//display all robots route
func allRobots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(robots)
}

//return a spesific robot route
func returnRobotByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	key := vars["id"]
	convd, err := strconv.Atoi(key)
	if err == nil {
		if convd > len(robots)-1 || convd < 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Out of range")
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(robots[convd])
		}
	} else {
		fmt.Fprintf(w, "parse error on key: "+key)
	}
}

func allBarriers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barriers)
}

func requestHandler() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", index)
	muxRouter.HandleFunc("/robots", allRobots)
	muxRouter.HandleFunc("/robot/{id}", returnRobotByID)
	muxRouter.HandleFunc("/barriers", allBarriers)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), muxRouter))
}

func main() {
	fmt.Println("Loading robots...")
	fmt.Println("Loading barriers...")
	fmt.Println("GESTRA is now running\nport:" +
		strconv.Itoa(port) + "\n---")
	requestHandler()
}
