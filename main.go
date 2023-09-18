package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"utils/queryBlockWithTime/service"
	"utils/queryBlockWithTime/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ethereum/go-ethereum/ethclient"
)

const TIME_LAYOUT = "2006.01.02-15:04:05"

var qs = []*survey.Question{
	{
		Name: "Endpoint",
		Prompt: &survey.Input{
			Message: "Write endpoint",
		},
		Validate: survey.Required,
	},
	{
		Name: "Time",
		Prompt: &survey.Input{
			Message: "Write UTC time \nex) 2022.04.13-19:50:34",
		},
		Validate: survey.Required,
	},
}

func main() {
	var (
		answers struct {
			Endpoint string
			Time     string
		}
	)

	if err := survey.Ask(qs, &answers); err != nil {
		log.Fatal(err.Error())
	}

	target := uint64(util.SeperateFatal(time.Parse(TIME_LAYOUT, answers.Time)).Unix())
	client := util.SeperateFatal(ethclient.Dial(answers.Endpoint))
	curbn := util.SeperateFatal(client.BlockNumber(context.Background()))
	qt := service.NewQueryTool(client)

	resbn, timestamp := qt.BinarySearch(target, 0, curbn)
	fmt.Printf("Block number: %v\nMined at: %v\n", resbn, time.Unix(int64(timestamp), 0).UTC().Format(TIME_LAYOUT))

}
