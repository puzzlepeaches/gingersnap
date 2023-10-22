package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// cfgFile is a string variable to hold the path of the configuration file
var cfgFile string

// rootCmd represents the base command when called without any subcommands
// It is the entry point for the cobra application
var rootCmd = &cobra.Command{
	Use:   "{{cookiecutter.app_name}}",
	Short: "Short description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:`

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//      Run: func(cmd *cobra.Command, args []string) { },

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// If there is an error during the execution, it will print the error and exit with status 1
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init function to initialize the cobra application
// It will be called only once before running the command
func init() {
	cobra.OnInitialize()

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
