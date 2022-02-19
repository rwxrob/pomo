# Pomodoro Timer Command in Go

**⚠ PORTING TO BONZAI CURRENTLY**

🎉 ***Bonzai shamelessly requires Go 1.18+*** 💋

[![Go Version](https://img.shields.io/github/go-mod/go-version/rwxrob/bonzai)](https://tip.golang.org/doc/go1.18)
[![GoDoc](https://godoc.org/bonzai-pomo?status.svg)](https://godoc.org/bonzai-pomo)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report
Card](https://goreportcard.com/badge/bonzai-pomo)](https://goreportcard.com/report/bonzai-pomo)

## Install `pomo` as Standalone

The `pomo` command can be used as a standalone program

```
go install github.com/rwxrob/bonzai-pomo@latest
```

## Usage

```
pomo help
```

## Add Pomodoro to TMUX

Here's an example of how to add `pomo` to your TMUX configuration. Your
mileage may vary.

```tmux
set -g status-interval 1
set -g status-right "#(pomo)"
```
