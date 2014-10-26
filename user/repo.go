package user

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	mgo "gopkg.in/mgo.v2"



)

type EventRepo struct {
	Collection *mgo.Collection
}

//MONGO FUNCTIONS
func (repo EventRepo) create(item *Event) error {
	//check the family
	query := bson.M{
		"child": item.Child,
		"name": item.Name,
	}

	exist, err := repo.exist(query)
	if !exist && err != nil {
		item.Created = time.Now()
		item.Updated = time.Now()
		err = repo.update(item)
		return err
	}

	return err
}

func (r EventRepo) exist(query bson.M) (bool, error) {
	query_response, error := r.Collection.Find(query).Count()
	return query_response < 1, error

}

func (r EventRepo) update(item *Event) (err error) {
	var id bson.ObjectId
	if item.Id.Hex() == "" {
		id = bson.NewObjectId()
	} else {
		id = item.Id
	}
	item.Updated = time.Now()
	_, err = r.Collection.UpsertId(id, item)

	return
}

func (r EventRepo) all() (items []Event, err error) {
	err = r.Collection.Find(bson.M{}).All(&items)
	return
}

func (r EventRepo) destroy(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}
