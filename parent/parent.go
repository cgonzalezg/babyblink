package parent

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
	"time"
	"io"
	"src/babyblick-backend/repos"
	json "src/babyblick-backend/deserialize"



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

type ParentRepo struct {
	Repo
}

func (repo ParentRepo) create(parent Parent) error {
	//check the family
	query := bson.M{
		"family_id": bson.ObjectIdHex(parent.Family),
		"name": parent.Name,
	}

	exis, err := repo.Exist(query)
	if exis && err != nil {
		err = repo.Update(parent.Id, parent)
		return err
	}

	return err
}

func (repo ParentRepo) ParentCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		item Parent
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

func (repo ParentRepo) ParentAll(w http.ResponseWriter, r *http.Request) {
	var (
		families []Parent
		err   error
	)
	if families, err = repo.All(); err != nil {
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
	if err := repo.Update(item.Id, item); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	json.WriteJson(w, item)

}


