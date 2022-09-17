/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

type SimpleStruct struct {
	A int
	B int
}

func StructReflect(in interface{}) error {
	val := reflect.ValueOf(in)
	if val.Type().Kind() != reflect.Ptr {
		return fmt.Errorf("you must pass in a pointer")
	}
	return nil
}

// func populateStructReflect(in interface{}) error {
// 	val := reflect.ValueOf(in)
// 	if val.Type().Kind() != reflect.Ptr {
// 		return fmt.Errorf("you must pass in a pointer")
// 	}
// 	elmv := val.Elem()
// 	if elmv.Type().Kind() != reflect.Struct {
// 		return fmt.Errorf("you must pass in a pointer to a struct")
// 	}

// 	fval := elmv.FieldByName("B")
// 	fval.SetInt(42)

// 	return nil
// }

// reflectCmd represents the reflect command
var reflectCmd = &cobra.Command{
	Use:   "reflect",
	Short: "reflect example",
	Long: `展示 go 反射的使用方法,原文连接:
	https://philpearl.github.io/post/aintnecessarilyslow/
	https://mp.weixin.qq.com/s/fzmN6zFVioQGedTdSDmyqQ`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reflect called")

	},
}

func init() {
	rootCmd.AddCommand(reflectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reflectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reflectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
