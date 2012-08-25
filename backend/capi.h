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
typedef struct {
	short Red;
	short Green;
	short Blue;
	short Alpha;
} Color;

inline void openDisplay(int, int, int);
inline void closeDisplay();
inline const char* displayTitle();
inline void setDisplayTitle(const char*);
inline int displayWidth();
inline int displayHeight();
inline void clear();
inline void flip();

inline int imageWidth(void*);
inline int imageHeight(void*);

inline void fillRect(int, int, int, int, Color);
inline void fillRoundedRect(int, int, int, int, int, Color);
inline void setClipRect(int, int, int, int);
