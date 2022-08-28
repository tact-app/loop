package command

import (
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"go.octolab.org/tact/loop/internal/pkg/unsafe"
	"go.octolab.org/tact/loop/internal/service/loom"
)

// Tree returns command to show whole org structure with looms.
func Tree() *cobra.Command {
	command := cobra.Command{
		Use:   "tree",
		Short: "show whole org structure with looms",
		Long:  "Show whole organization structure with looms.",

		RunE: func(cmd *cobra.Command, args []string) error {
			token := os.Getenv("LOOM_TOKEN")
			workspaceID := unsafe.ReturnInt(cmd.Flags().GetInt("workspace"))

			client, err := loom.NewClient(new(http.Client), "https://www.loom.com/graphql", token)
			if err != nil {
				return err
			}

			service := loom.New(client)
			tree, err := service.Tree(cmd.Context(), workspaceID)
			if err != nil {
				return err
			}

			return yaml.NewEncoder(cmd.OutOrStdout()).Encode(tree)
		},

		Args: cobra.NoArgs,
	}
	command.Flags().Int("workspace", 0, "workspace id")

	return &command
}
