//GAction
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
 * GAction
 */

// Action is a representation of GIO's GAction.
type Action struct {
	*glib.Object
}

// native returns a pointer to the underlying GAction.
func (v *Action) native() *C.GAction {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGAction(p)
}

func marshalAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapAction(obj), nil
}

func wrapAction(obj *glib.Object) *Action {
	return &Action{obj}
}

func (v *Action) toAction() *C.GAction {
	if v == nil {
		return nil
	}
	return C.toGAction(unsafe.Pointer(v.GObject))
}

func convertToAction(c *C.GAction) *Action {
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapAction(obj)
}

//gboolean
//g_action_name_is_valid (const gchar *action_name);
//Checks if action_name is valid.
//action_name is valid if it consists only of alphanumeric characters, plus '-' and '.'. The empty string is not a valid action name.
//It is an error to call this function with a non-utf8 action_name . action_name must not be NULL.
func ActionNameIsValid(action_name string) bool {
	cstr := C.CString(action_name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_action_name_is_valid((*C.gchar)(cstr))
	return gobool(c)
}

//const gchar *
//g_action_get_name (GAction *action);
//Queries the name of action .
func (v *Action) GetName() string {
	c := (*C.char)(C.g_action_get_name(v.native()))
	return C.GoString(c)
}

//const GVariantType *
//g_action_get_parameter_type (GAction *action);
//Queries the type of the parameter that must be given when activating action .
//When activating the action using g_action_activate(), the GVariant given to that function must be of the type returned by this function.
//In the case that this function returns NULL, you must not give any GVariant, but NULL instead.
func (v *Action) GetParameterType() *glib.VariantType {
	c := (C.g_action_get_parameter_type(v.native()))
	p := unsafe.Pointer(c)
	return glib.VariantTypeFromUnsafePointer(p)
}

//const GVariantType *
//g_action_get_state_type (GAction *action);
//Queries the type of the state of action .
//If the action is stateful (e.g. created with g_simple_action_new_stateful()) then this function returns the GVariantType of the state. This is the type of the initial value given as the state. All calls to g_action_change_state() must give a GVariant of this type and g_action_get_state() will return a GVariant of the same type.
//If the action is not stateful (e.g. created with g_simple_action_new()) then this function will return NULL. In that case, g_action_get_state() will return NULL and you must not call g_action_change_state().
func (v *Action) GetStateType() *glib.VariantType {
	c := (C.g_action_get_state_type(v.native()))
	p := unsafe.Pointer(c)
	return glib.VariantTypeFromUnsafePointer(p)
}

//GVariant *
//g_action_get_state_hint (GAction *action);
//Requests a hint about the valid range of values for the state of action .
//If NULL is returned it either means that the action is not stateful or that there is no hint about the valid range of values for the state of the action.
//If a GVariant array is returned then each item in the array is a possible value for the state. If a GVariant pair (ie: two-tuple) is returned then the tuple specifies the inclusive lower and upper bound of valid values for the state.
//In any case, the information is merely a hint. It may be possible to have a state value outside of the hinted range and setting a value within the range may fail.
//The return value (if non-NULL) should be freed with g_variant_unref() when it is no longer required.
func (v *Action) GetStateHint() *glib.Variant {
	c := (C.g_action_get_state_hint(v.native()))
	p := unsafe.Pointer(c)
	return glib.VariantFromUnsafePointer(p)
}

//gboolean
//g_action_get_enabled (GAction *action);
//Checks if action is currently enabled.
//An action must be enabled in order to be activated or in order to have its state changed from outside callers.
func (v *Action) GetEnabled() bool {
	c := C.g_action_get_enabled(v.native())
	return gobool(c)
}

//GVariant *
//g_action_get_state (GAction *action);
//Queries the current state of action .
//If the action is not stateful then NULL will be returned. If the action is stateful then the type of the return value is the type given by g_action_get_state_type().
//The return value (if non-NULL) should be freed with g_variant_unref() when it is no longer required.
func (v *Action) GetState() *glib.Variant {
	c := (C.g_action_get_state(v.native()))
	p := unsafe.Pointer(c)
	return glib.VariantFromUnsafePointer(p)
}

//void
//g_action_change_state (GAction *action,
//                       GVariant *value);
//Request for the state of action to be changed to value .
//The action must be stateful and value must be of the correct type. See g_action_get_state_type().
//This call merely requests a change. The action may refuse to change its state or may change its state to something other than value . See g_action_get_state_hint().
//If the value GVariant is floating, it is consumed.
func (v *Action) ChangeState(value *glib.Variant) {
	C.g_action_change_state(v.native(), (*C.GVariant)(value.ToGVariant()))
}

//void
//g_action_activate (GAction *action,
//                   GVariant *parameter);
//Activates the action.
//parameter must be the correct type of parameter for the action (ie: the parameter type given at construction time). If the parameter type was NULL then parameter must also be NULL.
//If the parameter GVariant is floating, it is consumed.
func (v *Action) Activate(parameter *glib.Variant) {
	C.g_action_activate(v.native(), (*C.GVariant)(parameter.ToGVariant()))
}

//gboolean
//g_action_parse_detailed_name (const gchar *detailed_name,
//                              gchar **action_name,
//                              GVariant **target_value,
//                              GError **error);
//Parses a detailed action name into its separate name and target components.
//Detailed action names can have three formats.
//The first format is used to represent an action name with no target value and consists of just an action name containing no whitespace nor the characters ':', '(' or ')'. For example: "app.action".
//The second format is used to represent an action with a target value that is a non-empty string consisting only of alphanumerics, plus '-' and '.'. In that case, the action name and target value are separated by a double colon ("::"). For example: "app.action::target".
//The third format is used to represent an action with any type of target value, including strings. The target value follows the action name, surrounded in parens. For example: "app.action(42)". The target value is parsed using g_variant_parse(). If a tuple-typed value is desired, it must be specified in the same way, resulting in two sets of parens, for example: "app.action((1,2,3))". A string target can be specified this way as well: "app.action('target')". For strings, this third format must be used if * target value is empty or contains characters other than alphanumerics, '-' and '.'.

//gchar *
//g_action_print_detailed_name (const gchar *action_name,
//                              GVariant *target_value);
//Formats a detailed action name from action_name and target_value .
//It is an error to call this function with an invalid action name.
//This function is the opposite of g_action_parse_detailed_action_name(). It will produce a string that can be parsed back to the action_name and target_value by that function.
//See that function for the types of strings that will be printed by this function.
func (v *Action) ActionPrintDetailedName(action_name string, target_value *glib.Variant) string {
	cstr := C.CString(action_name)
	defer C.free(unsafe.Pointer(cstr))
	cstrRes := (*C.char)(C.g_action_print_detailed_name((*C.gchar)(cstr), (*C.GVariant)(target_value.ToGVariant())))
	return C.GoString(cstrRes)
}
