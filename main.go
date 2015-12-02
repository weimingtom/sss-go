package main

import (
	"fmt"
	. "github.com/cwchiu/go-winapi"
	"unsafe"
	"syscall"
)

const (
	WINDOW_WIDTH = 800
	WINDOW_HEIGHT = 600
	WINDOW_TITLE = "SimpleScriptSystem version 0.0.1"
	BUFSIZE = 256

	WINDOW_BGCOLOR = 0xCCCCCC
	TIMER_ELAPSE = 1

	SSS_APPNAME = "SimpleScriptSystem"
	SSS_VERSION = "0.0.1"

	SSS_CLASS = "SimpleScriptSystem_GDI"
	SSS_TITLE = SSS_APPNAME + " GDI " + SSS_VERSION

	WINDOW_STYLE = WS_OVERLAPPED | WS_CAPTION | WS_MINIMIZEBOX | WS_SYSMENU;
)

var (
	s_hAppWnd HWND
	s_hCanvas HDC


	s_hBackBufferDC HDC
	s_hBitmap HBITMAP
	s_hOldBitmap HBITMAP
	s_hPen HPEN
	s_hOldPen HPEN
	s_hBrush HBRUSH
	s_hOldBrush HBRUSH
)

func KeyboardRelease() {

}

func MouseRelease() {

}

func WndProc(hWnd HWND, message uint32, wParam, lParam uintptr) (result uintptr) {
	var hdc HDC
	switch message {
		case WM_CREATE:
			s_hAppWnd = hWnd;
			SetWindowPos(hWnd, 0, 0, 0, -1, -1, SWP_NOSIZE);
			s_hCanvas = CanvasInit(hWnd);
			MouseInit();
			KeyboardInit();
			SetTimer(hWnd, 1, TIMER_ELAPSE, 0);
			break;

		case WM_PAINT:
		{
			var ps PAINTSTRUCT
			hdc = BeginPaint(hWnd, &ps);
			BitBlt(hdc,
				ps.RcPaint.Left, ps.RcPaint.Top,
				ps.RcPaint.Right - ps.RcPaint.Left,
				ps.RcPaint.Bottom - ps.RcPaint.Top,
				s_hCanvas,
				ps.RcPaint.Left, ps.RcPaint.Top,
				SRCCOPY);
			EndPaint(hWnd, &ps);
		}
		break;

		case WM_TIMER:
			{
				var nIDEvent UINT = UINT(wParam);
				switch (nIDEvent) {
					case 1:
						break;
				}
			}
			break;

		case WM_DESTROY:
			KeyboardRelease();
			MouseRelease();
			CanvasRelease();
			PostQuitMessage(0);
			break;
	}
	return DefWindowProc(hWnd, message, wParam, lParam)
}

func MiscAppInit() {

}

func MyRegisterClass(hInstance HINSTANCE) ATOM {
	var wcex WNDCLASSEX
	wcex.CbSize = uint32(unsafe.Sizeof(wcex)) //FIXME:
	wcex.Style = CS_HREDRAW | CS_VREDRAW;
	wcex.LpfnWndProc = syscall.NewCallback(WndProc);
	wcex.CbClsExtra = 0;
	wcex.CbWndExtra = 0;
	wcex.HInstance = hInstance;
	wcex.HIcon = LoadIcon(hInstance, (*uint16)(unsafe.Pointer(uintptr(0))));
	wcex.HCursor = LoadCursor(0, (*uint16)(unsafe.Pointer(uintptr(IDC_ARROW))));
	wcex.HbrBackground = (HBRUSH)(COLOR_WINDOW + 1);
	wcex.LpszMenuName = nil;
	wcex.LpszClassName = syscall.StringToUTF16Ptr(SSS_CLASS);
	wcex.HIconSm = LoadIcon(hInstance, (*uint16)(unsafe.Pointer(uintptr(0))));
	return RegisterClassEx(&wcex);
}

func MainFrameGetMsg() int {
	var msg MSG
	if GetMessage(&msg, 0, 0, 0) == FALSE {
		return 1;
	}
	TranslateMessage(&msg);
	DispatchMessage(&msg);
	return 0;
}

