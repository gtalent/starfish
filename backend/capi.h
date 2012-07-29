void openDisplay();

void closeDisplay();

void mainBlock();

const char* getDisplayTitle();

void setDisplayTitle(const char* title);

void setFullscreen(int fullscreen);

int getDisplayWidth();

int getDisplayHeight();

void setDisplayWidth(int w);

void setDisplayHeight(int h);

void setDisplaySize(int w, int h);

void* loadImage(const char* path);

void* resizeImage(void* img, int w, int h);

void setRGBA(int r, int g, int b, int a);

void drawRect(int x, int y, int w, int h);

void drawRoundedRect(int x, int y, int w, int h, int radius);

void fillRect(int x, int y, int w, int h);

void fillRoundedRect(int x, int y, int w, int h, int radius);

void invokeDraw();
