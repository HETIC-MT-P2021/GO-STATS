// Code generated by MockGen. DO NOT EDIT.
// Source: service/leagueoflegends/lolconfig.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "github.com/wyllisMonteiro/GO-STATS/service/config"
)

// MockLeagueOfLegends is a mock of LeagueOfLegends interface
type MockLeagueOfLegends struct {
	ctrl     *gomock.Controller
	recorder *MockLeagueOfLegendsMockRecorder
}

// MockLeagueOfLegendsMockRecorder is the mock recorder for MockLeagueOfLegends
type MockLeagueOfLegendsMockRecorder struct {
	mock *MockLeagueOfLegends
}

// NewMockLeagueOfLegends creates a new mock instance
func NewMockLeagueOfLegends(ctrl *gomock.Controller) *MockLeagueOfLegends {
	mock := &MockLeagueOfLegends{ctrl: ctrl}
	mock.recorder = &MockLeagueOfLegendsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLeagueOfLegends) EXPECT() *MockLeagueOfLegendsMockRecorder {
	return m.recorder
}

// MakeConfig mocks base method
func (m *MockLeagueOfLegends) MakeConfig(arg0 string) config.LeagueOfLegendsAPI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeConfig", arg0)
	ret0, _ := ret[0].(config.LeagueOfLegendsAPI)
	return ret0
}

// MakeConfig indicates an expected call of MakeConfig
func (mr *MockLeagueOfLegendsMockRecorder) MakeConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeConfig", reflect.TypeOf((*MockLeagueOfLegends)(nil).MakeConfig), arg0)
}
