// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/third_party/proxy (interfaces: Provider)

// Package testing is a generated GoMock package.
package testing

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/api/discovery/v1"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// OnEndpointSliceAdd mocks base method
func (m *MockProvider) OnEndpointSliceAdd(arg0 *v10.EndpointSlice) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEndpointSliceAdd", arg0)
}

// OnEndpointSliceAdd indicates an expected call of OnEndpointSliceAdd
func (mr *MockProviderMockRecorder) OnEndpointSliceAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEndpointSliceAdd", reflect.TypeOf((*MockProvider)(nil).OnEndpointSliceAdd), arg0)
}

// OnEndpointSliceDelete mocks base method
func (m *MockProvider) OnEndpointSliceDelete(arg0 *v10.EndpointSlice) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEndpointSliceDelete", arg0)
}

// OnEndpointSliceDelete indicates an expected call of OnEndpointSliceDelete
func (mr *MockProviderMockRecorder) OnEndpointSliceDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEndpointSliceDelete", reflect.TypeOf((*MockProvider)(nil).OnEndpointSliceDelete), arg0)
}

// OnEndpointSliceUpdate mocks base method
func (m *MockProvider) OnEndpointSliceUpdate(arg0, arg1 *v10.EndpointSlice) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEndpointSliceUpdate", arg0, arg1)
}

// OnEndpointSliceUpdate indicates an expected call of OnEndpointSliceUpdate
func (mr *MockProviderMockRecorder) OnEndpointSliceUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEndpointSliceUpdate", reflect.TypeOf((*MockProvider)(nil).OnEndpointSliceUpdate), arg0, arg1)
}

// OnEndpointSlicesSynced mocks base method
func (m *MockProvider) OnEndpointSlicesSynced() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEndpointSlicesSynced")
}

// OnEndpointSlicesSynced indicates an expected call of OnEndpointSlicesSynced
func (mr *MockProviderMockRecorder) OnEndpointSlicesSynced() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEndpointSlicesSynced", reflect.TypeOf((*MockProvider)(nil).OnEndpointSlicesSynced))
}

// OnEndpointsAdd mocks base method
func (m *MockProvider) OnEndpointsAdd(arg0 *v1.Endpoints) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEndpointsAdd", arg0)
}

// OnEndpointsAdd indicates an expected call of OnEndpointsAdd
func (mr *MockProviderMockRecorder) OnEndpointsAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEndpointsAdd", reflect.TypeOf((*MockProvider)(nil).OnEndpointsAdd), arg0)
}

// OnEndpointsDelete mocks base method
func (m *MockProvider) OnEndpointsDelete(arg0 *v1.Endpoints) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEndpointsDelete", arg0)
}

// OnEndpointsDelete indicates an expected call of OnEndpointsDelete
func (mr *MockProviderMockRecorder) OnEndpointsDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEndpointsDelete", reflect.TypeOf((*MockProvider)(nil).OnEndpointsDelete), arg0)
}

// OnEndpointsSynced mocks base method
func (m *MockProvider) OnEndpointsSynced() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEndpointsSynced")
}

// OnEndpointsSynced indicates an expected call of OnEndpointsSynced
func (mr *MockProviderMockRecorder) OnEndpointsSynced() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEndpointsSynced", reflect.TypeOf((*MockProvider)(nil).OnEndpointsSynced))
}

// OnEndpointsUpdate mocks base method
func (m *MockProvider) OnEndpointsUpdate(arg0, arg1 *v1.Endpoints) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEndpointsUpdate", arg0, arg1)
}

// OnEndpointsUpdate indicates an expected call of OnEndpointsUpdate
func (mr *MockProviderMockRecorder) OnEndpointsUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEndpointsUpdate", reflect.TypeOf((*MockProvider)(nil).OnEndpointsUpdate), arg0, arg1)
}

// OnNodeAdd mocks base method
func (m *MockProvider) OnNodeAdd(arg0 *v1.Node) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNodeAdd", arg0)
}

