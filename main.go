package main

import (
	"context"
	"fmt"
	"gobd/internal/build"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var buildAll bool
	var buildMain bool
	var buildNoDebug bool
	var buildNoCgo bool
	var buildOs string
	var buildArch string
	var buildOutputDirectory string
	var buildOutputName string
	var buildOpts []string
	var buildEnvs []string

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:        "all",
			Usage:       "set build all supported os and architecture",
			Destination: &buildAll,
		},
		&cli.BoolFlag{
			Name:        "main",
			Usage:       "set build all supported architecture for windows, macos, linux and freebsd",
			Destination: &buildMain,
		},
		&cli.StringFlag{
			Name:        "os",
			Usage:       "set build operating system",
			Destination: &buildOs,
		},
		&cli.StringFlag{
			Name:        "arch",
			Usage:       "set build architecture",
			Destination: &buildArch,
		},
		&cli.StringFlag{
			Name:        "name",
			Aliases:     []string{"n"},
			Usage:       "set build output name",
			Destination: &buildOutputName,
		},
		&cli.StringFlag{
			Name:        "dir",
			Aliases:     []string{"d"},
			Usage:       "set build output directory",
			Destination: &buildOutputDirectory,
		},
		&cli.BoolFlag{
			Name:        "no_debug",
			Usage:       "set build not using debug options to reduce compile size",
			Destination: &buildNoDebug,
		},
		&cli.BoolFlag{
			Name:        "no_cgo",
			Usage:       "set build not using cgo to avoid relying on the host operating system's native libraries",
			Destination: &buildNoCgo,
		},
		&cli.StringSliceFlag{
			Name:        "opts",
			Usage:       "set build opts",
			Destination: &buildOpts,
		},
		&cli.StringSliceFlag{
			Name:        "envs",
			Usage:       "set build envs",
			Destination: &buildEnvs,
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Printf("%s\n", cmd.Root().Version)
	}

	cmd := &cli.Command{
		Usage:   "Golang Build Tool",
		Version: "v2.10",
		Flags:   flags,
		Action: func(ctx context.Context, cmd *cli.Command) (err error) {
			var ps []build.Pair

			// 获取编译的操作系统/处理器架构对
			if buildMain {
				ps = build.GetMainPairs()
			} else if buildAll {
				ps, err = build.GetAllPairs()
				if err != nil {
					return err
				}
			} else {
				ps = build.GetSelectedPairs(buildOs, buildArch)
			}
			// 获取不到操作系统/处理器架构对则返回错误
			if len(ps) == 0 {
				return fmt.Errorf("can't find any pair")
			}

			// 遍历操作系统/处理器架构对进行编译
			for _, p := range ps {
				err = build.Build(p.OS, p.ARCH, buildOutputName, buildOutputDirectory, buildNoDebug, buildNoCgo, buildOpts, buildEnvs)
				if err != nil {
					log.Println(err)
				}
			}
			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
