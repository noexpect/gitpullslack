package main

import "fmt"
import (
	"flag"
	"github.com/codeskyblue/go-sh"
)

func main() {
	// get flags
	var (
		gitPath string
		gitOrigin string
		gitBranchMergeFrom string
		gitBranchMergeTo string
		slackToken string
		slackChannel string
		commandAfterGitPull string
	)
	flag.StringVar(&gitPath, "git.path", "blank", "git path")
	flag.StringVar(&gitOrigin, "git.origin", "blank", "git origin url")
	flag.StringVar(&gitBranchMergeFrom, "git.branch.merge.from", "blank", "git branch merge from")
	flag.StringVar(&gitBranchMergeTo, "git.branch.merge.to", "blank", "git branch merge to")
	flag.StringVar(&slackToken, "slack.token", "blank", "slack api token")
	flag.StringVar(&slackChannel, "slack.channel", "blank", "slack channel")
	flag.StringVar(&commandAfterGitPull, "command.after.git.pull", "blank", "command after git pull")

	flag.Parse()

	fmt.Printf("--git.path:%s, --git.origin:%s, --git.branch.merge.from:%s, --git.branch.merge.to:%s, --slack.token:%s, --slack.channel:%s, --command.after.git.pull:%s\n", gitPath, gitOrigin, gitBranchMergeFrom, gitBranchMergeTo, slackToken, slackChannel, commandAfterGitPull)

	session := sh.NewSession()
	session.SetDir("./gitpullslack")
	session.Command("git", "fetch",  "origin", "master").Run()
	session.ShowCMD = true
}
/*
TODO
[done]- get command line flags
-- sub option
-- git path
-- git local branch(ex. develop)
-- git remote branch(ex. local)
-- slack token
-- slack channel
-- after pull command(ex. supervisor restart)

- run system command
-- git branch develop
-- git fetch
-- git diff (develop) origin/(local)
-- git merge dry run
-- git pull origin local
-- command after git pull

- slack notify
-- send merge diff
-- show reaction button
-- merge by reaction callback

- err handle
-- logging
-- exception

- refactor
-- naming
-- package/file

- write tests

- include cross compiled binary

- example of run with supervisor
-- logging
-- restart process

 */

