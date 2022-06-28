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

package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"golang.org/x/net/proxy"
)

var httpClient *http.Client

func InitHttpClient() {
	httpProxy := beego.AppConfig.String("httpProxy")
	if httpProxy == "" {
		httpClient = &http.Client{}
		return
	}

	// https://stackoverflow.com/questions/33585587/creating-a-go-socks5-client
	dialer, err := proxy.SOCKS5("tcp", httpProxy, nil, proxy.Direct)
	if err != nil {
		panic(err)
	}

	tr := &http.Transport{Dial: dialer.Dial}
	httpClient = &http.Client{
		Transport: tr,
	}

	//resp, err2 := httpClient.Get("https://google.com")
	//if err2 != nil {
	//	panic(err2)
	//}
	//defer resp.Body.Close()
	//println("Response status: %s", resp.Status)
}

func (c *ApiController) ResponseError(error string) {
	c.Data["json"] = Response{Status: "error", Msg: error}
	c.ServeJSON()
}

func (c *ApiController) ResponseErrorWithData(error string, data interface{}) {
	c.Data["json"] = Response{Status: "error", Msg: error, Data: data}
	c.ServeJSON()
}

func (c *ApiController) RequireSignedIn() (string, bool) {
	userId := c.GetSessionUsername()
	if userId == "" {
		resp := Response{Status: "error", Msg: "Please sign in first"}
		c.Data["json"] = resp
		c.ServeJSON()
		return "", false
	}
	return userId, true
}
