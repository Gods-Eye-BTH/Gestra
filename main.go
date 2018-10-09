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
	Coords [2]int `json:"coords"`
	Size   [2]int `json:"size"`
}

//Barrier is a json object
type Barrier struct {
	ID     int    `json:"id"`
	Coords [2]int `json:"coords"`
	Size   [2]int `json:"size"`
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
		Coords: [2]int{1, 2},
		Size:   [2]int{3, 4},
	},
	Robot{
		ID:     1,
		Name:   "Tess",
		Coords: [2]int{5, 6},
		Size:   [2]int{7, 8},
	},
}

//Json object containing the barriers
var barriers = Barriers{
	Barrier{
		ID:     0,
		Coords: [2]int{1, 2},
		Size:   [2]int{3, 4},
	},
	Barrier{
		ID:     1,
		Coords: [2]int{5, 6},
		Size:   [2]int{7, 8},
	},
}

//index route
func index(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(robots)
}

//return a spesific robot route
func returnRobotByID(w http.ResponseWriter, r *http.Request) {
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
