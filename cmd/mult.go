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

// multCmd represents the mult command
var multCmd = &cobra.Command{
	Use:   "mult",
	Short: "Multiply a number by its own",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("please only one number such as 5 or 14")
		}

		// fmt.Println(mult(i))

		show, err := cmd.Flags().GetBool("show_steps")
		if err != nil {
			return err
		}

		fmt.Println(mult2(i, show))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(multCmd)
	multCmd.Flags().BoolP("show_steps", "s", false, "show every step of the multiplication")

}

func mult(i int) int {
	result := i * i
	return result
}

func mult2(i int, show bool) int {
	result := 0
	for x := i; x > 0; x-- {
		if show {
			fmt.Printf("Result %d: --> %d \n", x, result)
		}
		result += i
	}
	return result
}
