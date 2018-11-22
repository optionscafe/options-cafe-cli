//
// Date: 2018-11-22
// Author: Spicer Matthews (spicer@cloudmanic.com)
// Last Modified by: spicer
// Last Modified: 2018-11-22
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package main

import (
	"log"
	"os"
	"os/user"

	"github.com/joho/godotenv"
	"github.com/optionscafe/options-cafe-cli/actions"
	"github.com/optionscafe/options-cafe-cli/output"
)

//
// Main...
//
func main() {

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

	// Make sure we have at least one arg
	if len(os.Args) <= 1 {
		output.PrintHelp()
		return
	}

	// Switch based on the first argument
	switch os.Args[1] {

	// List Trade Groups
	case "trade-group-list":
		actions.DoListTradeGroups()

	// Spicer's Blog Trades
	case "spicers-blog-trades":
		actions.DoSpicersBlogTrades()

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
