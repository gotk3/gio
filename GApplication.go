//GApplication — Core application class
package gio

// #cgo pkg-config: gio-2.0 glib-2.0
// #include <gio/gio.h>
// #include "gio.go.h"
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GApplication
 */

// Handler type without params
type OnNoParamHandler func()

// Handler type with files and hint for open signal
type OnApplicationOpenFileHandler func([]*File, string)

// IWidget is an interface type implemented by all structs
// embedding a Widget.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkWidget.
//type IApplication interface {
//	toGApplication() *C.GApplication
//}

// Application is a representation of GIO's GApplication.
type Application struct {
	*glib.Object

	activateHandlers []OnNoParamHandler             //Slice of handlers to call when activate signal appends
	shutdownHandlers []OnNoParamHandler             //Slice of handlers to call when shutdown signal appends
	startupHandlers  []OnNoParamHandler             //Slice of handlers to call when startup signal appends
	openHandlers     []OnApplicationOpenFileHandler //Slice of handlers to call when open signal appends
}

// native returns a pointer to the underlying GApplication.
func (v *Application) native() *C.GApplication {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGApplication(p)
}

func marshalApplication(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapApplication(obj), nil
}

func wrapApplication(obj *glib.Object) *Application {
	return &Application{Object: obj}
}

func (v *Application) ToGApplication() *C.GApplication {
	if v == nil {
		return nil
	}
	return C.toGApplication(unsafe.Pointer(v.GObject))
}

//gboolean
//g_application_id_is_valid (const gchar *application_id);

//Checks if application_id is a valid application identifier.
//A valid ID is required for calls to g_application_new() and g_application_set_application_id().
//For convenience, the restrictions on application identifiers are reproduced here:
//*Application identifiers must contain only the ASCII characters "A-Z[0-9]_-." and must not begin with a digit.
//*Application identifiers must contain at least one '.' (period) character (and thus at least three elements).
//*Application identifiers must not begin or end with a '.' (period) character.
//*Application identifiers must not contain consecutive '.' (period) characters.
//*Application identifiers must not exceed 255 characters.
func ApplicationIdIsValid(application_id string) bool {
	cstr := C.CString(application_id)
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_application_id_is_valid((*C.gchar)(cstr))

	return gobool(c)
}

//GApplication *
//g_application_new (const gchar *application_id,
//                   GApplicationFlags flags);

//Creates a new GApplication instance.
//If non-NULL, the application id must be valid. See g_application_id_is_valid().
//If no application ID is given then some features of GApplication (most notably application uniqueness) will be disabled.
func ApplicationNew(application_id string, flags ApplicationFlags) (*Application, error) {
	cstr := C.CString(application_id)
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_application_new((*C.gchar)(cstr), C.GApplicationFlags(flags))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	app := wrapApplication(obj)
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return app, nil
}

//const gchar *
//g_application_get_application_id (GApplication *application);

