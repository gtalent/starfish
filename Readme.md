# What is starfish?
## What starfish is:
starfish is a simple 2D graphics and user input library for Go built on SDL.
## What starfish is not:
While it is built on SDL, but unlike many SDL wrappers, starfish is not meant to replicate the SDL API, only leverage it. Don't expect knowledge of SDL to be beneficial when using this library.

This is done for two reasons:

* If a better foundation comes along or SDL is no longer sufficient, it will be a simple matter to reimplement starfish with it.
* SDL has more options than most programmers care about for simple applications. A greater number of options leads to greater complexity and a higher learning curve.

# Getting starfish:
## Prerequisites:
* Install GCC, SDL, SDL_gfx, SDL_image, and SDL_ttf:
	* RHEL-like systems the following should work:
 
			yum install gcc SDL-devel SDL_gfx-devel SDL_image-devel SDL_ttf-devel

	* Anyone feel free to add instructions for other systems here.

## Installation:
* To install starfish, the following should pull and install all three starfish packages:

			go get github.com/gtalent/starfish/graphics
			go get github.com/gtalent/starfish/input
	
