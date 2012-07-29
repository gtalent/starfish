#include "stdio.h"
#include "math.h"
#include "gtk/gtk.h"
#include "cairo.h"
#include "capi.h"

cairo_t *c;
GtkWindow *win;
GtkWidget *canvas;

static gboolean draw(GtkWidget *w, GdkEventExpose *e, gpointer data);

void openDisplay() {
	gtk_init(0, NULL);
	win = (GtkWindow*) gtk_window_new(GTK_WINDOW_TOPLEVEL);

	canvas = gtk_drawing_area_new();
	gtk_container_add(GTK_CONTAINER(win), canvas);

	g_signal_connect(canvas, "draw", G_CALLBACK(draw), NULL);
	g_signal_connect(win, "destroy", G_CALLBACK(gtk_main_quit), NULL);
	gtk_widget_set_app_paintable((GtkWidget*) win, TRUE);
	gtk_widget_show_all((GtkWidget*) win);
}

void closeDisplay() {
	gtk_widget_destroy((GtkWidget*) win);
}

const char* getDisplayTitle() {
	return gtk_window_get_title(win);
}

void setDisplayTitle(const char* title) {
	gtk_window_set_title(win, title);
}

void setFullscreen(int fullscreen) {
	if (fullscreen) {
		gtk_window_fullscreen(win);
	} else {
		gtk_window_unfullscreen(win);
	}
}

int getDisplayWidth() {
	gint r;
	gtk_window_get_size(win, &r, NULL);
	return (int) r;
}

int getDisplayHeight() {
	gint r;
	gtk_window_get_size(win, NULL, &r);
	return (int) r;
}

void setDisplayWidth(int w) {
	gtk_window_resize(win, w, getDisplayHeight());
}

void setDisplayHeight(int h) {
	gtk_window_resize(win, getDisplayWidth(), h);
}

void setDisplaySize(int w, int h) {
	gtk_window_resize(win, w, h);
}

void mainBlock() {
	gtk_main();
}



void setRGBA(int r, int g, int b, int a) {
	inline double f(int x) {
		return ((double) x) / 255;
	}
	cairo_set_source_rgba(c, f(r), f(g), f(b), f(a));
}

void drawRect(int x, int y, int w, int h) {
	cairo_rectangle(c, x, y, w, h);
	cairo_stroke(c);
}

void fillRect(int x, int y, int w, int h) {
	cairo_rectangle(c, x, y, w, h);
	cairo_fill(c);
}

void drawRoundedRect(int x, int y, int w, int h, int r) {
	double d = M_PI / 180.0;
	cairo_new_sub_path(c);
	cairo_arc(c, x + w - r, y + r, r, -90 * d, 0);
	cairo_arc(c, x + w - r, y + h - r, r, 0 * d, 90 * d);
	cairo_arc(c, x + r, y + h - r, r, 90 * d, 180 * d);
	cairo_arc(c, x + r, y + r, r, 180 * d, 270 * d);
	cairo_close_path(c);
	cairo_stroke(c);
}

void fillRoundedRect(int x, int y, int w, int h, int r) {
	double d = M_PI / 180.0;
	cairo_new_sub_path(c);
	cairo_arc(c, x + w - r, y + r, r, -90 * d, 0);
	cairo_arc(c, x + w - r, y + h - r, r, 0 * d, 90 * d);
	cairo_arc(c, x + r, y + h - r, r, 90 * d, 180 * d);
	cairo_arc(c, x + r, y + r, r, 180 * d, 270 * d);
	cairo_close_path(c);
	cairo_fill(c);
}

static gboolean draw(GtkWidget *w, GdkEventExpose *e, gpointer data) {
	c = gdk_cairo_create(gtk_widget_get_window(w));
	setRGBA(0, 0, 255, 255);
	fillRoundedRect(10, 10, 75, 75, 10);
	cairo_destroy(c);
	return FALSE;
}
