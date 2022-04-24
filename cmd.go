package pomo

import (
	"math"
	"time"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
	"github.com/rwxrob/to"
	"github.com/rwxrob/vars"
)

var (
	Duration   = "52m" // max length of Twitch no-ad run
	Interval   = "20s" // default StreamLabs clip length
	Warn       = "1m"
	Prefix     = "🍅"
	PrefixWarn = "💢"
)

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
	Z.Dynamic[`dduration`] = func() string { return Duration }
	Z.Dynamic[`dinterval`] = func() string { return Interval }
	Z.Dynamic[`dwarn`] = func() string { return Warn }
	Z.Dynamic[`dprefix`] = func() string { return Prefix }
	Z.Dynamic[`dprefixwarn`] = func() string { return PrefixWarn }
}

var Cmd = &Z.Cmd{
	Name: `pomo`,
	Commands: []*Z.Cmd{
		printCmd,                     // default
		help.Cmd, vars.Cmd, conf.Cmd, // common
		initCmd, startCmd, stopCmd,
	},
	Shortcuts: Z.ArgMap{
		`started`:    {`var`, `started`},
		`duration`:   {`var`, `set`, `duration`},
		`warn`:       {`var`, `set`, `warn`},
		`prefix`:     {`var`, `set`, `prefix`},
		`prefixwarn`: {`var`, `set`, `prefixwarn`},
	},
	Summary:   `sets or prints a countdown timer (with tomato)`,
	Version:   `v0.1.1`,
	Copyright: `(c) Robert S. Muhlestein <rob@rwx.gg> (rwxrob.tv)`,
	License:   `Apache-2.0`,
	Source:    `https://github.com/rwxrob/pomo`,
	Issues:    `https://github.com/rwxrob/pomo/issues`,

	Description: `
		The {{cmd .Name}} command is a simple tool to help people follow the  
	  Pomodoro method of time boxing. Many add 
		{{print "#(" .Name ")" | pre }}
		to their {{cmd "tmux"}} status lines and turn up the
		refresh to one second.`,
}

var initCmd = &Z.Cmd{
	Name:     `init`,
	Summary:  `sets all values to defaults`,
	Commands: []*Z.Cmd{help.Cmd},

	Description: `
		The {{cmd .Name}} command sets all cached variables to their initial
		values. Any variable name from {{cmd "conf"}} will be used to
		initialize if defined.  Otherwise, the following hard-coded package
		globals will be used instead:

		    duration    - {{dduration}}
		    interval    - {{dinterval}}
		    warn        - {{dwarn}}
		    prefix      - {{dprefix}}
		    prefixwarn  - {{dprefixwarn}}
	`,

	Call: func(x *Z.Cmd, _ ...string) error {

		val, _ := x.Caller.C(`duration`)
		if val == "null" {
			val = Duration
		}
		x.Caller.Set(`duration`, val)

		val, _ = x.Caller.C(`interval`)
		if val == "null" {
			val = Interval
		}
		x.Caller.Set(`interval`, val)

		val, _ = x.Caller.C(`warn`)
		if val == "null" {
			val = Warn
		}
		x.Caller.Set(`warn`, val)

		val, _ = x.Caller.C(`prefix`)
		if val == "null" {
			val = Prefix
		}
		x.Caller.Set(`prefix`, val)

		val, _ = x.Caller.C(`prefixwarn`)
		if val == "null" {
			val = PrefixWarn
		}
		x.Caller.Set(`prefixwarn`, val)

		return nil
	},
}

var printCmd = &Z.Cmd{
	Name:     `print`,
	Aliases:  []string{`show`, `p`},
	Summary:  `print current to standard output (default)`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, _ ...string) error {

		started, err := x.Caller.Get(`started`)
		if err != nil {
			return err
		}
		if started == "" {
			return nil
		}

		endt, err := time.Parse(time.RFC3339, started)
		if err != nil {
			return err
		}

		var subt time.Duration
		subc, err := x.Caller.Get(`interval`)
		if err != nil {
			return err
		}

		if subc != "" {
			subt, err = time.ParseDuration(subc)
			if err != nil {
				return err
			}
		}

		prefix, err := x.Caller.Get(`prefix`)
		if err != nil {
			return err
		}

		prefixwarn, err := x.Caller.Get(`prefixwarn`)
		if err != nil {
			return err
		}

		warn, err := x.Caller.Get(`warn`)
		if err != nil {
			return err
		}
		warnt, err := time.ParseDuration(warn)
		if err != nil {
			return err
		}

		sec := time.Second
		left := endt.Sub(time.Now()).Round(sec)

		var sub float64
		if subc != "" {
			sub = math.Abs(math.Mod(left.Seconds(), subt.Seconds()))
		}

		if left < warnt && left%(sec*2) == 0 {
			prefix = prefixwarn
		}

		if subc != "" {
			term.Printf("%v%v(%02v)", prefix, to.StopWatch(left), sub)
		} else {
			term.Printf("%v%v", prefix, to.StopWatch(left))
		}

		return nil
	},
}

var startCmd = &Z.Cmd{
	Name:     `start`,
	Summary:  `start the current pomo timer`,
	Commands: []*Z.Cmd{help.Cmd},
	MaxArgs:  1,
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) > 0 {
			if err := x.Caller.Set("duration", args[0]); err != nil {
				return err
			}
		}
		s, err := x.Caller.Get("duration")
		if err != nil {
			return err
		}
		if s == "" {
			s = Duration
		}
		dur, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		started := time.Now().Add(dur).Format(time.RFC3339)
		return x.Caller.Set("started", started)
	},
}

var stopCmd = &Z.Cmd{
	Name:     `stop`,
	Summary:  `stop the current pomo timer`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		x.Caller.Del("started")
		return nil
	},
}
