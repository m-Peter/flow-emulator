// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onflow/flow-emulator/emulator (interfaces: Emulator)
//
// Generated by this command:
//
//	mockgen -destination=emulator/mocks/emulator.go -package=mocks github.com/onflow/flow-emulator/emulator Emulator
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	common "github.com/onflow/cadence/common"
	interpreter "github.com/onflow/cadence/interpreter"
	runtime "github.com/onflow/cadence/runtime"
	emulator "github.com/onflow/flow-emulator/emulator"
	types "github.com/onflow/flow-emulator/types"
	access "github.com/onflow/flow-go/model/access"
	flow "github.com/onflow/flow-go/model/flow"
	gomock "go.uber.org/mock/gomock"
)

// MockEmulator is a mock of Emulator interface.
type MockEmulator struct {
	ctrl     *gomock.Controller
	recorder *MockEmulatorMockRecorder
	isgomock struct{}
}

// MockEmulatorMockRecorder is the mock recorder for MockEmulator.
type MockEmulatorMockRecorder struct {
	mock *MockEmulator
}

// NewMockEmulator creates a new mock instance.
func NewMockEmulator(ctrl *gomock.Controller) *MockEmulator {
	mock := &MockEmulator{ctrl: ctrl}
	mock.recorder = &MockEmulatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmulator) EXPECT() *MockEmulatorMockRecorder {
	return m.recorder
}

// AddTransaction mocks base method.
func (m *MockEmulator) AddTransaction(tx flow.TransactionBody) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransaction", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTransaction indicates an expected call of AddTransaction.
func (mr *MockEmulatorMockRecorder) AddTransaction(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransaction", reflect.TypeOf((*MockEmulator)(nil).AddTransaction), tx)
}

// CommitBlock mocks base method.
func (m *MockEmulator) CommitBlock() (*flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitBlock")
	ret0, _ := ret[0].(*flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitBlock indicates an expected call of CommitBlock.
func (mr *MockEmulatorMockRecorder) CommitBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitBlock", reflect.TypeOf((*MockEmulator)(nil).CommitBlock))
}

// ComputationReport mocks base method.
func (m *MockEmulator) ComputationReport() *emulator.ComputationReport {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputationReport")
	ret0, _ := ret[0].(*emulator.ComputationReport)
	return ret0
}

// ComputationReport indicates an expected call of ComputationReport.
func (mr *MockEmulatorMockRecorder) ComputationReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputationReport", reflect.TypeOf((*MockEmulator)(nil).ComputationReport))
}

// CoverageReport mocks base method.
func (m *MockEmulator) CoverageReport() *runtime.CoverageReport {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoverageReport")
	ret0, _ := ret[0].(*runtime.CoverageReport)
	return ret0
}

// CoverageReport indicates an expected call of CoverageReport.
func (mr *MockEmulatorMockRecorder) CoverageReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoverageReport", reflect.TypeOf((*MockEmulator)(nil).CoverageReport))
}

// CreateSnapshot mocks base method.
func (m *MockEmulator) CreateSnapshot(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockEmulatorMockRecorder) CreateSnapshot(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockEmulator)(nil).CreateSnapshot), name)
}

// DisableAutoMine mocks base method.
func (m *MockEmulator) DisableAutoMine() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DisableAutoMine")
}

// DisableAutoMine indicates an expected call of DisableAutoMine.
func (mr *MockEmulatorMockRecorder) DisableAutoMine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableAutoMine", reflect.TypeOf((*MockEmulator)(nil).DisableAutoMine))
}

// EnableAutoMine mocks base method.
func (m *MockEmulator) EnableAutoMine() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnableAutoMine")
}

// EnableAutoMine indicates an expected call of EnableAutoMine.
func (mr *MockEmulatorMockRecorder) EnableAutoMine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAutoMine", reflect.TypeOf((*MockEmulator)(nil).EnableAutoMine))
}

// EndDebugging mocks base method.
func (m *MockEmulator) EndDebugging() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EndDebugging")
}

// EndDebugging indicates an expected call of EndDebugging.
func (mr *MockEmulatorMockRecorder) EndDebugging() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndDebugging", reflect.TypeOf((*MockEmulator)(nil).EndDebugging))
}

