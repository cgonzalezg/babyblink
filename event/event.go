package event

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
	"time"
	"io"
	json "github.com/cgonzalezg/babyblick-backend/deserialize"



)

type Event struct {
	Id              bson.ObjectId            `json:"id" 				bson:"_id"`
	Owner           bson.ObjectId            `json:"owner" 			bson:"owner_id"`
	Name            string                    `json:"name" 			bson:"name"`
	Child           []bson.ObjectId            `json:"childs" 			bson:"childs"`
	Location        []int                    `json:"location" 			bson:"location"`
	data            time.Time                `json:"data"		bson:"date"`
	Photo           string                `json:"photo"			bson:"photo"`
	Created         time.Time                `json:"c"           	bson:"c"`
	Updated         time.Time                `json:"u,omitempty" 	bson:"u,omitempty"`
}

//HANDLERS
func (repo EventRepo) Create(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		item Event
	)
	json.ReadJson(r, &item)
	if err = repo.update(&item); err == nil {
		log.Println(item.Created)
		json.WriteJson(w, item)
	}else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not allow, family already create")
		return
	}
}

func (repo EventRepo) All(w http.ResponseWriter, r *http.Request) {
	var (
		items []Event
		err   error
	)
	if items, err = repo.all(); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, items)
}

func (repo EventRepo) Update(w http.ResponseWriter, r *http.Request) {
	var (
		item Event
	)
	json.ReadJson(r, &item)
	if err := repo.update(&item); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, item)

}
