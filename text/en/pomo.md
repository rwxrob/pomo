sets or prints a countdown timer (with tomato)

The {{aka}} command is a simple tool to help people follow the  Pomodoro method of time boxing. Many add {{print "#(" .Name ")" | pre }} to their {{cmd "tmux"}} status lines and turn up the refresh to one second.

***Disable interval timer***

To disable the interval timer just delete the `interval` variable:

    {{aka}} delete interval

***Set duration to top of the hour***

Set the duration length to whatever remains of the current hour by using the special `hour` parameter instead of a duration to the {{cmd "start"}} command.

    {{aka}} start hour

***Automatically reset on every hour***

If setting the duration to the top of the hour becomes tedious this can be set to automatically reset by assigning the value of `hour` to the `autoreset` variable.

    {{aka}} autoreset hour

To disable it, just {{cmd "delete"}} (or {{cmd "unset"}}) it.

    {{aka}} delete autoreset
