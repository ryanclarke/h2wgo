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

package cmd

import (
	"github.com/ryanclarke/h2wgo/api"
	"github.com/spf13/cobra"
)

var (
	email    string
	password string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Health2Wealth",
	Long: `Login to Health2Wealth with an email and password.
   
Retrives a user token and stores it locally for use with other requests. By
default the token is stored in '$PWD/.h2wgotoken' file. This can be changed
with the --token-file flag. Command will prompt for email/password if either
or both are not provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Login(email, password, tokenFile)
	},
}

func init() {
	authCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&email, "email", "e", "", "User login email for Health2Wealth")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "User login password for Health2Wealth")
}
