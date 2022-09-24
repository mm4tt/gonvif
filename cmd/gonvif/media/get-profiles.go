package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
)

var getProfiles = &cobra.Command{
	Use:   "get-profiles",
	Short: "List Onvif device media profiles",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password)
		if err != nil {
			return nil
		}
		types, err := cmd.Flags().GetStringArray("types")
		if err != nil {
			return nil
		}
		return runGetProfiles(client, types)
	},
}

func init() {
	getProfiles.Flags().StringArrayP("types", "t", []string{"All"}, "Types of profile configurations to include")
}

func runGetProfiles(client wsdl.Media2, types []string) error {
	resp, err := client.GetProfiles(&wsdl.GetProfiles{
		Type: types,
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
