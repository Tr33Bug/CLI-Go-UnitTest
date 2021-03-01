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

	"github.com/spf13/cobra"
)

// bsp2Cmd represents the bsp2 command
var bsp2Cmd = &cobra.Command{
	Use:   "bsp2",
	Short: "Tipes \"Hello World\" on the terminal, but one word in one line",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(printHello())
		fmt.Println(printWorld())
	},
}

func init() {
	rootCmd.AddCommand(bsp2Cmd)
}

func printHello() string {
	return ("Hello")
}

func printWorld() string {
	return ("World")
}
