/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// bsp6Cmd represents the bsp6 command
var bsp6Cmd = &cobra.Command{
	Use:   "bsp6",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(randomReturnHello())

	},
}

func init() {
	rootCmd.AddCommand(bsp6Cmd)
}

func randomReturnHello() string {
	rand.Seed(time.Now().UnixNano())
	answers := []string{
		"...",
		"Hello",
		"Hello",
		"...",
		"...",
	}

	// return answers[rand.Intn(len(answers))]
	return answers[rand.Intn(len(answers))]

}
