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

func (r FamilyRepo) All() (todos Family, err error) {
	err = r.Collection.Find(bson.M{}).All(&todos)
	return
}

func (r FamilyRepo) Create(todo *Family) (err bool) {
	query := bson.M{
		"name": todo.Name,
	}
	query_response, _ := r.Collection.Find(query).Count()
	if (query_response < 1) {
		if todo.Id.Hex() == "" {
			todo.Id = bson.NewObjectId()
		}
		if todo.Created.IsZero() {
			todo.Created = time.Now()
		}
		todo.Updated = time.Now()
		_, _ = r.Collection.UpsertId(todo.Id, todo)
		return true
	}
	return false
}

func (r FamilyRepo) Update(todo *Family) (err error) {
	var change = mgo.Change{
ReturnNew: true,
Update: bson.M{
		"$set": bson.M{
			"name": todo.Id,
//			"d": todo.Due,
//			"t": todo.Task,
		}}}
_, err = r.Collection.FindId(todo.Id).Apply(change, todo)

return
}
func (r FamilyRepo) Destroy(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}

