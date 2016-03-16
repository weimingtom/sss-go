// Copyright 2010 The go-winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"syscall"
	"unsafe"
)

const MAX_PATH = 260

// Error codes
const (
	ERROR_SUCCESS             = 0
	ERROR_FILE_NOT_FOUND      = 2
	ERROR_INVALID_PARAMETER   = 87
	ERROR_INSUFFICIENT_BUFFER = 122
	ERROR_MORE_DATA           = 234
)

// GlobalAlloc flags
const (
	GHND          = 0x0042
	GMEM_FIXED    = 0x0000
	GMEM_MOVEABLE = 0x0002
	GMEM_ZEROINIT = 0x0040
	GPTR          = 0x004
)

// Predefined locale ids
const (
	LOCALE_INVARIANT      = 0x007f
    LOCALE_NOUSEROVERRIDE        = 0x80000000
    LOCALE_USE_CP_ACP            = 0x40000000
    LOCALE_RETURN_NUMBER         = 0x20000000
    LOCALE_ILANGUAGE             = 1
    LOCALE_SLANGUAGE             = 2
    LOCALE_SENGLANGUAGE          = 0x1001
    LOCALE_SABBREVLANGNAME       = 3
    LOCALE_SNATIVELANGNAME       = 4
    LOCALE_ICOUNTRY              = 5
    LOCALE_SCOUNTRY              = 6
    LOCALE_SENGCOUNTRY           = 0x1002
    LOCALE_SABBREVCTRYNAME       = 7
    LOCALE_SNATIVECTRYNAME       = 8
    LOCALE_IDEFAULTLANGUAGE      = 9
    LOCALE_IDEFAULTCOUNTRY       = 10
    LOCALE_IDEFAULTCODEPAGE      = 11
    LOCALE_IDEFAULTANSICODEPAGE  = 0x1004
    LOCALE_SLIST                 = 12
    LOCALE_IMEASURE              = 13
    LOCALE_SDECIMAL              = 14
    LOCALE_STHOUSAND             = 15
    LOCALE_SGROUPING             = 16
    LOCALE_IDIGITS               = 17
    LOCALE_ILZERO                = 18
    LOCALE_INEGNUMBER            = 0x1010
    LOCALE_SNATIVEDIGITS         = 19
    LOCALE_SCURRENCY             = 20
    LOCALE_SINTLSYMBOL           = 21
    LOCALE_SMONDECIMALSEP        = 22
    LOCALE_SMONTHOUSANDSEP       = 23
    LOCALE_SMONGROUPING          = 24
    LOCALE_ICURRDIGITS           = 25
    LOCALE_IINTLCURRDIGITS       = 26
    LOCALE_ICURRENCY             = 27
    LOCALE_INEGCURR              = 28
    LOCALE_SDATE                 = 29
    LOCALE_STIME                 = 30
    LOCALE_SSHORTDATE            = 31
    LOCALE_SLONGDATE             = 32
    LOCALE_STIMEFORMAT           = 0x1003
    LOCALE_IDATE                 = 33
    LOCALE_ILDATE                = 34
    LOCALE_ITIME                 = 35
    LOCALE_ITIMEMARKPOSN         = 0x1005
    LOCALE_ICENTURY              = 36
    LOCALE_ITLZERO               = 37
    LOCALE_IDAYLZERO             = 38
    LOCALE_IMONLZERO             = 39
    LOCALE_S1159                 = 40
    LOCALE_S2359                 = 41
    LOCALE_ICALENDARTYPE         = 0x1009
    LOCALE_IOPTIONALCALENDAR     = 0x100B
    LOCALE_IFIRSTDAYOFWEEK       = 0x100C
    LOCALE_IFIRSTWEEKOFYEAR      = 0x100D
    LOCALE_SDAYNAME1             = 42
    LOCALE_SDAYNAME2             = 43
    LOCALE_SDAYNAME3             = 44
    LOCALE_SDAYNAME4             = 45
    LOCALE_SDAYNAME5             = 46
    LOCALE_SDAYNAME6             = 47
    LOCALE_SDAYNAME7             = 48
    LOCALE_SABBREVDAYNAME1       = 49
    LOCALE_SABBREVDAYNAME2       = 50
    LOCALE_SABBREVDAYNAME3       = 51
    LOCALE_SABBREVDAYNAME4       = 52
    LOCALE_SABBREVDAYNAME5       = 53
    LOCALE_SABBREVDAYNAME6       = 54
    LOCALE_SABBREVDAYNAME7       = 55
    LOCALE_SMONTHNAME1           = 56
    LOCALE_SMONTHNAME2           = 57
    LOCALE_SMONTHNAME3           = 58
    LOCALE_SMONTHNAME4           = 59
    LOCALE_SMONTHNAME5           = 60
    LOCALE_SMONTHNAME6           = 61
    LOCALE_SMONTHNAME7           = 62
    LOCALE_SMONTHNAME8           = 63
    LOCALE_SMONTHNAME9           = 64
    LOCALE_SMONTHNAME10          = 65
    LOCALE_SMONTHNAME11          = 66
    LOCALE_SMONTHNAME12          = 67
    LOCALE_SMONTHNAME13          = 0x100E
    LOCALE_SABBREVMONTHNAME1     = 68
    LOCALE_SABBREVMONTHNAME2     = 69
    LOCALE_SABBREVMONTHNAME3     = 70
    LOCALE_SABBREVMONTHNAME4     = 71
    LOCALE_SABBREVMONTHNAME5     = 72
    LOCALE_SABBREVMONTHNAME6     = 73
    LOCALE_SABBREVMONTHNAME7     = 74
    LOCALE_SABBREVMONTHNAME8     = 75
    LOCALE_SABBREVMONTHNAME9     = 76
    LOCALE_SABBREVMONTHNAME10    = 77
    LOCALE_SABBREVMONTHNAME11    = 78
    LOCALE_SABBREVMONTHNAME12    = 79
    LOCALE_SABBREVMONTHNAME13    = 0x100F
    LOCALE_SPOSITIVESIGN         = 80
    LOCALE_SNEGATIVESIGN         = 81
    LOCALE_IPOSSIGNPOSN          = 82
    LOCALE_INEGSIGNPOSN          = 83
    LOCALE_IPOSSYMPRECEDES       = 84
    LOCALE_IPOSSEPBYSPACE        = 85
    LOCALE_INEGSYMPRECEDES       = 86
    LOCALE_INEGSEPBYSPACE        = 87
    LOCALE_FONTSIGNATURE         = 88
    LOCALE_SISO639LANGNAME       = 89
    LOCALE_SISO3166CTRYNAME      = 90
    LOCALE_SYSTEM_DEFAULT        = 0x800
    LOCALE_USER_DEFAULT          = 0x400
)

