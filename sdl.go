package gosdl_13

/*
struct private_hwdata{};
struct SDL_BlitMap{};
#define map _map

#cgo LDFLAGS: -lSDL

#define SDL_NO_COMPAT
#include <SDL/SDL.h>

*/
import "C"
import "unsafe"

type cast unsafe.Pointer

// Initializes SDL, returns non-zero if there is a problem
func Init(flags uint32) int {
	status := int(C.SDL_Init(C.Uint32(flags)))
	return status
}

// Shuts down SDL
func Quit() {
	C.SDL_Quit()
}

// Gets SDL error string
func GetError() string {
	return C.GoString(C.SDL_GetError())
}

// Clear the current SDL error
func ClearError() {
	C.SDL_ClearError()
}

// Polls for currently pending events
func PollEvent() Event {
	var cev C.SDL_Event
	ret := C.SDL_PollEvent(&cev)
	if ret == 0 {
		return nil
	}

	return goEvent((*cevent)(cast(&cev)))
}

func goEvent(cev *cevent) Event {
	switch cev.Type {
	case KEYUP, KEYDOWN:
		return (*KeyboardEvent)(cast(cev))
	case MOUSEMOTION:
		return (*MouseMotionEvent)(cast(cev))
	case MOUSEBUTTONUP, MOUSEBUTTONDOWN:
		return (*MouseButtonEvent)(cast(cev))
	case MOUSEWHEEL:
		return (*MouseWheelEvent)(cast(cev))
	case TEXTINPUT:
		return (*TextInputEvent)(cast(cev))
	case QUIT:
		return (*QuitEvent)(cast(cev))
	}

	return nil
}

func GL_SetAttribute(attr int, value int) int {
	status := int(C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(value)))
	return status
}

func CreateWindow(name string, x, y, w, h int, flags uint32) []uint8 {
	s := C.CString(name)

	win := C.SDL_CreateWindow(s, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	var ptr = make([]uint8, unsafe.Sizeof(win))
	*((**C.SDL_Window)(cast(&ptr))) = win

	return ptr
}

func DestroyWindow(win []uint8) {
	C.SDL_DestroyWindow((*C.SDL_Window)(cast(&win[0])))
}

func GL_CreateContext(win []uint8) []uint8 {
	context := C.SDL_GL_CreateContext((*C.SDL_Window)(cast(&win[0])))
	var ptr = make([]uint8, unsafe.Sizeof(context))
	*((*C.SDL_GLContext)(cast(&ptr))) = context
	return ptr
}

func GL_DeleteContext(context []uint8) {
	C.SDL_GL_DeleteContext((C.SDL_GLContext)(cast(&context[0])))
}

func GL_SetSwapInterval(interval int) int {
	return int(C.SDL_GL_SetSwapInterval(C.int(interval)))
}

func GL_SwapWindow(win []uint8) {
	C.SDL_GL_SwapWindow((*C.SDL_Window)(cast(&win[0])))
}
