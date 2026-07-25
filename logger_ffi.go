package coakka_logger

/*
#cgo linux LDFLAGS: -ldl
#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#if defined(_WIN32)
#include <windows.h>
#else
#include <dlfcn.h>
#endif

typedef struct coakka_logger_core_handle coakka_logger_core_handle_t;

typedef struct {
  size_t struct_size;
  uint32_t abi_version;
  const char *runtime_version;
  const char *git_commit;
  const char *docs_hint;
} coakka_logger_core_info_t;

typedef struct {
  size_t struct_size;
  const char *system_name;
  uint32_t queue_capacity;
  int min_level;
  int sink_mode;
  const char *output_path;
  uint64_t max_file_size_bytes;
  uint64_t max_total_size_bytes;
  int pressure_state;
  uint32_t max_archived_files;
  uint32_t category_capacity;
  uint32_t message_capacity;
  const char *emergency_output_path;
  uint64_t emergency_file_size_bytes;
  uint32_t category_override_capacity;
  uint32_t sink_targets;
  int32_t console_fd;
  int32_t console_stdout_fd;
  int32_t console_stderr_fd;
  uint32_t console_stdout_level_mask;
  uint32_t console_stderr_level_mask;
  const char *file_output_path;
  uint64_t file_max_size_bytes;
  uint64_t file_total_size_limit_bytes;
  uint32_t file_max_archived_files;
  const char *file_emergency_output_path;
  uint64_t file_emergency_size_bytes;
} coakka_logger_core_config_t;

typedef struct {
  size_t struct_size;
  const char *system_name;
  int state;
  uint32_t queue_capacity;
  int min_level;
  int sink_mode;
  const char *output_path;
  uint64_t max_file_size_bytes;
  uint64_t max_total_size_bytes;
  int pressure_state;
  uint32_t max_archived_files;
  uint32_t category_capacity;
  uint32_t message_capacity;
  const char *emergency_output_path;
  uint64_t emergency_file_size_bytes;
  uint32_t category_override_capacity;
  uint32_t sink_targets;
  int32_t console_fd;
  int32_t console_stdout_fd;
  int32_t console_stderr_fd;
  uint32_t console_stdout_level_mask;
  uint32_t console_stderr_level_mask;
  const char *file_output_path;
  uint64_t file_max_size_bytes;
  uint64_t file_total_size_limit_bytes;
  uint32_t file_max_archived_files;
  const char *file_emergency_output_path;
  uint64_t file_emergency_size_bytes;
} coakka_logger_core_config_view_t;

typedef struct {
  size_t struct_size;
  int state;
  uint32_t queue_capacity;
  uint32_t queue_depth;
  uint32_t queue_high_watermark;
  uint64_t next_sequence;
  uint64_t emitted_count;
  uint64_t delivered_count;
  uint64_t dropped_count;
  uint64_t dropped_below_level_count;
  uint64_t sink_write_count;
  uint64_t sink_bytes_written;
  uint64_t sink_write_failure_count;
  uint64_t sink_current_file_size_bytes;
  uint64_t sink_total_size_bytes;
  int pressure_state;
  uint64_t sink_roll_count;
  uint64_t sink_deleted_archive_count;
  uint64_t dropped_pressure_count;
  uint32_t emergency_active;
  int emergency_last_reason;
  uint64_t emergency_enter_count;
  uint64_t emergency_recovered_count;
  uint64_t emergency_bytes_written;
  uint64_t emergency_write_failure_count;
  uint64_t emergency_overwrite_count;
  int emergency_last_recovered_reason;
  uint64_t emergency_last_recovered_duration_ms;
  uint64_t emergency_last_recovered_dropped_total;
  uint64_t emergency_last_recovered_dropped_pressure;
  uint64_t emergency_last_recovered_deleted_archives;
  uint64_t reload_count;
  uint32_t category_override_count;
  uint64_t sink_reopen_failure_count;
  uint64_t sink_roll_failure_count;
  uint64_t sink_append_failure_count;
  uint64_t sink_cleanup_failure_count;
  uint64_t file_write_count;
  uint64_t file_bytes_written;
  uint64_t file_write_failure_count;
  uint64_t file_current_size_bytes;
  uint64_t file_total_size_bytes;
  uint64_t file_roll_count;
  uint64_t file_deleted_archive_count;
  uint64_t file_reopen_failure_count;
  uint64_t file_roll_failure_count;
  uint64_t file_append_failure_count;
  uint64_t file_cleanup_failure_count;
  uint64_t console_write_count;
  uint64_t console_bytes_written;
  uint64_t console_write_failure_count;
} coakka_logger_core_stats_t;

typedef struct {
  size_t struct_size;
  uint64_t sequence;
  uint64_t wall_time_unix_ms;
  uint64_t monotonic_time_ns;
  int level;
  char *category;
  size_t category_capacity;
  size_t category_length;
  char *message;
  size_t message_capacity;
  size_t message_length;
} coakka_logger_core_record_buffer_t;

typedef uint32_t (*coakka_logger_core_get_abi_version_fn)(void);
typedef int (*coakka_logger_core_get_info_fn)(coakka_logger_core_info_t *out_info);
typedef int (*coakka_logger_core_create_fn)(const coakka_logger_core_config_t *config, coakka_logger_core_handle_t **out_handle);
typedef int (*coakka_logger_core_start_fn)(coakka_logger_core_handle_t *handle);
typedef int (*coakka_logger_core_stop_fn)(coakka_logger_core_handle_t *handle);
typedef void (*coakka_logger_core_destroy_fn)(coakka_logger_core_handle_t *handle);
typedef int (*coakka_logger_core_get_config_fn)(coakka_logger_core_handle_t *handle, coakka_logger_core_config_view_t *out_view);
typedef int (*coakka_logger_core_get_stats_fn)(coakka_logger_core_handle_t *handle, coakka_logger_core_stats_t *out_stats);
typedef int (*coakka_logger_core_is_enabled_fn)(coakka_logger_core_handle_t *handle, int level);
typedef int (*coakka_logger_core_is_enabled_for_category_fn)(coakka_logger_core_handle_t *handle, const char *category, int level);
typedef int (*coakka_logger_core_log_fn)(coakka_logger_core_handle_t *handle, int level, const char *category, const char *message, uint64_t *out_sequence);
typedef int (*coakka_logger_core_read_next_fn)(coakka_logger_core_handle_t *handle, uint32_t timeout_ms, coakka_logger_core_record_buffer_t *out_record);
typedef const char *(*coakka_logger_name_fn)(int value);

typedef struct {
  void *handle;
  coakka_logger_core_get_abi_version_fn get_abi_version;
  coakka_logger_core_get_info_fn get_info;
  coakka_logger_core_create_fn create;
  coakka_logger_core_start_fn start;
  coakka_logger_core_stop_fn stop;
  coakka_logger_core_destroy_fn destroy;
  coakka_logger_core_get_config_fn get_config;
  coakka_logger_core_get_stats_fn get_stats;
  coakka_logger_core_is_enabled_fn is_enabled;
  coakka_logger_core_is_enabled_for_category_fn is_enabled_for_category;
  coakka_logger_core_log_fn log;
  coakka_logger_core_read_next_fn read_next;
  coakka_logger_name_fn status_name;
  coakka_logger_name_fn level_name;
  coakka_logger_name_fn state_name;
} coakka_logger_go_bindings_t;

#if defined(_WIN32)
static char *coakka_logger_go_strdup_owned(const char *text) {
  size_t length = strlen(text) + 1u;
  char *copy = (char *)malloc(length);
  if (copy != NULL) {
    memcpy(copy, text, length);
  }
  return copy;
}
#endif

static int coakka_logger_go_load_symbol(void *handle, void **out_symbol, const char *symbol_name, char **error_out) {
#if defined(_WIN32)
  FARPROC symbol = GetProcAddress((HMODULE)handle, symbol_name);
  if (symbol == NULL) {
    if (error_out != NULL) {
      char buffer[160];
      snprintf(buffer, sizeof(buffer), "GetProcAddress %s failed: %lu", symbol_name, (unsigned long)GetLastError());
      *error_out = coakka_logger_go_strdup_owned(buffer);
    }
    return -1;
  }
  *out_symbol = (void *)symbol;
  return 0;
#else
  dlerror();
  *out_symbol = dlsym(handle, symbol_name);
  const char *error = dlerror();
  if (error != NULL) {
    if (error_out != NULL) {
      *error_out = strdup(error);
    }
    return -1;
  }
  return 0;
#endif
}

static coakka_logger_go_bindings_t *coakka_logger_go_open_library(const char *path, char **error_out) {
#if defined(_WIN32)
  HMODULE handle = LoadLibraryA(path);
  if (handle == NULL) {
    if (error_out != NULL) {
      char buffer[160];
      snprintf(buffer, sizeof(buffer), "LoadLibrary failed: %lu", (unsigned long)GetLastError());
      *error_out = coakka_logger_go_strdup_owned(buffer);
    }
    return NULL;
  }
#else
  void *handle = dlopen(path, RTLD_NOW | RTLD_LOCAL);
  if (handle == NULL) {
    if (error_out != NULL) {
      const char *error = dlerror();
      *error_out = error == NULL ? strdup("dlopen failed") : strdup(error);
    }
    return NULL;
  }
#endif

  coakka_logger_go_bindings_t *bindings = (coakka_logger_go_bindings_t *)calloc(1, sizeof(coakka_logger_go_bindings_t));
  if (bindings == NULL) {
    if (error_out != NULL) {
      *error_out = strdup("calloc failed");
    }
    dlclose(handle);
    return NULL;
  }

  bindings->handle = handle;
  if (coakka_logger_go_load_symbol(handle, (void **)&bindings->get_abi_version, "coakka_logger_core_get_abi_version", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->get_info, "coakka_logger_core_get_info", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->create, "coakka_logger_core_create", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->start, "coakka_logger_core_start", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->stop, "coakka_logger_core_stop", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->destroy, "coakka_logger_core_destroy", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->get_config, "coakka_logger_core_get_config", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->get_stats, "coakka_logger_core_get_stats", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->is_enabled, "coakka_logger_core_is_enabled", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->is_enabled_for_category, "coakka_logger_core_is_enabled_for_category", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->log, "coakka_logger_core_log", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->read_next, "coakka_logger_core_read_next", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->status_name, "coakka_logger_status_name", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->level_name, "coakka_logger_level_name", error_out) != 0 ||
      coakka_logger_go_load_symbol(handle, (void **)&bindings->state_name, "coakka_logger_state_name", error_out) != 0) {
    dlclose(handle);
    free(bindings);
    return NULL;
  }

  return bindings;
}

static void coakka_logger_go_close_library(coakka_logger_go_bindings_t *bindings) {
  if (bindings == NULL) {
    return;
  }
  if (bindings->handle != NULL) {
#if defined(_WIN32)
    FreeLibrary((HMODULE)bindings->handle);
#else
    dlclose(bindings->handle);
#endif
  }
  free(bindings);
}

static uint32_t coakka_logger_go_get_abi_version(coakka_logger_go_bindings_t *bindings) {
  return bindings->get_abi_version();
}

static int coakka_logger_go_get_info(coakka_logger_go_bindings_t *bindings, coakka_logger_core_info_t *out_info) {
  return bindings->get_info(out_info);
}

static int coakka_logger_go_create(coakka_logger_go_bindings_t *bindings, const coakka_logger_core_config_t *config, coakka_logger_core_handle_t **out_handle) {
  return bindings->create(config, out_handle);
}

static int coakka_logger_go_start(coakka_logger_go_bindings_t *bindings, coakka_logger_core_handle_t *handle) {
  return bindings->start(handle);
}

static int coakka_logger_go_stop(coakka_logger_go_bindings_t *bindings, coakka_logger_core_handle_t *handle) {
  return bindings->stop(handle);
}

static void coakka_logger_go_destroy(coakka_logger_go_bindings_t *bindings, coakka_logger_core_handle_t *handle) {
  bindings->destroy(handle);
}

static int coakka_logger_go_get_config(coakka_logger_go_bindings_t *bindings, coakka_logger_core_handle_t *handle, coakka_logger_core_config_view_t *out_view) {
  return bindings->get_config(handle, out_view);
}

static int coakka_logger_go_get_stats(coakka_logger_go_bindings_t *bindings, coakka_logger_core_handle_t *handle, coakka_logger_core_stats_t *out_stats) {
  return bindings->get_stats(handle, out_stats);
}

static int coakka_logger_go_is_enabled(coakka_logger_go_bindings_t *bindings, coakka_logger_core_handle_t *handle, int level) {
  return bindings->is_enabled(handle, level);
}

static int coakka_logger_go_is_enabled_for_category(coakka_logger_go_bindings_t *bindings, coakka_logger_core_handle_t *handle, const char *category, int level) {
  return bindings->is_enabled_for_category(handle, category, level);
}

static int coakka_logger_go_log(coakka_logger_go_bindings_t *bindings, coakka_logger_core_handle_t *handle, int level, const char *category, const char *message, uint64_t *out_sequence) {
  return bindings->log(handle, level, category, message, out_sequence);
}

static int coakka_logger_go_read_next(coakka_logger_go_bindings_t *bindings, coakka_logger_core_handle_t *handle, uint32_t timeout_ms, coakka_logger_core_record_buffer_t *out_record) {
  return bindings->read_next(handle, timeout_ms, out_record);
}

static const char *coakka_logger_go_status_name(coakka_logger_go_bindings_t *bindings, int status) {
  return bindings->status_name(status);
}

static const char *coakka_logger_go_level_name(coakka_logger_go_bindings_t *bindings, int level) {
  return bindings->level_name(level);
}

static const char *coakka_logger_go_state_name(coakka_logger_go_bindings_t *bindings, int state) {
  return bindings->state_name(state);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type nativeBindings struct {
	ptr *C.coakka_logger_go_bindings_t
}

type nativeHandle struct {
	ptr *C.coakka_logger_core_handle_t
}

func openNativeBindings(path string) (*nativeBindings, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	var cError *C.char
	bindings := C.coakka_logger_go_open_library(cPath, &cError)
	if bindings == nil {
		defer C.free(unsafe.Pointer(cError))
		if cError != nil {
			return nil, fmt.Errorf("dlopen logger native: %s", C.GoString(cError))
		}
		return nil, fmt.Errorf("dlopen logger native failed")
	}
	return &nativeBindings{ptr: bindings}, nil
}

func (b *nativeBindings) close() {
	if b == nil || b.ptr == nil {
		return
	}
	C.coakka_logger_go_close_library(b.ptr)
	b.ptr = nil
}

func (b *nativeBindings) getAbiVersion() uint32 {
	return uint32(C.coakka_logger_go_get_abi_version(b.ptr))
}

func (b *nativeBindings) requireOK(status C.int, operation string) error {
	if int(status) == StatusOK {
		return nil
	}
	return &StatusError{
		Operation:  operation,
		Status:     int(status),
		StatusName: C.GoString(C.coakka_logger_go_status_name(b.ptr, status)),
	}
}

func (b *nativeBindings) readInfo() (LoggerInfoSnapshot, error) {
	var info C.coakka_logger_core_info_t
	info.struct_size = C.size_t(unsafe.Sizeof(info))
	if err := b.requireOK(C.coakka_logger_go_get_info(b.ptr, &info), "logger_core_get_info"); err != nil {
		return LoggerInfoSnapshot{}, err
	}
	return LoggerInfoSnapshot{
		ABIVersion:     uint32(info.abi_version),
		RuntimeVersion: C.GoString(info.runtime_version),
		GitCommit:      C.GoString(info.git_commit),
		DocsHint:       C.GoString(info.docs_hint),
	}, nil
}

func (b *nativeBindings) create(spec LoggerSpec) (nativeHandle, error) {
	cSystemName := C.CString(spec.SystemName)
	defer C.free(unsafe.Pointer(cSystemName))
	config := C.coakka_logger_core_config_t{
		struct_size:                C.size_t(unsafe.Sizeof(C.coakka_logger_core_config_t{})),
		system_name:                cSystemName,
		queue_capacity:             C.uint32_t(spec.QueueCapacity),
		min_level:                  C.int(spec.MinLevel),
		sink_mode:                  C.int(SinkModeManualDrain),
		pressure_state:             C.int(PressureStateNormal),
		category_capacity:          C.uint32_t(spec.CategoryCapacity),
		message_capacity:           C.uint32_t(spec.MessageCapacity),
		category_override_capacity: C.uint32_t(32),
		sink_targets:               C.uint32_t(SinkTargetNone),
		console_fd:                 C.int32_t(-1),
		console_stdout_fd:          C.int32_t(-1),
		console_stderr_fd:          C.int32_t(-1),
		console_stdout_level_mask:  C.uint32_t(levelMaskTrace | levelMaskDebug | levelMaskInfo),
		console_stderr_level_mask:  C.uint32_t(levelMaskWarn | levelMaskError | levelMaskFatal),
	}
	var handle *C.coakka_logger_core_handle_t
	if err := b.requireOK(C.coakka_logger_go_create(b.ptr, &config, &handle), "logger_core_create"); err != nil {
		return nativeHandle{}, err
	}
	if handle == nil {
		return nativeHandle{}, fmt.Errorf("logger_core_create returned nil handle")
	}
	return nativeHandle{ptr: handle}, nil
}

func (b *nativeBindings) start(handle nativeHandle) error {
	return b.requireOK(C.coakka_logger_go_start(b.ptr, handle.ptr), "logger_core_start")
}

func (b *nativeBindings) stop(handle nativeHandle) error {
	return b.requireOK(C.coakka_logger_go_stop(b.ptr, handle.ptr), "logger_core_stop")
}

func (b *nativeBindings) destroy(handle nativeHandle) {
	if handle.ptr != nil {
		C.coakka_logger_go_destroy(b.ptr, handle.ptr)
	}
}

func (b *nativeBindings) config(handle nativeHandle) (LoggerConfigSnapshot, error) {
	var view C.coakka_logger_core_config_view_t
	view.struct_size = C.size_t(unsafe.Sizeof(view))
	if err := b.requireOK(C.coakka_logger_go_get_config(b.ptr, handle.ptr, &view), "logger_core_get_config"); err != nil {
		return LoggerConfigSnapshot{}, err
	}
	state := int(view.state)
	return LoggerConfigSnapshot{
		SystemName:       C.GoString(view.system_name),
		State:            state,
		StateName:        C.GoString(C.coakka_logger_go_state_name(b.ptr, C.int(state))),
		QueueCapacity:    uint32(view.queue_capacity),
		CategoryCapacity: uint32(view.category_capacity),
		MessageCapacity:  uint32(view.message_capacity),
	}, nil
}

func (b *nativeBindings) stats(handle nativeHandle) (LoggerStatsSnapshot, error) {
	var stats C.coakka_logger_core_stats_t
	stats.struct_size = C.size_t(unsafe.Sizeof(stats))
	if err := b.requireOK(C.coakka_logger_go_get_stats(b.ptr, handle.ptr, &stats), "logger_core_get_stats"); err != nil {
		return LoggerStatsSnapshot{}, err
	}
	state := int(stats.state)
	return LoggerStatsSnapshot{
		State:              state,
		StateName:          C.GoString(C.coakka_logger_go_state_name(b.ptr, C.int(state))),
		QueueCapacity:      uint32(stats.queue_capacity),
		QueueDepth:         uint32(stats.queue_depth),
		QueueHighWatermark: uint32(stats.queue_high_watermark),
		NextSequence:       uint64(stats.next_sequence),
		EmittedCount:       uint64(stats.emitted_count),
		DeliveredCount:     uint64(stats.delivered_count),
		DroppedCount:       uint64(stats.dropped_count),
	}, nil
}

func (b *nativeBindings) isEnabled(handle nativeHandle, level int) bool {
	return C.coakka_logger_go_is_enabled(b.ptr, handle.ptr, C.int(level)) != 0
}

func (b *nativeBindings) isEnabledForCategory(handle nativeHandle, category string, level int) bool {
	cCategory := C.CString(category)
	defer C.free(unsafe.Pointer(cCategory))
	return C.coakka_logger_go_is_enabled_for_category(b.ptr, handle.ptr, cCategory, C.int(level)) != 0
}

func (b *nativeBindings) log(handle nativeHandle, level int, category string, message string) (uint64, error) {
	cCategory := C.CString(category)
	defer C.free(unsafe.Pointer(cCategory))
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	var sequence C.uint64_t
	if err := b.requireOK(C.coakka_logger_go_log(b.ptr, handle.ptr, C.int(level), cCategory, cMessage, &sequence), "logger_core_log"); err != nil {
		return 0, err
	}
	return uint64(sequence), nil
}

func (b *nativeBindings) readNext(handle nativeHandle, timeoutMs uint32, categoryCapacity uint32, messageCapacity uint32) (*LoggerRecordSnapshot, error) {
	category := C.malloc(C.size_t(categoryCapacity))
	if category == nil {
		return nil, fmt.Errorf("malloc category buffer failed")
	}
	defer C.free(category)
	message := C.malloc(C.size_t(messageCapacity))
	if message == nil {
		return nil, fmt.Errorf("malloc message buffer failed")
	}
	defer C.free(message)
	record := C.coakka_logger_core_record_buffer_t{
		struct_size:       C.size_t(unsafe.Sizeof(C.coakka_logger_core_record_buffer_t{})),
		category:          (*C.char)(category),
		category_capacity: C.size_t(categoryCapacity),
		message:           (*C.char)(message),
		message_capacity:  C.size_t(messageCapacity),
	}
	status := C.coakka_logger_go_read_next(b.ptr, handle.ptr, C.uint32_t(timeoutMs), &record)
	if int(status) == StatusTimedOut {
		return nil, nil
	}
	if err := b.requireOK(status, "logger_core_read_next"); err != nil {
		return nil, err
	}
	level := int(record.level)
	return &LoggerRecordSnapshot{
		Sequence:        uint64(record.sequence),
		WallTimeUnixMs:  uint64(record.wall_time_unix_ms),
		MonotonicTimeNs: uint64(record.monotonic_time_ns),
		Level:           level,
		LevelName:       C.GoString(C.coakka_logger_go_level_name(b.ptr, C.int(level))),
		Category:        C.GoStringN((*C.char)(category), C.int(record.category_length)),
		Message:         C.GoStringN((*C.char)(message), C.int(record.message_length)),
	}, nil
}
