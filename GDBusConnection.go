//GDBusConnection
package gio

// #cgo pkg-config: gio-2.0 glib-2.0
// #include <gio/gio.h>
// #include "gio.go.h"
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GDBusConnection
 */

// DBusConnection is a representation of GIO's GDBusConnection.
type DBusConnection struct {
	*glib.Object
}

// native returns a pointer to the underlying GDBusConnection.
func (v *DBusConnection) native() *C.GDBusConnection {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGDBusConnection(p)
}

func marshalDBusConnection(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapDBusConnection(obj), nil
}

func wrapDBusConnection(obj *glib.Object) *DBusConnection {
	return &DBusConnection{obj}
}

func (v *DBusConnection) toDBusConnection() *C.GDBusConnection {
	if v == nil {
		return nil
	}
	return C.toGDBusConnection(unsafe.Pointer(v.GObject))
}
