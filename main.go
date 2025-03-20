package main

import (
	"context"
	"fmt"
	"log"
	"os"

	//"github.com/bookshelfdave/gpredikit/cli"
	"github.com/bookshelfdave/gpredikit/commands"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "predikit",
		Usage: "A simple systems testing language",
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("Nothing to do.")
			fmt.Println("Run predikit -h for more info")
			return nil
		},
		Commands: []*cli.Command{
			commands.RunCommand(),
			{
				Name:    "parse",
				Aliases: []string{"p"},
				Usage:   "Parse input files without running tests",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("added task: ", cmd.Args().First())
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
