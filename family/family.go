package family

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
	"encoding/json"
	"time"
	"strconv"
	"io"
	"io/ioutil"
)

type Family struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Created   time.Time     `json:"c"            bson:"c"`
	Updated   time.Time     `json:"u,omitempty"  bson:"u,omitempty"`
}

func writeJson(w http.ResponseWriter, v interface{}) {
	// avoid json vulnerabilities, always wrap v in an object literal
	doc := map[string]interface{}{"data": v}

	if data, err := json.Marshal(doc); err != nil {
		log.Printf("Error marshalling json: %v", err)
	} else {
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func readJson(r *http.Request, v interface{}) bool {
	defer r.Body.Close()

	var (
		body []byte
		err  error
	)

	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("ReadJson couldn't read request body %v", err)
		return false
	}

	if err = json.Unmarshal(body, v); err != nil {
		log.Printf("ReadJson couldn't parse request body %v", err)
		return false
	}

	return true
}

func (repo FamilyRepo) FamilyCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
	)
	decoder := json.NewDecoder(r.Body)
	var family  map[string] string
	err = decoder.Decode(&family)
	if err != nil {
		log.Println("error" + family["name"])
	}
	item := Family{ Id: bson.NewObjectId(), Name: family["name"]}
	log.Println("print " + item.Name)
	if err = repo.Create(&item); err == nil {
		log.Println(item.Created)
		writeJson(w, item)
	}else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not allow, family already create")
		return
	}


}

func (repo FamilyRepo) FamilyAll(w http.ResponseWriter, r *http.Request) {
	var (
		families []Family
		err   error
	)
	if families, err = repo.All(); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	writeJson(w, families)

}

func (repo FamilyRepo) FamilyUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	var item Family
	readJson(r, &item)
	if err = repo.Update(&item); err == nil {
		log.Println(item.Created)
		writeJson(w, item)
	}else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println(err.Error())
		io.WriteString(w, "Not allow, family already create")
		return
	}

}
