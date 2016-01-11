//GSimpleAction — Interface for Action containers
package gio

// #cgo pkg-config: gio-2.0 glib-2.0
// #include <gio/gio.h>
// #include "gio.go.h"
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums

		// Objects/Interfaces

		{glib.Type(C.g_simple_action_get_type()), marshalSimpleAction},

		// Boxed

	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * GSimpleAction
 */

// SimpleAction is a representation of GIO's GSimpleAction.
type SimpleAction struct {
	*glib.Object
}

// native returns a pointer to the underlying GSimpleAction.
func (v *SimpleAction) native() *C.GSimpleAction {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGSimpleAction(p)
}

func marshalSimpleAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSimpleAction(obj), nil
}

func wrapSimpleAction(obj *glib.Object) *SimpleAction {
	return &SimpleAction{obj}
}

func (v *SimpleAction) toSimpleAction() *C.GSimpleAction {
	if v == nil {
		return nil
	}
	return C.toGSimpleAction(unsafe.Pointer(v.GObject))
}

func convertToSimpleAction(c *C.GSimpleAction) *SimpleAction {
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSimpleAction(obj)
}

//GSimpleAction *
//g_simple_action_new (const gchar *name,
//                     const GVariantType *parameter_type);
//Creates a new action.
//The created action is stateless. See g_simple_action_new_stateful().
func SimpleActionNew(name string, parameter_type *glib.VariantType) *SimpleAction {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_simple_action_new((*C.gchar)(cstr), (*C.GVariantType)(parameter_type.GVariantType))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSimpleAction(obj)
}

//g_simple_action_new_stateful ()
//GSimpleAction *
//g_simple_action_new_stateful (const gchar *name,
//                              const GVariantType *parameter_type,
//                              GVariant *state);
//Creates a new stateful action.
//state is the initial state of the action. All future state values must have the same GVariantType as the initial state.
//If the state GVariant is floating, it is consumed.
func SimpleActionNewStateful(name string, parameter_type *glib.VariantType, state *glib.Variant) *SimpleAction {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_simple_action_new_stateful((*C.gchar)(cstr), (*C.GVariantType)(parameter_type.GVariantType), (*C.GVariant)(state.GVariant))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSimpleAction(obj)
}

//void
//g_simple_action_set_enabled (GSimpleAction *simple,
//                             gboolean enabled);
//Sets the action as enabled or not.
//An action must be enabled in order to be activated or in order to have its state changed from outside callers.
//This should only be called by the implementor of the action. Users of the action should not attempt to modify its enabled flag.
func (v *SimpleAction) SetEnabled(enabled bool) {
	b := gbool(enabled)
	C.g_simple_action_set_enabled(v.native(), b)
}

//void
//g_simple_action_set_state (GSimpleAction *simple,
//                           GVariant *value);
//Sets the state of the action.
//This directly updates the 'state' property to the given value.
//This should only be called by the implementor of the action. Users of the action should not attempt to directly modify the 'state' property. Instead, they should call g_action_change_state() to request the change.
//If the value GVariant is floating, it is consumed.
func (v *SimpleAction) SetState(value *glib.Variant) {
	C.g_simple_action_set_state(v.native(), (*C.GVariant)(value.GVariant))
}

//void
//g_simple_action_set_state_hint (GSimpleAction *simple,
//                                GVariant *state_hint);
//Sets the state hint for the action.
//See g_action_get_state_hint() for more information about action state hints.
func (v *SimpleAction) SetStateHint(value *glib.Variant) {
	C.g_simple_action_set_state_hint(v.native(), (*C.GVariant)(value.GVariant))
}
