package models

import (
	"time"

	protobuffer "github.com/mxbikes/protobuf/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	CreatedAt     time.Time `json:"created_at"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	Identities    []struct {
		Connection string `json:"connection"`
		Provider   string `json:"provider"`
		UserID     string `json:"user_id"`
		IsSocial   bool   `json:"isSocial"`
	} `json:"identities"`
	Name        string    `json:"name"`
	Nickname    string    `json:"nickname"`
	Picture     string    `json:"picture"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      string    `json:"user_id"`
	Username    string    `json:"username"`
	LastIP      string    `json:"last_ip"`
	LastLogin   time.Time `json:"last_login"`
	LoginsCount int       `json:"logins_count"`
}

func UserToProto(user *User) *protobuffer.User {
	return &protobuffer.User{
		ID:       user.UserID,
		Name:     user.Name,
		Nickname: user.Nickname,
		Username: user.Username,
		Picture:  user.Picture,
		Email:    user.Email,
		CreateAt: timestamppb.New(user.CreatedAt),
	}
}
