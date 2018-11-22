//
// Date: 2018-11-22
// Author: Spicer Matthews (spicer@cloudmanic.com)
// Last Modified by: spicer
// Last Modified: 2018-11-22
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import(
  "time"
)

type Position struct {
  Id               uint      `json:"id"`
  TradeGroupId     uint      `json:"trade_group_id"`
  BrokerAccountRef string    `json:"broker_account_ref"`
  Status           string    `json:"status"`
  Qty              int       `json:"qty"`
  OrgQty           int       `json:"org_qty"`
  CostBasis        float64   `json:"cost_basis"`
  Proceeds         float64   `json:"proceeds"`
  Profit           float64   `json:"profit"`
  AvgOpenPrice     float64   `json:"avg_open_price"`
  AvgClosePrice    float64   `json:"avg_close_price"`
  Note             string    `json:"note"`
  OpenDate         time.Time `json:"open_date"`
  ClosedDate       time.Time `json:"close_date"`
  Symbol           Symbol    `json:"symbol"`
}

/* End File */
