package cmd

import (
	"errors"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"wxapkg/util"
)

var (
	err_  = color.New(color.FgRed)
	warn_ = color.New(color.FgYellow)
	info_ = color.New(color.FgCyan)
)

var (
	ErrInvalidWXAPkg = errors.New("invalid wxapkg file")
)

func New() *cli.App {
	app := cli.NewApp()
	app.Name = "wxapkg"
	app.Usage = "wxapkg analysis tool for macos"
	app.Version = "v0.0.1"
	app.Commands = []*cli.Command{
		{
			Name:  "unpack",
			Usage: "unpack .wxapkg file",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "in", Required: true, Usage: ".wxapkg file path"},
				&cli.StringFlag{Name: "out", Required: false, Value: "./unpack_out", Usage: "unpacked output path"},
				&cli.BoolFlag{Name: "format", Value: false, Usage: "format content (e.g. js html json)"},
				&cli.BoolFlag{Name: "v", Value: false, Usage: "more info"},
			},
			Action: util.Unpackall,
		},
		{
			Name:   "list",
			Usage:  "list macOS Wechat .wxapkg file",
			Flags:  []cli.Flag{},
			Action: util.List,
		},

		{
			Name:   "clean",
			Usage:  "Clean macOS Wechat .wxapkg file",
			Flags:  []cli.Flag{},
			Action: util.Clean,
		},

		{
			Name:   "auto",
			Usage:  "自动化执行list、unpack、clean",
			Flags:  []cli.Flag{},
			Action: util.Auto,
		},

		//{
		//	Name:  "search",
		//	Usage: "Search for files",
		//	Flags: []cli.Flag{
		//		&cli.StringFlag{Name: "dir", Required: false, Usage: "文件路径"},
		//	},
		//	Action: util.Search,
		//},
	}
	return app
}
