package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() {
	c := &cobra.Command{
		Use: "github-activity",
		Short: "Type Github username to fetch recent activity",
		Run: func(cmd *cobra.Command, args []string) {
			// 1. Provide the github username as an argument when running the CLI
			// e.g. github-activity <username>
			if len(args) < 1 {
				cmd.Help()
				return
			}
			// 2. Fetch the recent activity of the specified GitHub user using the GitHub API
			username := args[0]
			// 3. Display the recent activity of the specified GitHub user in the CLI
			activity, err := FetchActivity(username)
			if err != nil {
				cmd.Println("Error fetching activity: ", err)
				return
			}
			cmd.Println(activity)
		},
	}
	cobra.CheckErr(c.Execute())
}

func FetchActivity(username string) (string, error) {
	// Fetch the recent activity of the specified GitHub user using the GitHub API
	return "HARDCODED VALUE", nil
}