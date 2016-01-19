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
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/howeyc/gopass"
)

const (
	// RootURL https://h2w.cc/api/v1
	RootURL = "https://h2w.cc/api/v1"
	// GetTokenURL https://h2w.cc/api/v1/auth/token
	GetTokenURL = RootURL + "/auth/token"
	// VerifyTokenURL https://h2w.cc/api/v1/verify
	VerifyTokenURL = GetTokenURL + "/verify"
	//DashboardURL https://h2w.cc/api/v1/dashboard
	DashboardURL = RootURL + "/dashboard"
)

var (
	body  []byte
	jbody interface{}
)

// Login to Health2Wealth api
func Login(email string, password string, tokenFile string) {
	email, password = getUserCredentials(email, password)
	resp := hitEndpoint(GetTokenURL, email, password)

	readBody(resp)
	printDebug(resp)

	if resp.StatusCode == 200 {
		token := jbody.(map[string]interface{})["token"].(string)
		storeToken(tokenFile, token)
	}
}

// VerifyToken that token is valid
func VerifyToken(token string, tokenFile string) {
	email, password := getTokenCredentials(token, tokenFile)
	resp := hitEndpoint(VerifyTokenURL, email, password)

	readBody(resp)
	printDebug(resp)
}

// DeleteToken data
func DeleteToken(tokenFile string) {
	deleteToken(tokenFile)
}

// Dashboard data for the current day, week, and quarter
func Dashboard(tokenFile string) {
	email, password := getTokenCredentials("", tokenFile)
	resp := hitEndpoint(DashboardURL, email, password)

	readBody(resp)
	printDebug(resp)
}

// ShortStats one-line status update
func ShortStats(tokenFile string) {
	Dashboard(tokenFile)

	dashboard := jbody.(map[string]interface{})["dashboard"].(map[string]interface{})
	days := dashboard["days"].([]interface{})
	todayName := "today"
	for i := 0; i < len(days); i++ {
		day := days[i].(map[string]interface{})
		if day["is_today"].(bool) {
			todayName = day["abbr"].(string)
			break
		}
	}
	todaySteps := prettyNumber(dashboard["today_steps"].(float64))
	todayPct := prettyPct(100 * dashboard["today_steps"].(float64) / (dashboard["weekly_step_goal"].(float64) / 7.0))
	weekNumber := dashboard["week_number"].(float64)
	weekSteps := prettyNumber(dashboard["current_steps"].(float64))
	weekPct := prettyPct(dashboard["week_full_pct"].(float64))

	fmt.Printf("H2W - %v: %v (%v%%); week %v: %v (%v%%)", todayName, todaySteps, todayPct, weekNumber, weekSteps, weekPct)
}

func prettyNumber(f float64) string {
	s := fmt.Sprint(f)
	rev := reverse(s)
	chars := strings.Split(rev, "")
	rev = ""
	for i := 0; i < len(chars); i++ {
		if i%3 == 0 {
			rev += ","
		}
		rev += chars[i]
	}
	rev = strings.TrimPrefix(rev, ",")
	return reverse(rev)
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func roundPlaces(f float64, places int) float64 {
	shift := math.Pow10(places)
	return round(f*shift) / shift
}

func round(f float64) float64 {
	if f < 0 {
		return math.Ceil(f - 0.5)
	}
	return math.Floor(f + 0.5)
}

func prettyPct(f float64) string {
	num := prettyNumber(math.Trunc(f))
	rem := (f - math.Trunc(f)) * 100
	dec := round(rem)
	return fmt.Sprintf("%s.%02d", num, int32(dec))
}

func readBody(resp *http.Response) {
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		body = b
		json.Unmarshal(body, &jbody)
	}
}

func printDebug(resp *http.Response) {
	fmt.Printf("ENDPOINT: %s\n", resp.Request.URL)
	fmt.Printf("STATUS: %s\n", resp.Status)
	if jbody != nil {
		pp, _ := json.MarshalIndent(jbody, "", "  ")
		fmt.Printf("BODY: >>>>>>>>\n%s\n<<<<<<<<\n", pp)
	} else {
		fmt.Printf("BODY: >>>>>>>>\n%s\n<<<<<<<<\n", body)
	}
}

func hitEndpoint(urlString string, email string, password string) *http.Response {
	req, err := http.NewRequest("GET", urlString, nil)
	handle(err)

	req.Header.Add("X-H2W-Client-ID", "0")
	req.SetBasicAuth(email, password)

	resp, err := http.DefaultClient.Do(req)
	handle(err)

	return resp
}

func getTokenCredentials(token string, tokenFile string) (string, string) {
	token = strings.TrimSpace(token)
	if token == "" {
		token = retrieveToken(tokenFile)
	}
	return token, ""
}

func getUserCredentials(email string, password string) (string, string) {
	email = strings.TrimSpace(email)
	if email == "" {
		fmt.Print("Email: ")
		email, err := bufio.NewReader(os.Stdin).ReadString('\n')
		handle(err)
		email = strings.TrimSpace(email)
	}

	password = strings.TrimSpace(password)
	if password == "" {
		fmt.Print("Password: ")
		password = strings.TrimSpace(string(gopass.GetPasswdMasked()))
	}

	return email, password
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
