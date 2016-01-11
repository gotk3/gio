package gio

// #cgo pkg-config: gio-2.0 glib-2.0
// #include <gio/gio.h>
// #include "gio.go.h"
import "C"
import "unsafe"


/*
 * GApplicationFlags
 * Flags used to define the behaviour of a GApplication.
 */
type ApplicationFlags int

const (
	APPLICATION_FLAGS_NONE           ApplicationFlags = C.G_APPLICATION_FLAGS_NONE           //Default
	APPLICATION_IS_SERVICE           ApplicationFlags = C.G_APPLICATION_IS_SERVICE           //Run as a service. In this mode, registration fails if the service is already running, and the application will initially wait up to 10 seconds for an initial activation message to arrive.
	APPLICATION_IS_LAUNCHER          ApplicationFlags = C.G_APPLICATION_IS_LAUNCHER          //Don't try to become the primary instance.
	APPLICATION_HANDLES_OPEN         ApplicationFlags = C.G_APPLICATION_HANDLES_OPEN         //This application handles opening files (in the primary instance). Note that this flag only affects the default implementation of local_command_line(), and has no effect if G_APPLICATION_HANDLES_COMMAND_LINE is given. See g_application_run() for details.
	APPLICATION_HANDLES_COMMAND_LINE ApplicationFlags = C.G_APPLICATION_HANDLES_COMMAND_LINE //This application handles command line arguments (in the primary instance). Note that this flag only affect the default implementation of local_command_line(). See g_application_run() for details.
	APPLICATION_SEND_ENVIRONMENT     ApplicationFlags = C.G_APPLICATION_SEND_ENVIRONMENT     //Send the environment of the launching process to the primary instance. Set this flag if your application is expected to behave differently depending on certain environment variables. For instance, an editor might be expected to use the <envar>GIT_COMMITTER_NAME</envar> environment variable when editing a git commit message. The environment is available to the “command-line” signal handler, via g_application_command_line_getenv().
	APPLICATION_NON_UNIQUE           ApplicationFlags = C.G_APPLICATION_NON_UNIQUE           //Make no attempts to do any of the typical single-instance application negotiation, even if the application ID is given. The application neither attempts to become the owner of the application ID nor does it check if an existing owner already exists. Everything occurs in the local process. Since: 2.30.
)

func marshalApplicationFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ApplicationFlags(c), nil
}
