package familymember

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
	"time"
	"io"
	json "bitbucket.org/babyblick/babyblick-backend/deserialize"



)

type Member struct {
	Id            bson.ObjectId            `json:"id" 				bson:"_id"`
	User          bson.ObjectId            `json:"user" 			bson:"user_id"`
	Name          string                    `json:"name" 			bson:"name"`
	Family        bson.ObjectId            `json:"family" 			bson:"family_id"`
	Birthday      time.Time                `json:"birthday"		bson:"birthday"`
	Photo         string                `json:"photo"			bson:"photo"`
	Admin         bool                        `json:"admin" 	bson:"admin"`
	Created       time.Time                `json:"c"           	bson:"c"`
	Updated       time.Time                `json:"u,omitempty" 	bson:"u,omitempty"`
}

//HANDLERS
func (repo MemberRepo) Create(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		item Member
	)
	json.ReadJson(r, &item)
	if err = repo.create(item); err == nil {
		log.Println(item.Created)
		json.WriteJson(w, item)
	}else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not allow, family already create")
		return
	}
}

func (repo MemberRepo) All(w http.ResponseWriter, r *http.Request) {
	var (
		families []Member
		err   error
	)
	if families, err = repo.all(); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, families)
}

func (repo MemberRepo) Update(w http.ResponseWriter, r *http.Request) {
	var (
		item Member
	)
	json.ReadJson(r, &item)
	if err := repo.update(&item); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, item)

}
