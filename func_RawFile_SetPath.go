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

// RawFile.SetPath
package main

import (
	"errors"
	"os"
	"path"
)

func (rf *RawFile) SetPath(filePath string) (err error) {
	var isfile bool
	var fileinfo os.FileInfo
	rawfile := *rf

	if isfile, fileinfo, err = IsFile(filePath); err != nil {
		return
	}

	if isfile == false {
		err = errors.New(filePath + " is not a file.")
		return
	}

	rawfile.Path = filePath
	_, rawfile.Name = path.Split(filePath)
	rawfile.Extension = path.Ext(filePath)
	rawfile.FileInfo = fileinfo

	*rf = rawfile
	return
}
