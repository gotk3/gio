//GCancellable
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
 * GCancellable
 */

// Cancellable is a representation of GIO's GCancellable.
type Cancellable struct {
	*glib.Object
}

// native returns a pointer to the underlying GCancellable.
func (v *Cancellable) native() *C.GCancellable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGCancellable(p)
}

func marshalCancellable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapCancellable(obj), nil
}

func wrapCancellable(obj *glib.Object) *Cancellable {
	return &Cancellable{obj}
}

func (v *Cancellable) toCancellable() *C.GCancellable {
	if v == nil {
		return nil
	}
	return C.toGCancellable(unsafe.Pointer(v.GObject))
}
