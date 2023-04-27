package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version  string = "0.0.1"
	Revision string = "1111aaaa"
	verbose  bool
	GluCmd   = &cobra.Command{
		Use:     "git-local-util",
		Aliases: []string{"glu"},
		Short:   "Tool for migrating an origin from one git repo to another",
		Long:    ``,
		Example: "",
		Version: fmt.Sprintf("%s-%s", Version, Revision),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func Execute() {
	if err := GluCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func init() {
	GluCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}
