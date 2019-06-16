// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mradile/rssfeeder (interfaces: FeedStorage,FeedEntryStorage,AddingService,DeletingService,ViewingService)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	rssfeeder "github.com/mradile/rssfeeder"
	reflect "reflect"
)

// MockFeedStorage is a mock of FeedStorage interface
type MockFeedStorage struct {
	ctrl     *gomock.Controller
	recorder *MockFeedStorageMockRecorder
}

// MockFeedStorageMockRecorder is the mock recorder for MockFeedStorage
type MockFeedStorageMockRecorder struct {
	mock *MockFeedStorage
}

// NewMockFeedStorage creates a new mock instance
func NewMockFeedStorage(ctrl *gomock.Controller) *MockFeedStorage {
	mock := &MockFeedStorage{ctrl: ctrl}
	mock.recorder = &MockFeedStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFeedStorage) EXPECT() *MockFeedStorageMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockFeedStorage) Add(arg0 *rssfeeder.Feed) error {
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockFeedStorageMockRecorder) Add(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockFeedStorage)(nil).Add), arg0)
}

// Delete mocks base method
func (m *MockFeedStorage) Delete(arg0 int) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockFeedStorageMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFeedStorage)(nil).Delete), arg0)
}

// Exists mocks base method
func (m *MockFeedStorage) Exists(arg0, arg1 string) (bool, error) {
	ret := m.ctrl.Call(m, "Exists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockFeedStorageMockRecorder) Exists(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockFeedStorage)(nil).Exists), arg0, arg1)
}

// GetByNameAndLogin mocks base method
func (m *MockFeedStorage) GetByNameAndLogin(arg0, arg1 string) (*rssfeeder.Feed, error) {
	ret := m.ctrl.Call(m, "GetByNameAndLogin", arg0, arg1)
	ret0, _ := ret[0].(*rssfeeder.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNameAndLogin indicates an expected call of GetByNameAndLogin
func (mr *MockFeedStorageMockRecorder) GetByNameAndLogin(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNameAndLogin", reflect.TypeOf((*MockFeedStorage)(nil).GetByNameAndLogin), arg0, arg1)
}

// GetFeedsByLogin mocks base method
func (m *MockFeedStorage) GetFeedsByLogin(arg0 string) ([]*rssfeeder.Feed, error) {
	ret := m.ctrl.Call(m, "GetFeedsByLogin", arg0)
	ret0, _ := ret[0].([]*rssfeeder.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedsByLogin indicates an expected call of GetFeedsByLogin
func (mr *MockFeedStorageMockRecorder) GetFeedsByLogin(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedsByLogin", reflect.TypeOf((*MockFeedStorage)(nil).GetFeedsByLogin), arg0)
}

// MockFeedEntryStorage is a mock of FeedEntryStorage interface
type MockFeedEntryStorage struct {
	ctrl     *gomock.Controller
	recorder *MockFeedEntryStorageMockRecorder
}

// MockFeedEntryStorageMockRecorder is the mock recorder for MockFeedEntryStorage
type MockFeedEntryStorageMockRecorder struct {
	mock *MockFeedEntryStorage
}

// NewMockFeedEntryStorage creates a new mock instance
func NewMockFeedEntryStorage(ctrl *gomock.Controller) *MockFeedEntryStorage {
	mock := &MockFeedEntryStorage{ctrl: ctrl}
	mock.recorder = &MockFeedEntryStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFeedEntryStorage) EXPECT() *MockFeedEntryStorageMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockFeedEntryStorage) Add(arg0 *rssfeeder.FeedEntry) error {
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockFeedEntryStorageMockRecorder) Add(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockFeedEntryStorage)(nil).Add), arg0)
}

// AllByLoginAndFeedName mocks base method
func (m *MockFeedEntryStorage) AllByLoginAndFeedName(arg0, arg1 string) ([]*rssfeeder.FeedEntry, error) {
	ret := m.ctrl.Call(m, "AllByLoginAndFeedName", arg0, arg1)
	ret0, _ := ret[0].([]*rssfeeder.FeedEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllByLoginAndFeedName indicates an expected call of AllByLoginAndFeedName
func (mr *MockFeedEntryStorageMockRecorder) AllByLoginAndFeedName(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllByLoginAndFeedName", reflect.TypeOf((*MockFeedEntryStorage)(nil).AllByLoginAndFeedName), arg0, arg1)
}

// Delete mocks base method
func (m *MockFeedEntryStorage) Delete(arg0 int) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockFeedEntryStorageMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFeedEntryStorage)(nil).Delete), arg0)
}

