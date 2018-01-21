# slacksend

Send messages to an incoming webhook via the command line.
Useful if you want to automate sending messages via [Slack](https://slack.com).

## Installation

Download the latest release (Windows, Mac or Linux): https://github.com/ReneGa/slacksend/releases

If you have [Go](https://golang.org/) installed, you can also use `go get`.

```sh
go get github.com/ReneGa/slacksend
```

This downloads and compiles the latest source and installs the binary.

## Setup

Install [Incoming WebHook](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks) on your Slack workspace.

## Usage

### Examples

#### Mac / Linux

```Bash
./slacksend --secret=<secret> --channel "alerts" --text="this message is sent by slacksend."
```

#### Windows

```PowerShell
.\slacksend.exe --secret=<secret> --channel "alerts" --text="this message is sent by slacksend."
```

### Flags

- `secret`: Get your `<secret>` from the [API Guide](https://api.slack.com/incoming-webhooks). It's the part behind "services/" in `https://hooks.slack.com/services/<secret>`.
- `channel`: Optional. If not set, messages will be sent to channel `general`.
- `text`: the message's text
