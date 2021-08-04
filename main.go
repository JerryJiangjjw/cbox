package main

import (
	"fmt"
	"os"
	"strconv"

	"cbox/tools"

	"github.com/urfave/cli/v2"
)

func main() {

	var timeStampToData,
		dateTotimeStamp,
		urlEncode,
		urlDecode,
		base64Encode,
		base64Decode string

	fmt.Println(tools.Banner)

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
			&cli.StringFlag{
				Name:        "timestamptodate",
				Usage:       "timestamp to date",
				Aliases:     []string{"ttd"},
				Destination: &timeStampToData,
			},
			&cli.StringFlag{
				Name:        "datetotimestmap",
				Usage:       "date to timestamp",
				Aliases:     []string{"dtt"},
				Destination: &dateTotimeStamp,
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
			if base64Decode != "" {
				fmt.Println(tools.Base64Decode(base64Decode))
			}
			if timeStampToData != "" {
				int64, _ := strconv.ParseInt(timeStampToData, 10, 64)
				fmt.Println(tools.TimeStampToDate(int64, "2006-01-02 15:04:05"))
			}
			if dateTotimeStamp != "" {
				fmt.Println(dateTotimeStamp)
				int64, err := tools.DateToTimeStmap(dateTotimeStamp, "2006-01-02 15:04:05")
				if err == nil {
					fmt.Println(int64)
				} else {
					fmt.Println(err)
				}
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
