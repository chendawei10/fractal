// Copyright 2018 The Fractal Team Authors
// This file is part of the fractal project.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package dpos

import (
	"math"
	"math/big"
	"strings"
)

var (
	// LastBlockHeight latest
	LastBlockHeight = uint64(math.MaxUint64)
)

// IDB dpos database
type IDB interface {
	SetProducer(*producerInfo) error
	DelProducer(string) error
	GetProducer(string) (*producerInfo, error)
	Producers() ([]*producerInfo, error)
	ProducersSize() (uint64, error)

	SetVoter(*voterInfo) error
	DelVoter(string, string) error
	GetVoter(string) (*voterInfo, error)
	// GetDelegators(string) ([]string, error)

	SetState(*globalState) error
	DelState(uint64) error
	GetState(uint64) (*globalState, error)

	Delegate(string, *big.Int) error
	Undelegate(string, *big.Int) error
	IncAsset2Acct(string, string, *big.Int) error
}

type producerInfo struct {
	Name          string   // producer name
	URL           string   // producer url
	Quantity      *big.Int // producer stake quantity
	TotalQuantity *big.Int // producer total stake quantity
	Height        uint64   // timestamp
}

type voterInfo struct {
	Name     string   // voter name
	Producer string   // producer approved by this voter
	Quantity *big.Int // stake approved by this voter
	Height   uint64   // timestamp
}

type globalState struct {
	Height                          uint64   // block height
	ActivatedProducerScheduleUpdate uint64   // update time
	ActivatedProducerSchedule       []string // producers
	ActivatedTotalQuantity          *big.Int // the sum of activate producer votes
	TotalQuantity                   *big.Int // the sum of all producer votes
}

type producerInfoArray []*producerInfo

func (prods producerInfoArray) Len() int {
	return len(prods)
}
func (prods producerInfoArray) Less(i, j int) bool {
	val := prods[i].TotalQuantity.Cmp(prods[j].TotalQuantity)
	if val == 0 {
		if prods[i].Height == prods[j].Height {
			return strings.Compare(prods[i].Name, prods[j].Name) > 0
		}
		return prods[i].Height < prods[j].Height
	}
	return val > 0
}
func (prods producerInfoArray) Swap(i, j int) {
	prods[i], prods[j] = prods[j], prods[i]
}
