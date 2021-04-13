// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	logr "github.com/mattermost/logr"
	mock "github.com/stretchr/testify/mock"
)

// MetricsInterface is an autogenerated mock type for the MetricsInterface type
type MetricsInterface struct {
	mock.Mock
}

// AddMemCacheHitCounter provides a mock function with given fields: cacheName, amount
func (_m *MetricsInterface) AddMemCacheHitCounter(cacheName string, amount float64) {
	_m.Called(cacheName, amount)
}

// AddMemCacheMissCounter provides a mock function with given fields: cacheName, amount
func (_m *MetricsInterface) AddMemCacheMissCounter(cacheName string, amount float64) {
	_m.Called(cacheName, amount)
}

// DecrementJobActive provides a mock function with given fields: jobType
func (_m *MetricsInterface) DecrementJobActive(jobType string) {
	_m.Called(jobType)
}

// DecrementWebSocketBroadcastBufferSize provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) DecrementWebSocketBroadcastBufferSize(hub string, amount float64) {
	_m.Called(hub, amount)
}

// DecrementWebSocketBroadcastUsersRegistered provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) DecrementWebSocketBroadcastUsersRegistered(hub string, amount float64) {
	_m.Called(hub, amount)
}

// GetLoggerMetricsCollector provides a mock function with given fields:
func (_m *MetricsInterface) GetLoggerMetricsCollector() logr.MetricsCollector {
	ret := _m.Called()

	var r0 logr.MetricsCollector
	if rf, ok := ret.Get(0).(func() logr.MetricsCollector); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logr.MetricsCollector)
		}
	}

	return r0
}

// IncrementChannelIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementChannelIndexCounter() {
	_m.Called()
}

// IncrementClusterEventType provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementClusterEventType(eventType string) {
	_m.Called(eventType)
}

// IncrementClusterRequest provides a mock function with given fields:
func (_m *MetricsInterface) IncrementClusterRequest() {
	_m.Called()
}

// IncrementEtagHitCounter provides a mock function with given fields: route
func (_m *MetricsInterface) IncrementEtagHitCounter(route string) {
	_m.Called(route)
}

// IncrementEtagMissCounter provides a mock function with given fields: route
func (_m *MetricsInterface) IncrementEtagMissCounter(route string) {
	_m.Called(route)
}

// IncrementFileIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementFileIndexCounter() {
	_m.Called()
}

// IncrementFilesSearchCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementFilesSearchCounter() {
	_m.Called()
}

// IncrementHttpError provides a mock function with given fields:
func (_m *MetricsInterface) IncrementHttpError() {
	_m.Called()
}

// IncrementHttpRequest provides a mock function with given fields:
func (_m *MetricsInterface) IncrementHttpRequest() {
	_m.Called()
}

// IncrementJobActive provides a mock function with given fields: jobType
func (_m *MetricsInterface) IncrementJobActive(jobType string) {
	_m.Called(jobType)
}

// IncrementLogin provides a mock function with given fields:
func (_m *MetricsInterface) IncrementLogin() {
	_m.Called()
}

// IncrementLoginFail provides a mock function with given fields:
func (_m *MetricsInterface) IncrementLoginFail() {
	_m.Called()
}

// IncrementMemCacheHitCounter provides a mock function with given fields: cacheName
func (_m *MetricsInterface) IncrementMemCacheHitCounter(cacheName string) {
	_m.Called(cacheName)
}

// IncrementMemCacheHitCounterSession provides a mock function with given fields:
func (_m *MetricsInterface) IncrementMemCacheHitCounterSession() {
	_m.Called()
}

// IncrementMemCacheInvalidationCounter provides a mock function with given fields: cacheName
func (_m *MetricsInterface) IncrementMemCacheInvalidationCounter(cacheName string) {
	_m.Called(cacheName)
}

// IncrementMemCacheInvalidationCounterSession provides a mock function with given fields:
func (_m *MetricsInterface) IncrementMemCacheInvalidationCounterSession() {
	_m.Called()
}

// IncrementMemCacheMissCounter provides a mock function with given fields: cacheName
func (_m *MetricsInterface) IncrementMemCacheMissCounter(cacheName string) {
	_m.Called(cacheName)
}

// IncrementMemCacheMissCounterSession provides a mock function with given fields:
func (_m *MetricsInterface) IncrementMemCacheMissCounterSession() {
	_m.Called()
}

// IncrementPostBroadcast provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostBroadcast() {
	_m.Called()
}

// IncrementPostCreate provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostCreate() {
	_m.Called()
}

// IncrementPostFileAttachment provides a mock function with given fields: count
func (_m *MetricsInterface) IncrementPostFileAttachment(count int) {
	_m.Called(count)
}

// IncrementPostIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostIndexCounter() {
	_m.Called()
}

// IncrementPostSentEmail provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostSentEmail() {
	_m.Called()
}

