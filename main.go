//
// Date: 2018-11-22
// Author: Spicer Matthews (spicer@cloudmanic.com)
// Last Modified by: Spicer Matthews
// Last Modified: 2018-11-23
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package main

import (
	"flag"
	"log"
	"os/user"

	"github.com/joho/godotenv"
	"github.com/optionscafe/options-cafe-cli/actions"
	"github.com/optionscafe/options-cafe-cli/output"
)

//
// Main...
//
func main() {

	// Grab flags
	end := flag.String("end", "", "--end=xx/xx/xxxx")
	start := flag.String("start", "", "--start=xx/xx/xxxx")
	sort := flag.String("sort", "asc", "--sort={asc | desc}")
	order := flag.String("order", "id", "--order={field_name}")
	action := flag.String("action", "", "--action={action name}")
	flag.Parse()

	// Get the current user.
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	// Load .env file
	err = godotenv.Load(usr.HomeDir + "/.options-cafe")

	if err != nil {
		log.Fatal(err)
		//DoAuth()
		return
	}

	// Switch based on the first argument
	switch *action {

	// List Trade Groups
	case "trade-group-list":
		actions.DoListTradeGroups()

	// Spicer's Blog Trades
	case "spicers-blog-trades":
		actions.DoSpicersBlogTrades(*start, *end, *order, *sort)

	// // Do Auth
	// case "auth":
	// 	DoAuth()

	// Print Help
	case "help":
		output.PrintHelp()

	// Print Help
	default:
		output.PrintHelp()

	}

}

/* End File */
