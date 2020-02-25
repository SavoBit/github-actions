package main

import (
	"github.com/docker/github-actions/internal/command"
	"github.com/docker/github-actions/internal/options"
)

func buildPush(cmd command.Runner) error {
	github, err := options.GetGitHubOptions()
	if err != nil {
		return err
	}

	server := options.GetServer()
	tags := options.GetTags(server, github)

	login, err := options.GetLoginOptions()
	if err != nil {
		return err
	}
	if login.Username != "" && login.Password != "" {
		if err = runLogin(cmd, login, server); err != nil {
			return err
		}
	}

	build, err := options.GetBuildOptions()
	if err != nil {
		return err
	}
	if err = runBuild(cmd, build, github, tags); err != nil {
		return err
	}

	return runPush(cmd, tags)
}
