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

package tail

import (
	"context"
	"io"
	"log"

	"github.com/docker/docker/pkg/stdcopy"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerTailer struct {
	cli *client.Client
}

func NewLocalDockerTailer() (Tailer, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &DockerTailer{
		cli: cli,
	}, nil
}

func (t *DockerTailer) Tail(c string, writer io.Writer) error {
	logs, err := t.cli.ContainerLogs(context.Background(), c, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer logs.Close()

	container, err := t.cli.ContainerInspect(context.Background(), c)
	if err != nil {
		return err
	}

	// The stream format on the response will be in one of two formats:
	//
	// If the container is using a TTY, there is only a single stream (stdout),
	// and data is copied directly from the container output stream, no extra multiplexing or headers.
	//
	//If the container is *not* using a TTY, streams for stdout and stderr are multiplexed.
	if container.Config.Tty {
		_, err := io.Copy(writer, logs)
		if err != nil {
			return err
		}
	} else {
		_, err = stdcopy.StdCopy(writer, writer, logs)
		if err != nil {
			return err
		}
	}
	return nil
}
