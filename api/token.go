// Copyright © 2016 Ryan Clarke <ryan@ryanclarke.net>
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

package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func retrieveToken(tokenFile string) (token string) {
	data, err := ioutil.ReadFile(tokenFile)
	handle(err)
	if err == nil {
		token = strings.TrimSpace(string(data) + ":")
	}
	return
}

func storeToken(tokenFile string, token string) {
	fmt.Printf("TOKEN: %s\n", token)
	err := ioutil.WriteFile(tokenFile, []byte(token), 0644)
	handle(err)
}

func deleteToken(tokenFile string) {
	os.Remove(tokenFile)
}
