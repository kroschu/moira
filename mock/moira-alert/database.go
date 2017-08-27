// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira-alert (interfaces: Database)

package mock_moira_alert

import (
	gomock "github.com/golang/mock/gomock"
	moira_alert "github.com/moira-alert/moira-alert"
	tomb_v2 "gopkg.in/tomb.v2"
	time "time"
)

// MockDatabase is a mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return _m.recorder
}

// AcquireTriggerCheckLock mocks base method
func (_m *MockDatabase) AcquireTriggerCheckLock(_param0 string, _param1 int) error {
	ret := _m.ctrl.Call(_m, "AcquireTriggerCheckLock", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcquireTriggerCheckLock indicates an expected call of AcquireTriggerCheckLock
func (_mr *MockDatabaseMockRecorder) AcquireTriggerCheckLock(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AcquireTriggerCheckLock", arg0, arg1)
}

// AddNotification mocks base method
func (_m *MockDatabase) AddNotification(_param0 *moira_alert.ScheduledNotification) error {
	ret := _m.ctrl.Call(_m, "AddNotification", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNotification indicates an expected call of AddNotification
func (_mr *MockDatabaseMockRecorder) AddNotification(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddNotification", arg0)
}

// AddPatternMetric mocks base method
func (_m *MockDatabase) AddPatternMetric(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "AddPatternMetric", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPatternMetric indicates an expected call of AddPatternMetric
func (_mr *MockDatabaseMockRecorder) AddPatternMetric(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddPatternMetric", arg0, arg1)
}

// AddTriggerToCheck mocks base method
func (_m *MockDatabase) AddTriggerToCheck(_param0 string) error {
	ret := _m.ctrl.Call(_m, "AddTriggerToCheck", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTriggerToCheck indicates an expected call of AddTriggerToCheck
func (_mr *MockDatabaseMockRecorder) AddTriggerToCheck(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddTriggerToCheck", arg0)
}

// CleanupMetricValues mocks base method
func (_m *MockDatabase) CleanupMetricValues(_param0 string, _param1 int64) error {
	ret := _m.ctrl.Call(_m, "CleanupMetricValues", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupMetricValues indicates an expected call of CleanupMetricValues
func (_mr *MockDatabaseMockRecorder) CleanupMetricValues(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CleanupMetricValues", arg0, arg1)
}

// DeleteTag mocks base method
func (_m *MockDatabase) DeleteTag(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteTag", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTag indicates an expected call of DeleteTag
func (_mr *MockDatabaseMockRecorder) DeleteTag(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTag", arg0)
}

// DeleteTrigger mocks base method
func (_m *MockDatabase) DeleteTrigger(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteTrigger", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrigger indicates an expected call of DeleteTrigger
func (_mr *MockDatabaseMockRecorder) DeleteTrigger(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTrigger", arg0)
}

// DeleteTriggerCheckLock mocks base method
func (_m *MockDatabase) DeleteTriggerCheckLock(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteTriggerCheckLock", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTriggerCheckLock indicates an expected call of DeleteTriggerCheckLock
func (_mr *MockDatabaseMockRecorder) DeleteTriggerCheckLock(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTriggerCheckLock", arg0)
}

// DeleteTriggerThrottling mocks base method
func (_m *MockDatabase) DeleteTriggerThrottling(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteTriggerThrottling", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTriggerThrottling indicates an expected call of DeleteTriggerThrottling
func (_mr *MockDatabaseMockRecorder) DeleteTriggerThrottling(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTriggerThrottling", arg0)
}

// FetchEvent mocks base method
func (_m *MockDatabase) FetchEvent() (*moira_alert.EventData, error) {
	ret := _m.ctrl.Call(_m, "FetchEvent")
	ret0, _ := ret[0].(*moira_alert.EventData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchEvent indicates an expected call of FetchEvent
func (_mr *MockDatabaseMockRecorder) FetchEvent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchEvent")
}

// GetAllContacts mocks base method
func (_m *MockDatabase) GetAllContacts() ([]*moira_alert.ContactData, error) {
	ret := _m.ctrl.Call(_m, "GetAllContacts")
	ret0, _ := ret[0].([]*moira_alert.ContactData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllContacts indicates an expected call of GetAllContacts
func (_mr *MockDatabaseMockRecorder) GetAllContacts() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllContacts")
}

// GetChecksCount mocks base method
func (_m *MockDatabase) GetChecksCount() (int64, error) {
	ret := _m.ctrl.Call(_m, "GetChecksCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecksCount indicates an expected call of GetChecksCount
func (_mr *MockDatabaseMockRecorder) GetChecksCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetChecksCount")
}

// GetContact mocks base method
func (_m *MockDatabase) GetContact(_param0 string) (moira_alert.ContactData, error) {
	ret := _m.ctrl.Call(_m, "GetContact", _param0)
	ret0, _ := ret[0].(moira_alert.ContactData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContact indicates an expected call of GetContact
func (_mr *MockDatabaseMockRecorder) GetContact(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetContact", arg0)
}

// GetContacts mocks base method
func (_m *MockDatabase) GetContacts(_param0 []string) ([]*moira_alert.ContactData, error) {
	ret := _m.ctrl.Call(_m, "GetContacts", _param0)
	ret0, _ := ret[0].([]*moira_alert.ContactData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContacts indicates an expected call of GetContacts
func (_mr *MockDatabaseMockRecorder) GetContacts(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetContacts", arg0)
}

// GetEvents mocks base method
func (_m *MockDatabase) GetEvents(_param0 string, _param1 int64, _param2 int64) ([]*moira_alert.EventData, error) {
	ret := _m.ctrl.Call(_m, "GetEvents", _param0, _param1, _param2)
	ret0, _ := ret[0].([]*moira_alert.EventData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents
func (_mr *MockDatabaseMockRecorder) GetEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEvents", arg0, arg1, arg2)
}

// GetFilteredTriggerCheckIds mocks base method
func (_m *MockDatabase) GetFilteredTriggerCheckIds(_param0 []string, _param1 bool) ([]string, int64, error) {
	ret := _m.ctrl.Call(_m, "GetFilteredTriggerCheckIds", _param0, _param1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFilteredTriggerCheckIds indicates an expected call of GetFilteredTriggerCheckIds
func (_mr *MockDatabaseMockRecorder) GetFilteredTriggerCheckIds(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFilteredTriggerCheckIds", arg0, arg1)
}

// GetMetricRetention mocks base method
func (_m *MockDatabase) GetMetricRetention(_param0 string) (int64, error) {
	ret := _m.ctrl.Call(_m, "GetMetricRetention", _param0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricRetention indicates an expected call of GetMetricRetention
func (_mr *MockDatabaseMockRecorder) GetMetricRetention(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetricRetention", arg0)
}

// GetMetricsCount mocks base method
func (_m *MockDatabase) GetMetricsCount() (int64, error) {
	ret := _m.ctrl.Call(_m, "GetMetricsCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricsCount indicates an expected call of GetMetricsCount
func (_mr *MockDatabaseMockRecorder) GetMetricsCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetricsCount")
}

// GetMetricsValues mocks base method
func (_m *MockDatabase) GetMetricsValues(_param0 []string, _param1 int64, _param2 int64) (map[string][]*moira_alert.MetricValue, error) {
	ret := _m.ctrl.Call(_m, "GetMetricsValues", _param0, _param1, _param2)
	ret0, _ := ret[0].(map[string][]*moira_alert.MetricValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricsValues indicates an expected call of GetMetricsValues
func (_mr *MockDatabaseMockRecorder) GetMetricsValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetricsValues", arg0, arg1, arg2)
}

// GetNotificationTrigger mocks base method
func (_m *MockDatabase) GetNotificationTrigger(_param0 string) (moira_alert.TriggerData, error) {
	ret := _m.ctrl.Call(_m, "GetNotificationTrigger", _param0)
	ret0, _ := ret[0].(moira_alert.TriggerData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotificationTrigger indicates an expected call of GetNotificationTrigger
func (_mr *MockDatabaseMockRecorder) GetNotificationTrigger(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNotificationTrigger", arg0)
}

// GetNotifications mocks base method
func (_m *MockDatabase) GetNotifications(_param0 int64, _param1 int64) ([]*moira_alert.ScheduledNotification, int64, error) {
	ret := _m.ctrl.Call(_m, "GetNotifications", _param0, _param1)
	ret0, _ := ret[0].([]*moira_alert.ScheduledNotification)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNotifications indicates an expected call of GetNotifications
func (_mr *MockDatabaseMockRecorder) GetNotifications(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNotifications", arg0, arg1)
}

// GetNotificationsAndDelete mocks base method
func (_m *MockDatabase) GetNotificationsAndDelete(_param0 int64) ([]*moira_alert.ScheduledNotification, error) {
	ret := _m.ctrl.Call(_m, "GetNotificationsAndDelete", _param0)
	ret0, _ := ret[0].([]*moira_alert.ScheduledNotification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotificationsAndDelete indicates an expected call of GetNotificationsAndDelete
func (_mr *MockDatabaseMockRecorder) GetNotificationsAndDelete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNotificationsAndDelete", arg0)
}

// GetPatternMetrics mocks base method
func (_m *MockDatabase) GetPatternMetrics(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetPatternMetrics", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatternMetrics indicates an expected call of GetPatternMetrics
func (_mr *MockDatabaseMockRecorder) GetPatternMetrics(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPatternMetrics", arg0)
}

// GetPatternTriggerIds mocks base method
func (_m *MockDatabase) GetPatternTriggerIds(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetPatternTriggerIds", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatternTriggerIds indicates an expected call of GetPatternTriggerIds
func (_mr *MockDatabaseMockRecorder) GetPatternTriggerIds(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPatternTriggerIds", arg0)
}

// GetPatterns mocks base method
func (_m *MockDatabase) GetPatterns() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetPatterns")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatterns indicates an expected call of GetPatterns
func (_mr *MockDatabaseMockRecorder) GetPatterns() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPatterns")
}

// GetSubscription mocks base method
func (_m *MockDatabase) GetSubscription(_param0 string) (moira_alert.SubscriptionData, error) {
	ret := _m.ctrl.Call(_m, "GetSubscription", _param0)
	ret0, _ := ret[0].(moira_alert.SubscriptionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscription indicates an expected call of GetSubscription
func (_mr *MockDatabaseMockRecorder) GetSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscription", arg0)
}

// GetSubscriptions mocks base method
func (_m *MockDatabase) GetSubscriptions(_param0 []string) ([]*moira_alert.SubscriptionData, error) {
	ret := _m.ctrl.Call(_m, "GetSubscriptions", _param0)
	ret0, _ := ret[0].([]*moira_alert.SubscriptionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptions indicates an expected call of GetSubscriptions
func (_mr *MockDatabaseMockRecorder) GetSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscriptions", arg0)
}

// GetTagNames mocks base method
func (_m *MockDatabase) GetTagNames() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTagNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagNames indicates an expected call of GetTagNames
func (_mr *MockDatabaseMockRecorder) GetTagNames() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTagNames")
}

// GetTagTriggerIds mocks base method
func (_m *MockDatabase) GetTagTriggerIds(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTagTriggerIds", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagTriggerIds indicates an expected call of GetTagTriggerIds
func (_mr *MockDatabaseMockRecorder) GetTagTriggerIds(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTagTriggerIds", arg0)
}

// GetTagsSubscriptions mocks base method
func (_m *MockDatabase) GetTagsSubscriptions(_param0 []string) ([]moira_alert.SubscriptionData, error) {
	ret := _m.ctrl.Call(_m, "GetTagsSubscriptions", _param0)
	ret0, _ := ret[0].([]moira_alert.SubscriptionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsSubscriptions indicates an expected call of GetTagsSubscriptions
func (_mr *MockDatabaseMockRecorder) GetTagsSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTagsSubscriptions", arg0)
}

// GetTrigger mocks base method
func (_m *MockDatabase) GetTrigger(_param0 string) (*moira_alert.Trigger, error) {
	ret := _m.ctrl.Call(_m, "GetTrigger", _param0)
	ret0, _ := ret[0].(*moira_alert.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrigger indicates an expected call of GetTrigger
func (_mr *MockDatabaseMockRecorder) GetTrigger(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTrigger", arg0)
}

// GetTriggerCheckIds mocks base method
func (_m *MockDatabase) GetTriggerCheckIds() ([]string, int64, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerCheckIds")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTriggerCheckIds indicates an expected call of GetTriggerCheckIds
func (_mr *MockDatabaseMockRecorder) GetTriggerCheckIds() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerCheckIds")
}

// GetTriggerChecks mocks base method
func (_m *MockDatabase) GetTriggerChecks(_param0 []string) ([]moira_alert.TriggerChecks, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerChecks", _param0)
	ret0, _ := ret[0].([]moira_alert.TriggerChecks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerChecks indicates an expected call of GetTriggerChecks
func (_mr *MockDatabaseMockRecorder) GetTriggerChecks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerChecks", arg0)
}

// GetTriggerEventsCount mocks base method
func (_m *MockDatabase) GetTriggerEventsCount(_param0 string, _param1 int64) int64 {
	ret := _m.ctrl.Call(_m, "GetTriggerEventsCount", _param0, _param1)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetTriggerEventsCount indicates an expected call of GetTriggerEventsCount
func (_mr *MockDatabaseMockRecorder) GetTriggerEventsCount(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerEventsCount", arg0, arg1)
}

// GetTriggerIds mocks base method
func (_m *MockDatabase) GetTriggerIds() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerIds")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerIds indicates an expected call of GetTriggerIds
func (_mr *MockDatabaseMockRecorder) GetTriggerIds() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerIds")
}

// GetTriggerLastCheck mocks base method
func (_m *MockDatabase) GetTriggerLastCheck(_param0 string) (*moira_alert.CheckData, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerLastCheck", _param0)
	ret0, _ := ret[0].(*moira_alert.CheckData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerLastCheck indicates an expected call of GetTriggerLastCheck
func (_mr *MockDatabaseMockRecorder) GetTriggerLastCheck(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerLastCheck", arg0)
}

// GetTriggerTags mocks base method
func (_m *MockDatabase) GetTriggerTags(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerTags", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerTags indicates an expected call of GetTriggerTags
func (_mr *MockDatabaseMockRecorder) GetTriggerTags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerTags", arg0)
}

// GetTriggerThrottlingTimestamps mocks base method
func (_m *MockDatabase) GetTriggerThrottlingTimestamps(_param0 string) (time.Time, time.Time) {
	ret := _m.ctrl.Call(_m, "GetTriggerThrottlingTimestamps", _param0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(time.Time)
	return ret0, ret1
}

// GetTriggerThrottlingTimestamps indicates an expected call of GetTriggerThrottlingTimestamps
func (_mr *MockDatabaseMockRecorder) GetTriggerThrottlingTimestamps(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerThrottlingTimestamps", arg0)
}

// GetTriggerToCheck mocks base method
func (_m *MockDatabase) GetTriggerToCheck() (*string, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerToCheck")
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerToCheck indicates an expected call of GetTriggerToCheck
func (_mr *MockDatabaseMockRecorder) GetTriggerToCheck() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerToCheck")
}

// GetTriggers mocks base method
func (_m *MockDatabase) GetTriggers(_param0 []string) ([]*moira_alert.Trigger, error) {
	ret := _m.ctrl.Call(_m, "GetTriggers", _param0)
	ret0, _ := ret[0].([]*moira_alert.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggers indicates an expected call of GetTriggers
func (_mr *MockDatabaseMockRecorder) GetTriggers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggers", arg0)
}

// GetUserContactIDs mocks base method
func (_m *MockDatabase) GetUserContactIDs(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetUserContactIDs", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserContactIDs indicates an expected call of GetUserContactIDs
func (_mr *MockDatabaseMockRecorder) GetUserContactIDs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserContactIDs", arg0)
}

// GetUserSubscriptionIDs mocks base method
func (_m *MockDatabase) GetUserSubscriptionIDs(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetUserSubscriptionIDs", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSubscriptionIDs indicates an expected call of GetUserSubscriptionIDs
func (_mr *MockDatabaseMockRecorder) GetUserSubscriptionIDs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserSubscriptionIDs", arg0)
}

// PushEvent mocks base method
func (_m *MockDatabase) PushEvent(_param0 *moira_alert.EventData, _param1 bool) error {
	ret := _m.ctrl.Call(_m, "PushEvent", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushEvent indicates an expected call of PushEvent
func (_mr *MockDatabaseMockRecorder) PushEvent(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PushEvent", arg0, arg1)
}

// RemoveContact mocks base method
func (_m *MockDatabase) RemoveContact(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "RemoveContact", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContact indicates an expected call of RemoveContact
func (_mr *MockDatabaseMockRecorder) RemoveContact(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveContact", arg0, arg1)
}

// RemoveNotification mocks base method
func (_m *MockDatabase) RemoveNotification(_param0 string) (int64, error) {
	ret := _m.ctrl.Call(_m, "RemoveNotification", _param0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveNotification indicates an expected call of RemoveNotification
func (_mr *MockDatabaseMockRecorder) RemoveNotification(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveNotification", arg0)
}

// RemovePattern mocks base method
func (_m *MockDatabase) RemovePattern(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemovePattern", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePattern indicates an expected call of RemovePattern
func (_mr *MockDatabaseMockRecorder) RemovePattern(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePattern", arg0)
}

// RemovePatternTriggers mocks base method
func (_m *MockDatabase) RemovePatternTriggers(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemovePatternTriggers", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePatternTriggers indicates an expected call of RemovePatternTriggers
func (_mr *MockDatabaseMockRecorder) RemovePatternTriggers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePatternTriggers", arg0)
}

// RemovePatternWithMetrics mocks base method
func (_m *MockDatabase) RemovePatternWithMetrics(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemovePatternWithMetrics", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePatternWithMetrics indicates an expected call of RemovePatternWithMetrics
func (_mr *MockDatabaseMockRecorder) RemovePatternWithMetrics(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePatternWithMetrics", arg0)
}

// RemovePatternsMetrics mocks base method
func (_m *MockDatabase) RemovePatternsMetrics(_param0 []string) error {
	ret := _m.ctrl.Call(_m, "RemovePatternsMetrics", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePatternsMetrics indicates an expected call of RemovePatternsMetrics
func (_mr *MockDatabaseMockRecorder) RemovePatternsMetrics(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePatternsMetrics", arg0)
}

// RemoveSubscription mocks base method
func (_m *MockDatabase) RemoveSubscription(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "RemoveSubscription", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSubscription indicates an expected call of RemoveSubscription
func (_mr *MockDatabaseMockRecorder) RemoveSubscription(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveSubscription", arg0, arg1)
}

// SaveContact mocks base method
func (_m *MockDatabase) SaveContact(_param0 *moira_alert.ContactData) error {
	ret := _m.ctrl.Call(_m, "SaveContact", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveContact indicates an expected call of SaveContact
func (_mr *MockDatabaseMockRecorder) SaveContact(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveContact", arg0)
}

// SaveMetrics mocks base method
func (_m *MockDatabase) SaveMetrics(_param0 map[string]*moira_alert.MatchedMetric) error {
	ret := _m.ctrl.Call(_m, "SaveMetrics", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveMetrics indicates an expected call of SaveMetrics
func (_mr *MockDatabaseMockRecorder) SaveMetrics(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveMetrics", arg0)
}

// SaveSubscription mocks base method
func (_m *MockDatabase) SaveSubscription(_param0 *moira_alert.SubscriptionData) error {
	ret := _m.ctrl.Call(_m, "SaveSubscription", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSubscription indicates an expected call of SaveSubscription
func (_mr *MockDatabaseMockRecorder) SaveSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveSubscription", arg0)
}

// SaveTrigger mocks base method
func (_m *MockDatabase) SaveTrigger(_param0 string, _param1 *moira_alert.Trigger) error {
	ret := _m.ctrl.Call(_m, "SaveTrigger", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTrigger indicates an expected call of SaveTrigger
func (_mr *MockDatabaseMockRecorder) SaveTrigger(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveTrigger", arg0, arg1)
}

// SetTriggerCheckLock mocks base method
func (_m *MockDatabase) SetTriggerCheckLock(_param0 string) (bool, error) {
	ret := _m.ctrl.Call(_m, "SetTriggerCheckLock", _param0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTriggerCheckLock indicates an expected call of SetTriggerCheckLock
func (_mr *MockDatabaseMockRecorder) SetTriggerCheckLock(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerCheckLock", arg0)
}

// SetTriggerLastCheck mocks base method
func (_m *MockDatabase) SetTriggerLastCheck(_param0 string, _param1 *moira_alert.CheckData) error {
	ret := _m.ctrl.Call(_m, "SetTriggerLastCheck", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTriggerLastCheck indicates an expected call of SetTriggerLastCheck
func (_mr *MockDatabaseMockRecorder) SetTriggerLastCheck(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerLastCheck", arg0, arg1)
}

// SetTriggerMetricsMaintenance mocks base method
func (_m *MockDatabase) SetTriggerMetricsMaintenance(_param0 string, _param1 map[string]int64) error {
	ret := _m.ctrl.Call(_m, "SetTriggerMetricsMaintenance", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTriggerMetricsMaintenance indicates an expected call of SetTriggerMetricsMaintenance
func (_mr *MockDatabaseMockRecorder) SetTriggerMetricsMaintenance(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerMetricsMaintenance", arg0, arg1)
}

// SetTriggerThrottlingTimestamp mocks base method
func (_m *MockDatabase) SetTriggerThrottlingTimestamp(_param0 string, _param1 time.Time) error {
	ret := _m.ctrl.Call(_m, "SetTriggerThrottlingTimestamp", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTriggerThrottlingTimestamp indicates an expected call of SetTriggerThrottlingTimestamp
func (_mr *MockDatabaseMockRecorder) SetTriggerThrottlingTimestamp(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerThrottlingTimestamp", arg0, arg1)
}

// SubscribeMetricEvents mocks base method
func (_m *MockDatabase) SubscribeMetricEvents(_param0 *tomb_v2.Tomb) <-chan *moira_alert.MetricEvent {
	ret := _m.ctrl.Call(_m, "SubscribeMetricEvents", _param0)
	ret0, _ := ret[0].(<-chan *moira_alert.MetricEvent)
	return ret0
}

// SubscribeMetricEvents indicates an expected call of SubscribeMetricEvents
func (_mr *MockDatabaseMockRecorder) SubscribeMetricEvents(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SubscribeMetricEvents", arg0)
}

// UpdateMetricsHeartbeat mocks base method
func (_m *MockDatabase) UpdateMetricsHeartbeat() error {
	ret := _m.ctrl.Call(_m, "UpdateMetricsHeartbeat")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetricsHeartbeat indicates an expected call of UpdateMetricsHeartbeat
func (_mr *MockDatabaseMockRecorder) UpdateMetricsHeartbeat() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateMetricsHeartbeat")
}

// WriteContact mocks base method
func (_m *MockDatabase) WriteContact(_param0 *moira_alert.ContactData) error {
	ret := _m.ctrl.Call(_m, "WriteContact", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteContact indicates an expected call of WriteContact
func (_mr *MockDatabaseMockRecorder) WriteContact(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteContact", arg0)
}

// WriteSubscriptions mocks base method
func (_m *MockDatabase) WriteSubscriptions(_param0 []*moira_alert.SubscriptionData) error {
	ret := _m.ctrl.Call(_m, "WriteSubscriptions", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteSubscriptions indicates an expected call of WriteSubscriptions
func (_mr *MockDatabaseMockRecorder) WriteSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteSubscriptions", arg0)
}
