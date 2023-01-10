package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mxbikes/mxbikesclient.service.user/models"
)

var lastAuth models.AuthResponse

type Auth struct {
	auth0Connection string
}

func NewAuthHandler(auth0 string) *Auth {
	return &Auth{auth0Connection: auth0}
}

func (e Auth) GetAccessToken() (string, error) {
	if IsValid(lastAuth) {
		return lastAuth.AccessToken, nil
	}

	newAuth, err := e.requestToken()
	if err != nil {
		return "", err
	}

	return newAuth.AccessToken, err
}

func (e Auth) requestToken() (*models.AuthResponse, error) {
	url := "https://dev-tm250wxm.us.auth0.com/oauth/token"
	payload := strings.NewReader(e.auth0Connection)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 401 {
		return nil, errors.New("Unable to request access_token, unauthorized")
	}

	json.Unmarshal([]byte(body), &lastAuth)

	lastAuth.ExpiresAt = time.Now().Add(time.Duration(time.Duration(lastAuth.ExpiresIn)) * time.Second)
	return &lastAuth, nil
}

func IsValid(auth models.AuthResponse) bool {
	if time.Now().After(auth.ExpiresAt) {
		return false
	}

	return true
}
