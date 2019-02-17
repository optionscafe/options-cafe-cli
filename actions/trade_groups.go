//
// Date: 2018-11-22
// Author: Spicer Matthews (spicer@cloudmanic.com)
// Last Modified by: Spicer Matthews
// Last Modified: 2018-11-23
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/optionscafe/options-cafe-cli/helpers"
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
// ./app --action=spicers-blog-trades --start="11/01/2018" --end="11/30/2018" --order=closed_date --sort=asc --type="Put Credit Spread"
//
func DoSpicersBlogTrades(start string, end string, order string, sort string, optionsType string) {

	uri := "/api/v1/tradegroups?broker_account_id=" + os.Getenv("BROKER_ACCOUNT_ID") + "&order=" + order + "&sort=" + sort

	if len(optionsType) > 0 {
		uri += "&type=" + optionsType
	}

	// Make API request
	body, err := MakeGetRequest(uri)

	if err != nil {
		log.Fatal(err)
	}

	// Convert JSON to struct
	tg := []models.TradeGroup{}

	if err := json.Unmarshal([]byte(body), &tg); err != nil {
		panic(err)
	}

	// Print header of table
	fmt.Println("")
	fmt.Println("")
	fmt.Println(`<div class="table"><table><tr><th width="59"><span>Ticker</span></th><th width="50"><span>Type</span></th><th width="109"><span>Spread</span></th><th width="85"><span>Expire Date</span></th><th width="79"><span>Open Date</span></th><th width="90"><span>Closed Date</span></th><th width="87"><span>Open Credit</span></th><th width="94"><span>Closed Debit</span></th><th width="78"><span>Profit / Loss</span></th></tr>`)

	// Loop through the trade groups
	for _, row := range tg {

		// Only care about SPY
		if row.Positions[0].Symbol.OptionUnderlying != "SPY" {
			continue
		}

		// Only closed trades
		if row.Status != "Closed" {
			continue
		}

		// Filter out start date
		if (len(start) > 0) && (row.ClosedDate.Unix() < helpers.ParseDateNoError(start).Unix()) {
			continue
		}

		// Filter out end date
		if (len(end) > 0) && (row.ClosedDate.Unix() > helpers.ParseDateNoError(end).Unix()) {
			continue
		}

		// Percent gain
		y1 := row.Risked
		y2 := (row.Risked + row.Credit) - (row.Proceeds * -1.00)
		percent := ((y2 - y1) / y1) * 100

		// open credit amount
		openCredit := (row.Credit / float64(row.Positions[0].OrgQty)) / 100

		// Closed amount
		closedAmount := ((row.Proceeds * -1.00) / float64(row.Positions[0].OrgQty)) / 100

		// Figure out strong tag (make red if negative)
		strong := `<strong class="red-text">`

		if percent > 0 {
			strong = `<strong>`
		}

		// Print positions
		fmt.Println("<tr><td>"+row.Positions[0].Symbol.OptionUnderlying+"</td>",
			"<td>"+row.Positions[0].Symbol.OptionType+"</td>",
			"<td>"+fmt.Sprint(row.Positions[0].Symbol.OptionStrike)+"/"+fmt.Sprint(row.Positions[1].Symbol.OptionStrike)+"</td>",
			"<td>"+row.Positions[1].Symbol.OptionExpire.Format("01/02/2006")+"</td>",
			"<td>"+row.OpenDate.Format("01/02/06")+"</td>",
			"<td>"+row.ClosedDate.Format("01/02/06")+"</td>",
			"<td>$"+fmt.Sprint(openCredit)+"</td>",
			"<td>$"+fmt.Sprint(closedAmount)+"</td>",
			"<td>"+strong+fmt.Sprint(helpers.Round(percent, 2))+"%</strong></td></tr>",
		)
	}

	// Close table
	fmt.Println(`</table></div>`)
	fmt.Println("")
	fmt.Println("")

}

/* End File */
