// MIT License
//
// Copyright (c) 2018 Mark
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"log"
	"os"

	"github.com/djmarkoz/docker-tail/logger"
	"github.com/djmarkoz/docker-tail/pkg/tail"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-tail <CONTAINER>...",
	Short: "Tail multiple Docker containers at once",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
		run(args)
	},
}

func run(args []string) {
	tailer, err := tail.NewLocalDockerTailer()
	if err != nil {
		log.Fatal(err)
	}

	// start workers
	numWorkers := len(args)
	workers := make(chan int, numWorkers)
	for _, container := range args {
		go func(c string) {
			err := tailer.Tail(c, logger.NewLogWriter(c))
			if err != nil {
				log.Fatal(err)
			}
			workers <- 0
		}(container)
	}

	// wait for all workers to finish
	for i := 0; i < numWorkers; i++ {
		<-workers
	}
	close(workers)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
