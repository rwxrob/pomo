package pomo

import "github.com/rwxrob/bonzai"

var Cmd = &bonzai.Cmd{
	Name:      `pomo`,
	Hidden:    []string{`up`},
	Usage:     `[COMMAND]`,
	Summary:   `sets or prints a countdown timer (with tomato)`,
	Version:   `v2.0.0`,
	Copyright: `(c) Robert S. Muhlestein <rob@rwx.gg> (rwxrob.tv)`,
	License:   `Apache-2`,
	Source:    `https://github.com/rwxrob/cmdbox-pomo`,
	Issues:    `https://github.com/rwxrob/cmdbox-pomo/issues`,
	Description: `
		The *pomo* command assists those with creating scripts and other
		tools to help them follow the simple Pomodoro method of time boxing.
		When called without arguments *pomo show* is assumed.`,
}

/*

func init() {
	setDefaults := func() {
		v := config.Get("pomo.duration")
		if v == "" {
			v = "48m49s"
			config.Set("pomo.duration", v)
		}
		v = config.Get("pomo.sub")
		if v == "" {
			v = "20s"
			config.Set("pomo.sub", v)
		}
		v = config.Get("pomo.warning")
		if v == "" {
			v = "1m"
			config.Set("pomo.warning", v)
		}
		v = config.Get("pomo.emoji")
		if v == "" {
			v = "🍅"
			config.Set("pomo.emoji", v)
		}
		v = config.Get("pomo.warning.emoji")
		if v == "" {
			v = "💢"
			config.Set("pomo.warning.emoji", v)
		}
	}
}

	x := cmdbox.Add("pomo", "show", "start", "stop",
		"set", "sub", "emoji",
		"we|warning.emoji", "w|warn|warning",
		"f|file", "e|vi|ed|edit", "up",
		"defaults",
	)
	x.Default = "pomo show"

	x = cmdbox.Add("pomo file")
	x.Summary = `show full path to configuration file`
	x.Method = func(args ...string) error {
		fmt.Println(config.Path())
		return nil
	}

	x = cmdbox.Add("pomo up")
	x.Summary = `show date/time started`
	x.Method = func(args ...string) error {
		fmt.Println(config.Get("pomo.up"))
		return nil
	}

	x = cmdbox.Add("pomo warning.emoji")
	x.Usage = `[EMOJI]`
	x.Summary = `show/set warning emoji`
	x.Method = func(args ...string) error {
		if len(args) == 0 {
			fmt.Println(config.Get("pomo.warning.emoji"))
		} else {
			config.Set("pomo warning.emoji", args[0])
		}
		return nil
	}

	x = cmdbox.Add("pomo emoji")
	x.Usage = `[EMOJI]`
	x.Summary = `show/set current emoji`
	x.Method = func(args ...string) error {
		if len(args) == 0 {
			fmt.Println(config.Get("pomo.emoji"))
		} else {
			config.Set("pomo.emoji", args[0])
		}
		return nil
	}

	x = cmdbox.Add("pomo stop")
	x.Summary = `stop pomo timer without resetting`
	x.Method = func(args ...string) error {
		config.Set("pomo.up", "")
		return nil
	}

	x = cmdbox.Add("pomo show")
	x.Summary = `show the current pomo timer (default)`
	x.Method = func(args ...string) error {
		up := config.Get("pomo.up")
		if up == "" {
			return nil
		}
		endt, err := time.Parse(time.RFC3339, up)
		if err != nil {
			return err
		}
		subc := config.Get("pomo.sub")
		if subc == "0" || subc == "0s" {
			subc = "1s"
		}
		subt, err := time.ParseDuration(subc)
		if err != nil {
			return err
		}
		emoji := config.Get("pomo.emoji")
		warnEmoji := config.Get("pomo.warning.emoji")
		warning, err := time.ParseDuration(config.Get("pomo.warning"))
		if err != nil {
			return err
		}
		sec := time.Second
		timeLeft := endt.Sub(time.Now()).Round(sec)
		sub := int64(timeLeft.Seconds()) % int64(subt.Seconds())
		if timeLeft < warning && timeLeft%(sec*2) == 0 {
			fmt.Printf("%v%v(%v)", warnEmoji, timeLeft, sub)
			fmt.Printf("\n")
			return nil
		}
		fmt.Printf("%v%v(%v)", emoji, timeLeft, sub)
		fmt.Printf("\n")
		return nil
	}

	x = cmdbox.Add("pomo start")
	x.Summary = `start the current pomo timer (without showing)`
	x.Method = func(args ...string) error {
		setDefaults()
		s := config.Get("pomo.duration")
		dur, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		up := time.Now().Add(dur).Format(time.RFC3339)
		config.Set("pomo.up", up)
		return nil
	}

	x = cmdbox.Add("pomo set")
	x.Usage = "[DURATION]"
	x.Summary = `show/set duration and start over`
	x.Method = func(args ...string) error {
		if len(args) > 0 {
			config.Set("pomo.duration", args[0])
			return x.Call("pomo start")
		}
		fmt.Println(config.Get("pomo.duration"))
		return nil
	}

	x = cmdbox.Add("pomo warning")
	x.Usage = "[DURATION]"
	x.Summary = `show/set warning seconds remaining`
	x.Method = func(args ...string) error {
		if len(args) > 0 {
			config.Set("pomo.warning", args[0])
		}
		fmt.Println(config.Get("pomo.warning"))
		return nil
	}

	x = cmdbox.Add("pomo edit")
	x.Summary = `edit pomo configuration file`
	x.Method = func(args ...string) error {
		return config.Edit()
	}

	x = cmdbox.Add("pomo sub")
	x.Summary = `show/set repeating sub-duration interval`
	x.Description = `
		Set to 0 to keep from appearing. Defaults to 20s (the normal
		time for a highlight video).`
	x.Method = func(args ...string) error {
		if len(args) > 0 {
			config.Set("pomo.sub", args[0])
		}
		fmt.Println(config.Get("pomo.sub"))
		return nil
	}

	x = cmdbox.Add("pomo defaults")
	x.Summary = `reset to defaults`
	x.Method = func(args ...string) error {
		os.Remove(config.Path())
		return nil
	}

}
*/
