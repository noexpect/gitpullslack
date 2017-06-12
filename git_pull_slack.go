package main

import (
	"fmt"
	"flag"
	"github.com/codeskyblue/go-sh"
	"github.com/nlopes/slack"
	"log"
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
)
	//"github.com/nlopes/slack"

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
	branch, _ := session.Command("git", "branch").Output()
	fmt.Printf("branch:%s",branch)
	slackMessage := ""
	if session.Test("executable", "git fetch origin master"){
		session.Command("git", "fetch",  "origin", "master").Run()
		diff, _ := session.Command("git", "diff").Output()
		slackMessage = string(diff)
	}else  {
		slackMessage = "no updates"

	}
	session.ShowCMD = true // for debug
	fmt.Println(slackMessage)

	// load yaml
	buf, err := ioutil.ReadFile("./gitpullslack/conf.yml")
	if err != nil {
		fmt.Printf("yml_read%s\n", err)
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		fmt.Printf("yml_marsh:%s\n", err)
	}

	fmt.Printf("%s\n", m["slack_token"])
	api := slack.New(m["slack_token"].(string))
	os.Exit(run(api, slackMessage))
}

func run(api *slack.Client, mes string) int {
	rtm := api.NewRTM()
	rtm.SendMessage(rtm.NewOutgoingMessage(mes, "C52U4FUJF")) // input chnel as channel id
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				log.Printf("Channel: %v\n", ev.Channel)
				rtm.SendMessage(rtm.NewOutgoingMessage("hi", ev.Channel))

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}

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

[done]- run system command
-- git branch develop
-- git fetch
-- git diff (develop) origin/(local)
// do them later
-- git merge dry run
-- git pull origin local
-- command after git pull

[done]- load external yml
-- install lib
-- read yml
-- set gitignore
-- save slack api token

[doing]- slack notify
-- install lib
-- send merge diff
-- show reaction button
-- merge by reaction callback

[doing]- debug
-- install debugger(delve)
-- how to use it

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