func CanvasInit(hWnd HWND) HDC {
	var rect RECT
	var hDC HDC
	GetClientRect(hWnd, &rect);
	hDC = GetDC(hWnd);
	s_hBitmap = CreateBitmap(
		rect.Right - rect.Left,
		rect.Bottom - rect.Top,
		uint32(GetDeviceCaps(hDC, PLANES)),
		uint32(GetDeviceCaps(hDC, BITSPIXEL)),
		nil);
	s_hPen = CreatePen(PS_SOLID, 0, 0)
	s_hBackBufferDC = CreateCompatibleDC(hDC)
	s_hBrush = CreateSolidBrush(WINDOW_BGCOLOR)
	s_hOldBitmap = HBITMAP(SelectObject(s_hBackBufferDC, HGDIOBJ(s_hBitmap)))
	s_hOldPen = HPEN(SelectObject(s_hBackBufferDC, HGDIOBJ(s_hPen)))
	s_hOldBrush = HBRUSH(SelectObject(s_hBackBufferDC, HGDIOBJ(s_hBrush)))
	ReleaseDC(hWnd, hDC);
	return s_hBackBufferDC;
}

func CanvasRelease() {
	SelectObject(s_hBackBufferDC, HGDIOBJ(s_hOldPen))
	SelectObject(s_hBackBufferDC, HGDIOBJ(s_hOldBitmap))
	SelectObject(s_hBackBufferDC, HGDIOBJ(s_hOldBrush))
	DeleteObject(HGDIOBJ(s_hBackBufferDC))
	DeleteObject(HGDIOBJ(s_hBitmap))
	DeleteObject(HGDIOBJ(s_hPen))
	DeleteObject(HGDIOBJ(s_hBrush))
}

func CanvasSetPixel(x int, y int, color int) {
	SetPixel(s_hBackBufferDC, int32(x), int32(y), COLORREF(color))
}

func CanvasGetPixel(x int, y int) int {
	return int(GetPixel(s_hBackBufferDC, int32(x), int32(y)))
}

func CanvasSetColor(color int) {
	var hNewPen HPEN = CreatePen(PS_SOLID, 0, COLORREF(color))
	SelectObject(s_hBackBufferDC, HGDIOBJ(hNewPen))
	DeleteObject(HGDIOBJ(s_hPen))
	s_hPen = hNewPen
}

func CanvasMoveTo(x int, y int) {
	MoveToEx(s_hBackBufferDC, int32(x), int32(y), nil)
}

func CanvasLineTo(x int, y int) {
	LineTo(s_hBackBufferDC, int32(x), int32(y));
}

func CanvasDrawLine(x1 int, y1 int, x2 int, y2 int, color int) {
	CanvasSetColor(color);
	CanvasMoveTo(x1, y1);
	CanvasLineTo(x2, y2);
}

func CanvasLock() {
	Rectangle(s_hBackBufferDC, 0, 0, WINDOW_WIDTH, WINDOW_HEIGHT);
}

func CanvasUnlock() {

}

func CanvasRGB(r int, g int, b int)  int {
	return int(RGB(int32(r), int32(g), int32(b)));
}

func MainFrameRefresh() {
	InvalidateRect(s_hAppWnd, nil, false);
}

func InitInstance(hInstance HINSTANCE, nCmdShow int32) BOOL {
	var hWnd HWND
	var rectWindow RECT
	SetRect(&rectWindow, 0, 0, WINDOW_WIDTH, WINDOW_HEIGHT)
	AdjustWindowRect(&rectWindow, WINDOW_STYLE, false)

	hWnd = CreateWindowEx(0, syscall.StringToUTF16Ptr(SSS_CLASS),
		syscall.StringToUTF16Ptr(SSS_TITLE), WINDOW_STYLE,
		CW_USEDEFAULT, CW_USEDEFAULT,
		rectWindow.Right - rectWindow.Left,
		rectWindow.Bottom - rectWindow.Top,
		0, 0, hInstance, nil)

	if hWnd == 0 {
		return FALSE;
	}
	ShowWindow(hWnd, nCmdShow);
	UpdateWindow(hWnd);
	return TRUE;
}

func MainframeResize(width int, height int) {

}

func MouseInit() {

}

func KeyboardInit() {

}

func main() {
	fmt.Print("hello, world")

	var nCmdShow int32 = 1
	hInstance := GetModuleHandle(nil)
	if hInstance == 0 {
		panic("GetModuleHandle")
	}
	MiscAppInit();
	MyRegisterClass(hInstance);

	if InitInstance(hInstance, nCmdShow) == FALSE {
		return;
	}

	width := 800;
	height := 600;
	MainframeResize(width, height);
	for {
		if MainFrameGetMsg() != 0 {
			break;
		}
		CanvasLock();
		CanvasDrawLine(0, 0, width, height, 0xff0000);
		CanvasUnlock();
		MainFrameRefresh();
	}
}
