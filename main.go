package main

import (
	"fmt"
	. "github.com/cwchiu/go-winapi"
	. "./sssgo"
	"strconv"
)

/*
http://www.cnblogs.com/hustcat/p/4004889.html
rect3 := &Rect{0, 0, 100, 200}
rect4 := &Rect{width: 100, height: 200}
*/

func main() {
	//fmt.Print("SimpleScriptSystem start...\n")
	MiscAppInit();

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
		
		ks := KeyboardGetKeyboardStatus()
		
		if ks != 0 {
			fmt.Printf("keyboard status : " + strconv.Itoa(ks) + "\n")
			MainFrameSetTitle("keyboard status : " + strconv.Itoa(ks) + "\n")
		}
		MainFrameRefresh();
	}
	MiscAppExit();
}