//Gets the unique identifier for application .
func (v *Application) GetApplicationId() string {
	cstr := (*C.char)(C.g_application_get_application_id(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

//void
//g_application_set_application_id (GApplication *application,
//                                  const gchar *application_id);

//Sets the unique identifier for application .
//The application id can only be modified if application has not yet been registered.
//If non-NULL, the application id must be valid. See g_application_id_is_valid().
func (v *Application) SetApplicationId(application_id string) {
	cstr := C.CString(application_id)
	defer C.free(unsafe.Pointer(cstr))
	C.g_application_set_application_id(v.native(), (*C.gchar)(cstr))
}

//guint
//g_application_get_inactivity_timeout (GApplication *application);

//Gets the current inactivity timeout for the application.
//This is the amount of time (in milliseconds) after the last call to g_application_release() before the application stops running.
func (v *Application) GetInactivityTimeout() uint {
	c := C.g_application_get_inactivity_timeout(v.native())
	return uint(c)
}

//void
//g_application_set_inactivity_timeout (GApplication *application,
//                                      guint inactivity_timeout);

//Sets the current inactivity timeout for the application.
//This is the amount of time (in milliseconds) after the last call to g_application_release() before the application stops running.
//This call has no side effects of its own. The value set here is only used for next time g_application_release() drops the use count to zero. Any timeouts currently in progress are not impacted.
func (v *Application) SetInactivityTimeout(inactivity_timeout uint) {
	C.g_application_set_inactivity_timeout(v.native(), (C.guint)(inactivity_timeout))
}

//GApplicationFlags
//g_application_get_flags (GApplication *application);

//Gets the flags for application .
//See GApplicationFlags.
func (v *Application) GetFlags() ApplicationFlags {
	c := C.g_application_get_flags(v.native())
	return (ApplicationFlags)(c)
}

//void
//g_application_set_flags (GApplication *application,
//                         GApplicationFlags flags);

//Sets the flags for application .
//The flags can only be modified if application has not yet been registered.
//See GApplicationFlags.
func (v *Application) SetFlags(flags ApplicationFlags) {
	C.g_application_set_flags(v.native(), C.GApplicationFlags(flags))
}

//const gchar *
//g_application_get_resource_base_path (GApplication *application);

//Gets the resource base path of application .
//See g_application_set_resource_base_path() for more information.
func (v *Application) GetResourceBasePath() string {
	cstr := (*C.char)(C.g_application_get_resource_base_path(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

//void
//g_application_set_resource_base_path (GApplication *application,
//                                      const gchar *resource_path);

//Sets (or unsets) the base resource path of application .
//The path is used to automatically load various application resources such as menu layouts and File descriptions. The various types of resources will be found at fixed names relative to the given base path.
//By default, the resource base path is determined from the application ID by prefixing '/' and replacing each '.' with '/'. This is done at the time that the GApplication object is constructed. Changes to the application ID after that point will not have an impact on the resource base path.
//As an example, if the application has an ID of "org.example.app" then the default resource base path will be "/org/example/app". If this is a GtkApplication (and you have not manually changed the path) then Gtk will then search for the menus of the application at "/org/example/app/gtk/menus.ui".
//See GResource for more information about adding resources to your application.
//You can disable automatic resource loading functionality by setting the path to NULL.
//Changing the resource base path once the application is running is not recommended. The point at which the resource path is consulted for forming paths for various purposes is unspecified.
func (v *Application) SetResourceBasePath(resource_path string) {
	cstr := C.CString(resource_path)
	defer C.free(unsafe.Pointer(cstr))
	C.g_application_set_resource_base_path(v.native(), (*C.gchar)(cstr))
}

//GDBusConnection *
//g_application_get_dbus_connection (GApplication *application);

//Gets the GDBusConnection being used by the application, or NULL.
//If GApplication is using its D-Bus backend then this function will return the GDBusConnection being used for uniqueness and communication with the desktop environment and other instances of the application.
//If GApplication is not using D-Bus then this function will return NULL. This includes the situation where the D-Bus backend would normally be in use but we were unable to connect to the bus.
//This function must not be called before the application has been registered. See g_application_get_is_registered().
func (v *Application) GetDBusConnection() *DBusConnection {
	c := C.g_application_get_dbus_connection(v.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapDBusConnection(obj)
}

//const gchar *
//g_application_get_dbus_object_path (GApplication *application);

//Gets the D-Bus object path being used by the application, or NULL.
//If GApplication is using its D-Bus backend then this function will return the D-Bus object path that GApplication is using. If the application is the primary instance then there is an object published at this path. If the application is not the primary instance then the result of this function is undefined.
//If GApplication is not using D-Bus then this function will return NULL. This includes the situation where the D-Bus backend would normally be in use but we were unable to connect to the bus.
//This function must not be called before the application has been registered. See g_application_get_is_registered().
func (v *Application) GetDBusObjectPath() string {
	cstr := (*C.char)(C.g_application_get_dbus_object_path(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

//DEPRECATED	void	g_application_set_File_group ()

//gboolean
//g_application_get_is_registered (GApplication *application);

//Checks if application is registered.
//An application is registered if g_application_register() has been successfully called.
func (v *Application) GetIsRegistered() bool {
	c := C.g_application_get_is_registered(v.native())
	return gobool(c)
}

//gboolean
//g_application_get_is_remote (GApplication *application);

//Checks if application is remote.
//If application is remote then it means that another instance of application already exists (the 'primary' instance). Calls to perform Files on application will result in the Files being performed by the primary instance.
//The value of this property cannot be accessed before g_application_register() has been called. See g_application_get_is_registered().
func (v *Application) GetIsRemote() bool {
	c := C.g_application_get_is_remote(v.native())
	return gobool(c)
}

//gboolean
//g_application_register (GApplication *application,
//                        GCancellable *cancellable,
//                        GError **error);

//Attempts registration of the application.
//This is the point at which the application discovers if it is the primary instance or merely acting as a remote for an already-existing primary instance. This is implemented by attempting to acquire the application identifier as a unique bus name on the session bus using GDBus.
//If there is no application ID or if G_APPLICATION_NON_UNIQUE was given, then this process will always become the primary instance.
//Due to the internal architecture of GDBus, method calls can be dispatched at any time (even if a main loop is not running). For this reason, you must ensure that any object paths that you wish to register are registered before calling this function.
//If the application has already been registered then TRUE is returned with no work performed.
//The “startup” signal is emitted if registration succeeds and application is the primary instance (including the non-unique case).
//In the event of an error (such as cancellable being cancelled, or a failure to connect to the session bus), FALSE is returned and error is set appropriately.
//Note: the return value of this function is not an indicator that this instance is or is not the primary instance of the application. See g_application_get_is_remote() for that.
func (v *Application) Register(cancellable *Cancellable) bool {
	var gerror **C.GError
	c := C.g_application_register(v.native(), cancellable.native(), gerror)
	return gobool(c)
}

//void
//g_application_hold (GApplication *application);

//Increases the use count of application .
//Use this function to indicate that the application has a reason to continue to run. For example, g_application_hold() is called by GTK+ when a toplevel window is on the screen.
//To cancel the hold, call g_application_release().
func (v *Application) Hold() {
	C.g_application_hold(v.native())
}

//void
//g_application_release (GApplication *application);

//Decrease the use count of application .
//When the use count reaches zero, the application will stop running.
//Never call this function except to cancel the effect of a previous call to g_application_hold().
func (v *Application) Release() {
	C.g_application_release(v.native())
}

//void
//g_application_quit (GApplication *application);

//Immediately quits the application.
//Upon return to the mainloop, g_application_run() will return, calling only the 'shutdown' function before doing so.
//The hold count is ignored.
//The result of calling g_application_run() again after it returns is unspecified.
func (v *Application) Quit() {
	C.g_application_quit(v.native())
}

//void
//g_application_activate (GApplication *application);

//Activates the application.
//In essence, this results in the “activate” signal being emitted in the primary instance.
//The application must be registered before calling this function.
func (v *Application) Activate() {
	C.g_application_activate(v.native())
}

//void
//g_application_open (GApplication *application,
//                    GFile **files,
//                    gint n_files,
//                    const gchar *hint);

//Opens the given files.
//In essence, this results in the “open” signal being emitted in the primary instance.
//n_files must be greater than zero.
//hint is simply passed through to the ::open signal. It is intended to be used by applications that have multiple modes for opening files (eg: "view" vs "edit", etc). Unless you have a need for this functionality, you should use "".
//The application must be registered before calling this function and it must have the G_APPLICATION_HANDLES_OPEN flag set.
func (v *Application) Open(files []*File, hint string) {
	gfiles := C.alloc_files(C.int(len(files)))
	for n, val := range files {
		C.set_file(gfiles, C.int(n), val.native())
	}
	defer C.g_free(C.gpointer(gfiles))
	cstr_hint := C.CString(hint)
	defer C.free(unsafe.Pointer(cstr_hint))
	C.g_application_open(v.native(), gfiles, C.gint(len(files)), (*C.gchar)(cstr_hint))
}

//void
//g_application_send_notification (GApplication *application,
//                                 const gchar *id,
//                                 GNotification *notification);

//Sends a notification on behalf of application to the desktop shell. There is no guarantee that the notification is displayed immediately, or even at all.
//Notifications may persist after the application exits. It will be D-Bus-activated when the notification or one of its Files is activated.
//Modifying notification after this call has no effect. However, the object can be reused for a later call to this function.
//id may be any string that uniquely identifies the event for the application. It does not need to be in any special format. For example, "new-message" might be appropriate for a notification about new messages.
//If a previous notification was sent with the same id , it will be replaced with notification and shown again as if it was a new notification. This works even for notifications sent from a previous execution of the application, as long as id is the same string.
//id may be NULL, but it is impossible to replace or withdraw notifications without an id.
//If notification is no longer relevant, it can be withdrawn with g_application_withdraw_notification().
func (v *Application) SendNotification(id string, notification *Notification) {
	cstr := C.CString(id)
	defer C.free(unsafe.Pointer(cstr))
	C.g_application_send_notification(v.native(), (*C.gchar)(cstr), notification.native())
}

//void
//g_application_withdraw_notification (GApplication *application,
//                                     const gchar *id);

//Withdraws a notification that was sent with g_application_send_notification().
//This call does nothing if a notification with id doesn't exist or the notification was never sent.
//This function works even for notifications sent in previous executions of this application, as long id is the same as it was for the sent notification.
//Note that notifications are dismissed when the user clicks on one of the buttons in a notification or triggers its default File, so there is no need to explicitly withdraw the notification in that case.
func (v *Application) WithdrawNotification(id string) {
	cstr := C.CString(id)
	defer C.free(unsafe.Pointer(cstr))
	C.g_application_withdraw_notification(v.native(), (*C.gchar)(cstr))
}

//int
//g_application_run (GApplication *application,
//                   int argc,
//                   char **argv);

//Runs the application.
//This function is intended to be run from main() and its return value is intended to be returned by main(). Although you are expected to pass the argc , argv parameters from main() to this function, it is possible to pass NULL if argv is not available or commandline handling is not required. Note that on Windows, argc and argv are ignored, and g_win32_get_command_line() is called internally (for proper support of Unicode commandline arguments).
//GApplication will attempt to parse the commandline arguments. You can add commandline flags to the list of recognised options by way of g_application_add_main_option_entries(). After this, the “handle-local-options” signal is emitted, from which the application can inspect the values of its GOptionEntrys.
//“handle-local-options” is a good place to handle options such as --version, where an immediate reply from the local process is desired (instead of communicating with an already-running instance). A “handle-local-options” handler can stop further processing by returning a non-negative value, which then becomes the exit status of the process.
//What happens next depends on the flags: if G_APPLICATION_HANDLES_COMMAND_LINE was specified then the remaining commandline arguments are sent to the primary instance, where a “command-line” signal is emitted. Otherwise, the remaining commandline arguments are assumed to be a list of files. If there are no files listed, the application is activated via the “activate” signal. If there are one or more files, and G_APPLICATION_HANDLES_OPEN was specified then the files are opened via the “open” signal.
//If you are interested in doing more complicated local handling of the commandline then you should implement your own GApplication subclass and override local_command_line(). In this case, you most likely want to return TRUE from your local_command_line() implementation to suppress the default handling. See gapplication-example-cmdline2.c for an example.
//If, after the above is done, the use count of the application is zero then the exit status is returned immediately. If the use count is non-zero then the default main context is iterated until the use count falls to zero, at which point 0 is returned.
//If the G_APPLICATION_IS_SERVICE flag is set, then the service will run for as much as 10 seconds with a use count of zero while waiting for the message that caused the activation to arrive. After that, if the use count falls to zero the application will exit immediately, except in the case that g_application_set_inactivity_timeout() is in use.
//This function sets the prgname (g_set_prgname()), if not already set, to the basename of argv[0].
//Since 2.40, applications that are not explicitly flagged as services or launchers (ie: neither G_APPLICATION_IS_SERVICE or G_APPLICATION_IS_LAUNCHER are given as flags) will check (from the default handler for local_command_line) if "--gapplication-service" was given in the command line. If this flag is present then normal commandline processing is interrupted and the G_APPLICATION_IS_SERVICE flag is set. This provides a "compromise" solution whereby running an application directly from the commandline will invoke it in the normal way (which can be useful for debugging) while still allowing applications to be D-Bus activated in service mode. The D-Bus service file should invoke the executable with "--gapplication-service" as the sole commandline argument. This approach is suitable for use by most graphical applications but should not be used from applications like editors that need precise control over when processes invoked via the commandline will exit and what their exit status will be.
func (v *Application) Run(args []string) (ret int) {
	if args != nil {
		argc := C.int(len(args))
		argv := make([](*C.char), argc)
		for i, arg := range args {
			argv[i] = C.CString(arg)
			defer C.free(unsafe.Pointer(argv[i]))
		}
		ret = int(C.g_application_run(v.native(), argc, (**C.char)(unsafe.Pointer(&argv[0]))))
	} else {
		ret = int(C.g_application_run(v.native(), 0, nil))
	}
	return
}

//void	g_application_add_main_option_entries ()
//void	g_application_add_main_option ()
//void	g_application_add_option_group ()
//void	g_application_set_default ()
//GApplication *	g_application_get_default ()
//void	g_application_mark_busy ()
//void	g_application_unmark_busy ()
//gboolean	g_application_get_is_busy ()
//void	g_application_bind_busy_property ()
//void	g_application_unbind_busy_property ()

//---- EVENTS ----//

//The ::activate signal is emitted on the primary instance when an activation occurs. See g_application_activate().
func (v *Application) OnActivateAdd(handler OnNoParamHandler) {
	if len(v.activateHandlers) <= 0 {
		v.Connect("activate", func(app *Application) {
			for _, h := range v.activateHandlers {
				h()
			}
		})
	}
	v.activateHandlers = append(v.activateHandlers, handler)
}

//The ::shutdown signal is emitted only on the registered primary instance immediately after the main loop terminates.
func (v *Application) OnShutdownAdd(handler OnNoParamHandler) {
	if len(v.shutdownHandlers) <= 0 {
		v.Connect("activate", func(app *Application) {
			for _, h := range v.shutdownHandlers {
				h()
			}
		})
	}
	v.shutdownHandlers = append(v.shutdownHandlers, handler)
}

//The ::startup signal is emitted on the primary instance immediately after registration. See g_application_register().
func (v *Application) OnStartupAdd(handler OnNoParamHandler) {
	if len(v.startupHandlers) <= 0 {
		v.Connect("activate", func(app *Application) {
			for _, h := range v.startupHandlers {
				h()
			}
		})
	}
	v.startupHandlers = append(v.startupHandlers, handler)
}

//The ::open signal is emitted on the primary instance when there are files to open. See g_application_open() for more information.
func (v *Application) OnOpenAdd(handler OnApplicationOpenFileHandler) {
	if len(v.openHandlers) <= 0 {
		v.Connect("open", func(app *Application, gfiles unsafe.Pointer, nfiles int, hint string) {
			files := make([]*File, nfiles)
			for i := 0; i < nfiles; i++ {
				files[i] = convertToFile(C.get_file(gfiles, (C.int)(i)))
			}
			for _, h := range v.openHandlers {
				h(files, hint)
			}
		})
	}
	v.openHandlers = append(v.openHandlers, handler)
}
