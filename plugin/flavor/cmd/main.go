package main

import (
	"os"

	"github.com/FrenchBen/infrakit.k8s/plugin/flavor"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/infrakit/cli"
	flavor_plugin "github.com/docker/infrakit/rpc/flavor"
	"github.com/spf13/cobra"
)

func main() {

	logLevel := cli.DefaultLogLevel
	var name string
	var sslDir string

	cmd := &cobra.Command{
		Use:   os.Args[0],
		Short: "Kubernetes flavor plugin",
		Run: func(c *cobra.Command, args []string) {
			cli.SetLogLevel(logLevel)
			flavorPlugin := kubernetes.NewPlugin(sslDir)

			cli.SetLogLevel(logLevel)
			cli.RunPlugin(name, flavor_plugin.PluginServer(flavorPlugin))
		},
	}

	cmd.Flags().IntVar(&logLevel, "log", logLevel, "Logging level. 0 is least verbose. Max is 5")
	cmd.Flags().StringVar(&name, "name", "flavor-kubernetes", "Plugin name to advertise for discovery")
	defaultDir, erdir := os.Getwd()
	if erdir != nil {
		log.Error(erdir)
		os.Exit(1)
	}
	cmd.Flags().StringVar(&sslDir, "ssl-dir", defaultDir, "Kubernetes SSL directory")

	cmd.AddCommand(cli.VersionCommand())

	err := cmd.Execute()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
