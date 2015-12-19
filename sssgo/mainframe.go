package sssgo

import (
	"fmt"
	. "github.com/cwchiu/go-winapi"
	"syscall"
	"unsafe"
)

var (
	s_hAppWnd HWND
	s_hCanvas HDC
)

func WndProc(hWnd HWND, message uint32, wParam, lParam uintptr) (result uintptr) {
	var hdc HDC
	var point POINT
	switch message {
		case WM_CREATE:
			s_hAppWnd = hWnd;
			SetWindowPos(hWnd, 0, 0, 0, -1, -1, SWP_NOSIZE);
			s_hCanvas = CanvasInit(hWnd);
			MouseInit();
			KeyboardInit();
			SetTimer(hWnd, 1, TIMER_ELAPSE, 0);
		case WM_PAINT:
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
		case WM_TIMER:
			var nIDEvent UINT = UINT(wParam);
			switch (nIDEvent) {
				case 1:
			}
		case WM_CHAR:
			key := int(wParam);
			fmt.Printf("WM_CHAR: %d\n", key)
			KeyboardChar(key)
		case WM_MOUSEMOVE:
			point.X = int32(LOWORD(uint32(lParam)));
			point.Y = int32(HIWORD(uint32(lParam)));
			fmt.Printf("WM_MOUSEMOVE: %d, %d\n", point.X, point.Y);
			MouseMove(int(point.X), int(point.Y));
		case WM_LBUTTONDOWN:
			point.X = int32(LOWORD(uint32(lParam)));
			point.Y = int32(HIWORD(uint32(lParam)));
			fmt.Printf("WM_LBUTTONDOWN: %d, %d\n", point.X, point.Y);
			MouseLButtonDown(int(point.X), int(point.Y));
		case WM_LBUTTONUP:
			point.X = int32(LOWORD(uint32(lParam)));
			point.Y = int32(HIWORD(uint32(lParam)));
			fmt.Printf("WM_LBUTTONUP: %d, %d\n", point.X, point.Y);
			MouseLButtonUp(int(point.X), int(point.Y));
		case WM_DESTROY:
			KeyboardRelease();
			MouseRelease();
			CanvasRelease();
			PostQuitMessage(0);
	}
	return DefWindowProc(hWnd, message, wParam, lParam)
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

func MainFrameRefresh() {
	InvalidateRect(s_hAppWnd, nil, false);
}

func MainframeResize(w int, h int) {
	var rectWindow RECT
	rectWindow.Left = 0
    rectWindow.Top = 0
	rectWindow.Right = int32(w)
    rectWindow.Bottom = int32(h)
	AdjustWindowRect(&rectWindow, WINDOW_STYLE, false)
	SetWindowPos(s_hAppWnd, HWND(0), 
		0, 0, 
		rectWindow.Right - rectWindow.Left, 
		rectWindow.Bottom - rectWindow.Top, 
		SWP_NOMOVE)
	UpdateWindow(s_hAppWnd)
}

func MainFrameSetTitle(str string) {
	SetWindowText(s_hAppWnd, syscall.StringToUTF16Ptr(str));
}
