package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invopop/jsonschema"
	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/internal/types"
)

type schemaCmd struct {
	cmd    *cobra.Command
	output string
}

func newSchemaCmd() *schemaCmd {
	root := &schemaCmd{}
	cmd := &cobra.Command{
		Use:           "jsonschema",
		Aliases:       []string{"schema"},
		Short:         "outputs cxcli's JSON schema",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			schema := jsonschema.Reflect(&types.Suite{})
			schema.Definitions["Tests"] = jsonschema.Reflect(&types.Tests{})
			schema.Description = "cxcli configuration definition file"
			bts, err := json.MarshalIndent(schema, "	", "	")
			if err != nil {
				return fmt.Errorf("failed to create jsonschema: %w", err)
			}
			if root.output == "-" {
				fmt.Println(string(bts))
				return nil
			}
			if err := os.MkdirAll(filepath.Dir(root.output), 0o755); err != nil {
				return fmt.Errorf("failed to write jsonschema file: %w", err)
			}
			if err := os.WriteFile(root.output, bts, 0o666); err != nil {
				return fmt.Errorf("failed to write jsonschema file: %w", err)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&root.output, "output", "o", "-", "Where to save the JSONSchema file")
	_ = cmd.Flags().SetAnnotation("output", cobra.BashCompFilenameExt, []string{"json"})

	root.cmd = cmd
	return root
}

func init() {
	rootCmd.AddCommand(newSchemaCmd().cmd)
}