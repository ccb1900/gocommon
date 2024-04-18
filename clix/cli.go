package clix

import (
	"fmt"
	"os"
	"text/template"

	"github.com/ccb1900/gocommon/logger"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
)

type CliApp struct {
	instance *cli.App
}

func New(name string) *CliApp {
	app := cli.NewApp()
	app.Name = name
	return &CliApp{
		instance: app,
	}
}

func (c *CliApp) Run() {
	c.RunWithArgs(os.Args)
}

func (c *CliApp) RunWithArgs(args []string) {
	if err := c.instance.Run(args); err != nil {
		logger.Default().Info("boot fail", "err", err)
		os.Exit(-1)
	}
}

func execTpl(c *cli.Context, src, dst string) error {
	t := template.Must(template.ParseFiles(src))
	w, err := os.Create(dst)
	if err != nil {
		logger.Default().Info("fail to create service", "err", err)
		return err
	}

	if err := w.Chmod(0o755); err != nil {
		logger.Default().Info("fail to chmod", "err", err)
		return err
	}

	if err := t.Execute(w, map[string]string{
		"desc": c.String("desc"),
		"cmd":  c.String("cmd"),
		"boot": c.String("arg"),
	}); err != nil {
		logger.Default().Info("fail to execute template", "err", err)
		return err
	}

	return nil
}

func (c *CliApp) AddSysService() {
	c.instance.Commands = append(c.instance.Commands, &cli.Command{
		Name:  "gen:service",
		Usage: "gen linux service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "desc",
				Required: true,
				Usage:    "service description",
				Aliases:  []string{"d"},
			},
			&cli.StringFlag{
				Name:     "cmd",
				Required: true,
				Aliases:  []string{"c"},
				Usage:    "service name",
			},
			&cli.StringFlag{
				Name:     "arg",
				Required: true,
				Aliases:  []string{"a"},
				Usage:    "service name",
			},
		},
		Action: func(c *cli.Context) error {
			items := map[string]string{
				"app.service.template":       "app.service",
				"scripts/tpl/postinstall.sh": "scripts/postinstall.sh",
				"scripts/tpl/preremove.sh":   "scripts/preremove.sh",
				"scripts/tpl/postremove.sh":  "scripts/postremove.sh",
				"scripts/tpl/preinstall.sh":  "scripts/preinstall.sh",
			}

			for k, v := range items {
				if err := execTpl(c, k, v); err != nil {
					return err
				}
			}

			logger.Default().Info("create service success")

			return nil
		},
	})
}

func (c *CliApp) AddGenPassword() {
	c.instance.Commands = append(c.instance.Commands, &cli.Command{
		Name:  "gen:password",
		Usage: "gen password",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "src",
				Required: true,
				Usage:    "password",
				Aliases:  []string{"s"},
			},
		},
		Action: func(c *cli.Context) error {
			res, err := bcrypt.GenerateFromPassword([]byte(c.String("desc")), bcrypt.DefaultCost)
			if err != nil {
				return err
			}

			fmt.Println(string(res))
			return nil
		},
	})
}

func (c *CliApp) AddCommand(cmd *cli.Command) {
	c.instance.Commands = append(c.instance.Commands, cmd)
}
