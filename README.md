# Installation

```
go get -u github.com/cclin81922/pipelinecmds/cmd/pipelinecmds
export PATH=$PATH:~/go/bin
```

# Package Usage

```
import "github.com/cclin81922/pipelinecmds/pkg/pipelinecmds"

func demo(cmds ...*exec.Cmd) {
    finalStdout, finalStderr, anyErr := pipelinecmds.Pipeline(cmds...)
}
```

# For Developer

Run all tests

```
go test github.com/cclin81922/pipelinecmds/pkg/pipelinecmds
```

# Related Resources

* [Golang code gist which provides another pipeline implementation](https://gist.github.com/kylelemons/1525278)