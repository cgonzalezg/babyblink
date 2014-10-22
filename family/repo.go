package family

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	FamilyRepo struct {
		Collection *mgo.Collection
	}
)

//MONGO FUNCTIONS
func (repo FamilyRepo) create(member *Family) error {
	//check the family
	query := bson.M{
		"owner_id": member.Owner,
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

func (r FamilyRepo) exist(query bson.M) (bool, error) {
	query_response, error := r.Collection.Find(query).Count()
	return query_response < 1, error

}

func (r FamilyRepo) update(item *Family) (err error) {
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

func (r FamilyRepo) all() (items []Family, err error) {
	err = r.Collection.Find(bson.M{}).All(&items)
	return
}

func (r FamilyRepo) destroy(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}

