package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"luxcss/lib"
	"os"
)

func main() {
	var watch bool
	var prefix string
	var output string
	var include cli.StringSlice
	var extensions cli.StringSlice

	app := &cli.App{
		Name:  "luxcss",
		Usage: "CLI tool for LuxCSS",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "watch",
				Usage:       "Whether the included directories should be watched for changes",
				Value:       false,
				Aliases:     []string{"w"},
				Destination: &watch,
			},
			&cli.StringFlag{
				Name:        "prefix",
				Usage:       "Vendor prefix for class names",
				Value:       "",
				Aliases:     []string{"p"},
				Destination: &prefix,
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       "File to write the output to",
				Value:       "./output.css",
				Aliases:     []string{"o"},
				Destination: &output,
			},
			&cli.StringSliceFlag{
				Name:        "include",
				Usage:       "Directories to include",
				Value:       cli.NewStringSlice("."),
				Aliases:     []string{"i"},
				Destination: &include,
			},
			&cli.StringSliceFlag{
				Name:        "extensions",
				Usage:       "File extensions to include",
				Value:       cli.NewStringSlice("html", "js"),
				Aliases:     []string{"e"},
				Destination: &extensions,
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println("Hello world")

			fmt.Println("watch: ", watch)
			fmt.Println("prefix: ", prefix)
			fmt.Println("output: ", output)
			fmt.Println("include: ", include)
			fmt.Println("extensions: ", extensions)

			lib.Extract(include.Value(), extensions.Value(), prefix)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
