// Copyright 2024 Google LLC
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

package testutils

import (
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

// fileExists check if a give file exists
func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// filter select only values who match the condition for the field
func Filter(field string, value string, iamList []gjson.Result) []gjson.Result {
	var filtered []gjson.Result

	for _, iam := range iamList {
		if strings.Contains(iam.Get(field).String(), value) {
			filtered = append(filtered, iam)
		}
	}
	return filtered
}

// verify if gjson array of string contains another string
func Contains(slice []gjson.Result, item string) bool {
	for _, v := range slice {
		if v.String() == item {
			return true
		}
	}
	return false
}
