package models

import (
	"time"
)

type TelemetryClient interface {
	// (much omitted)

	// Log a user action with the specified name
	TrackEvent(name string)

	// Log a numeric value that is not specified with a specific event.
	// Typically used to send regular reports of performance indicators.
	TrackMetric(name string, value float64)

	// Log a trace message with the specified severity level.
	//	TrackTrace(name string, severity contracts.SeverityLevel)

	// Log an HTTP request with the specified method, URL, duration and
	// response code.
	TrackRequest(method, url string, duration time.Duration, responseCode string)

	// Log a dependency with the specified name, type, target, and
	// success status.
	TrackRemoteDependency(name, dependencyType, target string, success bool)

	// Log an availability test result with the specified test name,
	// duration, and success status.
	TrackAvailability(name string, duration time.Duration, success bool)

	// Log an exception with the specified error, which may be a string,
	// error or Stringer. The current callstack is collected
	// automatically.
	TrackException(err interface{})
}
