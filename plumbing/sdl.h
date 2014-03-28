/*
 * Copyright 2011-2014 starfish authors
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
#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>

SDL_Window* openDisplay(int w, int h, int full);

void closeDisplay();

int isMainThread();

void setEventType(SDL_Event *e, Uint32 type);

Uint32 eventType(SDL_Event *e);

int eventKey(SDL_Event *e);

int eventMouseButton(SDL_Event *e);

int eventMouseX(SDL_Event *e);

int eventMouseY(SDL_Event *e);

int eventMouseWheelY(SDL_Event *e);
