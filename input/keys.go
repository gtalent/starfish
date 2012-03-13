/*
   Copyright 2011-2012 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package input

/*
#cgo LDFLAGS: -lSDL
#include "SDL/SDL.h"
*/
import "C"

const (
	Key_a int = C.SDLK_a
	Key_b     = C.SDLK_b
	Key_c     = C.SDLK_c
	Key_d     = C.SDLK_d
	Key_e     = C.SDLK_e
	Key_f     = C.SDLK_f
	Key_g     = C.SDLK_g
	Key_h     = C.SDLK_h
	Key_i     = C.SDLK_i
	Key_j     = C.SDLK_j
	Key_k     = C.SDLK_k
	Key_l     = C.SDLK_l
	Key_m     = C.SDLK_m
	Key_n     = C.SDLK_n
	Key_o     = C.SDLK_o
	Key_p     = C.SDLK_p
	Key_q     = C.SDLK_q
	Key_r     = C.SDLK_r
	Key_s     = C.SDLK_s
	Key_t     = C.SDLK_t
	Key_u     = C.SDLK_u
	Key_v     = C.SDLK_v
	Key_w     = C.SDLK_w
	Key_x     = C.SDLK_x
	Key_y     = C.SDLK_y
	Key_z     = C.SDLK_z

	Key_0 = C.SDLK_0
	Key_1 = C.SDLK_1
	Key_2 = C.SDLK_2
	Key_3 = C.SDLK_3
	Key_4 = C.SDLK_4
	Key_5 = C.SDLK_5
	Key_6 = C.SDLK_6
	Key_7 = C.SDLK_7
	Key_8 = C.SDLK_8
	Key_9 = C.SDLK_9

	Key_Colon        = C.SDLK_COLON
	Key_SemiColon    = C.SDLK_SEMICOLON
	Key_LessThan     = C.SDLK_LESS
	Key_Equals       = C.SDLK_EQUALS
	Key_GreaterThan  = C.SDLK_GREATER
	Key_QuestionMark = C.SDLK_QUESTION
	Key_At           = C.SDLK_AT
	Key_LeftBracket  = C.SDLK_LEFTBRACKET
	Key_RightBracket = C.SDLK_RIGHTBRACKET
	Key_Caret        = C.SDLK_CARET
	Key_Underscore   = C.SDLK_UNDERSCORE
	Key_BackQuote    = C.SDLK_BACKQUOTE
	Key_Backspace    = C.SDLK_BACKSPACE
	Key_Tab          = C.SDLK_TAB
	Key_Clear        = C.SDLK_CLEAR
	Key_Enter        = C.SDLK_RETURN
	Key_Pause        = C.SDLK_PAUSE
	Key_Escape       = C.SDLK_ESCAPE
	Key_Space        = C.SDLK_SPACE
	Key_ExclaimMark  = C.SDLK_EXCLAIM
	Key_DoubleQuote  = C.SDLK_QUOTEDBL
	Key_Hash         = C.SDLK_HASH
	Key_Dollar       = C.SDLK_DOLLAR
	Key_LeftParen    = C.SDLK_LEFTPAREN
	Key_RightParen   = C.SDLK_RIGHTPAREN
	Key_Asterisk     = C.SDLK_ASTERISK
	Key_Plus         = C.SDLK_PLUS
	Key_Comma        = C.SDLK_COMMA
	Key_Minus        = C.SDLK_MINUS
	Key_Period       = C.SDLK_PERIOD
	Key_Slash        = C.SDLK_SLASH
	Key_Delete       = C.SDLK_DELETE
	Key_Up           = C.SDLK_UP
	Key_Down         = C.SDLK_DOWN
	Key_Left         = C.SDLK_LEFT
	Key_Right        = C.SDLK_RIGHT
	Key_RCtrl        = C.SDLK_RCTRL
	Key_LCtrl        = C.SDLK_LCTRL
)
