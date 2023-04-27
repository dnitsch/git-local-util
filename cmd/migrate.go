package cmd

import (
	"fmt"
	"os"

	"github.com/dnitsch/git-local-util/internal/migrate"
	log "github.com/dnitsch/simplelog"
	"github.com/spf13/cobra"
)

var (
	parentDir, find, replace string
	migrateCmd               = &cobra.Command{
		Use:   "migrate <options>",
		Short: "Tool for migrating an origin from one git repo to another",
		Long:  ``,
		RunE:  migrateOrigin,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(find)+len(replace) < 1 {
				return fmt.Errorf(`the value for find: "%v" and replace: "%v" are invalid: values must be non-empty strings`, find, replace)
			}
			return nil
		},
	}
)

func init() {
	migrateCmd.PersistentFlags().StringVarP(&parentDir, "dir", "d", ".", "Set the parent directory where you want the search to start.")
	migrateCmd.PersistentFlags().StringVarP(&find, "find", "f", "", "The value to find in the origin URL to replace.\nThis could be a portion of the ")
	migrateCmd.PersistentFlags().StringVarP(&replace, "replace", "r", "", "The value to replace in the origin URL.")
	GluCmd.AddCommand(migrateCmd)
}

func migrateOrigin(cmd *cobra.Command, args []string) error {
	logger := log.New(os.Stdout, log.ErrorLvl)
	if verbose {
		logger = log.New(os.Stdout, log.DebugLvl)
	}
	gm := migrate.New(find, replace, logger)
	return gm.ReplaceConfigOrigin(parentDir)
}
