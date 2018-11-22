//
// Date: 2018-11-22
// Author: Spicer Matthews (spicer@cloudmanic.com)
// Last Modified by: spicer
// Last Modified: 2018-11-22
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/optionscafe/options-cafe-cli/models"
	"github.com/optionscafe/options-cafe-cli/output"
)

//
// List Trade Groups
//
func DoListTradeGroups() {

	// Set output data.
	rows := [][]string{}

	// Make API request
	body, err := MakeGetRequest("/api/v1/tradegroups?broker_account_id=" + os.Getenv("BROKER_ACCOUNT_ID"))

	if err != nil {
		log.Fatal(err)
	}

	// Convert JSON to struct
	tg := []models.TradeGroup{}

	if err := json.Unmarshal([]byte(body), &tg); err != nil {
		panic(err)
	}

	// Loop through the trade groups
	for _, row := range tg {
		rows = append(rows, []string{strconv.Itoa(int(row.Id)), row.Name})
	}

	// Print table and return.
	output.PrintTable(rows, []string{"Id", "Name"})

}

//
// Print Spicer's Blog Trades
//
func DoSpicersBlogTrades() {

	// Make API request
	body, err := MakeGetRequest("/api/v1/tradegroups?broker_account_id=" + os.Getenv("BROKER_ACCOUNT_ID"))

	if err != nil {
		log.Fatal(err)
	}

	// Convert JSON to struct
	tg := []models.TradeGroup{}

	if err := json.Unmarshal([]byte(body), &tg); err != nil {
		panic(err)
	}

	// Loop through the trade groups
	for _, row := range tg {

		// Only care about put credit spreads
		if row.Type != "Put Credit Spread" {
			continue
		}

		// Only care about SPY
		if row.Positions[0].Symbol.OptionUnderlying != "SPY" {
			continue
		}

		// Print positions
		fmt.Println(row.Positions[0].Symbol.OptionUnderlying,
			row.Positions[0].Symbol.OptionType,
			row.Positions[0].Symbol.OptionStrike,
			"/",
			row.Positions[1].Symbol.OptionStrike,
			row.Positions[1].Symbol.OptionExpire.Format("01/02/2006"),
			row.OpenDate.Format("01/02/2006"),
			row.ClosedDate.Format("01/02/2006"),
		)

	}

}

/* End File */
