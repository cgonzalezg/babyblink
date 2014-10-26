package main

import (
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"net/http"
	family "babyblick-backend/family"
	event "babyblick-backend/event"
	child "babyblick-backend/child"
	member "babyblick-backend/familymember"
	"log"
)

var (
	database     *mgo.Database
	mongoSession *mgo.Session
	childLogic   *child.ChildRepo
	eventLogic   *event.EventRepo
	familyLogic  *family.FamilyRepo
	memberLogic  *member.MemberRepo
)

func initDB() {
	var err error
	if mongoSession, err = mgo.Dial("admin:mongo@ds063769.mongolab.com:63769/baby"); err != nil {
		panic(err)
	}
	log.Println("Connected to mongodb")

	database = mongoSession.DB("baby")
	childLogic = &child.ChildRepo {Collection: database.C("child")}
	eventLogic = &event.EventRepo {Collection: database.C("event")}
	familyLogic = &family.FamilyRepo {Collection: database.C("family")}
	memberLogic = &member.MemberRepo {Collection: database.C("familymember")}

}

func main() {
	r := mux.NewRouter()
	initDB()

	//Family
	r.HandleFunc("/family/create", familyLogic.FamilyCreate).Methods("POST")
	r.HandleFunc("/family/update", familyLogic.FamilyUpdate).Methods("POST")
	r.HandleFunc("/family/all", familyLogic.FamilyAll).Methods("GET")

	//Child
	r.HandleFunc("/child/create", childLogic.ChildCreate).Methods("POST")
	r.HandleFunc("/child/update", childLogic.ChildUpdate).Methods("POST")
	r.HandleFunc("/child/all", childLogic.ChildAll).Methods("GET")

	//Event
	r.HandleFunc("/event/create", eventLogic.Create).Methods("POST")
	r.HandleFunc("/event/update", eventLogic.Update).Methods("POST")
	r.HandleFunc("/event/all", eventLogic.All).Methods("GET")

	//Member
	r.HandleFunc("/member/create", memberLogic.Create).Methods("POST")
	r.HandleFunc("/member/update", memberLogic.Update).Methods("POST")
	r.HandleFunc("/member/all", memberLogic.All).Methods("GET")


	//	r.HandleFunc("/pictures", PictureHandler)
	http.Handle("/", r)
	http.ListenAndServe(":6969", nil)
}
