//GMenuModel : GMenuModel — An abstract class representing the contents of a menu
package gio

// #cgo pkg-config: gio-2.0 glib-2.0
// #include <gio/gio.h>
// #include "gio.go.h"
import "C"

import (
	"unsafe"

	"github.com/terrak/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums

		// Objects/Interfaces

		{glib.Type(C.g_menu_model_get_type()), marshalMenuModel},
		{glib.Type(C.g_menu_link_iter_get_type()), marshalMenuLinkIter},
		{glib.Type(C.g_menu_attribute_iter_get_type()), marshalMenuAttributeIter},
		
		// Boxed

	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * GMenuModel
 */

type MenuModel struct {
	*glib.Object
}

// native returns a pointer to the underlying GMenuModel.
func (v *MenuModel) native() *C.GMenuModel {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGMenuModel(p)
}

func marshalMenuModel(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuModel(obj), nil
}

func wrapMenuModel(obj *glib.Object) *MenuModel {
	return &MenuModel{obj}
}

func (v *MenuModel) ToGMenuModel() *C.GMenuModel {
	return v.native()
}

// IsMutable is a wrapper around g_menu_model_is_mutable().
/*
Queries if model is mutable.
An immutable GMenuModel will never emit the “items-changed” signal. Consumers of the model may make optimisations accordingly.
*/
func (v *MenuModel) IsMutable() bool {
	c := C.g_menu_model_is_mutable(v.native())
	return gobool(c)
}

// GetNItems is a wrapper around g_menu_model_get_n_items().
/*
Query the number of items in model .
*/
func (v *MenuModel) GetNItems() int {
	c := C.g_menu_model_get_n_items(v.native())
	return int(c)
}

// GetItemAttributeValue is a wrapper around g_menu_model_get_item_attribute_value().
/*
Queries the item at position item_index in model for the attribute specified by attribute .
If expected_type is non-NULL then it specifies the expected type of the attribute. If it is NULL then any type will be accepted.
If the attribute exists and matches expected_type (or if the expected type is unspecified) then the value is returned.
If the attribute does not exist, or does not match the expected type then NULL is returned.
*/
/*
func (v *MenuModel) GetItemAttributeValue(item_index int, attribute string, expected_type interface{}) interface{} { //FIXME Variant
	cstrattribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstrattribute))
	c := C.g_menu_model_get_item_attribute_value(v.native(), (C.gint)(item_index), (*C.gchar)(cstrattribute), expected_type.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapVariant(obj)
}
*/

// GetItemAttribute is a wrapper around g_menu_model_get_item_attribute().
/*
Queries item at position item_index in model for the attribute specified by attribute .
If the attribute exists and matches the GVariantType corresponding to format_string then format_string is used to deconstruct the value into the positional parameters and TRUE is returned.
If the attribute does not exist, or it does exist but has the wrong type, then the positional parameters are ignored and FALSE is returned.
This function is a mix of g_menu_model_get_item_attribute_value() and g_variant_get(), followed by a g_variant_unref(). As such, format_string must make a complete copy of the data (since the GVariant may go away after the call to g_variant_unref()). In particular, no '&' characters are allowed in format_string .
*/
/*
func (v *MenuModel) GetItemAttribute(item_index int, attribute, format string ,a ...interface{}) bool { //FIXME ... vargs
	cstrattribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstrattribute))
	cstrformat := C.CString(format)
	defer C.free(unsafe.Pointer(cstrformat))
	s := fmt.Sprintf(format, a...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_menu_model_get_item_attribute(v.native(), (C.gint)(item_index), (*C.gchar)(cstrattribute), (*C.gchar)(cstrformat) ,cstr)
	return gobool(c)
}
*/

// GetItemLink is a wrapper around g_menu_model_get_item_link().
/*
Queries the item at position item_index in model for the link specified by link .
If the link exists, the linked GMenuModel is returned. If the link does not exist, NULL is returned.
*/
func (v *MenuModel) GetItemLink(item_index int, link string) *MenuModel {
	cstrlink := C.CString(link)
	defer C.free(unsafe.Pointer(cstrlink))
	c := C.g_menu_model_get_item_link(v.native(), (C.gint)(item_index), (*C.gchar)(cstrlink))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuModel(obj)
}

// IterateItemAttributes is a wrapper around g_menu_model_iterate_item_attributes().
/*
Creates a GMenuAttributeIter to iterate over the attributes of the item at position item_index in model .
You must free the iterator with g_object_unref() when you are done.
*/
func (v *MenuModel) IterateItemAttributes(item_index int) *MenuAttributeIter {
	c := C.g_menu_model_iterate_item_attributes(v.native(), (C.gint)(item_index))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuAttributeIter(obj)
}

// IterateItemLinks is a wrapper around g_menu_model_iterate_item_links().
/*
Creates a GMenuLinkIter to iterate over the links of the item at position item_index in model .
You must free the iterator with g_object_unref() when you are done.
*/
func (v *MenuModel) IterateItemLinks(item_index int) *MenuLinkIter {
	c := C.g_menu_model_iterate_item_links(v.native(), (C.gint)(item_index))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuLinkIter(obj)
}

// ItemsChanged is a wrapper around g_menu_model_items_changed().
/*
Requests emission of the “items-changed” signal on model .
This function should never be called except by GMenuModel subclasses. Any other calls to this function will very likely lead to a violation of the interface of the model.
The implementation should update its internal representation of the menu before emitting the signal. The implementation should further expect to receive queries about the new state of the menu (and particularly added menu items) while signal handlers are running.
The implementation must dispatch this call directly from a mainloop entry and not in response to calls -- particularly those from the GMenuModel API. Said another way: the menu must not change while user code is running without returning to the mainloop.
*/
func (v *MenuModel) ItemsChanged(position, removed, added int) {
	C.g_menu_model_items_changed(v.native(), (C.gint)(position), (C.gint)(removed), (C.gint)(added))
}

/*
 * GMenuAttributeIter
 */

type MenuAttributeIter struct {
	*glib.Object
}

// native returns a pointer to the underlying GMenuModel.
func (v *MenuAttributeIter) native() *C.GMenuAttributeIter {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGMenuAttributeIter(p)
}

func marshalMenuAttributeIter(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuAttributeIter(obj), nil
}

func wrapMenuAttributeIter(obj *glib.Object) *MenuAttributeIter {
	return &MenuAttributeIter{obj}
}

func (v *MenuAttributeIter) toMenuAttributeIter() *C.GMenuAttributeIter {
	if v == nil {
		return nil
	}
	return C.toGMenuAttributeIter(unsafe.Pointer(v.GObject))
}

// GetNext is a wrapper around g_menu_attribute_iter_get_next().
/*
This function combines g_menu_attribute_iter_next() with g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value().
First the iterator is advanced to the next (possibly first) attribute. If that fails, then FALSE is returned and there are no other effects.
If successful, name and value are set to the name and value of the attribute that has just been advanced to. At this point, g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value() will return the same values again.
The value returned in name remains valid for as long as the iterator remains at the current position. The value returned in value must be unreffed using g_variant_unref() when it is no longer in use.
*/
/*
func (v *MenuAttributeIter) GetNext() ([]string, []interface{}, bool) { //FIXME Variant
	c := C.g_menu_attribute_iter_get_next(v.native(), out_name, value)
	return out_name, value, gobool(c)
}
*/
// GetName is a wrapper around g_menu_attribute_iter_get_name().
/*
Gets the name of the attribute at the current iterator position, as a string.
The iterator is not advanced.
*/
func (v *MenuAttributeIter) GetName() string {
	cstr := (*C.char)(C.g_menu_attribute_iter_get_name(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// GetValue is a wrapper around g_menu_attribute_iter_get_value().
/*
Gets the value of the attribute at the current iterator position.
The iterator is not advanced.
*/
/*
func (v *MenuAttributeIter) GetValue() interface{} { //FIXME Variant
	c := C.g_menu_attribute_iter_get_value(v.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapVariant(obj)
}
*/

// IterNext is a wrapper around g_menu_attribute_iter_next().
/*
Attempts to advance the iterator to the next (possibly first) attribute.
TRUE is returned on success, or FALSE if there are no more attributes.
You must call this function when you first acquire the iterator to advance it to the first attribute (and determine if the first attribute exists at all).
*/
func (v *MenuAttributeIter) IterNext() bool {
	c := C.g_menu_attribute_iter_next(v.native())
	return gobool(c)
}

/*
 * GMenuLinkIterMenuAttributeIter
 */

type MenuLinkIter struct {
	*glib.Object
}

// native returns a pointer to the underlying GMenuModel.
func (v *MenuLinkIter) native() *C.GMenuLinkIter {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGMenuLinkIter(p)
}

func marshalMenuLinkIter(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenuLinkIter(obj), nil
}

func wrapMenuLinkIter(obj *glib.Object) *MenuLinkIter {
	return &MenuLinkIter{obj}
}

func (v *MenuLinkIter) toMenuLinkIter() *C.GMenuLinkIter {
	if v == nil {
		return nil
	}
	return C.toGMenuLinkIter(unsafe.Pointer(v.GObject))
}

// GetNext is a wrapper around g_menu_link_iter_get_next().
/*
This function combines g_menu_link_iter_next() with g_menu_link_iter_get_name() and g_menu_link_iter_get_value().
First the iterator is advanced to the next (possibly first) link. If that fails, then FALSE is returned and there are no other effects.
If successful, out_link and value are set to the name and GMenuModel of the link that has just been advanced to. At this point, g_menu_link_iter_get_name() and g_menu_link_iter_get_value() will return the same values again.
The value returned in out_link remains valid for as long as the iterator remains at the current position. The value returned in value must be unreffed using g_object_unref() when it is no longer in use.
*/
/*
func (v *MenuLinkIter) GetNext() ([]string, []interface{}, bool) { //FIXME Variant
	c := C.g_menu_link_iter_get_next(v.native(), out_name, value)
	return out_name, value, gobool(c)
}
*/
// GetName is a wrapper around g_menu_link_iter_get_name().
/*
Gets the name of the link at the current iterator position.
The iterator is not advanced.
*/
func (v *MenuLinkIter) GetName() string {
	cstr := (*C.char)(C.g_menu_link_iter_get_name(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// GetValue is a wrapper around g_menu_link_iter_get_value().
/*
Gets the linked GMenuModel at the current iterator position.
The iterator is not advanced.
*/
/*
func (v *MenuLinkIter) GetValue() interface{} { //FIXME Variant
	c := C.g_menu_link_iter_get_value(v.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapVariant(obj)
}
*/

// IterNext is a wrapper around g_menu_link_iter_next().
/*
Attempts to advance the iterator to the next (possibly first) link.
TRUE is returned on success, or FALSE if there are no more links.
You must call this function when you first acquire the iterator to advance it to the first link (and determine if the first link exists at all).
*/
func (v *MenuLinkIter) IterNext() bool {
	c := C.g_menu_link_iter_next(v.native())
	return gobool(c)
}
