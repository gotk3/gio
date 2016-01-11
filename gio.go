package gio

// #cgo pkg-config: gio-2.0 glib-2.0
// #include <gio/gio.h>
// #include "gio.go.h"
import "C"
import (
"github.com/terrak/gotk3/glib"
"errors"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.g_application_flags_get_type()), marshalApplicationFlags},

		// Objects/Interfaces
		{glib.Type(C.g_application_get_type()), marshalApplication},
		{glib.Type(C.g_cancellable_get_type()), marshalCancellable},
		{glib.Type(C.g_dbus_connection_get_type()), marshalDBusConnection},
		{glib.Type(C.g_file_get_type()), marshalFile},
		{glib.Type(C.g_notification_get_type()), marshalNotification},
		{glib.Type(C.g_type_module_get_type()), marshalTypeModule},

		// Boxed
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Type conversions
 */

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")
