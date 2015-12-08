package sssgo

import (
	. "github.com/cwchiu/go-winapi"
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