var (
	// Library
	libkernel32 uintptr

	// Functions
	closeHandle            uintptr
	fileTimeToSystemTime   uintptr
	getLastError           uintptr
	getLogicalDriveStrings uintptr
	getModuleHandle        uintptr
	getNumberFormat        uintptr
	getThreadLocale        uintptr
	getVersion             uintptr
	globalAlloc            uintptr
	globalFree             uintptr
	globalLock             uintptr
	globalUnlock           uintptr
	moveMemory             uintptr
	mulDiv                 uintptr
	setLastError           uintptr
	systemTimeToFileTime   uintptr
	getProfileString       uintptr
)

type (
    ATOM      uint16
    HANDLE    uintptr
    HGLOBAL   HANDLE
    HINSTANCE HANDLE
    HMODULE   HANDLE
    HRSRC     HANDLE
    LCID      uint32
    LCTYPE    uint32
    DWORD     uint32
    DWORD32   uint32
    DWORDLONG uint64   
    DWORD64   uint64   
    LPTSTR    *uint16
    LPCTSTR   *uint16    
    INT       int32
    UINT      uint32
    UINT_PTR  uint32
    LPVOID    uintptr
)

const MAXDWORD DWORD = 4294967295

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

type NUMBERFMT struct {
	NumDigits     uint32
	LeadingZero   uint32
	Grouping      uint32
	LpDecimalSep  *uint16
	LpThousandSep *uint16
	NegativeOrder uint32
}

type SYSTEMTIME struct {
	WYear         uint16
	WMonth        uint16
	WDayOfWeek    uint16
	WDay          uint16
	WHour         uint16
	WMinute       uint16
	WSecond       uint16
	WMilliseconds uint16
}

