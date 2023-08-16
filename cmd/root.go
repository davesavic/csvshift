package cmd

import (
	"csvshift/application"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "csvshift",
	Short: "An intuitive DSL that simplifies and elevates CSV data transformation from ingestion to output.",
	Long: `A powerful DSL tailored for seamless CSV file management. 
With intuitive syntax, it handles everything from data ingestion to transformation, refinement, and final output. 
Elevate your CSV data processing with unparalleled ease and precision.`,

	Run: func(cmd *cobra.Command, args []string) {
		source, err := cmd.Flags().GetString("source")
		if err != nil {
			panic(err)
		}

		destination, err := cmd.Flags().GetString("destination")
		if err != nil {
			panic(err)
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			panic(err)
		}

		if strings.TrimSpace(path) == "" {
			p, err := findCsvShiftFile()
			if err != nil {
				panic(err)
			}

			path = *p
		}

		application.Run(source, destination, path)
	},
}

func findCsvShiftFile() (*string, error) {
	files, err := os.ReadDir("./")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".csvshift") {
			path := file.Name()
			return &path, nil
		}
	}

	return nil, fmt.Errorf("no .csvshift file found in current directory")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("source", "s", "", "path to source .csv file")
	rootCmd.PersistentFlags().StringP("destination", "d", "", "path to destination .csv file")
	rootCmd.PersistentFlags().StringP("path", "p", "", "path to .csvshift dsl file (defaults to .csvshift in current directory)")
}
