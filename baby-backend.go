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
)

func main() {
	var err error
	r := mux.NewRouter()
	//	r.HandleFunc("/", HomeHandler)
	if mongoSession, err = mgo.Dial("admin:mongo@ds063769.mongolab.com:63769/baby"); err != nil {
		panic(err)
	}
	log.Println("Connected to mongodb")

	database = mongoSession.DB("baby")
	//	repo.Collection = database.C("family")
	//Family
	familyRepo := family.FamilyRepo {Collection: database.C("family")}
	r.HandleFunc("/family/create", familyRepo.FamilyCreate).Methods("POST")
	r.HandleFunc("/family/update", familyRepo.FamilyUpdate).Methods("POST")
	r.HandleFunc("/family/all", familyRepo.FamilyAll).Methods("GET")
	// r.HandleFunc("/family/child", FamilyCreate).Methods("POST")
	// r.HandleFunc("/family/parent", FamilyCreate).Methods("POST")




	//	r.HandleFunc("/pictures", PictureHandler)
	http.Handle("/", r)
	http.ListenAndServe(":6969", nil)
}
