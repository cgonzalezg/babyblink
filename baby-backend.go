package main

import (
	"log"
	"net/http"
	"os"

	auth "github.com/cgonzalezg/babyblick-backend/auth"
	child "github.com/cgonzalezg/babyblick-backend/child"
	event "github.com/cgonzalezg/babyblick-backend/event"
	family "github.com/cgonzalezg/babyblick-backend/family"
	member "github.com/cgonzalezg/babyblick-backend/familymember"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
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
	if mongoSession, err = mgo.Dial(os.Getenv("mongoUri")); err != nil {
		panic(err)
	}
	log.Println("Connected to mongodb")

	database = mongoSession.DB("baby")
	childLogic = &child.ChildRepo{Collection: database.C("child")}
	eventLogic = &event.EventRepo{Collection: database.C("event")}
	familyLogic = &family.FamilyRepo{Collection: database.C("family")}
	memberLogic = &member.MemberRepo{Collection: database.C("familymember")}

}

func endPoints(w http.ResponseWriter, r *http.Request) {

	// t := template.Must(template.ParseFiles("index.html"))

	// fmt.Println("got:", r.URL.Query())

	http.ServeFile(w, r, "index.html")

}

func main() {
	r := mux.NewRouter()
	initDB()

	r.HandleFunc("/auth", auth.Login).Methods("GET")

	//endpoints url
	r.HandleFunc("/", endPoints).Methods("GET")
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
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
