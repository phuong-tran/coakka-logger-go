package coakka_logger

import "fmt"

const LoggerCoreABIVersion uint32 = 10

const (
	StatusOK              = 0
	StatusInvalidArgument = 1
	StatusBadState        = 2
	StatusQueueFull       = 3
	StatusRecordTooLarge  = 4
	StatusTimedOut        = 5
	StatusBufferTooSmall  = 6
	StatusInternalError   = 7
)

const (
	LevelTrace = 0
	LevelDebug = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelError = 4
	LevelFatal = 5
)

const (
	SinkModeManualDrain = 0
	SinkModeFile        = 1
	SinkModeConsole     = 2
	SinkModeMulti       = 3
)

const (
	SinkTargetNone    = 0
	SinkTargetFile    = 1 << 0
	SinkTargetConsole = 1 << 1
)

const (
	PressureStateNormal          = 0
	PressureStateQuotaPressure   = 1
	PressureStateDropLowPriority = 2
	PressureStateEmergencyOnly   = 3
)

const (
	levelMaskTrace = 1 << LevelTrace
	levelMaskDebug = 1 << LevelDebug
	levelMaskInfo  = 1 << LevelInfo
	levelMaskWarn  = 1 << LevelWarn
	levelMaskError = 1 << LevelError
	levelMaskFatal = 1 << LevelFatal
)

type LoggerSpec struct {
	SystemName       string
	QueueCapacity    uint32
	CategoryCapacity uint32
	MessageCapacity  uint32
	MinLevel         int
}

func DefaultLoggerSpec() LoggerSpec {
	return LoggerSpec{
		SystemName:       "goLogger",
		QueueCapacity:    256,
		CategoryCapacity: 64,
		MessageCapacity:  512,
		MinLevel:         LevelTrace,
	}
}

func (s LoggerSpec) withDefaults() LoggerSpec {
	defaults := DefaultLoggerSpec()
	if s.SystemName == "" {
		s.SystemName = defaults.SystemName
	}
	if s.QueueCapacity == 0 {
		s.QueueCapacity = defaults.QueueCapacity
	}
	if s.CategoryCapacity == 0 {
		s.CategoryCapacity = defaults.CategoryCapacity
	}
	if s.MessageCapacity == 0 {
		s.MessageCapacity = defaults.MessageCapacity
	}
	if s.MinLevel == 0 {
		s.MinLevel = defaults.MinLevel
	}
	return s
}

type LoggerInfoSnapshot struct {
	ABIVersion     uint32 `json:"abiVersion"`
	RuntimeVersion string `json:"runtimeVersion"`
	GitCommit      string `json:"gitCommit"`
	DocsHint       string `json:"docsHint"`
}

type LoggerConfigSnapshot struct {
	SystemName       string `json:"systemName"`
	State            int    `json:"state"`
	StateName        string `json:"stateName"`
	QueueCapacity    uint32 `json:"queueCapacity"`
	CategoryCapacity uint32 `json:"categoryCapacity"`
	MessageCapacity  uint32 `json:"messageCapacity"`
}

type LoggerStatsSnapshot struct {
	State              int    `json:"state"`
	StateName          string `json:"stateName"`
	QueueCapacity      uint32 `json:"queueCapacity"`
	QueueDepth         uint32 `json:"queueDepth"`
	QueueHighWatermark uint32 `json:"queueHighWatermark"`
	NextSequence       uint64 `json:"nextSequence"`
	EmittedCount       uint64 `json:"emittedCount"`
	DeliveredCount     uint64 `json:"deliveredCount"`
	DroppedCount       uint64 `json:"droppedCount"`
}

type LoggerRecordSnapshot struct {
	Sequence        uint64 `json:"sequence"`
	WallTimeUnixMs  uint64 `json:"wallTimeUnixMs"`
	MonotonicTimeNs uint64 `json:"monotonicTimeNs"`
	Level           int    `json:"level"`
	LevelName       string `json:"levelName"`
	Category        string `json:"category"`
	Message         string `json:"message"`
}

type StatusError struct {
	Operation  string
	Status     int
	StatusName string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("%s failed: %s (%d)", e.Operation, e.StatusName, e.Status)
}
