// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/aggregator/aggregator (interfaces: ElectionManager,FlushTimesManager,PlacementManager)

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package aggregator is a generated GoMock package.
package aggregator

import (
	"context"
	"reflect"

	"github.com/m3db/m3/src/aggregator/generated/proto/flush"
	"github.com/m3db/m3/src/cluster/placement"
	"github.com/m3db/m3/src/cluster/shard"
	"github.com/m3db/m3/src/x/watch"

	"github.com/golang/mock/gomock"
)

// MockElectionManager is a mock of ElectionManager interface
type MockElectionManager struct {
	ctrl     *gomock.Controller
	recorder *MockElectionManagerMockRecorder
}

// MockElectionManagerMockRecorder is the mock recorder for MockElectionManager
type MockElectionManagerMockRecorder struct {
	mock *MockElectionManager
}

// NewMockElectionManager creates a new mock instance
func NewMockElectionManager(ctrl *gomock.Controller) *MockElectionManager {
	mock := &MockElectionManager{ctrl: ctrl}
	mock.recorder = &MockElectionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockElectionManager) EXPECT() *MockElectionManagerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockElectionManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockElectionManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockElectionManager)(nil).Close))
}

// ElectionState mocks base method
func (m *MockElectionManager) ElectionState() ElectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ElectionState")
	ret0, _ := ret[0].(ElectionState)
	return ret0
}

// ElectionState indicates an expected call of ElectionState
func (mr *MockElectionManagerMockRecorder) ElectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ElectionState", reflect.TypeOf((*MockElectionManager)(nil).ElectionState))
}

// IsCampaigning mocks base method
func (m *MockElectionManager) IsCampaigning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCampaigning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCampaigning indicates an expected call of IsCampaigning
func (mr *MockElectionManagerMockRecorder) IsCampaigning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCampaigning", reflect.TypeOf((*MockElectionManager)(nil).IsCampaigning))
}

// Open mocks base method
func (m *MockElectionManager) Open(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockElectionManagerMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockElectionManager)(nil).Open), arg0)
}

// Reset mocks base method
func (m *MockElectionManager) Reset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset
func (mr *MockElectionManagerMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockElectionManager)(nil).Reset))
}

// Resign mocks base method
func (m *MockElectionManager) Resign(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resign", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resign indicates an expected call of Resign
func (mr *MockElectionManagerMockRecorder) Resign(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resign", reflect.TypeOf((*MockElectionManager)(nil).Resign), arg0)
}

// MockFlushTimesManager is a mock of FlushTimesManager interface
type MockFlushTimesManager struct {
	ctrl     *gomock.Controller
	recorder *MockFlushTimesManagerMockRecorder
}

// MockFlushTimesManagerMockRecorder is the mock recorder for MockFlushTimesManager
type MockFlushTimesManagerMockRecorder struct {
	mock *MockFlushTimesManager
}

// NewMockFlushTimesManager creates a new mock instance
func NewMockFlushTimesManager(ctrl *gomock.Controller) *MockFlushTimesManager {
	mock := &MockFlushTimesManager{ctrl: ctrl}
	mock.recorder = &MockFlushTimesManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFlushTimesManager) EXPECT() *MockFlushTimesManagerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockFlushTimesManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockFlushTimesManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFlushTimesManager)(nil).Close))
}

// Get mocks base method
func (m *MockFlushTimesManager) Get() (*flush.ShardSetFlushTimes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*flush.ShardSetFlushTimes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockFlushTimesManagerMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFlushTimesManager)(nil).Get))
}

// Open mocks base method
func (m *MockFlushTimesManager) Open(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockFlushTimesManagerMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockFlushTimesManager)(nil).Open), arg0)
}

// Reset mocks base method
func (m *MockFlushTimesManager) Reset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset
func (mr *MockFlushTimesManagerMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockFlushTimesManager)(nil).Reset))
}

// StoreAsync mocks base method
func (m *MockFlushTimesManager) StoreAsync(arg0 *flush.ShardSetFlushTimes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreAsync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreAsync indicates an expected call of StoreAsync
func (mr *MockFlushTimesManagerMockRecorder) StoreAsync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAsync", reflect.TypeOf((*MockFlushTimesManager)(nil).StoreAsync), arg0)
}

// Watch mocks base method
func (m *MockFlushTimesManager) Watch() (watch.Watch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(watch.Watch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockFlushTimesManagerMockRecorder) Watch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockFlushTimesManager)(nil).Watch))
}

// MockPlacementManager is a mock of PlacementManager interface
type MockPlacementManager struct {
	ctrl     *gomock.Controller
	recorder *MockPlacementManagerMockRecorder
}

// MockPlacementManagerMockRecorder is the mock recorder for MockPlacementManager
type MockPlacementManagerMockRecorder struct {
	mock *MockPlacementManager
}

// NewMockPlacementManager creates a new mock instance
func NewMockPlacementManager(ctrl *gomock.Controller) *MockPlacementManager {
	mock := &MockPlacementManager{ctrl: ctrl}
	mock.recorder = &MockPlacementManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlacementManager) EXPECT() *MockPlacementManagerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockPlacementManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockPlacementManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPlacementManager)(nil).Close))
}

// HasReplacementInstance mocks base method
func (m *MockPlacementManager) HasReplacementInstance() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasReplacementInstance")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasReplacementInstance indicates an expected call of HasReplacementInstance
func (mr *MockPlacementManagerMockRecorder) HasReplacementInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasReplacementInstance", reflect.TypeOf((*MockPlacementManager)(nil).HasReplacementInstance))
}

// Instance mocks base method
func (m *MockPlacementManager) Instance() (placement.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Instance")
	ret0, _ := ret[0].(placement.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Instance indicates an expected call of Instance
func (mr *MockPlacementManagerMockRecorder) Instance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Instance", reflect.TypeOf((*MockPlacementManager)(nil).Instance))
}

// InstanceFrom mocks base method
func (m *MockPlacementManager) InstanceFrom(arg0 placement.Placement) (placement.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceFrom", arg0)
	ret0, _ := ret[0].(placement.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceFrom indicates an expected call of InstanceFrom
func (mr *MockPlacementManagerMockRecorder) InstanceFrom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceFrom", reflect.TypeOf((*MockPlacementManager)(nil).InstanceFrom), arg0)
}

// InstanceID mocks base method
func (m *MockPlacementManager) InstanceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// InstanceID indicates an expected call of InstanceID
func (mr *MockPlacementManagerMockRecorder) InstanceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceID", reflect.TypeOf((*MockPlacementManager)(nil).InstanceID))
}

// Open mocks base method
func (m *MockPlacementManager) Open() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockPlacementManagerMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockPlacementManager)(nil).Open))
}

// Placement mocks base method
func (m *MockPlacementManager) Placement() (placement.ActiveStagedPlacement, placement.Placement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Placement")
	ret0, _ := ret[0].(placement.ActiveStagedPlacement)
	ret1, _ := ret[1].(placement.Placement)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Placement indicates an expected call of Placement
func (mr *MockPlacementManagerMockRecorder) Placement() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Placement", reflect.TypeOf((*MockPlacementManager)(nil).Placement))
}

// Shards mocks base method
func (m *MockPlacementManager) Shards() (shard.Shards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shards")
	ret0, _ := ret[0].(shard.Shards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shards indicates an expected call of Shards
func (mr *MockPlacementManagerMockRecorder) Shards() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shards", reflect.TypeOf((*MockPlacementManager)(nil).Shards))
}