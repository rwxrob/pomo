package pomo

import _ "embed"

// Keep all documentation in this file as embeds in order to easily
// support compiles to other target languages by simply changing the
// language identifier before compilation.

//go:embed text/en/pomo.md
var _pomo string

//go:embed text/en/init.md
var _init string

//go:embed text/en/start.md
var _start string

//go:embed text/en/stop.md
var _stop string

//go:embed text/en/print.md
var _print string
