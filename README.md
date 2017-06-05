# gitpullslack
Simple deploy tool by git pull and slack notifications.
Now developing.

## Description
sync (local development environment) -> (remote development environment)

In small development, the deployment means only sync of local and remote server.
So it is enough fo sync by `git pull origin (local baranch)` on (develop branch) in remote server after local branch's push.

## Usage

```

$ gitpullslack --git.path ""/bin/git" --git.origin "https://github.com/foo/bar.git" --git.branch.merge.from "local" --git.branch.merge.to "develop" --slack.token "XXXXXX" --slack.channel "general" --command.after.git.pull "supervisorctl restart target_app"
# watching local development branch push after run

```