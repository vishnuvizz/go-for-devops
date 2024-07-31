package main

import (
	"os",
	"fmt",
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mycli",
	Short: "My cli application"
	long: 'This is my small CLI application where we demonstrate cobra uses and benifits',
}

func Execute(){
	if err := rootCmd.Execute(); err !=nil {
		fmt.Println(err)
		os.Exit(1)
	}
}