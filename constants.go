// MACHINE GENERATED - DO NOT EDIT.

package gosdl_13

// Constants
const (
	INIT_AUDIO               = 0x10
	INIT_VIDEO               = 0x20
	INIT_TIMER               = 0x1
	INIT_JOYSTICK            = 0x200
	INIT_NOPARACHUTE         = 0x100000
	INIT_EVERYTHING          = 0xffff
	RLEACCEL                 = 0x2
	PREALLOC                 = 0x1
	WINDOW_FULLSCREEN        = 0x1
	WINDOW_OPENGL            = 0x2
	WINDOW_SHOWN             = 0x4
	WINDOW_HIDDEN            = 0x8
	WINDOW_BORDERLESS        = 0x10
	WINDOW_RESIZABLE         = 0x20
	WINDOW_MINIMIZED         = 0x40
	WINDOW_MAXIMIZED         = 0x80
	WINDOW_INPUT_GRABBED     = 0x100
	WINDOW_INPUT_FOCUS       = 0x200
	WINDOW_MOUSE_FOCUS       = 0x400
	WINDOW_FOREIGN           = 0x800
	GL_RED_SIZE              = 0
	GL_GREEN_SIZE            = 0x1
	GL_BLUE_SIZE             = 0x2
	GL_ALPHA_SIZE            = 0x3
	GL_BUFFER_SIZE           = 0x4
	GL_DOUBLEBUFFER          = 0x5
	GL_DEPTH_SIZE            = 0x6
	GL_STENCIL_SIZE          = 0x7
	GL_ACCUM_RED_SIZE        = 0x8
	GL_ACCUM_GREEN_SIZE      = 0x9
	GL_ACCUM_BLUE_SIZE       = 0xa
	GL_ACCUM_ALPHA_SIZE      = 0xb
	GL_STEREO                = 0xc
	GL_MULTISAMPLEBUFFERS    = 0xd
	GL_MULTISAMPLES          = 0xe
	GL_ACCELERATED_VISUAL    = 0xf
	GL_RETAINED_BACKING      = 0x10
	GL_CONTEXT_MAJOR_VERSION = 0x11
	GL_CONTEXT_MINOR_VERSION = 0x12
	KEYDOWN                  = 0x300
	KEYUP                    = 0x301
	TEXTEDITING              = 0x302
	TEXTINPUT                = 0x303
	MOUSEMOTION              = 0x400
	MOUSEBUTTONDOWN          = 0x401
	MOUSEBUTTONUP            = 0x402
	MOUSEWHEEL               = 0x403
	JOYAXISMOTION            = 0x600
	JOYBALLMOTION            = 0x601
	JOYHATMOTION             = 0x602
	JOYBUTTONDOWN            = 0x603
	JOYBUTTONUP              = 0x604
	QUIT                     = 0x100
	SYSWMEVENT               = 0x201
	USEREVENT                = 0x8000
	LASTEVENT                = 0xffff
	QUERY                    = -0x1
	DISABLE                  = 0
	ENABLE                   = 0x1
	K_UNKNOWN                = 0
	K_BACKSPACE              = 0x8
	K_TAB                    = 0x9
	K_CLEAR                  = 0x4000009c
	K_RETURN                 = 0xd
	K_PAUSE                  = 0x40000048
	K_ESCAPE                 = 0x1b
	K_SPACE                  = 0x20
	K_EXCLAIM                = 0x21
	K_QUOTEDBL               = 0x22
	K_HASH                   = 0x23
	K_DOLLAR                 = 0x24
	K_AMPERSAND              = 0x26
	K_QUOTE                  = 0x27
	K_LEFTPAREN              = 0x28
	K_RIGHTPAREN             = 0x29
	K_ASTERISK               = 0x2a
	K_PLUS                   = 0x2b
	K_COMMA                  = 0x2c
	K_MINUS                  = 0x2d
	K_PERIOD                 = 0x2e
	K_SLASH                  = 0x2f
	K_0                      = 0x30
	K_1                      = 0x31
	K_2                      = 0x32
	K_3                      = 0x33
	K_4                      = 0x34
	K_5                      = 0x35
	K_6                      = 0x36
	K_7                      = 0x37
	K_8                      = 0x38
	K_9                      = 0x39
	K_COLON                  = 0x3a
	K_SEMICOLON              = 0x3b
	K_LESS                   = 0x3c
	K_EQUALS                 = 0x3d
	K_GREATER                = 0x3e
	K_QUESTION               = 0x3f
	K_AT                     = 0x40
	K_LEFTBRACKET            = 0x5b
	K_BACKSLASH              = 0x5c
	K_RIGHTBRACKET           = 0x5d
	K_CARET                  = 0x5e
	K_UNDERSCORE             = 0x5f
	K_BACKQUOTE              = 0x60
	K_a                      = 0x61
	K_b                      = 0x62
	K_c                      = 0x63
	K_d                      = 0x64
	K_e                      = 0x65
	K_f                      = 0x66
	K_g                      = 0x67
	K_h                      = 0x68
	K_i                      = 0x69
	K_j                      = 0x6a
	K_k                      = 0x6b
	K_l                      = 0x6c
	K_m                      = 0x6d
	K_n                      = 0x6e
	K_o                      = 0x6f
	K_p                      = 0x70
	K_q                      = 0x71
	K_r                      = 0x72
	K_s                      = 0x73
	K_t                      = 0x74
	K_u                      = 0x75
	K_v                      = 0x76
	K_w                      = 0x77
	K_x                      = 0x78
	K_y                      = 0x79
	K_z                      = 0x7a
	K_DELETE                 = 0x7f
	K_KP_PERIOD              = 0x40000063
	K_KP_DIVIDE              = 0x40000054
	K_KP_MULTIPLY            = 0x40000055
	K_KP_MINUS               = 0x40000056
	K_KP_PLUS                = 0x40000057
	K_KP_ENTER               = 0x40000058
	K_KP_EQUALS              = 0x40000067
	K_UP                     = 0x40000052
	K_DOWN                   = 0x40000051
	K_RIGHT                  = 0x4000004f
	K_LEFT                   = 0x40000050
	K_INSERT                 = 0x40000049
	K_HOME                   = 0x4000004a
	K_END                    = 0x4000004d
	K_PAGEUP                 = 0x4000004b
	K_PAGEDOWN               = 0x4000004e
	K_F1                     = 0x4000003a
	K_F2                     = 0x4000003b
	K_F3                     = 0x4000003c
	K_F4                     = 0x4000003d
	K_F5                     = 0x4000003e
	K_F6                     = 0x4000003f
	K_F7                     = 0x40000040
	K_F8                     = 0x40000041
	K_F9                     = 0x40000042
	K_F10                    = 0x40000043
	K_F11                    = 0x40000044
	K_F12                    = 0x40000045
	K_F13                    = 0x40000068
	K_F14                    = 0x40000069
	K_F15                    = 0x4000006a
	K_CAPSLOCK               = 0x40000039
	K_RSHIFT                 = 0x400000e5
	K_LSHIFT                 = 0x400000e1
	K_RCTRL                  = 0x400000e4
	K_LCTRL                  = 0x400000e0
	K_RALT                   = 0x400000e6
	K_LALT                   = 0x400000e2
	K_MODE                   = 0x40000101
	K_HELP                   = 0x40000075
	K_SYSREQ                 = 0x4000009a
	K_MENU                   = 0x40000076
	K_POWER                  = 0x40000066
	K_UNDO                   = 0x4000007a
	KMOD_NONE                = 0
	KMOD_LSHIFT              = 0x1
	KMOD_RSHIFT              = 0x2
	KMOD_LCTRL               = 0x40
	KMOD_RCTRL               = 0x80
	KMOD_LALT                = 0x100
	KMOD_RALT                = 0x200
	KMOD_NUM                 = 0x1000
	KMOD_CAPS                = 0x2000
	KMOD_MODE                = 0x4000
	KMOD_RESERVED            = 0x8000
	HAT_CENTERED             = 0
	HAT_UP                   = 0x1
	HAT_RIGHT                = 0x2
	HAT_DOWN                 = 0x4
	HAT_LEFT                 = 0x8
	HAT_RIGHTUP              = 0x3
	HAT_RIGHTDOWN            = 0x6
	HAT_LEFTUP               = 0x9
	HAT_LEFTDOWN             = 0xc
)

// Types
