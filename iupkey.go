package iup

/* from 32 to 126, all character sets are equal,
   the key code is the same as the ASCii character code. */

const (
	K_SP           = int(' ')  /* 32 (0x20) */
	K_exclam       = int('!')  /* 33 */
	K_quotedbl     = int('"')  /* 34 */
	K_numbersign   = int('#')  /* 35 */
	K_dollar       = int('$')  /* 36 */
	K_percent      = int('%')  /* 37 */
	K_ampersand    = int('&')  /* 38 */
	K_apostrophe   = int('\'') /* 39 */
	K_parentleft   = int('(')  /* 40 */
	K_parentright  = int(')')  /* 41 */
	K_asterisk     = int('*')  /* 42 */
	K_plus         = int('+')  /* 43 */
	K_comma        = int(',')  /* 44 */
	K_minus        = int('-')  /* 45 */
	K_period       = int('.')  /* 46 */
	K_slash        = int('/')  /* 47 */
	K_0            = int('0')  /* 48 (0x30) */
	K_1            = int('1')  /* 49 */
	K_2            = int('2')  /* 50 */
	K_3            = int('3')  /* 51 */
	K_4            = int('4')  /* 52 */
	K_5            = int('5')  /* 53 */
	K_6            = int('6')  /* 54 */
	K_7            = int('7')  /* 55 */
	K_8            = int('8')  /* 56 */
	K_9            = int('9')  /* 57 */
	K_colon        = int(':')  /* 58 */
	K_semicolon    = int(';')  /* 59 */
	K_less         = int('<')  /* 60 */
	K_equal        = int('=')  /* 61 */
	K_greater      = int('>')  /* 62 */
	K_question     = int('?')  /* 63 */
	K_at           = int('@')  /* 64 */
	K_A            = int('A')  /* 65 (0x41) */
	K_B            = int('B')  /* 66 */
	K_C            = int('C')  /* 67 */
	K_D            = int('D')  /* 68 */
	K_E            = int('E')  /* 69 */
	K_F            = int('F')  /* 70 */
	K_G            = int('G')  /* 71 */
	K_H            = int('H')  /* 72 */
	K_I            = int('I')  /* 73 */
	K_J            = int('J')  /* 74 */
	K_K            = int('K')  /* 75 */
	K_L            = int('L')  /* 76 */
	K_M            = int('M')  /* 77 */
	K_N            = int('N')  /* 78 */
	K_O            = int('O')  /* 79 */
	K_P            = int('P')  /* 80 */
	K_Q            = int('Q')  /* 81 */
	K_R            = int('R')  /* 82 */
	K_S            = int('S')  /* 83 */
	K_T            = int('T')  /* 84 */
	K_U            = int('U')  /* 85 */
	K_V            = int('V')  /* 86 */
	K_W            = int('W')  /* 87 */
	K_X            = int('X')  /* 88 */
	K_Y            = int('Y')  /* 89 */
	K_Z            = int('Z')  /* 90 */
	K_bracketleft  = int('[')  /* 91 */
	K_backslash    = int('\\') /* 92 */
	K_bracketright = int(']')  /* 93 */
	K_circum       = int('^')  /* 94 */
	K_underscore   = int('_')  /* 95 */
	K_grave        = int('`')  /* 96 */
	K_a            = int('a')  /* 97 (0x61) */
	K_b            = int('b')  /* 98 */
	K_c            = int('c')  /* 99 */
	K_d            = int('d')  /* 100 */
	K_e            = int('e')  /* 101 */
	K_f            = int('f')  /* 102 */
	K_g            = int('g')  /* 103 */
	K_h            = int('h')  /* 104 */
	K_i            = int('i')  /* 105 */
	K_j            = int('j')  /* 106 */
	K_k            = int('k')  /* 107 */
	K_l            = int('l')  /* 108 */
	K_m            = int('m')  /* 109 */
	K_n            = int('n')  /* 110 */
	K_o            = int('o')  /* 111 */
	K_p            = int('p')  /* 112 */
	K_q            = int('q')  /* 113 */
	K_r            = int('r')  /* 114 */
	K_s            = int('s')  /* 115 */
	K_t            = int('t')  /* 116 */
	K_u            = int('u')  /* 117 */
	K_v            = int('v')  /* 118 */
	K_w            = int('w')  /* 119 */
	K_x            = int('x')  /* 120 */
	K_y            = int('y')  /* 121 */
	K_z            = int('z')  /* 122 */
	K_braceleft    = int('{')  /* 123 */
	K_bar          = int('|')  /* 124 */
	K_braceright   = int('}')  /* 125 */
	K_tilde        = int('~')  /* 126 (0x7E) */
)

