package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mxbikes/mxbikesclient.service.user/models"
)

var lastAuth models.AuthResponse

type Auth struct {
}

func NewAuthHandler() *Auth {
	return &Auth{}
}

func (e Auth) GetAccessToken() string {
	if valid(lastAuth) {
		return lastAuth.AccessToken
	}

	return requestToken().AccessToken
}

func requestToken() models.AuthResponse {
	url := "https://dev-tm250wxm.us.auth0.com/oauth/token"
	payload := strings.NewReader("{\"client_id\":\"s0TL2fboqrS7P8kVFOJI30Yip6805zT1\",\"client_secret\":\"5yxd9phctrfIh8C2w-jn-1V0vu8_R2VvfAvHeRUyqbv8JhOUnkRPlF2W5-7mTYCA\",\"audience\":\"https://dev-tm250wxm.us.auth0.com/api/v2/\",\"grant_type\":\"client_credentials\"}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal([]byte(body), &lastAuth)

	lastAuth.ExpiresAt = time.Now().Add(time.Duration(time.Duration(lastAuth.ExpiresIn)) * time.Second)
	return lastAuth
}

func valid(auth models.AuthResponse) bool {
	if time.Now().After(auth.ExpiresAt) {
		return false
	}

	return true
}
