package child

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	mgo "gopkg.in/mgo.v2"



)

type ChildRepo struct {
	Collection *mgo.Collection
}

//MONGO FUNCTIONS
func (repo ChildRepo) create(member *Child) error {
	//check the family
	query := bson.M{
		"family_id": member.Family,
		"name": member.Name,
	}

	exist, err := repo.exist(query)
	if !exist && err != nil {
		member.Created = time.Now()
		member.Updated = time.Now()
		err = repo.update(member)
		return err
	}

	return err
}

func (r ChildRepo) exist(query bson.M) (bool, error) {
	query_response, error := r.Collection.Find(query).Count()
	return query_response < 1, error

}

func (r ChildRepo) update(item *Child) (err error) {
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

func (r ChildRepo) all() (items []Child, err error) {
	err = r.Collection.Find(bson.M{}).All(&items)
	return
}

func (r ChildRepo) destroy(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}