// IncrementPostSentPush provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostSentPush() {
	_m.Called()
}

// IncrementPostsSearchCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostsSearchCounter() {
	_m.Called()
}

// IncrementRemoteClusterConnStateChangeCounter provides a mock function with given fields: remoteID, online
func (_m *MetricsInterface) IncrementRemoteClusterConnStateChangeCounter(remoteID string, online bool) {
	_m.Called(remoteID, online)
}

// IncrementRemoteClusterMsgErrorsCounter provides a mock function with given fields: remoteID, timeout
func (_m *MetricsInterface) IncrementRemoteClusterMsgErrorsCounter(remoteID string, timeout bool) {
	_m.Called(remoteID, timeout)
}

// IncrementRemoteClusterMsgReceivedCounter provides a mock function with given fields: remoteID
func (_m *MetricsInterface) IncrementRemoteClusterMsgReceivedCounter(remoteID string) {
	_m.Called(remoteID)
}

// IncrementRemoteClusterMsgSentCounter provides a mock function with given fields: remoteID
func (_m *MetricsInterface) IncrementRemoteClusterMsgSentCounter(remoteID string) {
	_m.Called(remoteID)
}

// IncrementUserIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementUserIndexCounter() {
	_m.Called()
}

// IncrementWebSocketBroadcast provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementWebSocketBroadcast(eventType string) {
	_m.Called(eventType)
}

// IncrementWebSocketBroadcastBufferSize provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) IncrementWebSocketBroadcastBufferSize(hub string, amount float64) {
	_m.Called(hub, amount)
}

// IncrementWebSocketBroadcastUsersRegistered provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) IncrementWebSocketBroadcastUsersRegistered(hub string, amount float64) {
	_m.Called(hub, amount)
}

// IncrementWebhookPost provides a mock function with given fields:
func (_m *MetricsInterface) IncrementWebhookPost() {
	_m.Called()
}

// IncrementWebsocketEvent provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementWebsocketEvent(eventType string) {
	_m.Called(eventType)
}

// ObserveApiEndpointDuration provides a mock function with given fields: endpoint, method, statusCode, elapsed
func (_m *MetricsInterface) ObserveApiEndpointDuration(endpoint string, method string, statusCode string, elapsed float64) {
	_m.Called(endpoint, method, statusCode, elapsed)
}

// ObserveClusterRequestDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObserveClusterRequestDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObserveEnabledUsers provides a mock function with given fields: users
func (_m *MetricsInterface) ObserveEnabledUsers(users int64) {
	_m.Called(users)
}

// ObserveFilesSearchDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObserveFilesSearchDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObservePluginApiDuration provides a mock function with given fields: pluginID, apiName, success, elapsed
func (_m *MetricsInterface) ObservePluginApiDuration(pluginID string, apiName string, success bool, elapsed float64) {
	_m.Called(pluginID, apiName, success, elapsed)
}

// ObservePluginHookDuration provides a mock function with given fields: pluginID, hookName, success, elapsed
func (_m *MetricsInterface) ObservePluginHookDuration(pluginID string, hookName string, success bool, elapsed float64) {
	_m.Called(pluginID, hookName, success, elapsed)
}

// ObservePluginMultiHookDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObservePluginMultiHookDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObservePluginMultiHookIterationDuration provides a mock function with given fields: pluginID, elapsed
func (_m *MetricsInterface) ObservePluginMultiHookIterationDuration(pluginID string, elapsed float64) {
	_m.Called(pluginID, elapsed)
}

// ObservePostsSearchDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObservePostsSearchDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObserveRemoteClusterClockSkew provides a mock function with given fields: remoteID, skew
func (_m *MetricsInterface) ObserveRemoteClusterClockSkew(remoteID string, skew float64) {
	_m.Called(remoteID, skew)
}

// ObserveRemoteClusterPingDuration provides a mock function with given fields: remoteID, elapsed
func (_m *MetricsInterface) ObserveRemoteClusterPingDuration(remoteID string, elapsed float64) {
	_m.Called(remoteID, elapsed)
}

// ObserveStoreMethodDuration provides a mock function with given fields: method, success, elapsed
func (_m *MetricsInterface) ObserveStoreMethodDuration(method string, success string, elapsed float64) {
	_m.Called(method, success, elapsed)
}

// Register provides a mock function with given fields:
func (_m *MetricsInterface) Register() {
	_m.Called()
}

// SetReplicaLagAbsolute provides a mock function with given fields: node, value
func (_m *MetricsInterface) SetReplicaLagAbsolute(node string, value float64) {
	_m.Called(node, value)
}

// SetReplicaLagTime provides a mock function with given fields: node, value
func (_m *MetricsInterface) SetReplicaLagTime(node string, value float64) {
	_m.Called(node, value)
}
