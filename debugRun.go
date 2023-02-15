// terraform-provider-solacebroker
//
// Copyright 2023 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !windows

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func findDebugPort() (string, error) {
	entries, err := os.ReadDir(os.TempDir())
	if err != nil {
		return "", err
	}
	var sockets []os.DirEntry
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), "plugin") && e.Type() == os.ModeSocket {
			sockets = append(sockets, e)
		}
	}
	if len(sockets) == 0 {
		return "", errors.New("could not find debug reattach socket")
	}
	socket := sockets[0]
	info, err := socket.Info()
	if err != nil {
		return "", err
	}
	for _, s := range sockets[1:] {
		i, err := s.Info()
		if err != nil {
			return "", err
		}
		if i.ModTime().After(info.ModTime()) {
			socket, info = s, i
		}
	}
	return filepath.Join(os.TempDir(), socket.Name()), nil
}

func debugRun(debugRun, address string) {
	if debugRun == "" {
		return
	}
	time.Sleep(time.Second)
	socket, err := findDebugPort()
	if err != nil {
		fmt.Printf("error finding debug reattach socket: %v\n", err)
		return
	}
	cmdLine := strings.Split(debugRun, " ")
	cmd := exec.Command(cmdLine[0], cmdLine[1:]...)
	data, err := json.Marshal(map[string]any{
		address: map[string]any{
			"Protocol":        "grpc",
			"ProtocolVersion": 6,
			"Pid":             os.Getpid(),
			"Test":            true,
			"Addr": map[string]any{
				"Network": "unix",
				"String":  socket,
			},
		},
	})
	//fmt.Printf("debug run computed TF_REATTACH_PROVIDERS=%s\n", data)
	cmd.Env = append(
		[]string{"TF_REATTACH_PROVIDERS=" + string(data)},
		os.Environ()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("error running %v: %v\n", debugRun, err)
	} else {
		fmt.Printf("%v successful.\n", debugRun)
	}
	// send an interrupt rather than just os.Exit so that proper cleanup will be done
	syscall.Kill(os.Getpid(), syscall.SIGINT)
	// close stderr so that we don't get spurious error messages when debugging
	os.Stderr.Close()
}
