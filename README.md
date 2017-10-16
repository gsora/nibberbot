# `nibberbot` - a Telegram bot to write like a ni🅱️🅱️a

## Usage

`nibberbot` is go get-able:

```
$ go get -u github.com/gsora/nibberbot/...
$ nibberbot
Usage of nibberbot:
  -apikey string
    	required, Telegram bot API key
  -cert string
    	required, TLS certificate path
  -debug
    	debug Telegram bot interactions
  -domain string
    	required, domain associated to the TLS cert+key and the server where this bot will be running
  -key string
    	required, TLS key path
  -numspaces int
    	number of spaces between each word (default 1)
  -port string
    	port to run on, must be 443, 80, 88, 8443 (default "88")
```

`nibberbot` works with both inline request and direct messages.

Letsencrypt users must provide `fullchain.pem` as `cert`.

## Hacking

To add more emojis, one should add the emoji UTF-8 hex representation to the `Emojis` map contained in the `nibber/unleash.go` file.
You can get the hex [here](https://apps.timwhitlock.info/emoji/tables/unicode).

For example, if one want to add the '🅱️' emoji:

```go

var Emojis = map[string]string {
	...
	"b": "\xF0\x9F\x85\xB1",
	...
}

```

The `OrderedSubstitution` type will handle uppercase and the ordering.
