/*
   Copyright 2016 Jonathan Windle

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// RawFile.Metadata
package main

import (
	"os/exec"
	"regexp"
	"strings"
)

var reFindMeta = regexp.MustCompile(`(\S[\S ]*): (\S[\S ]*)`)

func (rf *RawFile) Metadata() (metadata string, err error) {
	var out []byte
	rawfile := *rf

	out, err = exec.Command("dcraw", "-i", "-v", rawfile.Path).Output()
	metadata = string(out)
	if err != nil {
		return
	}

	metaStrings := reFindMeta.FindAllStringSubmatch(metadata, -1)
	rawfile.Meta = make(map[string]string, len(metaStrings))
	for index := range metaStrings {
		key := strings.ToLower(metaStrings[index][1])
		rawfile.Meta[key] = metaStrings[index][2]
	}

	*rf = rawfile
	return
}
