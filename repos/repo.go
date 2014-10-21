package repos

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Repo struct {
		Collection *mgo.Collection
	}
)

func (r Repo) All() (items []interface{}, err error) {
	err = r.Collection.Find(bson.M{}).All(&items)
	return
}

func (r Repo) Update(id bson.ObjectId , item *interface{}) (err error) {

	if id.Hex() == "" {
		id = bson.NewObjectId()
	}
	_, err = r.Collection.UpsertId(id, item)

	return
}

func (r Repo) Exist(query bson.M) (bool, error) {


	query_response, error := r.Collection.Find(query).Count()

	return query_response < 1, error

}

func (r Repo) Destroy(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}