/* Printable ASCii keys */

//#define iup_isprint(_c) ((_c) > 31 && (_c) < 127)
func IsPrint(c int) bool {
	return c > 31 && c < 127
}

/* also define the escape sequences that have keys associated */

const (
	K_BS  = int('\b') /* 8 */
	K_TAB = int('\t') /* 9 */
	K_LF  = int('\n') /* 10 (0x0A) not a real key, is a combination of CR with a modifier, just to document */
	K_CR  = int('\r') /* 13 (0x0D) */
)

/* backward compatible definitions */

const (
	K_quoteleft  = K_grave
	K_quoteright = K_apostrophe

//	isxkey       = iup_isXkey
)

/* IUP Extended Key Codes, range start at 128      */

//#define iup_isXkey(_c)      ((_c) >= 128)
func IsXKey(c int) bool {
	return (c >= 128)
}

/* These use the same definition as X11 and GDK.
   This also means that any X11 or GDK definition can also be used. */

const (
	K_PAUSE  = 0xFF13
	K_ESC    = 0xFF1B
	K_HOME   = 0xFF50
	K_LEFT   = 0xFF51
	K_UP     = 0xFF52
	K_RIGHT  = 0xFF53
	K_DOWN   = 0xFF54
	K_PGUP   = 0xFF55
	K_PGDN   = 0xFF56
	K_END    = 0xFF57
	K_MIDDLE = 0xFF0B
	K_Print  = 0xFF61
	K_INS    = 0xFF63
	K_Menu   = 0xFF67
	K_DEL    = 0xFFFF
	K_F1     = 0xFFBE
	K_F2     = 0xFFBF
	K_F3     = 0xFFC0
	K_F4     = 0xFFC1
	K_F5     = 0xFFC2
	K_F6     = 0xFFC3
	K_F7     = 0xFFC4
	K_F8     = 0xFFC5
	K_F9     = 0xFFC6
	K_F10    = 0xFFC7
	K_F11    = 0xFFC8
	K_F12    = 0xFFC9
)

/* no Shift/Ctrl/Alt */

const (
	K_LSHIFT = 0xFFE1
	K_RSHIFT = 0xFFE2
	K_LCTRL  = 0xFFE3
	K_RCTRL  = 0xFFE4
	K_LALT   = 0xFFE9
	K_RALT   = 0xFFEA

	K_NUM    = 0xFF7F
	K_SCROLL = 0xFF14
	K_CAPS   = 0xFFE5
)

/* Also, these are the same as the Latin-1 definition */

const (
	K_ccedilla  = 0x00E7
	K_Ccedilla  = 0x00C7
	K_acute     = 0x00B4 /* no Shift/Ctrl/Alt */
	K_diaeresis = 0x00A8
)

/******************************************************/
/* Modifiers use last 4 bits. Since IUP 3.9           */
/* These modifiers definitions are specific to IUP    */
/******************************************************/

//#define iup_isShiftXkey(_c) ((_c) & 0x10000000)
func IsShiftXKey(c int) bool {
	return (uint(c) & 0x10000000) != 0
}

//#define iup_isCtrlXkey(_c)  ((_c) & 0x20000000)
func IsCtrlXKey(c int) bool {
	return (uint(c) & 0x20000000) != 0
}

//#define iup_isAltXkey(_c)   ((_c) & 0x40000000)
func IsAltXKey(c int) bool {
	return (uint(c) & 0x40000000) != 0
}

//#define iup_isSysXkey(_c)   ((_c) & 0x80000000)
func IsSysXKey(c int) bool {
	return (uint(c) & 0x80000000) != 0
}

//#define iup_XkeyBase(_c)  ((_c) & 0x0FFFFFFF)
func XKeyBase(c int) int {
	return int(uint(c) & 0x0FFFFFFF)
}

//#define iup_XkeyShift(_c) ((_c) | 0x10000000)   /* Shift  */
func XKeyShift(c int) int {
	return int(uint(c) | 0x10000000)
}

//#define iup_XkeyCtrl(_c)  ((_c) | 0x20000000)   /* Ctrl   */
func XKeyCtrl(c int) int {
	return int(uint(c) | 0x20000000)
}

//#define iup_XkeyAlt(_c)   ((_c) | 0x40000000)   /* Alt    */
func XKeyAlt(c int) int {
	return int(uint(c) | 0x40000000)
}

