package cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	peernodeaddressFlag = &cli.StringFlag{
		Name:	"peernodeaddress",
		Usage:	"partner node address, either addr:port or just port",
		Required: true,
	}
	nodeaddressFlag = &cli.StringFlag{
		Name:	"nodeaddress",
		Usage:	"what address is this node reachable under, addr:port or port (localhost)",
		Required: true,
	}
	apiportFlag = &cli.IntFlag{
		Name:	"apiport",
		Usage:	"apiport",
		Value:	8080,
	}
	minerFlag = &cli.StringFlag{
		Name:	"miner",
		Usage:  "miner address",
	}
	walletnameFlag = &cli.StringFlag{
		Name:	"walletname",
		Usage:  "wallet name",
		Required: true,
	}
	rawFlag = &cli.StringFlag{
		Name:	"raw",
		Usage: 	"need send raw",
	}
	isminingFlag = &cli.BoolFlag{
		Name:	"ismining",
		Usage:  "ismining",
		Value: false,
	}
	toFlag = &cli.StringFlag{
		Name:	"to",
		Usage:	"to address",
		Required: true,

	}
	amountFlag = &cli.IntFlag{
		Name:	"amount",
		Usage:	"send amount",
		Required: true,
	}

)