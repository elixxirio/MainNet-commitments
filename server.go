///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

// +build linux darwin
// +build amd64

package main

import "git.xx.network/elixxir/mainnet-commitments/cmd"

func main() {
	cmd.ExecuteServer()
}
