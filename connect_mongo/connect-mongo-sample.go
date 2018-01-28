package main

  import (
	"fmt"
	"log"
    "time"

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
	Tile string
	Description string
	Company string
	Salary string
}

  func main() {

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

	col := session.DB(database).C(collection)
	
	err = col.Insert(&Job{"DevOps Engineer","Should be familiar with Jenkins Pipeline","Company XYZ","$10,000"},
					 &Job{"Senior Software Engineer","Should be familiar with golang","Company XYZ","$12,000"})

	if err != nil {
		log.Fatal(err)
	}

	count, err := col.Count()
    if err != nil {
        panic(err)
    }
	fmt.Println(fmt.Sprintf("Messages count: %d", count))
	
  }	