/*
   Copyright 2011 gtalent2@gmail.com

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
#cgo LDFLAGS: -lWombat -lallegro -lallegro_font -lallegro_image -lallegro_primitives -lallegro_ttf
#include "allegro5/allegro.h"
*/
import "C"

const (
	Key_a int = C.ALLEGRO_KEY_A
	Key_b     = C.ALLEGRO_KEY_B
	Key_c     = C.ALLEGRO_KEY_C
	Key_d     = C.ALLEGRO_KEY_D
	Key_e     = C.ALLEGRO_KEY_E
	Key_f     = C.ALLEGRO_KEY_F
	Key_g     = C.ALLEGRO_KEY_G
	Key_h     = C.ALLEGRO_KEY_H
	Key_i     = C.ALLEGRO_KEY_I
	Key_j     = C.ALLEGRO_KEY_J
	Key_k     = C.ALLEGRO_KEY_K
	Key_l     = C.ALLEGRO_KEY_L
	Key_m     = C.ALLEGRO_KEY_M
	Key_n     = C.ALLEGRO_KEY_N
	Key_o     = C.ALLEGRO_KEY_O
	Key_p     = C.ALLEGRO_KEY_P
	Key_q     = C.ALLEGRO_KEY_Q
	Key_r     = C.ALLEGRO_KEY_R
	Key_s     = C.ALLEGRO_KEY_S
	Key_t     = C.ALLEGRO_KEY_T
	Key_u     = C.ALLEGRO_KEY_U
	Key_v     = C.ALLEGRO_KEY_V
	Key_w     = C.ALLEGRO_KEY_W
	Key_x     = C.ALLEGRO_KEY_X
	Key_y     = C.ALLEGRO_KEY_Y
	Key_z     = C.ALLEGRO_KEY_Z

	Key_0 = C.ALLEGRO_KEY_0
	Key_1 = C.ALLEGRO_KEY_1
	Key_2 = C.ALLEGRO_KEY_2
	Key_3 = C.ALLEGRO_KEY_3
	Key_4 = C.ALLEGRO_KEY_4
	Key_5 = C.ALLEGRO_KEY_5
	Key_6 = C.ALLEGRO_KEY_6
	Key_7 = C.ALLEGRO_KEY_7
	Key_8 = C.ALLEGRO_KEY_8
	Key_9 = C.ALLEGRO_KEY_9

	Key_SemiColon  = C.ALLEGRO_KEY_SEMICOLON
	Key_Equals     = C.ALLEGRO_KEY_EQUALS
	Key_At         = C.ALLEGRO_KEY_AT
	Key_BackQuote  = C.ALLEGRO_KEY_BACKQUOTE
	Key_Backspace  = C.ALLEGRO_KEY_BACKSPACE
	Key_Tab        = C.ALLEGRO_KEY_TAB
	Key_Enter      = C.ALLEGRO_KEY_ENTER
	Key_Pause      = C.ALLEGRO_KEY_PAUSE
	Key_Escape     = C.ALLEGRO_KEY_ESCAPE
	Key_Space      = C.ALLEGRO_KEY_SPACE
	Key_Quote      = C.ALLEGRO_KEY_QUOTE
	Key_Home       = C.ALLEGRO_KEY_HOME
	Key_End        = C.ALLEGRO_KEY_END
	Key_PageUp     = C.ALLEGRO_KEY_PGUP
	Key_PageDown   = C.ALLEGRO_KEY_PGDN
	Key_Comma      = C.ALLEGRO_KEY_COMMA
	Key_Minus      = C.ALLEGRO_KEY_MINUS
	Key_Slash      = C.ALLEGRO_KEY_SLASH
	Key_Delete     = C.ALLEGRO_KEY_DELETE
	Key_Up         = C.ALLEGRO_KEY_UP
	Key_Down       = C.ALLEGRO_KEY_DOWN
	Key_Left       = C.ALLEGRO_KEY_LEFT
	Key_Right      = C.ALLEGRO_KEY_RIGHT
	Key_RightCtrl  = C.ALLEGRO_KEY_RCTRL
	Key_LeftCtrl   = C.ALLEGRO_KEY_LCTRL
	Key_Alt        = C.ALLEGRO_KEY_ALT
	Key_RightBrace = C.ALLEGRO_KEY_CLOSEBRACE
	Key_LeftBrace  = C.ALLEGRO_KEY_OPENBRACE
	Key_RightShift = C.ALLEGRO_KEY_RSHIFT
	Key_LeftShift  = C.ALLEGRO_KEY_LSHIFT
)
