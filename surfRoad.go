package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Structure du champ "fields"
type Fields struct {
	SurfBreak       []string `json:"Surf Break"`
	DifficultyLevel int      `json:"Difficulty Level"`
	Destination     string   `json:"Destination"`
	Photos          []Photo  `json:"Photos"`
}

// Structure pour les photos
type Photo struct {
	ID         string     `json:"id"`
	URL        string     `json:"url"`
	Filename   string     `json:"filename"`
	Size       int        `json:"size"`
	Type       string     `json:"type"`
	Thumbnails Thumbnail  `json:"thumbnails"`
}

type Thumbnail struct {
	Small ThumbnailSize `json:"small"`
	Large ThumbnailSize `json:"large"`
	Full  ThumbnailSize `json:"full"`
}

type ThumbnailSize struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Structure pour un enregistrement (record)
type Record struct {
	ID     string `json:"id"`
	Fields Fields `json:"fields"`
}

type AllRecords []Record

var events = AllRecords{
	{
		ID: "rec5aF9TjMjBicXCK",
		Fields: Fields{
			SurfBreak:       []string{"Reef Break"},
			DifficultyLevel: 4,
			Destination:     "Pipeline",
		},
	},
	{
		ID: "recT98Z2El7YYwmc4",
		Fields: Fields{
			SurfBreak:       []string{"Point Break"},
			DifficultyLevel: 5,
			Destination:     "Skeleton Bay",
		},
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Record
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
            }

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent Record

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events[i] = updatedEvent
			json.NewEncoder(w).Encode(updatedEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}

func main() {
	//initEvents()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
