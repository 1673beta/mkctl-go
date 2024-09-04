package cmd

import (
	"fmt"
	"mkctl/cmd/util"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var deleteDays int
var remoteOnly bool

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove from database.",
}

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Remove old notes from database.",
	Run: func(cmd *cobra.Command, args []string) {
		cutoff := time.Now().AddDate(0, 0, -deleteDays)
		ms := cutoff.Sub(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)).Milliseconds()

		encoded := strconv.FormatInt(ms, 36)

		db, err := util.ConnectToDb()
		if err != nil {
			fmt.Printf("Error while connecting to database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		if remoteOnly {
			query := fmt.Sprintf(`DELETE FROM note WHERE SUBSTRING(note."id", 1, 8) <= %s AND note."userHost" IS NOT NULL`, encoded)
			_, err := db.Exec(query)
			if err != nil {
				fmt.Printf("Error while deleting notes: %v\n", err)
				os.Exit(1)
			}
		} else {
			query := fmt.Sprintf(`DELETE FROM note WHERE SUBSTRING(note."id", 1, 8) <= %s`, encoded)
			_, err := db.Exec(query)
			if err != nil {
				fmt.Printf("Error while deleting notes: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	removeCmd.Flags().IntVarP(&deleteDays, "days", "d", 120, "How old notes have to be before they are deleted. Defaults to 120.")
	removeCmd.Flags().BoolVarP(&remoteOnly, "remote", "r", false, "Only delete notes from remote server.")
	removeCmd.AddCommand(notesCmd)
	rootCmd.AddCommand(removeCmd)
}
