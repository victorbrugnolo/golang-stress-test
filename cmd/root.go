/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/victorbrugnolo/golang-stress-test/internal/usecase"
)

var rootCmd = &cobra.Command{
	Use:   "golang-stress-test",
	Short: "A tool to make stress tests using golang",

	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		err := usecase.Execute(url, requests, concurrency)

		if err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL where the stress tests will be performed")
	rootCmd.Flags().IntP("requests", "r", 0, "Number of requests")
	rootCmd.Flags().IntP("concurrency", "c", 0, "Number of concurrent requests")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")
}
