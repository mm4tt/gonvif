package ptz

import (
	"fmt"
	"time"

	tt "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/schema"
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
)

var relativeMove = &cobra.Command{
	Use:   "relative-move [token]",
	Short: "Move PTZ device",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := tt.ReferenceToken(args[0])
		fmt.Printf("USING TOKEN: %v\n", token)
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runRelativeMove(client, &token)
	},
}

func runRelativeMove(client wsdl.PTZ, token *tt.ReferenceToken) error {
	resp, err := client.ContinuousMove(&wsdl.ContinuousMove{
		ProfileToken: token,

		/*Translation: &tt.PTZVector{
			PanTilt: &tt.Vector2D{
				X:     100,
				Y:     100,
				Space: "https://www.onvif.org/ver10/tptz/PanTiltSpaces/PositionGenericSpace",
			},
		},*/
		Velocity: &tt.PTZSpeed{
			PanTilt: &tt.Vector2D{
				X: -1,
				//Y: -1,
			},
		},
	})
	if err != nil {
		return err
	}
	<-time.After(1000 * time.Millisecond)
	if _, err := client.Stop(&wsdl.Stop{ProfileToken: token}); err != nil {
		return err
	}

	return root.OutputJSON(resp)
}
