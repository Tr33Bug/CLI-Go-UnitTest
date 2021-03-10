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
	"time"

	"github.com/spf13/cobra"
)

var countRace int

// bsp5Cmd represents the bsp5 command
var bsp5Cmd = &cobra.Command{
	Use:   "bsp5",
	Short: "Showcase the Racecondition",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting race...")
		startRace()
		fmt.Println("Race finished!")
	},
}

func init() {
	rootCmd.AddCommand(bsp5Cmd)
}

func startRace() {
	go addOne()
	go addFive()
	time.Sleep(5 * time.Second)
}

func addFive() {
	countRace += 5
}

func addOne() {
	countRace *= 1
}
