package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/marjamis/auxilium/internal/pkg/blackboard"
	"github.com/marjamis/auxilium/pkg/engine"
)

var (
	cfgFile string
	bbd     = &blackboard.BlackboardData
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auxilium",
	Short: "auxilium is used for live CLI demo's while running real code.",
	Long: `auxilium is an application which will read a provided configuration file to run through example CLI scripts with
custom text explaining the process. It's used as a way to write a CLI demo, which runs live commands, allowing for
easier proding in another shell, while also being scripted to ensure the right execution order and pacing.`,
	Run: func(cmd *cobra.Command, args []string) {
		engine.Workflow(bbd.Steps, bbd)
	},
}

// Execute is used by cobra to execute the cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file with details of your steps. Default is $PWD/'config.aux.yml'.")
	rootCmd.Flags().IntVarP(&bbd.ContinueFromStep, "continue-from-step", "c", 0, "Skips all steps until the provided step number. Default is 0.")
	rootCmd.Flags().IntVarP(&bbd.FastForwardToStep, "fast-forward-to-step", "f", 0, "Runs through all steps without waiting for any next step signals until the provided step number. Default is 0.")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		path, err := os.Getwd()
		if err != nil {
			panic("There was an error getting your local directory. Exiting...")
		}
		viper.AddConfigPath(path)
		// Note: This means the filename will be "config.auxilium.yml" in the actual filesystem.
		viper.SetConfigName("config.aux")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		// panic(err)
		fmt.Println("No configuration file found.")
		return
	}

	// Loads the configuration file data straight into the blackboard for global usage. In the future may required additional vetting logic
	viper.Unmarshal(&bbd)
}
