/*
Copyright 2021 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	cecommands "github.com/google/container-explorer/cmd/commands"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	VERSION = "0.0.2"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	app := cli.NewApp()

	app.Name = "container-explorer"
	app.Version = VERSION
	app.Usage = "A standalone utility to explore container details"
	app.Description = `A standalone utility to exploer container details.
	
	Container explorer supports exploring containers managed using containerd and
	docker. The utility also supports exploring containers created and managed using
	Kubernetes.
	`
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug messages",
		},
		cli.StringFlag{
			Name:  "containerd-root, c",
			Usage: "specify containerd root directory",
			Value: "/var/lib/containerd",
		},
		cli.StringFlag{
			Name:  "image-root, i",
			Usage: "specify mount point for a disk image",
		},
		cli.StringFlag{
			Name:  "metadata-file, m",
			Usage: "specify the path to containerd metadata file i.e. meta.db",
		},
		cli.StringFlag{
			Name:  "snapshot-metadata-file, s",
			Usage: "specify the path to containerd snapshot metadata file i.e. metadata.db.",
		},
		cli.StringFlag{
			Name:  "namespace, n",
			Usage: "specify container namespace",
			Value: "default",
		},
		cli.BoolFlag{
			Name:  "docker-managed",
			Usage: "specify docker manages standalone or Kubernetes containers",
		},
		cli.StringFlag{
			Name:  "docker-root",
			Usage: "specify docker root directory. This is only used with flag --docker-managed",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		cecommands.ListCommand,
		cecommands.InfoCommand,
		cecommands.MountCommand,
	}

	app.Before = func(context *cli.Context) error {
		if context.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