func init() {
	// Library
	libkernel32 = MustLoadLibrary("kernel32.dll")

	// Functions
	closeHandle = MustGetProcAddress(libkernel32, "CloseHandle")
	fileTimeToSystemTime = MustGetProcAddress(libkernel32, "FileTimeToSystemTime")
	getLastError = MustGetProcAddress(libkernel32, "GetLastError")
	getLogicalDriveStrings = MustGetProcAddress(libkernel32, "GetLogicalDriveStringsW")
	getModuleHandle = MustGetProcAddress(libkernel32, "GetModuleHandleW")
	getNumberFormat = MustGetProcAddress(libkernel32, "GetNumberFormatW")
	getProfileString = MustGetProcAddress(libkernel32, "GetProfileStringW")
	getThreadLocale = MustGetProcAddress(libkernel32, "GetThreadLocale")
	getVersion = MustGetProcAddress(libkernel32, "GetVersion")
	globalAlloc = MustGetProcAddress(libkernel32, "GlobalAlloc")
	globalFree = MustGetProcAddress(libkernel32, "GlobalFree")
	globalLock = MustGetProcAddress(libkernel32, "GlobalLock")
	globalUnlock = MustGetProcAddress(libkernel32, "GlobalUnlock")
	moveMemory = MustGetProcAddress(libkernel32, "RtlMoveMemory")
	mulDiv = MustGetProcAddress(libkernel32, "MulDiv")
	setLastError = MustGetProcAddress(libkernel32, "SetLastError")
	systemTimeToFileTime = MustGetProcAddress(libkernel32, "SystemTimeToFileTime")

}

func CloseHandle(hObject HANDLE) bool {
	ret, _, _ := syscall.Syscall(closeHandle, 1,
		uintptr(hObject),
		0,
		0)

	return ret != 0
}

func FileTimeToSystemTime(lpFileTime *FILETIME, lpSystemTime *SYSTEMTIME) bool {
	ret, _, _ := syscall.Syscall(fileTimeToSystemTime, 2,
		uintptr(unsafe.Pointer(lpFileTime)),
		uintptr(unsafe.Pointer(lpSystemTime)),
		0)

	return ret != 0
}

func GetLastError() uint32 {
	ret, _, _ := syscall.Syscall(getLastError, 0,
		0,
		0,
		0)

	return uint32(ret)
}

func GetLogicalDriveStrings(nBufferLength uint32, lpBuffer *uint16) uint32 {
	ret, _, _ := syscall.Syscall(getLogicalDriveStrings, 2,
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		0)

	return uint32(ret)
}

func GetModuleHandle(lpModuleName *uint16) HINSTANCE {
	ret, _, _ := syscall.Syscall(getModuleHandle, 1,
		uintptr(unsafe.Pointer(lpModuleName)),
		0,
		0)

	return HINSTANCE(ret)
}

func GetNumberFormat(Locale LCID, dwFlags uint32, lpValue *uint16, lpFormat *NUMBERFMT, lpNumberStr *uint16, cchNumber int32) int32 {
	ret, _, _ := syscall.Syscall6(getNumberFormat, 6,
		uintptr(Locale),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpValue)),
		uintptr(unsafe.Pointer(lpFormat)),
		uintptr(unsafe.Pointer(lpNumberStr)),
		uintptr(cchNumber))

	return int32(ret)
}

func GetProfileString(lpAppName, lpKeyName, lpDefault *uint16, lpReturnedString uintptr, nSize uint32) bool {
	ret, _, _ := syscall.Syscall6(getProfileString, 5,
		uintptr(unsafe.Pointer(lpAppName)),
		uintptr(unsafe.Pointer(lpKeyName)),
		uintptr(unsafe.Pointer(lpDefault)),
		lpReturnedString,
		uintptr(nSize),
		0)
	return ret != 0
}

func GetThreadLocale() LCID {
	ret, _, _ := syscall.Syscall(getThreadLocale, 0,
		0,
		0,
		0)

	return LCID(ret)
}

func GetVersion() int64 {
	ret, _, _ := syscall.Syscall(getVersion, 0,
		0,
		0,
		0)
	return int64(ret)
}

func GlobalAlloc(uFlags uint32, dwBytes uintptr) HGLOBAL {
	ret, _, _ := syscall.Syscall(globalAlloc, 2,
		uintptr(uFlags),
		dwBytes,
		0)

	return HGLOBAL(ret)
}

func GlobalFree(hMem HGLOBAL) HGLOBAL {
	ret, _, _ := syscall.Syscall(globalFree, 1,
		uintptr(hMem),
		0,
		0)

	return HGLOBAL(ret)
}

func GlobalLock(hMem HGLOBAL) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(globalLock, 1,
		uintptr(hMem),
		0,
		0)

	return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem HGLOBAL) bool {
	ret, _, _ := syscall.Syscall(globalUnlock, 1,
		uintptr(hMem),
		0,
		0)

	return ret != 0
}

