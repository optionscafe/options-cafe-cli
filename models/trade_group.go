//
// Date: 2018-11-22
// Author: Spicer Matthews (spicer@cloudmanic.com)
// Last Modified by: spicer
// Last Modified: 2018-11-22
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import "time"

type TradeGroup struct {
	Id               uint       `json:"id"`
	Name             string     `json:"name"`
	BrokerAccountRef string     `json:"broker_account_ref"`
	Status           string     `json:"status"`
	Type             string     `json:"type"`
	Risked           float64    `json:"risked"`
	Credit           float64    `json:"credit"`
	Proceeds         float64    `json:"proceeds"`
	Profit           float64    `json:"profit"`
	PercentGain      float64    `json:"percent_gain"`
	Commission       float64    `json:"commission"`
	Note             string     `json:"note"`
	Positions        []Position `json:"positions"`
	OpenDate         time.Time  `json:"open_date"`
	ClosedDate       time.Time  `json:"closed_date"`
}

/* End File */