// ExecuteAndCommitBlock mocks base method.
func (m *MockEmulator) ExecuteAndCommitBlock() (*flow.Block, []*types.TransactionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteAndCommitBlock")
	ret0, _ := ret[0].(*flow.Block)
	ret1, _ := ret[1].([]*types.TransactionResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExecuteAndCommitBlock indicates an expected call of ExecuteAndCommitBlock.
func (mr *MockEmulatorMockRecorder) ExecuteAndCommitBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAndCommitBlock", reflect.TypeOf((*MockEmulator)(nil).ExecuteAndCommitBlock))
}

// ExecuteBlock mocks base method.
func (m *MockEmulator) ExecuteBlock() ([]*types.TransactionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteBlock")
	ret0, _ := ret[0].([]*types.TransactionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteBlock indicates an expected call of ExecuteBlock.
func (mr *MockEmulatorMockRecorder) ExecuteBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteBlock", reflect.TypeOf((*MockEmulator)(nil).ExecuteBlock))
}

// ExecuteNextTransaction mocks base method.
func (m *MockEmulator) ExecuteNextTransaction() (*types.TransactionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteNextTransaction")
	ret0, _ := ret[0].(*types.TransactionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteNextTransaction indicates an expected call of ExecuteNextTransaction.
func (mr *MockEmulatorMockRecorder) ExecuteNextTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteNextTransaction", reflect.TypeOf((*MockEmulator)(nil).ExecuteNextTransaction))
}

// ExecuteScript mocks base method.
func (m *MockEmulator) ExecuteScript(script []byte, arguments [][]byte) (*types.ScriptResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteScript", script, arguments)
	ret0, _ := ret[0].(*types.ScriptResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteScript indicates an expected call of ExecuteScript.
func (mr *MockEmulatorMockRecorder) ExecuteScript(script, arguments any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteScript", reflect.TypeOf((*MockEmulator)(nil).ExecuteScript), script, arguments)
}

// ExecuteScriptAtBlockHeight mocks base method.
func (m *MockEmulator) ExecuteScriptAtBlockHeight(script []byte, arguments [][]byte, blockHeight uint64) (*types.ScriptResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteScriptAtBlockHeight", script, arguments, blockHeight)
	ret0, _ := ret[0].(*types.ScriptResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteScriptAtBlockHeight indicates an expected call of ExecuteScriptAtBlockHeight.
func (mr *MockEmulatorMockRecorder) ExecuteScriptAtBlockHeight(script, arguments, blockHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteScriptAtBlockHeight", reflect.TypeOf((*MockEmulator)(nil).ExecuteScriptAtBlockHeight), script, arguments, blockHeight)
}

// ExecuteScriptAtBlockID mocks base method.
func (m *MockEmulator) ExecuteScriptAtBlockID(script []byte, arguments [][]byte, id flow.Identifier) (*types.ScriptResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteScriptAtBlockID", script, arguments, id)
	ret0, _ := ret[0].(*types.ScriptResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteScriptAtBlockID indicates an expected call of ExecuteScriptAtBlockID.
func (mr *MockEmulatorMockRecorder) ExecuteScriptAtBlockID(script, arguments, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteScriptAtBlockID", reflect.TypeOf((*MockEmulator)(nil).ExecuteScriptAtBlockID), script, arguments, id)
}

// GetAccount mocks base method.
func (m *MockEmulator) GetAccount(address flow.Address) (*flow.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", address)
	ret0, _ := ret[0].(*flow.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockEmulatorMockRecorder) GetAccount(address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockEmulator)(nil).GetAccount), address)
}

// GetAccountAtBlockHeight mocks base method.
func (m *MockEmulator) GetAccountAtBlockHeight(address flow.Address, blockHeight uint64) (*flow.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountAtBlockHeight", address, blockHeight)
	ret0, _ := ret[0].(*flow.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountAtBlockHeight indicates an expected call of GetAccountAtBlockHeight.
func (mr *MockEmulatorMockRecorder) GetAccountAtBlockHeight(address, blockHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountAtBlockHeight", reflect.TypeOf((*MockEmulator)(nil).GetAccountAtBlockHeight), address, blockHeight)
}

// GetAccountByIndex mocks base method.
func (m *MockEmulator) GetAccountByIndex(arg0 uint) (*flow.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByIndex", arg0)
	ret0, _ := ret[0].(*flow.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByIndex indicates an expected call of GetAccountByIndex.
func (mr *MockEmulatorMockRecorder) GetAccountByIndex(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByIndex", reflect.TypeOf((*MockEmulator)(nil).GetAccountByIndex), arg0)
}

// GetAccountUnsafe mocks base method.
func (m *MockEmulator) GetAccountUnsafe(address flow.Address) (*flow.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountUnsafe", address)
	ret0, _ := ret[0].(*flow.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountUnsafe indicates an expected call of GetAccountUnsafe.
func (mr *MockEmulatorMockRecorder) GetAccountUnsafe(address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountUnsafe", reflect.TypeOf((*MockEmulator)(nil).GetAccountUnsafe), address)
}

// GetBlockByHeight mocks base method.
func (m *MockEmulator) GetBlockByHeight(height uint64) (*flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHeight", height)
	ret0, _ := ret[0].(*flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHeight indicates an expected call of GetBlockByHeight.
func (mr *MockEmulatorMockRecorder) GetBlockByHeight(height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHeight", reflect.TypeOf((*MockEmulator)(nil).GetBlockByHeight), height)
}

// GetBlockByID mocks base method.
func (m *MockEmulator) GetBlockByID(id flow.Identifier) (*flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByID", id)
	ret0, _ := ret[0].(*flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByID indicates an expected call of GetBlockByID.
func (mr *MockEmulatorMockRecorder) GetBlockByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByID", reflect.TypeOf((*MockEmulator)(nil).GetBlockByID), id)
}

// GetCollectionByID mocks base method.
func (m *MockEmulator) GetCollectionByID(colID flow.Identifier) (*flow.LightCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionByID", colID)
	ret0, _ := ret[0].(*flow.LightCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionByID indicates an expected call of GetCollectionByID.
func (mr *MockEmulatorMockRecorder) GetCollectionByID(colID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionByID", reflect.TypeOf((*MockEmulator)(nil).GetCollectionByID), colID)
}

// GetEventsByHeight mocks base method.
func (m *MockEmulator) GetEventsByHeight(blockHeight uint64, eventType string) ([]flow.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsByHeight", blockHeight, eventType)
	ret0, _ := ret[0].([]flow.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsByHeight indicates an expected call of GetEventsByHeight.
func (mr *MockEmulatorMockRecorder) GetEventsByHeight(blockHeight, eventType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsByHeight", reflect.TypeOf((*MockEmulator)(nil).GetEventsByHeight), blockHeight, eventType)
}

// GetEventsForBlockIDs mocks base method.
func (m *MockEmulator) GetEventsForBlockIDs(eventType string, blockIDs []flow.Identifier) ([]flow.BlockEvents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsForBlockIDs", eventType, blockIDs)
	ret0, _ := ret[0].([]flow.BlockEvents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsForBlockIDs indicates an expected call of GetEventsForBlockIDs.
func (mr *MockEmulatorMockRecorder) GetEventsForBlockIDs(eventType, blockIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsForBlockIDs", reflect.TypeOf((*MockEmulator)(nil).GetEventsForBlockIDs), eventType, blockIDs)
}

// GetEventsForHeightRange mocks base method.
func (m *MockEmulator) GetEventsForHeightRange(eventType string, startHeight, endHeight uint64) ([]flow.BlockEvents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsForHeightRange", eventType, startHeight, endHeight)
	ret0, _ := ret[0].([]flow.BlockEvents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsForHeightRange indicates an expected call of GetEventsForHeightRange.
func (mr *MockEmulatorMockRecorder) GetEventsForHeightRange(eventType, startHeight, endHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsForHeightRange", reflect.TypeOf((*MockEmulator)(nil).GetEventsForHeightRange), eventType, startHeight, endHeight)
}

// GetFullCollectionByID mocks base method.
func (m *MockEmulator) GetFullCollectionByID(colID flow.Identifier) (*flow.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullCollectionByID", colID)
	ret0, _ := ret[0].(*flow.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullCollectionByID indicates an expected call of GetFullCollectionByID.
func (mr *MockEmulatorMockRecorder) GetFullCollectionByID(colID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullCollectionByID", reflect.TypeOf((*MockEmulator)(nil).GetFullCollectionByID), colID)
}

// GetLatestBlock mocks base method.
func (m *MockEmulator) GetLatestBlock() (*flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestBlock")
	ret0, _ := ret[0].(*flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestBlock indicates an expected call of GetLatestBlock.
func (mr *MockEmulatorMockRecorder) GetLatestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestBlock", reflect.TypeOf((*MockEmulator)(nil).GetLatestBlock))
}

// GetLogs mocks base method.
func (m *MockEmulator) GetLogs(arg0 flow.Identifier) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockEmulatorMockRecorder) GetLogs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockEmulator)(nil).GetLogs), arg0)
}

// GetNetworkParameters mocks base method.
func (m *MockEmulator) GetNetworkParameters() access.NetworkParameters {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkParameters")
	ret0, _ := ret[0].(access.NetworkParameters)
	return ret0
}

// GetNetworkParameters indicates an expected call of GetNetworkParameters.
func (mr *MockEmulatorMockRecorder) GetNetworkParameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkParameters", reflect.TypeOf((*MockEmulator)(nil).GetNetworkParameters))
}

// GetSourceFile mocks base method.
func (m *MockEmulator) GetSourceFile(location common.Location) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSourceFile", location)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSourceFile indicates an expected call of GetSourceFile.
func (mr *MockEmulatorMockRecorder) GetSourceFile(location any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceFile", reflect.TypeOf((*MockEmulator)(nil).GetSourceFile), location)
}

// GetTransaction mocks base method.
func (m *MockEmulator) GetTransaction(txID flow.Identifier) (*flow.TransactionBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", txID)
	ret0, _ := ret[0].(*flow.TransactionBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockEmulatorMockRecorder) GetTransaction(txID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockEmulator)(nil).GetTransaction), txID)
}

// GetTransactionResult mocks base method.
func (m *MockEmulator) GetTransactionResult(txID flow.Identifier) (*access.TransactionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionResult", txID)
	ret0, _ := ret[0].(*access.TransactionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionResult indicates an expected call of GetTransactionResult.
func (mr *MockEmulatorMockRecorder) GetTransactionResult(txID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionResult", reflect.TypeOf((*MockEmulator)(nil).GetTransactionResult), txID)
}

// GetTransactionResultsByBlockID mocks base method.
func (m *MockEmulator) GetTransactionResultsByBlockID(blockID flow.Identifier) ([]*access.TransactionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionResultsByBlockID", blockID)
	ret0, _ := ret[0].([]*access.TransactionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionResultsByBlockID indicates an expected call of GetTransactionResultsByBlockID.
func (mr *MockEmulatorMockRecorder) GetTransactionResultsByBlockID(blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionResultsByBlockID", reflect.TypeOf((*MockEmulator)(nil).GetTransactionResultsByBlockID), blockID)
}

// GetTransactionsByBlockID mocks base method.
func (m *MockEmulator) GetTransactionsByBlockID(blockID flow.Identifier) ([]*flow.TransactionBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsByBlockID", blockID)
	ret0, _ := ret[0].([]*flow.TransactionBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsByBlockID indicates an expected call of GetTransactionsByBlockID.
func (mr *MockEmulatorMockRecorder) GetTransactionsByBlockID(blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsByBlockID", reflect.TypeOf((*MockEmulator)(nil).GetTransactionsByBlockID), blockID)
}

// LoadSnapshot mocks base method.
func (m *MockEmulator) LoadSnapshot(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadSnapshot", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadSnapshot indicates an expected call of LoadSnapshot.
func (mr *MockEmulatorMockRecorder) LoadSnapshot(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadSnapshot", reflect.TypeOf((*MockEmulator)(nil).LoadSnapshot), name)
}

// Ping mocks base method.
func (m *MockEmulator) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockEmulatorMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockEmulator)(nil).Ping))
}

// ResetCoverageReport mocks base method.
func (m *MockEmulator) ResetCoverageReport() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetCoverageReport")
}

// ResetCoverageReport indicates an expected call of ResetCoverageReport.
func (mr *MockEmulatorMockRecorder) ResetCoverageReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetCoverageReport", reflect.TypeOf((*MockEmulator)(nil).ResetCoverageReport))
}

// RollbackToBlockHeight mocks base method.
func (m *MockEmulator) RollbackToBlockHeight(height uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackToBlockHeight", height)
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackToBlockHeight indicates an expected call of RollbackToBlockHeight.
func (mr *MockEmulatorMockRecorder) RollbackToBlockHeight(height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackToBlockHeight", reflect.TypeOf((*MockEmulator)(nil).RollbackToBlockHeight), height)
}

// SendTransaction mocks base method.
func (m *MockEmulator) SendTransaction(tx *flow.TransactionBody) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockEmulatorMockRecorder) SendTransaction(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockEmulator)(nil).SendTransaction), tx)
}

// ServiceKey mocks base method.
func (m *MockEmulator) ServiceKey() emulator.ServiceKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceKey")
	ret0, _ := ret[0].(emulator.ServiceKey)
	return ret0
}

// ServiceKey indicates an expected call of ServiceKey.
func (mr *MockEmulatorMockRecorder) ServiceKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceKey", reflect.TypeOf((*MockEmulator)(nil).ServiceKey))
}

// Snapshots mocks base method.
func (m *MockEmulator) Snapshots() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshots indicates an expected call of Snapshots.
func (mr *MockEmulatorMockRecorder) Snapshots() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockEmulator)(nil).Snapshots))
}

// StartDebugger mocks base method.
func (m *MockEmulator) StartDebugger() *interpreter.Debugger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDebugger")
	ret0, _ := ret[0].(*interpreter.Debugger)
	return ret0
}

// StartDebugger indicates an expected call of StartDebugger.
func (mr *MockEmulatorMockRecorder) StartDebugger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDebugger", reflect.TypeOf((*MockEmulator)(nil).StartDebugger))
}
