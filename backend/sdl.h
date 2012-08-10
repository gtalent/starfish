#include <SDL/SDL.h>

SDL_Surface* openDisplay(int w, int h, int full);

void closeDisplay();

int eventType(SDL_Event *e);

int eventKey(SDL_Event *e);

int eventMouseButton(SDL_Event *e);

int eventMouseX(SDL_Event *e);

int eventMouseY(SDL_Event *e);
