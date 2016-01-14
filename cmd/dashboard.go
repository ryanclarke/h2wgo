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

var shortStats bool

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "All the user's H2W data",
	Long: `A json result with all the user's Health2Weatlth data. Includes current steps,
points, and goals for the day, week and quarter. Also has stats on other users
in the user's org.`,
	Run: func(cmd *cobra.Command, args []string) {
		if shortStats {
			api.ShortStats(tokenFile)
		} else {
			api.Dashboard(tokenFile)
		}
	},
}

func init() {
	RootCmd.AddCommand(dashboardCmd)

	dashboardCmd.Flags().BoolVarP(&shortStats, "short-stats", "s", true, "One-line summary of day and week step counts")
}
