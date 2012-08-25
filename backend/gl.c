/*
 * Copyright 2011-2012 starfish authors
 * 
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * 
 *   http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <stdio.h>
#include <math.h>
#include <stdlib.h>

#include <GL/glfw.h>
#include "capi.h"
#include "gl.h"

inline void drawFunc() {
	Draw();
}

void GLFWCALL keyEvent(int key, int action) {
}


inline void openDisplay(int w, int h, int full) {
	glfwInit();
	int flags;
	if (full) {
		flags = GLFW_FULLSCREEN;
	} else {
		flags = GLFW_WINDOW;
	}
	glfwOpenWindow(800, 600, 1, 1, 0, 0, 0, 0, flags);
	glfwSetKeyCallback(keyEvent);
	ilInit();
}

inline void closeDisplay() {
	glfwCloseWindow();
}

inline void setDisplayTitle(const char* title) {
	glfwSetWindowTitle(title);
}

inline int displayWidth() {
	int w, h;
	glfwGetWindowSize(&w, &h);
	return w;
}

inline int displayHeight() {
	int w, h;
	glfwGetWindowSize(&w, &h);
	return h;
}

inline void clear() {
	glClearColor(1.0f, 0.0f, 0.0f, 0.0f);
	glClear(GL_COLOR_BUFFER_BIT| GL_DEPTH_BUFFER_BIT);
	glMatrixMode(GL_PROJECTION);
	glLoadIdentity();
}

inline void flip() {
	glfwSwapBuffers();
}

//GFX

inline void drawImage(Image *i, int x, int y) {
	Image *img = (Image*) i;
	if (!img->loaded) {
		ilBindImage(img->il_img);
		ilConvertImage(IL_RGBA, IL_UNSIGNED_BYTE);
		glGenTextures(1, &img->img);
		glBindTexture(GL_TEXTURE_2D, img->img);
		glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);
		glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
		glTexImage2D(GL_TEXTURE_2D, 0, ilGetInteger(IL_IMAGE_BPP),
			img->w, img->h, 0, ilGetInteger(IL_IMAGE_FORMAT),
			GL_UNSIGNED_BYTE, ilGetData());
		img->loaded = 1;
	}
	glPushMatrix();
	glEnable(GL_BLEND);
	glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA);

	glEnable(GL_ALPHA_TEST);
	glAlphaFunc(GL_GREATER, 0);

	glEnable(GL_TEXTURE_2D);
	glTexEnvf(GL_TEXTURE_2D, GL_TEXTURE_ENV_MODE, GL_REPLACE);
	glColor3d(1, 1, 1);

	glBindTexture(GL_TEXTURE_2D, img->img);

	glBegin(GL_QUADS);
	//Bottom Left
	glTexCoord2f(0, 0);
	glVertex3i(x, y, 0);
	//Bottom Right
	glTexCoord2f(1, 0);
	glVertex3i(x+img->w, y, 0);
	//Top Right
	glTexCoord2f(1, 1);
	glVertex3i(x+img->w, y+img->h, 0);
	//Top Left
	glTexCoord2f(0, 1);
	glVertex3i(x, y+img->h, 0);
	glEnd();

	glDisable(GL_TEXTURE_2D);
	glTexEnvi(GL_TEXTURE_ENV, GL_TEXTURE_ENV_MODE, GL_MODULATE);
	glDisable(GL_ALPHA);
	glDisable(GL_BLEND);
	glPopMatrix();
}

inline void setClipRect(int x, int y, int w, int h) {
	y = (displayHeight() - h) - y;
	glViewport(x, y, w, h);
	glLoadIdentity();
	glOrtho(0.0, w, h, 0.0, -1.0, 1.0);
	glTranslated(-x, -y, 0);
}

inline void fillRoundedRect(int x, int y, int w, int h, int radius, Color c) {
	glColor4ub(c.Red, c.Green, c.Blue, c.Alpha);
	glMatrixMode(GL_MODELVIEW);
	glPushMatrix();
	double ang=0;
	glBegin(GL_FILL);
		glVertex2f(x, y + radius);
		glVertex2f(x, y + h - radius);

		glVertex2f(x + radius, y);
		glVertex2f(x + w - radius, y);

		glVertex2f(x + w, y + radius);
		glVertex2f(x + w, y + h - radius);

		glVertex2f(x + radius, y + h);
		glVertex2f(x + w - radius, y + h);
	glEnd();

	float cX= x+radius, cY = y+radius;
	glBegin(GL_FILL);
		for(ang = M_PI; ang <= 1.5*M_PI; ang += 0.05) {
			glVertex2d(radius* cos(ang) + cX, radius * sin(ang) + cY); //Top Left
		}
		cX = x+w-radius;
	glEnd();
	glBegin(GL_FILL);
		for(ang = 1.5 * M_PI; ang <= 2 * M_PI; ang += 0.05) { 
			glVertex2d(radius* cos(ang) + cX, radius * sin(ang) + cY); //Top Right 
		}
	glEnd();
	glBegin(GL_FILL);
		cY = y+h-radius;
		for(ang = 0; ang <= 0.5 * M_PI; ang += 0.05) {
			glVertex2d(radius* cos(ang) + cX, radius * sin(ang) + cY); //Bottom Right
		}
	glEnd();
	glBegin(GL_FILL);
		cX = x+radius;
		for(ang = 0.5 * M_PI; ang <= M_PI; ang += 0.05) {
			glVertex2d(radius* cos(ang) + cX, radius * sin(ang) + cY);//Bottom Left
		}
	glEnd();
	glPopMatrix();
}

inline void fillRect(int x, int y, int w, int h, Color c) {
	glMatrixMode(GL_PROJECTION);
	glColor4ub(c.Red, c.Green, c.Blue, c.Alpha);
	glRecti(x, y, w, h);
}

inline Image loadImage(const char *path) {
	Image i;
	ilGenImages(1, &i.il_img);
	ilBindImage(i.il_img);
	ilLoadImage(path);
	i.w = ilGetInteger(IL_IMAGE_WIDTH);
	i.h = ilGetInteger(IL_IMAGE_HEIGHT);
	i.loaded = 0;
	return i;
}

inline void freeImage(Image *img) {
	ilDeleteImage(img->il_img);
	glDeleteTextures(1, &img->img);
}

inline int imageWidth(void*);
inline int imageHeight(void*);
