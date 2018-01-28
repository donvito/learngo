package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

const (
	hosts      = "localhost:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "jobs"
)

type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Salary      string `json:"salary"`
}

type DataStore struct {
    session *mgo.Session
}

func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/jobs", jobsHandler)

	log.Fatal(http.ListenAndServe(":9090", router))

}

func connectToMongo() (session *mgo.Session){

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	return

}

func jobsHandler(w http.ResponseWriter, r *http.Request) {

	session := connectToMongo()

	col := session.DB(database).C(collection)

	result := []Job{}

	col.Find(bson.M{"title": bson.RegEx{"Engineer", ""}}).All(&result)

	jsonString, err := json.Marshal(result)

	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(jsonString))

}