func MoveMemory(destination, source unsafe.Pointer, length uintptr) {
	syscall.Syscall(moveMemory, 3,
		uintptr(unsafe.Pointer(destination)),
		uintptr(source),
		uintptr(length))
}

func MulDiv(nNumber, nNumerator, nDenominator int32) int32 {
	ret, _, _ := syscall.Syscall(mulDiv, 3,
		uintptr(nNumber),
		uintptr(nNumerator),
		uintptr(nDenominator))

	return int32(ret)
}

func SetLastError(dwErrorCode uint32) {
	syscall.Syscall(setLastError, 1,
		uintptr(dwErrorCode),
		0,
		0)
}

func SystemTimeToFileTime(lpSystemTime *SYSTEMTIME, lpFileTime *FILETIME) bool {
	ret, _, _ := syscall.Syscall(systemTimeToFileTime, 2,
		uintptr(unsafe.Pointer(lpSystemTime)),
		uintptr(unsafe.Pointer(lpFileTime)),
		0)

	return ret != 0
}

func GetLocalTime(lpSystemTime *SYSTEMTIME){
	syscall.Syscall(MustGetProcAddress(libkernel32, "GetLocalTime"), 1,
		uintptr(unsafe.Pointer(lpSystemTime)),
		0,
		0)
}

func GetLocaleInfo(Locale LCID,  LCType LCTYPE, lpLCData *uint16, cchData int32) int32{
	ret, _, _ := syscall.Syscall6(MustGetProcAddress(libkernel32, "GetLocaleInfoW"), 4,
		uintptr(Locale),
		uintptr(LCType),
		uintptr(unsafe.Pointer(lpLCData)),
		uintptr(cchData),
		0,
		0)
        
	return int32(ret)
}

func GetLocaleInfoA(Locale LCID,  LCType LCTYPE, lpLCData *uint16, cchData int32) int32{
	ret, _, _ := syscall.Syscall6(MustGetProcAddress(libkernel32, "GetLocaleInfoA"), 4,
		uintptr(Locale),
		uintptr(LCType),
		uintptr(unsafe.Pointer(lpLCData)),
		uintptr(cchData),
		0,
		0)
        
	return int32(ret)
}

func GetCurrentDirectory(nBufferLength DWORD, lpBuffer uintptr) DWORD{
	ret, _, _ := syscall.Syscall(MustGetProcAddress(libkernel32, "GetCurrentDirectoryW"), 2,
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		0)
        
	return DWORD(ret)    
}

func SetCurrentDirectory(lpPathName LPCTSTR) BOOL{
	ret, _, _ := syscall.Syscall(MustGetProcAddress(libkernel32, "SetCurrentDirectoryW"), 1,
		uintptr(unsafe.Pointer(lpPathName)),
        0,
		0)
        
	return BOOL(ret)    
}

func FindResource(hModule HMODULE, lpName, lpType LPCTSTR) HRSRC {
	ret, _, _ := syscall.Syscall(MustGetProcAddress(libkernel32, "FindResourceW"), 3,
		uintptr(hModule),
        uintptr(unsafe.Pointer(lpName)),
		uintptr(unsafe.Pointer(lpType)))
        
	return HRSRC(ret) 
}

func LoadResource(hModule HMODULE, hResInfo HRSRC) HGLOBAL {
	ret, _, _ := syscall.Syscall(MustGetProcAddress(libkernel32, "LoadResource"), 2,
		uintptr(hModule),
        uintptr(hResInfo),
		0)
        
	return HGLOBAL(ret) 
}

func LockResource(hResData HGLOBAL)LPVOID {
	ret, _, _ := syscall.Syscall(MustGetProcAddress(libkernel32, "LockResource"), 2,
		uintptr(hResData),
        0,
		0)
        
	return LPVOID(ret) 
}

func SizeofResource(hModule HMODULE, hResInfo HRSRC)DWORD{
	ret, _, _ := syscall.Syscall(MustGetProcAddress(libkernel32, "SizeofResource"), 2,
		uintptr(hModule),
		uintptr(hResInfo),
		0)
        
	return DWORD(ret) 
}

func FreeResource(hglbResource HGLOBAL) BOOL{
	ret, _, _ := syscall.Syscall(MustGetProcAddress(libkernel32, "FreeResource"), 1,
		uintptr(hglbResource),
		0,
		0)
        
	return BOOL(ret) 
}