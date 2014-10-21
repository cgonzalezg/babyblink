package family

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"errors"
	"log"
)

type (
	FamilyRepo struct {
		Collection *mgo.Collection
	}
)

func (r FamilyRepo) All() (todos []Family, err error) {
	err = r.Collection.Find(bson.M{}).All(&todos)
	return
}

func (r FamilyRepo) Create(fam *Family) (err error) {
	query := bson.M{
		"name": fam.Name,
	}

	log.Printf("name", fam.Name)

	query_response, error := r.Collection.Find(query).Count()
	if error != nil {
		log.Printf("error query", fam.Name)
	}
	log.Println(query_response)
	if (query_response < 1) {
		if fam.Id.Hex() == "" {
			fam.Id = bson.NewObjectId()
		}
		if fam.Created.IsZero() {
			fam.Created = time.Now()
		}
		fam.Updated = time.Now()
		_, err = r.Collection.UpsertId(fam.Id, fam)

		if err != nil {
			log.Println("error to insert")
		}
		return
	}
	log.Printf("mierdaca")
	err = errors.New("Not allow")
	return err
}

func (r FamilyRepo) Update(todo *Family) (err error) {
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"name": todo.Name,
				"u": time.Now(),
			}}}
	_, err = r.Collection.FindId(todo.Id).Apply(change, todo)

	return
}
func (r FamilyRepo) Destroy(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}

