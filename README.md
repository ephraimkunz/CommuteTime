# CommuteTime [![Build Status](https://travis-ci.org/ephraimkunz/CommuteTime.svg?branch=master)](https://travis-ci.org/ephraimkunz/CommuteTime)
![Example graph](https://github.com/ephraimkunz/CommuteTime/blob/master/from-work.png "Example of generate graph")

Find out when the optimal time is to start your commute, taking into account traffic. Get reports emailed to you in the morning and in the afternoon.

## Description
This is a simple, command-line Go program that generates graphs of possible commute times and emails them to you. Which graphs to generate and where to email them are configured with environment variables in `commute.sh`. This is called by `towork.sh` or `fromwork.sh`. Cron should be configured to call these in the morning or the night. Make these scripts executable with `chmod +x <scriptname>`.

## Costs
Be aware that the price of the Google Distance Matrix API calls used to generate the graphs are high. Each call costs 1 cent, meaning that to generate a 6 hour graph with 15 minute resolution and best, worst and likely outcomes will cost around 72 cents (6 * 4 * 3).

## Alternatives
If you don't want to set up a whole Go environment on the system where this cron job will run (i.e. Raspberry Pi), here's another way.
1. Change `commute.sh` by removing the `go build` line.
2. Cross-compile for the target machine on another machine. For Raspberry Pi, run `env GOOS=linux GOARCH=arm GOARM=5 go build
`.
3. Scp this executable and the scripts to the target machine. You should be able to run the scripts there and everything should work.

## Other stuff
* You should use [Cronhub](https://cronhub.io) for monitoring your cron jobs.
