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
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/tidwall/gjson"
)

// getResultFieldStrSlice parses a field of a results list into a string slice
func GetResultFieldStrSlice(rs []gjson.Result, field string) []string {
	s := make([]string, 0)
	for _, r := range rs {
		s = append(s, r.Get(field).String())
	}
	return s
}

// GetBptOutputStrSlice parses a TFBlueprintTest result field into a string slice
func GetBptOutputStrSlice(bpt *tft.TFBlueprintTest, field string) []string {
	s := make([]string, 0)
	for _, r := range bpt.GetJsonOutput(field).Array() {
		s = append(s, r.String())
	}
	return s
}
