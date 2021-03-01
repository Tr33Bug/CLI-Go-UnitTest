/*
Copyright © 2021 Tr33bug <Tr33@gmx.com>

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
	"time"

	"github.com/spf13/cobra"
)

// bsp3Cmd represents the bsp3 command
var bsp3Cmd = &cobra.Command{
	Use:   "bsp3",
	Short: "Waits for 5 Seconds and prints Hello World",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(waitAndPrint())
	},
}

func init() {
	rootCmd.AddCommand(bsp3Cmd)
}

func waitAndPrint() string {
	time.Sleep(5 * time.Second)
	return "Hello World"
}

func failTest(x, y int) int {
	return x + y
}
