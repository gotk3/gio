//GFile
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
 * GFile
 */

// File is a representation of GIO's GFile.
type File struct {
	*glib.Object
}

// native returns a pointer to the underlying GFile.
func (v *File) native() *C.GFile {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFile(p)
}

func marshalFile(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapFile(obj), nil
}

func wrapFile(obj *glib.Object) *File {
	return &File{obj}
}

func (v *File) toFile() *C.GFile {
	if v == nil {
		return nil
	}
	return C.toGFile(unsafe.Pointer(v.GObject))
}

func convertToFile(c *C.GFile) *File {
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapFile(obj)
}

//void	(*GFileProgressCallback) ()
//gboolean	(*GFileReadMoreCallback) ()
//void	(*GFileMeasureProgressCallback) ()
//GFile *	g_file_new_for_path ()
//GFile *	g_file_new_for_uri ()
//GFile *	g_file_new_for_commandline_arg ()
//GFile *	g_file_new_for_commandline_arg_and_cwd ()
//GFile *	g_file_new_tmp ()
//GFile *	g_file_parse_name ()
//GFile *	g_file_dup ()
//guint	g_file_hash ()
//gboolean	g_file_equal ()

//char *
//g_file_get_basename (GFile *file);
//Gets the base name (the last component of the path) for a given GFile.
//If called for the top level of a system (such as the filesystem root or a uri like sftp://host/) it will return a single directory separator (and on Windows, possibly a drive letter).
//The base name is a byte string (not UTF-8). It has no defined encoding or rules other than it may not contain zero bytes. If you want to use filenames in a user interface you should use the display name that you can get by requesting the G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME attribute with g_file_query_info().
//This call does no blocking I/O.
func (v *File) GetBasename() string {
	if v == nil {
		return ""
	}
	cstr := (*C.char)(C.g_file_get_basename(v.native()))
	return C.GoString(cstr)
}

//char *
//g_file_get_path (GFile *file);
//Gets the local pathname for GFile, if one exists. If non-NULL, this is guaranteed to be an absolute, canonical path. It might contain symlinks.
//This call does no blocking I/O.
func (v *File) GetPath() string {
	if v == nil {
		return ""
	}
	cstr := (*C.char)(C.g_file_get_path(v.native()))
	return C.GoString(cstr)
}

//char *
//g_file_get_uri (GFile *file);
//Gets the URI for the file .
//This call does no blocking I/O.
func (v *File) GetUri() string {
	if v == nil {
		return ""
	}
	cstr := (*C.char)(C.g_file_get_uri(v.native()))
	return C.GoString(cstr)
}

//char *
//g_file_get_parse_name (GFile *file);
//Gets the parse name of the file . A parse name is a UTF-8 string that describes the file such that one can get the GFile back using g_file_parse_name().
//This is generally used to show the GFile as a nice full-pathname kind of string in a user interface, like in a location entry.
//For local files with names that can safely be converted to UTF-8 the pathname is used, otherwise the IRI is used (a form of URI that allows UTF-8 characters unescaped).
//This call does no blocking I/O.
func (v *File) GetParseName() string {
	if v == nil {
		return ""
	}
	cstr := (*C.char)(C.g_file_get_parse_name(v.native()))
	return C.GoString(cstr)
}

