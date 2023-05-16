package getstarted

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "go-micro-service-example.exe",
	Short: "This is the short description",
	Long: `A longer description
   
	for the first command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the first cobra example")
		fmt.Println("configFile=", configFile)
	},
}

var command1 = &cobra.Command{
	Use:     "command1",
	Short:   "This is command1",
	Aliases: []string{"c1"},
	Args:    cobra.ExactArgs(1), // only 1 parameter for command1
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing command1 ...")
		fmt.Println("configFile=", configFile)
	},
}

var command2 = &cobra.Command{
	Use:     "command2",
	Short:   "This is command2",
	Aliases: []string{"c2"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing command1 ...")
		fmt.Println("configFile=", configFile)
	},
}

func init() {
	rootCmd.AddCommand(command1)
	rootCmd.AddCommand(command2)
	rootCmd.PersistentFlags().StringVarP(&configFile, "conf", "c", "defaultPath", "config file path")
	//rootCmd.Flags().StringVarP(&configFile, "conf", "c", "defaultPath", "config file path")
	// There are two types of flags: 1. Local flags  2.Persistent flags
	// 1. Local flags: assigned to a single command
	// 2. Persistent flags: assigned to a command and all its sub-commands
}

func StartCobra() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

// usage:
// go build

// .\go-micro-service-example.exe -c=c:/mypath
//  This is the first cobra example
//  configFile= c:/mypath

// .\go-micro-service-example.exe command1 aaa -c=c:/mypath
//  executing command1 ...
//  configFile= c:/mypath

// .\go-micro-service-example.exe command2 aaa -c=c:/mypath
//  executing command1 ...
//  configFile= c:/mypath

// See more details at: https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/
