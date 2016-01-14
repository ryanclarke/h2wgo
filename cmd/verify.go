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
	"github.com/spf13/cobra"
	"github.com/ryanclarke/h2wgo/api"
)

var token string

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a user token is valid",
	Long: `Verify that a user token for Health2Wealth is valid.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.VerifyToken(token, tokenFile)
	},
}

func init() {
	authCmd.AddCommand(verifyCmd)

    verifyCmd.Flags().StringVarP(&token, "token", "t", "", "User token for Health2Wealth")
}
