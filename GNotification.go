//GNotification
package gio

// #cgo pkg-config: gio-2.0 glib-2.0
// #include <gio/gio.h>
// #include "gio.go.h"
import "C"

import (
	"unsafe"

	"github.com/terrak/gotk3/glib"
)


/*
 * GNotification
 */

// Notification is a representation of GIO's GNotification.
type Notification struct {
	*glib.Object
}

// native returns a pointer to the underlying GNotification.
func (v *Notification) native() *C.GNotification {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGNotification(p)
}

func marshalNotification(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapNotification(obj), nil
}

func wrapNotification(obj *glib.Object) *Notification {
	return &Notification{obj}
}

func (v *Notification) toNotification() *C.GNotification {
	if v == nil {
		return nil
	}
	return C.toGNotification(unsafe.Pointer(v.GObject))
}
