// 1. Provide the github username as an argument when running the CLI
// e.g. github-activity <username>

// 2. Fetch the recent activity of the specified GitHub user using the GitHub API

// 3. Display the recent activity of the specified GitHub user in the CLI

package main

import "github.com/livanjimenez/github-user-activity/cmd"

func main() {
	cmd.RootCmd()
}