package sssgo

import (
	. "github.com/cwchiu/go-winapi"
)

var (
	s_hBackBufferDC HDC
	s_hBitmap HBITMAP
	s_hOldBitmap HBITMAP
	s_hPen HPEN
	s_hOldPen HPEN
	s_hBrush HBRUSH
	s_hOldBrush HBRUSH
)

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

func CanvasSetColor(color int) {
	var hNewPen HPEN = CreatePen(PS_SOLID, 0, COLORREF(color))
	SelectObject(s_hBackBufferDC, HGDIOBJ(hNewPen))
	DeleteObject(HGDIOBJ(s_hPen))
	s_hPen = hNewPen
}

func CanvasRGB(r int, g int, b int)  int {
	return int(RGB(int32(r), int32(g), int32(b)));
}


func CanvasSetPixel(x int, y int, color int) {
	SetPixel(s_hBackBufferDC, int32(x), int32(y), COLORREF(color))
}

func CanvasGetPixel(x int, y int) int {
	return int(GetPixel(s_hBackBufferDC, int32(x), int32(y)))
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
