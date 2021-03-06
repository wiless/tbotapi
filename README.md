# TBotAPI - Telegram Bot-API for Go #

This is a simple wrapper for the Telegram Bot-API for Go. It provides a high-level API in Go that works with the Telegram REST API.

[![GoDoc](https://godoc.org/bitbucket.org/mrd0ll4r/tbotapi?status.svg)](https://godoc.org/bitbucket.org/mrd0ll4r/tbotapi)

The implementation is pretty raw, i.e. you will just send and receive messages - you have to handle any command parsing or stuff yourself.

### How do I get set up? ###

A simple

    go get bitbucket.org/mrd0ll4r/tbotapi

should do it.

### Example ###

See `cmd/example.go` for some simple bots.

### API-stableness ###

Is the API stable? **No**

Why is the API not stable? Because Telegram changes its bot API frequently. I try to only add things, not remove them.
So everything you did with this library before should also work after API changes.

### What do we use? ###

We use

* `github.com/go-resty/resty` for REST calls

### Contribution guidelines ###

First of all, if you find any un-idiomatic go code, please let me know! I'm always eager to learn.

If you want to contribute code, just write me on Telegram (see below) or notify me on bitbucket. If everything goes right:

* Writing tests
* Code review
* Write nice Go code

### Things that need to be done ###

* Clean up the model package, especially the outgoing types - we need a better way of doing all that.

### Who do I talk to? ###

* Repo owner or admin (mrd0ll4r)

### License
This work is licensed under the MIT License. A copy of the MIT License can be found in `license.txt`.

Feel free to use this library for any bot whatsoever. If you find any bugs, have any ideas about improvements or just
want to show me what you've done with this, please contact me at [Telegram](https://telegram.me/tbotapibot).