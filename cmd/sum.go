/*
Copyright © 2021 Tr33Bug <Tr33@gmx.com>

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
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "Sum up two numbers",
	Long: `Use the sum command to add two numbers.
	
	
	sum 2 3`,
	Run: func(cmd *cobra.Command, args []string) {
		sum()
		fmt.Println(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)

}

func sum() {
	fmt.Println(1 + 2)
}
