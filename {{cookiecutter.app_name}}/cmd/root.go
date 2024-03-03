package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"
)

var rootCmd = &cobra.Command{
	Use:   "{{cookiecutter.app_name}}",
	Short: "Short description of your application",
	Long: `Long description of your application.`
	PersistentPreRun: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) { },

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// rootCmd.PersistentFlags().StringVarP(&example, "example", "e", "", "Example flag")
	// rootCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 10, "Number of threads to use for the requests")
	// rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	// cobra.MarkFlagRequired(rootCmd.PersistentFlags(), "example")


}
