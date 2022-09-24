package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getImagingSettings = &cobra.Command{
	Use:   "get-imaging-settings",
	Short: "Show Onvif device imaging settings",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password)
		if err != nil {
			return nil
		}
		videoSourceToken, err := cmd.Flags().GetString("video_source_token")
		if err != nil {
			return nil
		}
		return runGetImagingSettings(client, videoSourceToken)
	},
}

func init() {
	getImagingSettings.Flags().StringP("video_source_token", "t", "", "Video source token")
	getImagingSettings.MarkFlagRequired("video_source_token")
}

func runGetImagingSettings(client wsdl.ImagingPort, token string) error {
	resp, err := client.GetImagingSettings(&wsdl.GetImagingSettings{
		VideoSourceToken: util.NewReferenceTokenPtr(token),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
