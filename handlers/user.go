package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gogo/status"
	"github.com/mxbikes/mxbikesclient.service.user/models"
	protobuffer "github.com/mxbikes/protobuf/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type User struct {
	protobuffer.UnimplementedUserServiceServer
	auth   *Auth
	logger logrus.Logger
}

// Return a new handler
func NewUserHandler(logger logrus.Logger) *User {
	return &User{logger: logger, auth: NewAuthHandler()}
}

func (e *User) GetUserByID(ctx context.Context, req *protobuffer.GetUserByIDRequest) (*protobuffer.GetUserByIDResponse, error) {
	if len(strings.TrimSpace(req.ID)) == 0 {
		e.logger.WithFields(logrus.Fields{"prefix": "SERVICE.User_GetUserByID"}).Errorf("request ID cannot be empty: {%s}", req.ID)
		return nil, status.Error(codes.Internal, "Error request value ID, is empty!")
	}

	url := fmt.Sprintf("https://dev-tm250wxm.us.auth0.com/api/v2/users/%s", req.ID)

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("authorization", "Bearer "+e.auth.GetAccessToken())
	res, _ := http.DefaultClient.Do(request)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// If response 404 user not found
	if res.StatusCode == 404 {
		e.logger.WithFields(logrus.Fields{"prefix": "SERVICE.User_GetUserByID"}).Errorf("no user found on request ID: {%s}", req.ID)
		return nil, status.Error(codes.Internal, "Error no user found on requested ID!")
	}

	var user models.User
	json.Unmarshal([]byte(body), &user)

	e.logger.WithFields(logrus.Fields{"prefix": "SERVICE.User_GetUserByID"}).Infof("user with id: {%s} ", req.ID)

	return &protobuffer.GetUserByIDResponse{User: models.UserToProto(&user)}, nil
}
