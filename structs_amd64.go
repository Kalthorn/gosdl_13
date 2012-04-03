package gosdl_13

type Event interface{}

// C-style union of event types (must be size of largest possible)
type cevent struct {
	Type uint32
	Pad  [44]byte
}

type WindowEvent struct {
	Type  uint32 // WINDOWEVENT
	Win   uint32 // the window where the event occurred
	Event uint8  // WindowEventID
	Pad   [3]byte
	Data1 int32 // depends on the event
	Data2 int32 // depends on the event
}

type KeyboardEvent struct {
	Type   uint32 // KEYDOWN or KEYUP
	Which  uint32
	State  uint8 // PRESSED or RELEASED
	Repeat uint8 // Non-zero if this is a held key
	Pad    [2]byte
	Keysym Keysym // the key itself
}

type Keysym struct {
	Scancode uint16 // the keyboard key pressed
	Sym      uint32 // the character code
	Mod      uint16 // key modifiers
	Unicode  uint32 // deprecated (use TextInputEvent)
}

type MouseMotionEvent struct {
	Type  uint32 // MOUSEMOTION
	Win   uint32 // the window where the event occurred
	State uint8  // PRESSED or RELEASED
	Pad   [3]byte
	X     int32
	Y     int32
	Xrel  int32
	Yrel  int32
}

type MouseButtonEvent struct {
	Type   uint32 // MOUSEBUTTONDOWN or MOUSEBUTTONUP
	Win    uint32 // the window where the event occurred
	Button uint8  // the button index
	State  uint8  // PRESSED or RELEASED
	Pad    [2]byte
	X      int32
	Y      int32
}

type MouseWheelEvent struct {
	Type uint32
	Win  uint32
	X    int32
	Y    int32
}

type TextInputEvent struct {
	Type  uint32
	Win   uint32
	Chars [4]byte
}

type QuitEvent struct {
	Type uint32 // QUIT
}
