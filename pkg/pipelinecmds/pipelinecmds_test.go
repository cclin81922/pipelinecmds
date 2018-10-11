//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package pipelinecmds

import (
	"os/exec"
	"testing"
)

func TestPipeline(t *testing.T) {
	testcases := []struct {
		name     string
		pipeline []*exec.Cmd
		stdout   string
		stderr   string
		exitcode int
	}{
		{
			name: "echo",
			pipeline: []*exec.Cmd{
				&exec.Cmd{Path: "echo", Args: []string{"ok"}},
			},
			stdout:   "ok",
			stderr:   "",
			exitcode: 0,
		},
		{
			name: "echo then cat",
			pipeline: []*exec.Cmd{
				&exec.Cmd{Path: "echo", Args: []string{"ok"}},
				&exec.Cmd{Path: "cat"},
			},
			stdout:   "ok",
			stderr:   "",
			exitcode: 0,
		},
		{
			name: "echo then grep",
			pipeline: []*exec.Cmd{
				&exec.Cmd{Path: "echo", Args: []string{"-e", "hello\nworld"}},
				&exec.Cmd{Path: "grep", Args: []string{"hello"}},
			},
			stdout:   "hello",
			stderr:   "",
			exitcode: 0,
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			finalStdout, finalStderr, anyError := Pipeline(tc.pipeline...)
			if anyError != nil {
				t.Fatal(anyError)
			}
			if finalStdout != tc.stdout {
				t.Fatalf("stdout error | expected %s | got %s", tc.stdout, finalStdout)
			}
			if finalStderr != tc.stderr {
				t.Fatalf("stderr error | expected %s | got %s", tc.stderr, finalStderr)
			}
		})
	}
}
