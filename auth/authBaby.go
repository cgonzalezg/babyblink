package auth

import (
	"bytes"
	"log"
	"net/http"
	"net/url"

	json "bitbucket.org/babyblick/babyblick-backend/deserialize"
)

type AccesToken struct {
	Token string `json:"access_token"`
	Type  string `json:"token_type"`
	ID    string `json:"id_token"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	// var jsonStr = []byte("client_id=CLIENT_ID&redirect_uri=REDIRECT_URI&client_secret=CLIENT_SECRET&code=AUTHORIZATION_CODE&grant_type=authorization_code")

	code := r.URL.Query()["code"][0]

	data := url.Values{}
	data.Set("client_id", "rlzj1WdHHXYWgoZJWRMT4yGhVIJZkKme")
	data.Add("redirect_uri", "bar")
	data.Add("client_secret", "qQA_CKfg-k6LR2phMSqWOgqw9H2oJDboapmvUo64zfU6n2pB-YCL1Qd3FHg4RhtU")
	data.Add("code", code)
	data.Add("grant_type", "authorization_code")

	req, err := http.Post("https://babyblick.auth0.com/oauth/token", "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))

	var token AccesToken

	json.ReadJsonRes(req, &token)
	//code take the user id

	if err != nil {
		log.Println("forbiden")
	}

	client := &http.Client{}

	req1, _ := http.NewRequest("GET", "https://babyblick.auth0.com/userinfo/", nil)

	// ...
	req.Header.Set("Authorization", "'Bearer "+token.Token+"'")
	resp, _ := client.Do(req1)
	var user interface{}
	json.ReadJsonRes(resp, user)

	json.WriteJson(w, user)
}
