# Raspberry Pi Examples

## Installation

Examples rely on [Gobot](https://gobot.io/). You can install necessary requirements on your Raspi by using:

```bash
go get -d -u gobot.io/x/gobot \ 
             gobot.io/x/drivers/gpio \
             gobot.io/x/gobot/platforms/raspi
```

Now you should be able to build examples using regular `go build` command.