//#define iup_XkeySys(_c)   ((_c) | 0x80000000)   /* Sys (Win or Apple) */
func XKeySys(c int) int {
	return int(uint(c) | 0x80000000)
}

/* These definitions are here for backward compatibility
   and to simplify some key combination usage.
   But since IUP 3.9, modifiers can be combined with any key
   and they can be mixed togheter. */

/*
#define K_sHOME    iup_XkeyShift(K_HOME   )
#define K_sUP      iup_XkeyShift(K_UP     )
#define K_sPGUP    iup_XkeyShift(K_PGUP   )
#define K_sLEFT    iup_XkeyShift(K_LEFT   )
#define K_sMIDDLE  iup_XkeyShift(K_MIDDLE )
#define K_sRIGHT   iup_XkeyShift(K_RIGHT  )
#define K_sEND     iup_XkeyShift(K_END    )
#define K_sDOWN    iup_XkeyShift(K_DOWN   )
#define K_sPGDN    iup_XkeyShift(K_PGDN   )
#define K_sINS     iup_XkeyShift(K_INS    )
#define K_sDEL     iup_XkeyShift(K_DEL    )
#define K_sSP      iup_XkeyShift(K_SP     )
#define K_sTAB     iup_XkeyShift(K_TAB    )
#define K_sCR      iup_XkeyShift(K_CR     )
#define K_sBS      iup_XkeyShift(K_BS     )
#define K_sPAUSE   iup_XkeyShift(K_PAUSE  )
#define K_sESC     iup_XkeyShift(K_ESC    )
#define K_sF1      iup_XkeyShift(K_F1     )
#define K_sF2      iup_XkeyShift(K_F2     )
#define K_sF3      iup_XkeyShift(K_F3     )
#define K_sF4      iup_XkeyShift(K_F4     )
#define K_sF5      iup_XkeyShift(K_F5     )
#define K_sF6      iup_XkeyShift(K_F6     )
#define K_sF7      iup_XkeyShift(K_F7     )
#define K_sF8      iup_XkeyShift(K_F8     )
#define K_sF9      iup_XkeyShift(K_F9     )
#define K_sF10     iup_XkeyShift(K_F10    )
#define K_sF11     iup_XkeyShift(K_F11    )
#define K_sF12     iup_XkeyShift(K_F12    )
#define K_sPrint   iup_XkeyShift(K_Print  )
#define K_sMenu    iup_XkeyShift(K_Menu   )

#define K_cHOME     iup_XkeyCtrl(K_HOME    )
#define K_cUP       iup_XkeyCtrl(K_UP      )
#define K_cPGUP     iup_XkeyCtrl(K_PGUP    )
#define K_cLEFT     iup_XkeyCtrl(K_LEFT    )
#define K_cMIDDLE   iup_XkeyCtrl(K_MIDDLE  )
#define K_cRIGHT    iup_XkeyCtrl(K_RIGHT   )
#define K_cEND      iup_XkeyCtrl(K_END     )
#define K_cDOWN     iup_XkeyCtrl(K_DOWN    )
#define K_cPGDN     iup_XkeyCtrl(K_PGDN    )
#define K_cINS      iup_XkeyCtrl(K_INS     )
#define K_cDEL      iup_XkeyCtrl(K_DEL     )
#define K_cSP       iup_XkeyCtrl(K_SP      )
#define K_cTAB      iup_XkeyCtrl(K_TAB     )
#define K_cCR       iup_XkeyCtrl(K_CR      )
#define K_cBS       iup_XkeyCtrl(K_BS      )
#define K_cPAUSE    iup_XkeyCtrl(K_PAUSE   )
#define K_cESC      iup_XkeyCtrl(K_ESC     )
#define K_cCcedilla iup_XkeyCtrl(K_Ccedilla)
#define K_cF1       iup_XkeyCtrl(K_F1      )
#define K_cF2       iup_XkeyCtrl(K_F2      )
#define K_cF3       iup_XkeyCtrl(K_F3      )
#define K_cF4       iup_XkeyCtrl(K_F4      )
#define K_cF5       iup_XkeyCtrl(K_F5      )
#define K_cF6       iup_XkeyCtrl(K_F6      )
#define K_cF7       iup_XkeyCtrl(K_F7      )
#define K_cF8       iup_XkeyCtrl(K_F8      )
#define K_cF9       iup_XkeyCtrl(K_F9      )
#define K_cF10      iup_XkeyCtrl(K_F10     )
#define K_cF11      iup_XkeyCtrl(K_F11     )
#define K_cF12      iup_XkeyCtrl(K_F12     )
#define K_cPrint    iup_XkeyCtrl(K_Print   )
#define K_cMenu     iup_XkeyCtrl(K_Menu    )

#define K_mHOME     iup_XkeyAlt(K_HOME    )
#define K_mUP       iup_XkeyAlt(K_UP      )
#define K_mPGUP     iup_XkeyAlt(K_PGUP    )
#define K_mLEFT     iup_XkeyAlt(K_LEFT    )
#define K_mMIDDLE   iup_XkeyAlt(K_MIDDLE  )
#define K_mRIGHT    iup_XkeyAlt(K_RIGHT   )
#define K_mEND      iup_XkeyAlt(K_END     )
#define K_mDOWN     iup_XkeyAlt(K_DOWN    )
#define K_mPGDN     iup_XkeyAlt(K_PGDN    )
#define K_mINS      iup_XkeyAlt(K_INS     )
#define K_mDEL      iup_XkeyAlt(K_DEL     )
#define K_mSP       iup_XkeyAlt(K_SP      )
#define K_mTAB      iup_XkeyAlt(K_TAB     )
#define K_mCR       iup_XkeyAlt(K_CR      )
#define K_mBS       iup_XkeyAlt(K_BS      )
#define K_mPAUSE    iup_XkeyAlt(K_PAUSE   )
#define K_mESC      iup_XkeyAlt(K_ESC     )
#define K_mCcedilla iup_XkeyAlt(K_Ccedilla)
#define K_mF1       iup_XkeyAlt(K_F1      )
#define K_mF2       iup_XkeyAlt(K_F2      )
#define K_mF3       iup_XkeyAlt(K_F3      )
#define K_mF4       iup_XkeyAlt(K_F4      )
#define K_mF5       iup_XkeyAlt(K_F5      )
#define K_mF6       iup_XkeyAlt(K_F6      )
#define K_mF7       iup_XkeyAlt(K_F7      )
#define K_mF8       iup_XkeyAlt(K_F8      )
#define K_mF9       iup_XkeyAlt(K_F9      )
#define K_mF10      iup_XkeyAlt(K_F10     )
#define K_mF11      iup_XkeyAlt(K_F11     )
#define K_mF12      iup_XkeyAlt(K_F12     )
#define K_mPrint    iup_XkeyAlt(K_Print   )
#define K_mMenu     iup_XkeyAlt(K_Menu    )

#define K_yHOME     iup_XkeySys(K_HOME    )
#define K_yUP       iup_XkeySys(K_UP      )
#define K_yPGUP     iup_XkeySys(K_PGUP    )
#define K_yLEFT     iup_XkeySys(K_LEFT    )
#define K_yMIDDLE   iup_XkeySys(K_MIDDLE  )
#define K_yRIGHT    iup_XkeySys(K_RIGHT   )
#define K_yEND      iup_XkeySys(K_END     )
#define K_yDOWN     iup_XkeySys(K_DOWN    )
#define K_yPGDN     iup_XkeySys(K_PGDN    )
#define K_yINS      iup_XkeySys(K_INS     )
#define K_yDEL      iup_XkeySys(K_DEL     )
#define K_ySP       iup_XkeySys(K_SP      )
#define K_yTAB      iup_XkeySys(K_TAB     )
#define K_yCR       iup_XkeySys(K_CR      )
#define K_yBS       iup_XkeySys(K_BS      )
#define K_yPAUSE    iup_XkeySys(K_PAUSE   )
#define K_yESC      iup_XkeySys(K_ESC     )
#define K_yCcedilla iup_XkeySys(K_Ccedilla)
#define K_yF1       iup_XkeySys(K_F1      )
#define K_yF2       iup_XkeySys(K_F2      )
#define K_yF3       iup_XkeySys(K_F3      )
#define K_yF4       iup_XkeySys(K_F4      )
#define K_yF5       iup_XkeySys(K_F5      )
#define K_yF6       iup_XkeySys(K_F6      )
#define K_yF7       iup_XkeySys(K_F7      )
#define K_yF8       iup_XkeySys(K_F8      )
#define K_yF9       iup_XkeySys(K_F9      )
#define K_yF10      iup_XkeySys(K_F10     )
#define K_yF11      iup_XkeySys(K_F11     )
#define K_yF12      iup_XkeySys(K_F12     )
#define K_yPrint    iup_XkeySys(K_Print   )
#define K_yMenu     iup_XkeySys(K_Menu    )

#define K_sPlus         iup_XkeyShift(K_plus    )
#define K_sComma        iup_XkeyShift(K_comma   )
#define K_sMinus        iup_XkeyShift(K_minus   )
#define K_sPeriod       iup_XkeyShift(K_period  )
#define K_sSlash        iup_XkeyShift(K_slash   )
#define K_sAsterisk     iup_XkeyShift(K_asterisk)

#define K_cA     iup_XkeyCtrl(K_A)
#define K_cB     iup_XkeyCtrl(K_B)
#define K_cC     iup_XkeyCtrl(K_C)
#define K_cD     iup_XkeyCtrl(K_D)
#define K_cE     iup_XkeyCtrl(K_E)
#define K_cF     iup_XkeyCtrl(K_F)
#define K_cG     iup_XkeyCtrl(K_G)
#define K_cH     iup_XkeyCtrl(K_H)
#define K_cI     iup_XkeyCtrl(K_I)
#define K_cJ     iup_XkeyCtrl(K_J)
#define K_cK     iup_XkeyCtrl(K_K)
#define K_cL     iup_XkeyCtrl(K_L)
#define K_cM     iup_XkeyCtrl(K_M)
#define K_cN     iup_XkeyCtrl(K_N)
#define K_cO     iup_XkeyCtrl(K_O)
#define K_cP     iup_XkeyCtrl(K_P)
#define K_cQ     iup_XkeyCtrl(K_Q)
#define K_cR     iup_XkeyCtrl(K_R)
#define K_cS     iup_XkeyCtrl(K_S)
#define K_cT     iup_XkeyCtrl(K_T)
#define K_cU     iup_XkeyCtrl(K_U)
#define K_cV     iup_XkeyCtrl(K_V)
#define K_cW     iup_XkeyCtrl(K_W)
#define K_cX     iup_XkeyCtrl(K_X)
#define K_cY     iup_XkeyCtrl(K_Y)
#define K_cZ     iup_XkeyCtrl(K_Z)
#define K_c1     iup_XkeyCtrl(K_1)
#define K_c2     iup_XkeyCtrl(K_2)
#define K_c3     iup_XkeyCtrl(K_3)
#define K_c4     iup_XkeyCtrl(K_4)
#define K_c5     iup_XkeyCtrl(K_5)
#define K_c6     iup_XkeyCtrl(K_6)
#define K_c7     iup_XkeyCtrl(K_7)
#define K_c8     iup_XkeyCtrl(K_8)
#define K_c9     iup_XkeyCtrl(K_9)
#define K_c0     iup_XkeyCtrl(K_0)
#define K_cPlus         iup_XkeyCtrl(K_plus        )
#define K_cComma        iup_XkeyCtrl(K_comma       )
#define K_cMinus        iup_XkeyCtrl(K_minus       )
#define K_cPeriod       iup_XkeyCtrl(K_period      )
#define K_cSlash        iup_XkeyCtrl(K_slash       )
#define K_cSemicolon    iup_XkeyCtrl(K_semicolon   )
#define K_cEqual        iup_XkeyCtrl(K_equal       )
#define K_cBracketleft  iup_XkeyCtrl(K_bracketleft )
#define K_cBracketright iup_XkeyCtrl(K_bracketright)
#define K_cBackslash    iup_XkeyCtrl(K_backslash   )
#define K_cAsterisk     iup_XkeyCtrl(K_asterisk    )

#define K_mA     iup_XkeyAlt(K_A)
#define K_mB     iup_XkeyAlt(K_B)
#define K_mC     iup_XkeyAlt(K_C)
#define K_mD     iup_XkeyAlt(K_D)
#define K_mE     iup_XkeyAlt(K_E)
#define K_mF     iup_XkeyAlt(K_F)
#define K_mG     iup_XkeyAlt(K_G)
#define K_mH     iup_XkeyAlt(K_H)
#define K_mI     iup_XkeyAlt(K_I)
#define K_mJ     iup_XkeyAlt(K_J)
#define K_mK     iup_XkeyAlt(K_K)
#define K_mL     iup_XkeyAlt(K_L)
#define K_mM     iup_XkeyAlt(K_M)
#define K_mN     iup_XkeyAlt(K_N)
#define K_mO     iup_XkeyAlt(K_O)
#define K_mP     iup_XkeyAlt(K_P)
#define K_mQ     iup_XkeyAlt(K_Q)
#define K_mR     iup_XkeyAlt(K_R)
#define K_mS     iup_XkeyAlt(K_S)
#define K_mT     iup_XkeyAlt(K_T)
#define K_mU     iup_XkeyAlt(K_U)
#define K_mV     iup_XkeyAlt(K_V)
#define K_mW     iup_XkeyAlt(K_W)
#define K_mX     iup_XkeyAlt(K_X)
#define K_mY     iup_XkeyAlt(K_Y)
#define K_mZ     iup_XkeyAlt(K_Z)
#define K_m1     iup_XkeyAlt(K_1)
#define K_m2     iup_XkeyAlt(K_2)
#define K_m3     iup_XkeyAlt(K_3)
#define K_m4     iup_XkeyAlt(K_4)
#define K_m5     iup_XkeyAlt(K_5)
#define K_m6     iup_XkeyAlt(K_6)
#define K_m7     iup_XkeyAlt(K_7)
#define K_m8     iup_XkeyAlt(K_8)
#define K_m9     iup_XkeyAlt(K_9)
#define K_m0     iup_XkeyAlt(K_0)
#define K_mPlus         iup_XkeyAlt(K_plus        )
#define K_mComma        iup_XkeyAlt(K_comma       )
#define K_mMinus        iup_XkeyAlt(K_minus       )
#define K_mPeriod       iup_XkeyAlt(K_period      )
#define K_mSlash        iup_XkeyAlt(K_slash       )
#define K_mSemicolon    iup_XkeyAlt(K_semicolon   )
#define K_mEqual        iup_XkeyAlt(K_equal       )
#define K_mBracketleft  iup_XkeyAlt(K_bracketleft )
#define K_mBracketright iup_XkeyAlt(K_bracketright)
#define K_mBackslash    iup_XkeyAlt(K_backslash   )
#define K_mAsterisk     iup_XkeyAlt(K_asterisk    )

#define K_yA     iup_XkeySys(K_A)
#define K_yB     iup_XkeySys(K_B)
#define K_yC     iup_XkeySys(K_C)
#define K_yD     iup_XkeySys(K_D)
#define K_yE     iup_XkeySys(K_E)
#define K_yF     iup_XkeySys(K_F)
#define K_yG     iup_XkeySys(K_G)
#define K_yH     iup_XkeySys(K_H)
#define K_yI     iup_XkeySys(K_I)
#define K_yJ     iup_XkeySys(K_J)
#define K_yK     iup_XkeySys(K_K)
#define K_yL     iup_XkeySys(K_L)
#define K_yM     iup_XkeySys(K_M)
#define K_yN     iup_XkeySys(K_N)
#define K_yO     iup_XkeySys(K_O)
#define K_yP     iup_XkeySys(K_P)
#define K_yQ     iup_XkeySys(K_Q)
#define K_yR     iup_XkeySys(K_R)
#define K_yS     iup_XkeySys(K_S)
#define K_yT     iup_XkeySys(K_T)
#define K_yU     iup_XkeySys(K_U)
#define K_yV     iup_XkeySys(K_V)
#define K_yW     iup_XkeySys(K_W)
#define K_yX     iup_XkeySys(K_X)
#define K_yY     iup_XkeySys(K_Y)
#define K_yZ     iup_XkeySys(K_Z)
#define K_y1     iup_XkeySys(K_1)
#define K_y2     iup_XkeySys(K_2)
#define K_y3     iup_XkeySys(K_3)
#define K_y4     iup_XkeySys(K_4)
#define K_y5     iup_XkeySys(K_5)
#define K_y6     iup_XkeySys(K_6)
#define K_y7     iup_XkeySys(K_7)
#define K_y8     iup_XkeySys(K_8)
#define K_y9     iup_XkeySys(K_9)
#define K_y0     iup_XkeySys(K_0)
#define K_yPlus         iup_XkeySys(K_plus        )
#define K_yComma        iup_XkeySys(K_comma       )
#define K_yMinus        iup_XkeySys(K_minus       )
#define K_yPeriod       iup_XkeySys(K_period      )
#define K_ySlash        iup_XkeySys(K_slash       )
#define K_ySemicolon    iup_XkeySys(K_semicolon   )
#define K_yEqual        iup_XkeySys(K_equal       )
#define K_yBracketleft  iup_XkeySys(K_bracketleft )
#define K_yBracketright iup_XkeySys(K_bracketright)
#define K_yBackslash    iup_XkeySys(K_backslash   )
#define K_yAsterisk     iup_XkeySys(K_asterisk    )
*/
