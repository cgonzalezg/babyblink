package child

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
	"time"
	"io"
	"src/babyblick-backend/repos"
	json "src/babyblick-backend/deserialize"



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

type ChildRepo struct {
	Repo
}

func (repo ChildRepo) create(child Child) error {
	//check the family
	query := bson.M{
		"family_id": bson.ObjectIdHex(child.Family),
		"name":
	}
	exis, err := repo.Exist(query)
	if exis && err != nil {
		err = repo.Update(child)
		return err
	}

	return err
}

func (repo ChildRepo) ChildCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		item Child
	)
	json.ReadJson(r, &item)
	if err = repo.Create(&item); err == nil {
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
		families []Family
		err   error
	)
	if families, err = repo.All(); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, families)
}

func (repo ChildRepo) ChildUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	var item Family
	json.ReadJson(r, &item)
	if err = repo.Update(&item); err == nil {
		log.Println(item.Created)
		json.WriteJson(w, item)
	}else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println(err.Error())
		io.WriteString(w, "Not allow, family already create")
		return
	}
}


