# slacksend

Send messages to an incoming webhook via the command line.
Useful if you want to automate sending messages via [Slack](https://slack.com).

## Usage

### Examples

#### Mac / Linux

```Bash
./slacksend --secret <secret> --channel "alerts" --text "this message is sent by slacksend."
```

#### Windows

```PowerShell
.\slacksend.exe --secret <secret> --channel "alerts" --text "this message is sent by slacksend."
```

### Flags

- `secret`: The secret part of your Incoming Webhook URL. You can get it from the [API Guide](https://api.slack.com/incoming-webhooks). Format: `T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX`
- `channel`: Optional. If not set, messages will be sent to channel `general`. `#` prefix is optional.
- `text`: the message's text
## Installation

Download the latest release (Windows, Mac or Linux): https://github.com/ReneGa/slacksend/releases

If you have [Go](https://golang.org/) installed, you can also use `go get`.

```sh
go get -u github.com/ReneGa/slacksend
```

This downloads and compiles the latest source and installs the binary.

## Setup

Install [Incoming WebHook](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks) on your Slack workspace.
