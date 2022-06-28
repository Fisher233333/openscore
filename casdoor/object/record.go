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
	"github.com/casdoor/casdoor/util"
)

type Records struct {
	Id     	int		`xorm:"int notnull pk autoincr" json:"id"`
	Record 	util.Record	`xorm:"extends"`
}

func AddRecord(record *util.Record) bool {
	records := new(Records)
	records.Record = *record

	affected, err := adapter.Engine.Insert(records)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func GetRecordCount() int {
	count, err := adapter.Engine.Count(&Records{})
	if err != nil {
		panic(err)
	}

	return int(count)
}

func GetRecords() []*Records {
	records := []*Records{}
	err := adapter.Engine.Desc("id").Find(&records)
	if err != nil {
		panic(err)
	}

	return records
}

func GetRecordsByField(record *Records) []*Records {
	records := []*Records{}
	err := adapter.Engine.Find(&records, record)
	if err != nil {
		panic(err)
	}

	return records
}
