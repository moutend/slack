`slack` -- CLI Slack Viewer
===========================

`slack` is a command line tool for reading messages in Slack.

## Install

```console
go get -u github.com/moutend/slack/cmd/slack
```

## Setup

Before starting, you need prepare API tokens for this tool.

1. Visit [api.slack.com/apps](https://api.slack.com/apps).
2. Create new app.
3. Add the following OAuth scopes:
  - Bot scope
    - `users:read`
  - User scope
    - `channels:history`
    - `channels:read`
    - `groups:history`
    - `groups:read`
    - `im:history`
    - `im:read`
    - `mpim:history`
    - `mpim:read`
4. After setting the scopes, install the app.
5. Your bot and user API tokens are issued.
6. Set the environement variables like this:

```console
export SLACK_BOT_API_TOKEN="xxxxxxxx"
export SLACK_USER_API_TOKEN="zzzzzzzz"
```

That's it!

## Usage

### Print about yourself

```console
slack whoami
```

### Print active users

```console
slack users
```

### Print active channels

```console
slack channels
```

### Print public conversations

```console
slack messages general
```

### Print private conversations

```console
slack messages @someone
```

## LICENSE

MIT

## Author

`Yoshiyuki Koyanagi <moutend@gmail.com>`
