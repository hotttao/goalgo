/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// errorHandleCmd represents the errorHandle command
var errorHandleCmd = &cobra.Command{
	Use:   "errorHandle",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("errorHandle called")
	},
}

func Readfile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()
	return io.ReadAll(f)
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := Readfile(filepath.Join(home, "setting.yaml"))
	return config, errors.WithMessage(err, "can not read config")
}

func ReadMe() {
	_, err := ReadConfig()
	
	if err != nil {
		fmt.Printf("origin error %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n %+v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(errorHandleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// errorHandleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// errorHandleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
