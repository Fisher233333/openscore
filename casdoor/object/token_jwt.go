// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("CasdoorSecret")

type Claims struct {
	Organization string `json:"organization"`
	Username     string `json:"username"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Affiliation  string `json:"affiliation"`
	Tag          string `json:"tag"`
	Language     string `json:"language"`
	Score        int    `json:"score"`
	IsAdmin      bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func generateJwtToken(application *Application, user *User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(application.ExpireInHours) * time.Hour)

	claims := Claims{
		Organization: user.Owner,
		Username:     user.Name,
		Type:         user.Type,
		Name:         user.DisplayName,
		Avatar:       user.Avatar,
		Email:        user.Email,
		Phone:        user.Phone,
		Affiliation:  user.Affiliation,
		Tag:          user.Tag,
		Language:     user.Language,
		Score:        user.Score,
		IsAdmin:      user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			Audience:  application.ClientId,
			ExpiresAt: expireTime.Unix(),
			Id:        "",
			IssuedAt:  nowTime.Unix(),
			Issuer:    "casdoor",
			NotBefore: nowTime.Unix(),
			Subject:   user.Id,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseJwtToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
