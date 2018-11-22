//
// Date: 2018-11-22
// Author: Spicer Matthews (spicer@cloudmanic.com)
// Last Modified by: spicer
// Last Modified: 2018-11-22
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

type Symbol struct {
  Id               uint      `json:"id"`
  ShortName        string    `json:"short_name"`
  Name             string    `json:"name"`
  Type             string    `json:"type"`
  OptionUnderlying string    `json:"option_underlying"`
  OptionType       string    `json:"option_type"`
  OptionExpire     Date      `json:"option_expire"`
  OptionStrike     float64   `json:"option_strike"`
}

/* End File */
