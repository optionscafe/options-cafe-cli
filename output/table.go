//
// Date: 2018-11-22
// Author: Spicer Matthews (spicer@cloudmanic.com)
// Last Modified by: spicer
// Last Modified: 2018-11-22
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package output

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

//
// Print table.
//
func PrintTable(rows [][]string, headers []string) {

	// Print table to screen.
	table := tablewriter.NewWriter(os.Stdout)

	// Build table headers
	table.SetHeader(headers)

	// Build table rows
	for _, v := range rows {
		table.Append(v)
	}

	// Send output
	table.Render()
}

/* End File */
