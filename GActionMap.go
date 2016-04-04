//GActionMap — Interface for Action containers

// +build
package gio_2_32

// #cgo pkg-config: gio-2.0 glib-2.0
// #include <gio/gio.h>
// #include "gio.go.h"
import "C"

import (
	"log"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums

		// Objects/Interfaces

		{glib.Type(C.g_action_map_get_type()), marshalActionMap},

		// Boxed

	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * GActionMap
 */

// ActionMap is a representation of GIO's GActionMap.
type ActionMap struct {
	*glib.Object
}

// native returns a pointer to the underlying GActionMap.
func (v *ActionMap) native() *C.GActionMap {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGActionMap(p)
}

func marshalActionMap(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapActionMap(obj), nil
}

func wrapActionMap(obj *glib.Object) *ActionMap {
	return &ActionMap{obj}
}

func (v *ActionMap) toActionMap() *C.GActionMap {
	if v == nil {
		return nil
	}
	return C.toGActionMap(unsafe.Pointer(v.GObject))
}

func convertToActionMap(c *C.GActionMap) *ActionMap {
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapActionMap(obj)
}

//GAction *
//g_action_map_lookup_action (GActionMap *action_map,
//                            const gchar *action_name);
//Looks up the action with the name action_name in action_map .
//If no such action exists, returns NULL.
func (v *ActionMap) LookAction(action_name string) *Action {
	cstr := C.CString(action_name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_action_map_lookup_action(v.native(), (*C.gchar)(cstr))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapAction(obj)
}

//g_action_map_add_action_entries ()
//void
//g_action_map_add_action_entries (GActionMap *action_map,
//                                 const GActionEntry *entries,
//                                 gint n_entries,
//                                 gpointer user_data);
//A convenience function for creating multiple GSimpleAction instances and adding them to a GActionMap.
//Each action is constructed as per one GActionEntry.
func (v *ActionMap) AddActionEntries(entries []ActionEntry, user_data interface{}) {
	gEntries := make([]C.GActionEntry, len(entries))
	for i := 0; i < len(entries); i++ {
		gEntries[i].name = (*C.gchar)(C.CString(entries[i].Name))
		defer C.free(unsafe.Pointer(gEntries[i].name))
		if len(entries[i].ParameterType) == 0 {
			gEntries[i].parameter_type = nil
		} else {
			gEntries[i].parameter_type = (*C.gchar)(C.CString(entries[i].ParameterType))
			defer C.free(unsafe.Pointer(gEntries[i].parameter_type))
		}
		if len(entries[i].State) == 0 {
			gEntries[i].state = nil
		} else {
			gEntries[i].state = (*C.gchar)(C.CString(entries[i].State))
			defer C.free(unsafe.Pointer(gEntries[i].state))
		}
		gEntries[i].activate = nil /*func (a *C.GSimpleAction, p *C.GVariant, u C.gpointer){
			entries[i].Activate(nil,nil,nil)
		} //entries[i].Activate //FIXME*/
		gEntries[i].change_state = nil //entries[i].ChangeState //FIXME

		defer C.free(unsafe.Pointer(gEntries[i].state))
	}
	//FIXME copy entries into gEntries
	C.g_action_map_add_action_entries(v.native(), (*C.GActionEntry)(unsafe.Pointer(&gEntries[0])), (C.gint)(len(entries)), (C.gpointer)(unsafe.Pointer(&user_data)))

	for i := 0; i < len(entries); i++ {
		action := v.LookAction(entries[i].Name)
		if action == nil {
			log.Println(entries[i].Name, " not an action -> not connected")
		} else {
			if entries[i].Activate != nil {
				_, err := action.Connect("activate", entries[i].Activate, nil)
				if err != nil {
					log.Println(entries[i].Name, " connect to activate signal error ", err.Error())
				}
			}
			if entries[i].ChangeState != nil {
				_, err := action.Connect("change_state", entries[i].ChangeState, nil)
				if err != nil {
					log.Println(entries[i].Name, " connect to change_state signal error ", err.Error())
				}
			}
		}
	}

}

//void
//g_action_map_add_action (GActionMap *action_map,
//                         GAction *action);
//Adds an action to the action_map .
//If the action map already contains an action with the same name as action then the old action is dropped from the action map.
//The action map takes its own reference on action .
func (v *ActionMap) AddAction(action *Action) {
	C.g_action_map_add_action(v.native(), action.native())
}

//void
//g_action_map_remove_action (GActionMap *action_map,
//                            const gchar *action_name);
//Removes the named action from the action map.
//If no action of this name is in the map then nothing happens.
func (v *ActionMap) RemoveAction(action_name string) {
	cstr := C.CString(action_name)
	defer C.free(unsafe.Pointer(cstr))
	C.g_action_map_remove_action(v.native(), (*C.gchar)(cstr))
}

/*
 * GActionEntry
 */

type ActionEntry struct {
	Name          string                                              //the name of the action
	Activate      func(action *SimpleAction, parameter *glib.Variant) //the callback to connect to the "activate" signal of the action. Since GLib 2.40, this can be NULL for stateful actions, in which case the default handler is used. For boolean-stated actions with no parameter, this is a toggle. For other state types (and parameter type equal to the state type) this will be a function that just calls change_state (which you should provide).
	ParameterType string                                              //the type of the parameter that must be passed to the activate function for this action, given as a single GVariant type string (or NULL for no parameter)
	State         string                                              //the initial state for this action, given in GVariant text format. The state is parsed with no extra type information, so type tags must be added to the string if they are necessary. Stateless actions should give NULL here.
	ChangeState   func(action *SimpleAction, parameter *glib.Variant) //the callback to connect to the "change-state" signal of the action. All stateful actions should provide a handler here; stateless actions should not.
}
