package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/client"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
)

var cmd = &cobra.Command{
	Use:   "media",
	Short: "Manipulate Onvif device media features.",
}

func init() {
	root.RequireAuthFlags(cmd)
	root.Command.AddCommand(cmd)
	cmd.AddCommand(
		getAnalyticsConfigurations,
		getProfiles,
		getVideoSourceConfigurations,
	)
}

func ServiceClient(url, username, password string, verbose bool) (wsdl.Media2, error) {
	onvif, err := client.New(url, username, password, verbose)
	if err != nil {
		return nil, err
	}
	return onvif.Media2()
}
