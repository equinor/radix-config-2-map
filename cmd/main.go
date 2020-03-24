package main

import (
	"github.com/equinor/radix-config-2-map/pkg/configmap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// Command line flags
	namespace     string
	configMapName string
	file          string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "clone-config",
		Short: "Radix clone config",
		Run:   root}

	pf := rootCmd.PersistentFlags()

	pf.StringVar(&namespace, "namespace", "",
		"namespace configmap should be applied to")
	pf.StringVar(&configMapName, "configmap-name", "",
		"name to give to configmap")
	pf.StringVar(&file, "file", "",
		"absolute path to file")

	cobra.MarkFlagRequired(pf, "namespace")
	cobra.MarkFlagRequired(pf, "configmap-name")
	cobra.MarkFlagRequired(pf, "file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func root(cmd *cobra.Command, args []string) {
	log.Infof("Copying radixconfig.yaml file (%s) into namespace (%s) and configmap (%s)", file, namespace, configMapName)

	err := configmap.CreateFromFile(namespace, configMapName, file)
	if err != nil {
		log.Fatal("Error copying radixconfig.yaml and creating config map from file: %v", err)
	} else {
		log.Infof("Successfully found and copied radixconfig.yaml file")
	}
}
