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
	"bytes"
	"os/exec"
)

// FailFastPipeline ...
func FailFastPipeline(cmds ...*exec.Cmd) (finalStdout []byte, finalStderr []byte, firstError error) {
	last := len(cmds) - 1
	for i, cmd := range cmds[:last] {

		var out bytes.Buffer
		cmd.Stdout = &out
		cmds[i+1].Stdin = &out

		if err := cmd.Run(); err != nil {
			firstError = err
			return
		}

	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmds[last].Stdout, cmds[last].Stderr = &stdout, &stderr
	if err := cmds[last].Run(); err != nil {
		firstError = err
		return
	}

	return stdout.Bytes(), stderr.Bytes(), nil
}
