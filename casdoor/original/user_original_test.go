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

package original

import (
	"fmt"
	"testing"
	"time"

	"github.com/casdoor/casdoor/object"
)

func TestGetUsers(t *testing.T) {
	initConfig()
	initAdapter()

	users := getUsersOriginal()
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
}

func TestSyncUsers(t *testing.T) {
	initConfig()
	initAdapter()
	object.InitAdapter()

	syncUsers()

	// run at every minute
	schedule := "* * * * *"
	err := ctab.AddJob(schedule, syncUsers)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Duration(1<<63 - 1))
}
