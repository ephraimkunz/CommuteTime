# CommuteTime
Find out when the optimal time is to start your commute, taking into account traffic

## Description
This is a simple, command-line Go program that generates graphs of possible commute times and emails them to you. Which graphs to generate and where to email them are configured with environment variables in `commute.sh`. This is called by `towork.sh` or `fromwork.sh`. Cron should be configured to call these in the morning or the night.

## Costs
Be aware that the price of the Google Distance Matrix API calls used to generate the graphs are high. Each call costs 1 cent, meaning that to generate a 6 hour graph with 15 minute resolution and best, worst and likely outcomes will cost around 72 cents (6 * 4 * 3).
