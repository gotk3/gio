//GTypeModule
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
 * GTypeModule
 */

// TypeModule is a representation of GIO's GTypeModule.
type TypeModule struct {
	*glib.Object
}

// native returns a pointer to the underlying GTypeModule.
func (v *TypeModule) native() *C.GTypeModule {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGTypeModule(p)
}

func marshalTypeModule(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTypeModule(obj), nil
}

func wrapTypeModule(obj *glib.Object) *TypeModule {
	return &TypeModule{obj}
}

func (v *TypeModule) toTypeModule() *C.GTypeModule {
	if v == nil {
		return nil
	}
	return C.toGTypeModule(unsafe.Pointer(v.GObject))
}
