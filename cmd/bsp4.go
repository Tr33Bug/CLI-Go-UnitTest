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
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// bsp4Cmd represents the bsp4 command
var bsp4Cmd = &cobra.Command{
	Use:   "bsp4",
	Short: "Summs two numbers",

	RunE: func(cmd *cobra.Command, args []string) error {
		// parse user argument input
		a, err := parseAndIntConv(args, 0)
		if err != nil {
			return err
		}
		b, err := parseAndIntConv(args, 1)
		if err != nil {
			return err
		}

		fmt.Println(sumUp(a, b))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(bsp4Cmd)
}

func parseAndIntConv(args []string, place int) (int, error) {
	errorTxt := "Please only use numbers such as 5 or 200"
	a, err := strconv.Atoi(args[place])
	if err != nil {
		err = errors.New(errorTxt)
		return a, err
	}
	return a, err
}

func sumUp(x, y int) int {
	return x + y
}

func retTrue() bool {
	return true
}
