package child

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
	"time"
	"io"
	json "babyblick-backend/deserialize"
)

type Child struct {
	Id            bson.ObjectId    `json:"id" 			bson:"_id"`
	Name          string           `json:"name" 		bson:"name"`
	Family        bson.ObjectId    `json:"family" 		bson:"family_id"`
	Birthday      time.Time        `json:"birthday"	bson:"birthday"`
	Photo         string           `json:"photo"		bson:"photo"`
	Created       time.Time        `json:"c"           bson:"c"`
	Updated       time.Time        `json:"u,omitempty" bson:"u,omitempty"`
}



func (repo ChildRepo) ChildCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		item Child
	)

	json.ReadJson(r, &item)
	if err = repo.create(&item); err == nil {
		log.Println(item.Created)
		json.WriteJson(w, item)
	}else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not allow, family already create")
		return
	}
}

func (repo ChildRepo) ChildAll(w http.ResponseWriter, r *http.Request) {
	var (
		childs []Child
		err   error
	)
	if childs, err = repo.all(); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, childs)
}

func (repo ChildRepo) ChildUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		item Child
	)
	json.ReadJson(r, &item)
	item.Updated = time.Now()
	if err := repo.update(&item); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, item)

}


