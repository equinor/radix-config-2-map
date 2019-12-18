package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// Command line flags
	configMapName string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "clone-config",
		Short: "Radix clone config",
		Run:   root}

	pf := rootCmd.PersistentFlags()
	pf.StringVar(&configMapName, "configmap-name", "",
		"name to give to configmap")
	cobra.MarkFlagRequired(pf, "configmap-name")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func root(cmd *cobra.Command, args []string) {
	log.Info("Running root command")
	log.Infof("Config map: %s", configMapName)
}
