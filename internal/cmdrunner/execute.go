package cmdrunner

import (
	"github.com/git-duet/git-duet"
	"github.com/git-duet/git-duet/internal/cmd"
)

func Execute(commands ...cmd.Command) error {
	println("***** executing ***** ")
	configuration, err := duet.NewConfiguration()
	if err != nil {
		return err
	}

	var gitConfig *duet.GitConfig
	if configuration.Global {
		gitConfig = &duet.GitConfig{
			Namespace:     configuration.Namespace,
			Scope:         duet.Global,
			SetUserConfig: configuration.SetGitUserConfig,
		}
	} else {
		gitConfig, err = duet.GetAuthorConfig(configuration.Namespace, configuration.SetGitUserConfig)
		if err != nil {
			return err
		}
	}

	var amending = false
	for _, command := range commands {
		if contains(command.Args, "--amend") {
			amending = true
		}
		if err := command.Execute(); err != nil {
			return err
		}
	}

	if !amending && configuration.RotateAuthor {
		println("**** rotating ****")
		if err := gitConfig.RotateAuthor(); err != nil {
			return err
		}
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}