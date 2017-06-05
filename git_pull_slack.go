package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

/*
TODO
- get command line flags
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
-- git merge dry run (also get merge diff)
-- git pull origin local

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