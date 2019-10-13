package cmd

import (
	"fmt"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/appLoader"
	"github.com/dev2choiz/f7k/command"
	"github.com/dev2choiz/f7k/configurator"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/spf13/cobra"
)

func VersionCmd() interfaces.Command {
	c := &command.Command{}
	c.CobraCmd = &cobra.Command {
		Use:   "version",
		Short: "F7k version",
		Long:  `F7k version`,
		Run: func(cmd *cobra.Command, args []string) {
			appLoader.GetCliServerLoader(&configurator.Config{}, "", "",false, false).Load()
			fmt.Println(string(pictures[0]))
			fmt.Printf("%s version %s\n", f7k.PrettyName, f7k.Version)
		},
	}

	return c
}

var pictures = [][]byte{
	[]byte("\x0a\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x95\x97\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x95\x97\xe2\x96\x88\xe2\x96\x88\xe2\x95\x97\x20\x20\xe2\x96\x88\xe2\x96\x88\xe2\x95\x97\x0a\xe2\x96\x88\xe2\x96\x88\xe2\x95\x94\xe2\x95\x90\xe2\x95\x90\xe2\x95\x90\xe2\x95\x90\xe2\x95\x9d\xe2\x95\x9a\xe2\x95\x90\xe2\x95\x90\xe2\x95\x90\xe2\x95\x90\xe2\x96\x88\xe2\x96\x88\xe2\x95\x91\xe2\x96\x88\xe2\x96\x88\xe2\x95\x91\x20\xe2\x96\x88\xe2\x96\x88\xe2\x95\x94\xe2\x95\x9d\x0a\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x95\x97\x20\x20\x20\x20\x20\x20\xe2\x96\x88\xe2\x96\x88\xe2\x95\x94\xe2\x95\x9d\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x96\x88\xe2\x95\x94\xe2\x95\x9d\x0a\xe2\x96\x88\xe2\x96\x88\xe2\x95\x94\xe2\x95\x90\xe2\x95\x90\xe2\x95\x9d\x20\x20\x20\x20\x20\xe2\x96\x88\xe2\x96\x88\xe2\x95\x94\xe2\x95\x9d\x20\xe2\x96\x88\xe2\x96\x88\xe2\x95\x94\xe2\x95\x90\xe2\x96\x88\xe2\x96\x88\xe2\x95\x97\x0a\xe2\x96\x88\xe2\x96\x88\xe2\x95\x91\x20\x20\x20\x20\x20\x20\x20\x20\xe2\x96\x88\xe2\x96\x88\xe2\x95\x91\x20\x20\xe2\x96\x88\xe2\x96\x88\xe2\x95\x91\x20\x20\xe2\x96\x88\xe2\x96\x88\xe2\x95\x97\x0a\xe2\x95\x9a\xe2\x95\x90\xe2\x95\x9d\x20\x20\x20\x20\x20\x20\x20\x20\xe2\x95\x9a\xe2\x95\x90\xe2\x95\x9d\x20\x20\xe2\x95\x9a\xe2\x95\x90\xe2\x95\x9d\x20\x20\xe2\x95\x9a\xe2\x95\x90\xe2\x95\x9d\x0a"),
}
