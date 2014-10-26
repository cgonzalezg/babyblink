package family

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
	"time"
	"io"
	json "bitbucket.org/babyblick/babyblick-backend/deserialize"
)

type Family struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Created   time.Time     `json:"c"            bson:"c"`
	Updated   time.Time   	`json:"u,omitempty"  bson:"u,omitempty"`
	Owner	bson.ObjectId 	`json:"owner" bson:"owner_id"`
}


func (repo FamilyRepo) FamilyCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		item Family
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

func (repo FamilyRepo) FamilyAll(w http.ResponseWriter, r *http.Request) {
	var (
		families []Family
		err   error
	)
	if families, err = repo.all(); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, families)

}

func (repo FamilyRepo) FamilyUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	var item Family
	json.ReadJson(r, &item)
	if err = repo.update(&item); err == nil {
		log.Println(item.Created)
		json.WriteJson(w, item)
	}else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println(err.Error())
		io.WriteString(w, "Not allow, family already create")
		return
	}

}
