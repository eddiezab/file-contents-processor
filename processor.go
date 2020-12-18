// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package file_contents

import (
	"io/ioutil"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/processors"
)

func init() {
	processors.RegisterPlugin("file_contents", New)
}

// New constructs a new script processor.
func New(c *common.Config) (processors.Processor, error) {
	var config = struct {
		File string `config:"file" validate:"required"`
	}{}
	if err := c.Unpack(&config); err != nil {
		return nil, err
	}

	p := &addID{
		config,
		gen,
	}

	return p, nil
}

func (p *addID) Run(event *beat.Event) (*beat.Event, error) {
	content, err := ioutil.ReadFile(config.File)
	check(err)

	event.PutValue("exfil.data", content)

	return event, nil
}
