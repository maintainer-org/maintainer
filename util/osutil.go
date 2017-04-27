// Copyright © 2017 Ce Gao <ce.gao@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import "os"

// OpenFile opens file from 'name', and create one if not exist.
func OpenFile(fileName string, flag int, perm os.FileMode) (*os.File, error) {
	var file *os.File
	var err error

	file, err = os.OpenFile(fileName, flag, perm)
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
	}

	return file, err
}
