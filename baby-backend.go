package main

import (
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"net/http"
	family "babyblick-backend/family"
	"log"
)

var (
	mongoSession *mgo.Session
	database     *mgo.Database
	familyLogic  *family.FamilyRepo
)


func initDB() {
	var err error
	if mongoSession, err = mgo.Dial("admin:mongo@ds063769.mongolab.com:63769/baby"); err != nil {
		panic(err)
	}
	log.Println("Connected to mongodb")

	database = mongoSession.DB("baby")
	familyLogic = &family.FamilyRepo {Collection: database.C("family")}
}

func main() {
	r := mux.NewRouter()
	initDB()
	//Family
	r.HandleFunc("/family/create", familyLogic.FamilyCreate).Methods("POST")
	r.HandleFunc("/family/update", familyLogic.FamilyUpdate).Methods("POST")
	r.HandleFunc("/family/all", familyLogic.FamilyAll).Methods("GET")


	// r.HandleFunc("/family/child", FamilyCreate).Methods("POST")
	// r.HandleFunc("/family/parent", FamilyCreate).Methods("POST")




	//	r.HandleFunc("/pictures", PictureHandler)
	http.Handle("/", r)
	http.ListenAndServe(":6969", nil)
}
