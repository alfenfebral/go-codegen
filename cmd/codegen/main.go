package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alfenfebral/go-codegen/generators/basic"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gogen",
	Short: "Go code generator",
	Long:  "A tool for generating Go code templates",
}

var basicCmd = &cobra.Command{
	Use:   "basic [name]",
	Short: "Generate a basic Go file",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		dir, _ := cmd.Flags().GetString("dir")

		if name == "" {
			log.Fatal("--name flag is required")
		}

		if err := basic.Generate(name, dir); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	basicCmd.Flags().StringP("name", "n", "", "Name of the generated file (required)")
	basicCmd.Flags().StringP("dir", "d", ".", "Output directory")
	basicCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(basicCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