// EntryBelongsToLogin mocks base method
func (m *MockFeedEntryStorage) EntryBelongsToLogin(arg0 int, arg1 string) (bool, error) {
	ret := m.ctrl.Call(m, "EntryBelongsToLogin", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EntryBelongsToLogin indicates an expected call of EntryBelongsToLogin
func (mr *MockFeedEntryStorageMockRecorder) EntryBelongsToLogin(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntryBelongsToLogin", reflect.TypeOf((*MockFeedEntryStorage)(nil).EntryBelongsToLogin), arg0, arg1)
}

// ExistsEntry mocks base method
func (m *MockFeedEntryStorage) ExistsEntry(arg0 int) (bool, error) {
	ret := m.ctrl.Call(m, "ExistsEntry", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsEntry indicates an expected call of ExistsEntry
func (mr *MockFeedEntryStorageMockRecorder) ExistsEntry(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsEntry", reflect.TypeOf((*MockFeedEntryStorage)(nil).ExistsEntry), arg0)
}

// GetCategories mocks base method
func (m *MockFeedEntryStorage) GetCategories(arg0 string) ([]string, error) {
	ret := m.ctrl.Call(m, "GetCategories", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories
func (mr *MockFeedEntryStorageMockRecorder) GetCategories(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockFeedEntryStorage)(nil).GetCategories), arg0)
}

// MockAddingService is a mock of AddingService interface
type MockAddingService struct {
	ctrl     *gomock.Controller
	recorder *MockAddingServiceMockRecorder
}

// MockAddingServiceMockRecorder is the mock recorder for MockAddingService
type MockAddingServiceMockRecorder struct {
	mock *MockAddingService
}

// NewMockAddingService creates a new mock instance
func NewMockAddingService(ctrl *gomock.Controller) *MockAddingService {
	mock := &MockAddingService{ctrl: ctrl}
	mock.recorder = &MockAddingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAddingService) EXPECT() *MockAddingServiceMockRecorder {
	return m.recorder
}

// AddFeedEntry mocks base method
func (m *MockAddingService) AddFeedEntry(arg0 *rssfeeder.FeedEntry) error {
	ret := m.ctrl.Call(m, "AddFeedEntry", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFeedEntry indicates an expected call of AddFeedEntry
func (mr *MockAddingServiceMockRecorder) AddFeedEntry(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFeedEntry", reflect.TypeOf((*MockAddingService)(nil).AddFeedEntry), arg0)
}

// MockDeletingService is a mock of DeletingService interface
type MockDeletingService struct {
	ctrl     *gomock.Controller
	recorder *MockDeletingServiceMockRecorder
}

// MockDeletingServiceMockRecorder is the mock recorder for MockDeletingService
type MockDeletingServiceMockRecorder struct {
	mock *MockDeletingService
}

// NewMockDeletingService creates a new mock instance
func NewMockDeletingService(ctrl *gomock.Controller) *MockDeletingService {
	mock := &MockDeletingService{ctrl: ctrl}
	mock.recorder = &MockDeletingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeletingService) EXPECT() *MockDeletingServiceMockRecorder {
	return m.recorder
}

// DeleteFeedEntry mocks base method
func (m *MockDeletingService) DeleteFeedEntry(arg0 int, arg1 string) error {
	ret := m.ctrl.Call(m, "DeleteFeedEntry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFeedEntry indicates an expected call of DeleteFeedEntry
func (mr *MockDeletingServiceMockRecorder) DeleteFeedEntry(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFeedEntry", reflect.TypeOf((*MockDeletingService)(nil).DeleteFeedEntry), arg0, arg1)
}

// MockViewingService is a mock of ViewingService interface
type MockViewingService struct {
	ctrl     *gomock.Controller
	recorder *MockViewingServiceMockRecorder
}

// MockViewingServiceMockRecorder is the mock recorder for MockViewingService
type MockViewingServiceMockRecorder struct {
	mock *MockViewingService
}

// NewMockViewingService creates a new mock instance
func NewMockViewingService(ctrl *gomock.Controller) *MockViewingService {
	mock := &MockViewingService{ctrl: ctrl}
	mock.recorder = &MockViewingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockViewingService) EXPECT() *MockViewingServiceMockRecorder {
	return m.recorder
}

// GetFeed mocks base method
func (m *MockViewingService) GetFeed(arg0, arg1 string) ([]*rssfeeder.FeedEntry, error) {
	ret := m.ctrl.Call(m, "GetFeed", arg0, arg1)
	ret0, _ := ret[0].([]*rssfeeder.FeedEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeed indicates an expected call of GetFeed
func (mr *MockViewingServiceMockRecorder) GetFeed(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeed", reflect.TypeOf((*MockViewingService)(nil).GetFeed), arg0, arg1)
}

// GetFeeds mocks base method
func (m *MockViewingService) GetFeeds(arg0 string) ([]*rssfeeder.Feed, error) {
	ret := m.ctrl.Call(m, "GetFeeds", arg0)
	ret0, _ := ret[0].([]*rssfeeder.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeeds indicates an expected call of GetFeeds
func (mr *MockViewingServiceMockRecorder) GetFeeds(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeeds", reflect.TypeOf((*MockViewingService)(nil).GetFeeds), arg0)
}