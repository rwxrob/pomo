# 🌳 Go Pomodoro Command Line Timer

[![GoDoc](https://godoc.org/pomo?status.svg)](https://godoc.org/pomo)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report
Card](https://goreportcard.com/badge/pomo)](https://goreportcard.com/report/pomo)

## Install

You can just grab the latest binary [release](https://github.com/rwxrob/pomo/releases).

This command can be installed as a standalone program or composed into a Bonzai command tree.

Standalone

```
go install github.com/rwxrob/pomo/cmd/pomo@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/pomo"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, pomo.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C pomo pomo
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.

## Add Pomodoro to TMUX

Here's an example of how to add `pomo` to your TMUX configuration. Your
mileage may vary.

```tmux
set -g status-interval 1
set -g status-right "#(pomo)"
```
