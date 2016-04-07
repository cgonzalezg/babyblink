package parent

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
	"time"
	"io"
	json "github.com/cgonzalezg/babyblick-backend/deserialize"



)

type Parent struct {
	Id            bson.ObjectId        	`json:"id" 				bson:"_id"`
	Name          string            	`json:"name" 			bson:"name"`
	User          bson.ObjectId        	`json:"user" 			bson:"user_id"`
	Birthday      time.Time            	`json:"birthday"		bson:"birthday"`
	Photo         string            	`json:"photo"			bson:"photo"`
	Created       time.Time            	`json:"c"           	bson:"c"`
	Updated       time.Time        		`json:"u,omitempty" 	bson:"u,omitempty"`

}


func (repo ParentRepo) ParentCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		item Parent
	)
	json.ReadJson(r, &item)
	if err = repo.create(&item); err == nil {
		json.WriteJson(w, item)
	}else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not allow, family already create")
		return
	}
}

func (repo ParentRepo) ParentAll(w http.ResponseWriter, r *http.Request) {
	var (
		families []Parent
		err   error
	)
	if families, err = repo.all(); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, families)
}

func (repo ParentRepo) ParentUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		item Parent
	)
	json.ReadJson(r, &item)
	if err := repo.update(&item); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, item)

}
