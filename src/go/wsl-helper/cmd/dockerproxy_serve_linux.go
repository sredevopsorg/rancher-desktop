/*
Copyright © 2021 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rancher-sandbox/rancher-desktop/src/wsl-helper/pkg/dockerproxy"
)

var dockerproxyServeViper = viper.New()

// dockerproxyServeCmd is the `wsl-helper docker-proxy serve` command.
var dockerproxyServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the docker socket proxy server",
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := dockerproxyServeViper.GetString("endpoint")
		proxyEndpoint := dockerproxyServeViper.GetString("proxy-endpoint")
		err := dockerproxy.Serve(endpoint, proxyEndpoint)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	dockerproxyServeCmd.Flags().String("endpoint", dockerproxy.DefaultEndpoint, "Endpoint to listen on")
	dockerproxyServeCmd.Flags().String("proxy-endpoint", dockerproxy.DefaultProxyEndpoint, "Endpoint dockerd is listening on")
	dockerproxyServeViper.AutomaticEnv()
	dockerproxyServeViper.BindPFlags(dockerproxyServeCmd.Flags())
	dockerproxyCmd.AddCommand(dockerproxyServeCmd)
}