# slacksend

Send messages to an incoming webhook via the command line. Useful if you want to automatically send messages via Slack.

## Setup

Install [Incoming WebHook](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks) on your Slack workspace.

## Flags

- secret: Get your `<secret>` from the [API Guide](https://api.slack.com/incoming-webhooks). It's the part behind "services/" in `https://hooks.slack.com/services/<secret>`.
- channel: Optional. If not set, messages will be sent to channel `general`.
- text: the message's text

## Example usage

### Mac

```Bash
./slacksend --secret=<secret> --channel "alerts" --text="this message is sent by sendslack."
```

### Windows

```PowerShell
.\slacksend.exe --secret=<secret> --channel "alerts" --text="this message is sent by sendslack."
```
