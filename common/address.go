// Copyright (c) 2018 The transformer developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>


package common
//This package defines some fundamental data structures, some of which are derived from the Ethereum existing data structures.

import (
	"github.com/ethereum/go-ethereum/common"
)
const (
	// AddressLength length of address in bytes.
	AddressLength = common.AddressLength
)

// Address address of account.
type Address common.Address