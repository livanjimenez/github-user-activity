package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() {
	c := &cobra.Command{
		Use: "github-activity",
		Short: "Type Github username to fetch recent activity",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Help()
				return
			}
			username := args[0]
			activity, err := FetchActivity(username)
			if err != nil {
				cmd.Println("Error fetching activity: ", err)
				return
			}
			cmd.Println(activity)
		},
	}
	c.Execute()
}

func FetchActivity(username string) (string, error) {
	newMessage := "Fetching activity for " + username
	return newMessage, nil
}