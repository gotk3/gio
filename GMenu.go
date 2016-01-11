//GMenu : GMenu — A simple implementation of GMenuModel
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

		{glib.Type(C.g_menu_get_type()), marshalMenu},

		// Boxed

	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * GMenu
 */

// Menu is a representation of GIO's GMenu.
type Menu struct {
	MenuModel
}

// native returns a pointer to the underlying GMenu.
func (v *Menu) native() *C.GMenu {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGMenu(p)
}

func marshalMenu(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenu(obj), nil
}

func wrapMenu(obj *glib.Object) *Menu {
	return &Menu{MenuModel{obj}}
}

func (v *Menu) toMenu() *C.GMenu {
	if v == nil {
		return nil
	}
	return C.toGMenu(unsafe.Pointer(v.GObject))
}

func convertToMenu(c *C.GMenu) *Menu {
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMenu(obj)
}

//GMenu *	g_menu_new ()
//void	g_menu_freeze ()
//void	g_menu_insert ()
//void	g_menu_prepend ()
//void	g_menu_append ()
//void	g_menu_insert_item ()
//void	g_menu_append_item ()
//void	g_menu_prepend_item ()
//void	g_menu_insert_section ()
//void	g_menu_prepend_section ()
//void	g_menu_append_section ()
//void	g_menu_append_submenu ()
//void	g_menu_insert_submenu ()
//void	g_menu_prepend_submenu ()
//void	g_menu_remove ()
//void	g_menu_remove_all ()
//GMenuItem *	g_menu_item_new ()
//GMenuItem *	g_menu_item_new_section ()
//GMenuItem *	g_menu_item_new_submenu ()
//GMenuItem *	g_menu_item_new_from_model ()
//void	g_menu_item_set_label ()
//void	g_menu_item_set_icon ()
//void	g_menu_item_set_action_and_target_value ()
//void	g_menu_item_set_action_and_target ()
//void	g_menu_item_set_detailed_action ()
//void	g_menu_item_set_section ()
//void	g_menu_item_set_submenu ()
//GVariant *	g_menu_item_get_attribute_value ()
//gboolean	g_menu_item_get_attribute ()
//GMenuModel *	g_menu_item_get_link ()
//void	g_menu_item_set_attribute_value ()
//void	g_menu_item_set_attribute ()
//void	g_menu_item_set_link ()
