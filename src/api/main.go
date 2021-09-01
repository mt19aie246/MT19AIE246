package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Scrap struct {
	Text      string    `json:"text" bson:"text"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
}

var scraps *mgo.Collection

func main() {
	session, err := mgo.Dial("mongo:27017")
	if err != nil {
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	scraps = session.DB("app").C("scraps")
	router := mux.NewRouter()
	router.HandleFunc("/scraps", saveScrap).
		Methods("POST")
	router.HandleFunc("/scraps", readScrap).
		Methods("GET")
	http.ListenAndServe(":8080", cors.AllowAll().Handler(router))
}

func saveScrap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	scrap := &Scrap{}
	err = json.Unmarshal(data, scrap)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	scrap.CreatedAt = time.Now().UTC()
	if err := scraps.Insert(scrap); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseJSON(w, scrap)
}

func responseError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func readScrap(w http.ResponseWriter, r *http.Request) {
	result := []Scrap{}
	if err := scraps.Find(nil).Sort("-created_at").All(&result); err != nil {
		responseError(w, err.Error(), http.StatusPreconditionFailed)
	} else {
		responseJSON(w, result)
	}
}