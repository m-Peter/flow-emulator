/*
 * Flow Emulator
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package emulator

import (
	"github.com/onflow/flow-emulator/types"
)

type ScriptReport struct {
	ID              string   `json:"ID"`
	ComputationUsed uint64   `json:"computation"`
	MemoryEstimate  uint64   `json:"memory"`
	Code            string   `json:"source"`
	Arguments       []string `json:"arguments"`
}

type TransactionReport struct {
	ID              string   `json:"ID"`
	ComputationUsed uint64   `json:"computation"`
	MemoryEstimate  uint64   `json:"memory"`
	Code            string   `json:"source"`
	Arguments       []string `json:"arguments"`
}

type ProfilingReport struct {
	Scripts      []ScriptReport      `json:"scripts"`
	Transactions []TransactionReport `json:"transactions"`
}

func (pr *ProfilingReport) AddScript(
	scriptResult types.ScriptResult,
	code string,
	arguments []string,
) {
	scriptReport := ScriptReport{
		ID:              scriptResult.ScriptID.String(),
		ComputationUsed: scriptResult.ComputationUsed,
		MemoryEstimate:  scriptResult.MemoryEstimate,
		Code:            code,
		Arguments:       arguments,
	}
	pr.Scripts = append(pr.Scripts, scriptReport)
}

func (pr *ProfilingReport) AddTransaction(
	txResult types.TransactionResult,
	code string,
	arguments []string,
) {
	txReport := TransactionReport{
		ID:              txResult.TransactionID.String(),
		ComputationUsed: txResult.ComputationUsed,
		MemoryEstimate:  txResult.MemoryEstimate,
		Code:            code,
		Arguments:       arguments,
	}
	pr.Transactions = append(pr.Transactions, txReport)
}
