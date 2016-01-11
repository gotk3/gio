#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>


//static int callGApplicationRun(GApplication *app, int argc, char **argv){
//	int i, res;
//	printf ("-- Tableau de %d éléments : --\n",argc);
//	for (i=0;i<argc;i++)
//		printf("%d - @ %x : %s\n",i, &argv[i], argv[i]);
//	printf ("-- Call Gio Application RUN --\n",argc);
//	res = g_application_run(app, argc, argv);
//	return res;
//}

static GAction *
toGAction(void *p)
{
	return (G_ACTION(p));
}

static GApplication *
toGApplication(void *p)
{
	return (G_APPLICATION(p));
}

static GCancellable *
toGCancellable(void *p)
{
	return (G_CANCELLABLE(p));
}

static GDBusConnection *
toGDBusConnection(void *p)
{
	return (G_DBUS_CONNECTION(p));
}

static GFile *
toGFile(void *p)
{
	return (G_FILE(p));
}

static GFile ** 
alloc_files(int n) {
	return ((GFile **)g_new0(GFile *, n));
}

static void
set_file(GFile **files, int n, GFile *f)
{
	files[n] = f;
}

static GFile *
get_file(void *files ,int n)
{
	GFile *f = ((GFile **)files)[n];
	return f;
}

static GNotification *
toGNotification(void *p)
{
	return (G_NOTIFICATION(p));
}

static GTypeModule *
toGTypeModule(void *p)
{
	return (G_TYPE_MODULE(p));
}