// OnNodeAdd indicates an expected call of OnNodeAdd
func (mr *MockProviderMockRecorder) OnNodeAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNodeAdd", reflect.TypeOf((*MockProvider)(nil).OnNodeAdd), arg0)
}

// OnNodeDelete mocks base method
func (m *MockProvider) OnNodeDelete(arg0 *v1.Node) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNodeDelete", arg0)
}

// OnNodeDelete indicates an expected call of OnNodeDelete
func (mr *MockProviderMockRecorder) OnNodeDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNodeDelete", reflect.TypeOf((*MockProvider)(nil).OnNodeDelete), arg0)
}

// OnNodeSynced mocks base method
func (m *MockProvider) OnNodeSynced() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNodeSynced")
}

// OnNodeSynced indicates an expected call of OnNodeSynced
func (mr *MockProviderMockRecorder) OnNodeSynced() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNodeSynced", reflect.TypeOf((*MockProvider)(nil).OnNodeSynced))
}

// OnNodeUpdate mocks base method
func (m *MockProvider) OnNodeUpdate(arg0, arg1 *v1.Node) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNodeUpdate", arg0, arg1)
}

// OnNodeUpdate indicates an expected call of OnNodeUpdate
func (mr *MockProviderMockRecorder) OnNodeUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNodeUpdate", reflect.TypeOf((*MockProvider)(nil).OnNodeUpdate), arg0, arg1)
}

// OnServiceAdd mocks base method
func (m *MockProvider) OnServiceAdd(arg0 *v1.Service) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnServiceAdd", arg0)
}

// OnServiceAdd indicates an expected call of OnServiceAdd
func (mr *MockProviderMockRecorder) OnServiceAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnServiceAdd", reflect.TypeOf((*MockProvider)(nil).OnServiceAdd), arg0)
}

// OnServiceDelete mocks base method
func (m *MockProvider) OnServiceDelete(arg0 *v1.Service) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnServiceDelete", arg0)
}

// OnServiceDelete indicates an expected call of OnServiceDelete
func (mr *MockProviderMockRecorder) OnServiceDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnServiceDelete", reflect.TypeOf((*MockProvider)(nil).OnServiceDelete), arg0)
}

// OnServiceSynced mocks base method
func (m *MockProvider) OnServiceSynced() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnServiceSynced")
}

// OnServiceSynced indicates an expected call of OnServiceSynced
func (mr *MockProviderMockRecorder) OnServiceSynced() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnServiceSynced", reflect.TypeOf((*MockProvider)(nil).OnServiceSynced))
}

// OnServiceUpdate mocks base method
func (m *MockProvider) OnServiceUpdate(arg0, arg1 *v1.Service) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnServiceUpdate", arg0, arg1)
}

// OnServiceUpdate indicates an expected call of OnServiceUpdate
func (mr *MockProviderMockRecorder) OnServiceUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnServiceUpdate", reflect.TypeOf((*MockProvider)(nil).OnServiceUpdate), arg0, arg1)
}

// Run mocks base method
func (m *MockProvider) Run(arg0 <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", arg0)
}

// Run indicates an expected call of Run
func (mr *MockProviderMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockProvider)(nil).Run), arg0)
}

// SyncLoop mocks base method
func (m *MockProvider) SyncLoop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SyncLoop")
}

// SyncLoop indicates an expected call of SyncLoop
func (mr *MockProviderMockRecorder) SyncLoop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncLoop", reflect.TypeOf((*MockProvider)(nil).SyncLoop))
}

// SyncedOnce mocks base method
func (m *MockProvider) SyncedOnce() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncedOnce")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SyncedOnce indicates an expected call of SyncedOnce
func (mr *MockProviderMockRecorder) SyncedOnce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncedOnce", reflect.TypeOf((*MockProvider)(nil).SyncedOnce))
}
