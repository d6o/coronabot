# CoronaBot ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/coronabot) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

## What is CoronaBot?
CoronaBot is a simple Slack/Telegram bot to update users about the Corona virus spread progression periodically.

## Where does the data come from?
Data comes from ArcGIS REST API.

## Project Status
CoronaBot is on beta. Pull Requests [are welcome](https://github.com/DiSiqueira/coronabot#social-coding)

# How to use this image

Start an instance of Slack bot

```bash 
$ docker run -e SLACK_TOKEN=xoxp-1111111-22222-3333-444 -e SLACK_CHANNEL_ID=C5P11AABB22 diegosiqueira/coronabot
```

Start an instance of Telegram bot. Bot will start sending notifications after recieving `/start` command. `/stop` command stops sending of notifications.

```bash 
$ docker run -e TELEGRAM_TOKEN=123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11 diegosiqueira/coronabot
```

Use a custom notify interval (same with Telegram bot)
```bash
$ docker run -e SLACK_TOKEN=xoxp-1111111-22222-3333-444 -e SLACK_CHANNEL_ID=C5P11AABB22 -e NOTIFY_INTERVAL_MINUTES=30 diegosiqueira/coronabot
```

Slack message example:
# ![CoronaBot](https://i.imgur.com/oYKRPHL.png)

Telegram message example:
# ![CoronaBot](https://i.imgur.com/ShwKeOh.jpg)

## Building from source

```bash
$ go install github.com/DiSiqueira/coronabot
$ coronabot
$ export SLACK_TOKEN="xoxp-1111111-22222-3333-444"
$ export SLACK_CHANNEL_ID="C5P11AABB22"
$ coronabot
```

## Environment Variables

*SLACK_TOKEN*

Slack token with permissions to post on a channel. Hot to generate a Slack Token: https://slack.com/help/articles/215770388

*SLACK_CHANNEL_ID*

Slack channel id can be found as the last argument on the channel url. Example channel url: https://app.slack.com/client/T0LC9999F/C5P111QZB5 Example channel id: C5P111QZB5

*TELEGRAM_TOKEN*

Telegram bot token (use `@botfather` in Telegram to create a bot).

*NOTIFY_INTERVAL_MINUTES*

Interval that the bot will update the chat with the latest CoronaVirus updates in minutes. Defaults to 60.

## Roadmap
* Listen to termination signals
* Unit tests
* Integration tests
* Improve documentation
* Improve message design
* Monitor a single Country/Region
* Live updates for changes in a Country/Region

## Developing

PRs are welcome. To begin developing, do this:

```bash
$ git clone git@github.com:DiSiqueira/corobabot.git
$ cd corobanot/
$ go mod vendor
$ go run cmd/coronabot/main.go
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/DiSiqueira/coronabot/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2020 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.