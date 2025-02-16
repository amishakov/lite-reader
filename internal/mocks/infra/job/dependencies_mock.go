// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubny/lite-reader/internal/infra/job (interfaces: FeedService,ItemService)
//
// Generated by this command:
//
//	mockgen -destination=./infra/job/dependencies_mock.go -package=mocks -mock_names=FeedService=FeedService,ItemService=ItemService github.com/cubny/lite-reader/internal/infra/job FeedService,ItemService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	feed "github.com/cubny/lite-reader/internal/app/feed"
	item "github.com/cubny/lite-reader/internal/app/item"
)

// FeedService is a mock of FeedService interface.
type FeedService struct {
	ctrl     *gomock.Controller
	recorder *FeedServiceMockRecorder
}

// FeedServiceMockRecorder is the mock recorder for FeedService.
type FeedServiceMockRecorder struct {
	mock *FeedService
}

// NewFeedService creates a new mock instance.
func NewFeedService(ctrl *gomock.Controller) *FeedService {
	mock := &FeedService{ctrl: ctrl}
	mock.recorder = &FeedServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FeedService) EXPECT() *FeedServiceMockRecorder {
	return m.recorder
}

// FetchItems mocks base method.
func (m *FeedService) FetchItems(arg0 int) ([]*item.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchItems", arg0)
	ret0, _ := ret[0].([]*item.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchItems indicates an expected call of FetchItems.
func (mr *FeedServiceMockRecorder) FetchItems(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchItems", reflect.TypeOf((*FeedService)(nil).FetchItems), arg0)
}

// ListFeeds mocks base method.
func (m *FeedService) ListFeeds() ([]*feed.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeeds")
	ret0, _ := ret[0].([]*feed.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFeeds indicates an expected call of ListFeeds.
func (mr *FeedServiceMockRecorder) ListFeeds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeeds", reflect.TypeOf((*FeedService)(nil).ListFeeds))
}

// ItemService is a mock of ItemService interface.
type ItemService struct {
	ctrl     *gomock.Controller
	recorder *ItemServiceMockRecorder
}

// ItemServiceMockRecorder is the mock recorder for ItemService.
type ItemServiceMockRecorder struct {
	mock *ItemService
}

// NewItemService creates a new mock instance.
func NewItemService(ctrl *gomock.Controller) *ItemService {
	mock := &ItemService{ctrl: ctrl}
	mock.recorder = &ItemServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *ItemService) EXPECT() *ItemServiceMockRecorder {
	return m.recorder
}

// UpsertItems mocks base method.
func (m *ItemService) UpsertItems(arg0 *item.UpsertItemsCommand) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertItems", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertItems indicates an expected call of UpsertItems.
func (mr *ItemServiceMockRecorder) UpsertItems(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertItems", reflect.TypeOf((*ItemService)(nil).UpsertItems), arg0)
}
