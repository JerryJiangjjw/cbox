package main

import (
	"fmt"
	"os"

	"cbox/tools"

	"github.com/urfave/cli/v2"
)

func main() {

	var urlEncode string
	var urlDecode string
	var base64Encode string
	var base64Decode string

	// fmt.Println(tools.Banner)

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "urlencode",
				Usage:       "url encode",
				Aliases:     []string{"ue"},
				Destination: &urlEncode,
			},
			&cli.StringFlag{
				Name:        "urldecode",
				Usage:       "url decode",
				Aliases:     []string{"ud"},
				Destination: &urlDecode,
			},
			&cli.StringFlag{
				Name:        "base64encode",
				Usage:       "base64 encode",
				Aliases:     []string{"b64e"},
				Destination: &base64Encode,
			},
			&cli.StringFlag{
				Name:        "base64decode",
				Usage:       "base64 decode",
				Aliases:     []string{"b64d"},
				Destination: &base64Decode,
			},
		},
		Action: func(c *cli.Context) error {

			if urlEncode != "" {
				fmt.Println(tools.UrlEncode(urlEncode))
			}
			if urlDecode != "" {
				fmt.Println(tools.UrlDecode(urlDecode))
			}
			if base64Encode != "" {
				fmt.Println(tools.Base64Encode(base64Encode))
			}
			if base64Encode != "" {
				fmt.Println(tools.Base64Decode(base64Decode))
			}
			if len(os.Args) < 2 || c.NumFlags() < 2 {
				cli.ShowAppHelp(c)
			}

			return nil
		},
		Name:  "CBox",
		Usage: "Converter Tool Box",
	}

	app.Run(os.Args)
}
