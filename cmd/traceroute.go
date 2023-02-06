package cmd

import (
	"fmt"

	"github.com/jsdelivr/globalping-cli/client"
	"github.com/jsdelivr/globalping-cli/model"
	"github.com/spf13/cobra"
)

// tracerouteCmd represents the traceroute command
var tracerouteCmd = &cobra.Command{
	Use:   "traceroute [target]",
	Short: "Implementation of the native traceroute command",
	Long: `traceroute tracks the route packets taken from an IP network on their way to a given host. It utilizes the IP protocol's time to live (TTL) field and attempts to elicit an ICMP TIME_EXCEEDED response from each gateway along the path to the host.
	
		Examples:
		# Traceroute google.com from a probe in the network
		globalping traceroute google.com --from "New York" --limit 2`,
	Args: requireTarget(),
	Run: func(cmd *cobra.Command, args []string) {
		// Make post struct
		opts = model.PostMeasurement{
			Type:      "traceroute",
			Target:    args[0],
			Locations: createLocations(from),
			Limit:     limit,
			Options: &model.MeasurementOptions{
				Protocol: protocol,
				Port:     port,
			},
		}

		res, err := client.PostAPI(opts)
		if err != nil {
			fmt.Println(err)
			return
		}

		client.OutputResults(res.ID, ctx)
	},
}

func init() {
	rootCmd.AddCommand(tracerouteCmd)

	// traceroute specific flags
	tracerouteCmd.Flags().StringVar(&protocol, "protocol", "", "Specifies the protocol used for tracerouting (ICMP, TCP or UDP). (default \"icmp\")")
	tracerouteCmd.Flags().IntVar(&port, "port", 0, "Specifies the port to use for the traceroute. Only applicable for TCP protocol. (default 80)")
}