//GFile *	g_file_get_parent ()
//gboolean	g_file_has_parent ()
//GFile *	g_file_get_child ()
//GFile *	g_file_get_child_for_display_name ()
//gboolean	g_file_has_prefix ()
//char *	g_file_get_relative_path ()
//GFile *	g_file_resolve_relative_path ()
//gboolean	g_file_is_native ()
//gboolean	g_file_has_uri_scheme ()
//char *	g_file_get_uri_scheme ()
//GFileInputStream *	g_file_read ()
//void	g_file_read_async ()
//GFileInputStream *	g_file_read_finish ()
//GFileOutputStream *	g_file_append_to ()
//GFileOutputStream *	g_file_create ()
//GFileOutputStream *	g_file_replace ()
//void	g_file_append_to_async ()
//GFileOutputStream *	g_file_append_to_finish ()
//void	g_file_create_async ()
//GFileOutputStream *	g_file_create_finish ()
//void	g_file_replace_async ()
//GFileOutputStream *	g_file_replace_finish ()
//GFileInfo *	g_file_query_info ()
//void	g_file_query_info_async ()
//GFileInfo *	g_file_query_info_finish ()
//gboolean	g_file_query_exists ()
//GFileType	g_file_query_file_type ()
//GFileInfo *	g_file_query_filesystem_info ()
//void	g_file_query_filesystem_info_async ()
//GFileInfo *	g_file_query_filesystem_info_finish ()
//GAppInfo *	g_file_query_default_handler ()
//gboolean	g_file_measure_disk_usage ()
//void	g_file_measure_disk_usage_async ()
//gboolean	g_file_measure_disk_usage_finish ()
//GMount *	g_file_find_enclosing_mount ()
//void	g_file_find_enclosing_mount_async ()
//GMount *	g_file_find_enclosing_mount_finish ()
//GFileEnumerator *	g_file_enumerate_children ()
//void	g_file_enumerate_children_async ()
//GFileEnumerator *	g_file_enumerate_children_finish ()
//GFile *	g_file_set_display_name ()
//void	g_file_set_display_name_async ()
//GFile *	g_file_set_display_name_finish ()
//gboolean	g_file_delete ()
//void	g_file_delete_async ()
//gboolean	g_file_delete_finish ()
//gboolean	g_file_trash ()
//void	g_file_trash_async ()
//gboolean	g_file_trash_finish ()
//gboolean	g_file_copy ()
//void	g_file_copy_async ()
//gboolean	g_file_copy_finish ()
//gboolean	g_file_move ()
//gboolean	g_file_make_directory ()
//void	g_file_make_directory_async ()
//gboolean	g_file_make_directory_finish ()
//gboolean	g_file_make_directory_with_parents ()
//gboolean	g_file_make_symbolic_link ()
//GFileAttributeInfoList *	g_file_query_settable_attributes ()
//GFileAttributeInfoList *	g_file_query_writable_namespaces ()
//gboolean	g_file_set_attribute ()
//gboolean	g_file_set_attributes_from_info ()
//void	g_file_set_attributes_async ()
//gboolean	g_file_set_attributes_finish ()
//gboolean	g_file_set_attribute_string ()
//gboolean	g_file_set_attribute_byte_string ()
//gboolean	g_file_set_attribute_uint32 ()
//gboolean	g_file_set_attribute_int32 ()
//gboolean	g_file_set_attribute_uint64 ()
//gboolean	g_file_set_attribute_int64 ()
//void	g_file_mount_mountable ()
//GFile *	g_file_mount_mountable_finish ()
//void	g_file_unmount_mountable ()
//gboolean	g_file_unmount_mountable_finish ()
//void	g_file_unmount_mountable_with_operation ()
//gboolean	g_file_unmount_mountable_with_operation_finish ()
//void	g_file_eject_mountable ()
//gboolean	g_file_eject_mountable_finish ()
//void	g_file_eject_mountable_with_operation ()
//gboolean	g_file_eject_mountable_with_operation_finish ()
//void	g_file_start_mountable ()
//gboolean	g_file_start_mountable_finish ()
//void	g_file_stop_mountable ()
//gboolean	g_file_stop_mountable_finish ()
//void	g_file_poll_mountable ()
//gboolean	g_file_poll_mountable_finish ()
//void	g_file_mount_enclosing_volume ()
//gboolean	g_file_mount_enclosing_volume_finish ()
//GFileMonitor *	g_file_monitor_directory ()
//GFileMonitor *	g_file_monitor_file ()
//GFileMonitor *	g_file_monitor ()

//gboolean
//g_file_load_contents (GFile *file,
//                      GCancellable *cancellable,
//                      char **contents,
//                      gsize *length,
//                      char **etag_out,
//                      GError **error);

//Loads the content of the file into memory. The data is always zero-terminated, but this is not included in the resultant length . The returned content should be freed with g_free() when no longer needed.
//If cancellable is not NULL, then the operation can be cancelled by triggering the cancellable object from another thread. If the operation was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (v *File) LoadContents(cancellable *Cancellable) (text string, ok bool) {
	contents := new(*C.char)
	length := new(C.gsize)
	c := C.g_file_load_contents(v.native(), cancellable.native(), contents, length, nil, nil)

	ok = gobool(c)
	text = C.GoStringN((*C.char)(*contents), (C.int)(*length))
	C.g_free((C.gpointer)(contents))
	return
}

//void	g_file_load_contents_async ()
//gboolean	g_file_load_contents_finish ()
//void	g_file_load_partial_contents_async ()
//gboolean	g_file_load_partial_contents_finish ()
//gboolean	g_file_replace_contents ()
//void	g_file_replace_contents_async ()
//void	g_file_replace_contents_bytes_async ()
//gboolean	g_file_replace_contents_finish ()
//gboolean	g_file_copy_attributes ()
//GFileIOStream *	g_file_create_readwrite ()
//void	g_file_create_readwrite_async ()
//GFileIOStream *	g_file_create_readwrite_finish ()
//GFileIOStream *	g_file_open_readwrite ()
//void	g_file_open_readwrite_async ()
//GFileIOStream *	g_file_open_readwrite_finish ()
//GFileIOStream *	g_file_replace_readwrite ()
//void	g_file_replace_readwrite_async ()
//GFileIOStream *	g_file_replace_readwrite_finish ()
//gboolean	g_file_supports_thread_contexts ()
