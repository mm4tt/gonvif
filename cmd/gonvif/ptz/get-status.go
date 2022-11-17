package ptz

import (
	"fmt"

	tt "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/schema"
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
)

var getStatus = &cobra.Command{
	Use:   "get-status [token]",
	Short: "Get status about a PTZ device",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := tt.ReferenceToken(args[0])
		fmt.Printf("USING TOKEN: %v\n", token)
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetStatus(client, &token)
	},
}

func runGetStatus(client wsdl.PTZ, token *tt.ReferenceToken) error {
	resp, err := client.GetStatus(&wsdl.GetStatus{
		ProfileToken: token,
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
