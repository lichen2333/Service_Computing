/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"../src"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register -n [username] -p [password] -e [email] -t [phone]",
	Short: "Register new use",
	Long: `Eg : register -n yourname -p 123456 -e xxxxxxxxxx@qq.com -t 139xxxxxxxx`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Registered.........")
		//TODO
		name2, _ := cmd.Flags().GetString("name")
		password2, _ := cmd.Flags().GetString("password")
		email2, _ := cmd.Flags().GetString("email")
		phone2, _ := cmd.Flags().GetString("phone")
		if src.IsLogin() == false {
			src.RegisterUser(name2, password2, email2, phone2)
		} else {
			fmt.Println("This account is already logged in!")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("name", "n", "", "user name")
	registerCmd.Flags().StringP("password", "p", "", "user password")
	registerCmd.Flags().StringP("email", "e", "", "user email")
	registerCmd.Flags().StringP("phone", "t", "", "user phone")
}
