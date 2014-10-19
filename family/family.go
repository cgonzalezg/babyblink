package family

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
	"encoding/json"
	"time"
//	"fmt"
)

type Family struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Created   time.Time     `json:"c"            bson:"c"`
	Updated   time.Time     `json:"u,omitempty"  bson:"u,omitempty"`
}

func (repo FamilyRepo) FamilyCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var family  map[string] string
	err := decoder.Decode(&family)
	if err != nil {
		log.Println("error" + family["name"])
	}
	log.Println("print " + family["name"])

	item := &Family{ Id: bson.NewObjectId(), Name: family["name"]}

	var response string = ""
	if (repo.Create(item)) {
		response = string(item.Id)

	}else {
		log.Println("erroe")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Hello " + response))
}

func FamilyUpdate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var family  map[string] string
	err := decoder.Decode(&family)
	if err != nil {
		log.Println("error" + family["name"])
	}
	log.Println("print " + family["name"])

	session, err := mgo.Dial("admin:mongo@ds063769.mongolab.com:63769/baby")
	defer session.Close()



	c := session.DB("baby").C("family")
	query := bson.M{
		"name": family["name"],
	}
	query_response, _ := c.Find(query).Count()
	var response string = ""
	if (query_response < 1) {

		entry := Family{ Id: bson.NewObjectId(), Name: family["name"]}
		c.Insert(entry)
		if err != nil {
			log.Fatal(err)
		}
		response = string(entry.Id.Hex())
	}else {
		response = "mierda"
	}

	w.Write([]byte("Hello " + response))
}
