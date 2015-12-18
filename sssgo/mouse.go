package sssgo

const (
	PEN_DOWN = 0xC0
	PEN_MOVE = 0x90
	PEN_UP = 0x20
	PEN_LEAVE = 0x8
)

type mouse_buffer_struct struct {
	status int
	buffer [BUFSIZE]int
	head int
	tail int
	reading bool
	writing bool
}

var (
	s_MBuffer mouse_buffer_struct
)

func MouseGetBufNum(head int, tail int) int {
	if (tail > head) {
		return tail - head
	} else if (tail < head) {
		return BUFSIZE - head + tail
	} else {
		return 0
	}
}

func MouseInit() {
    s_MBuffer.status = 0
    s_MBuffer.writing = false
    s_MBuffer.reading = false
    s_MBuffer.head = 0
    s_MBuffer.tail = 0
}

func MouseRelease() {
    s_MBuffer.head = 0
    s_MBuffer.tail = 1
    s_MBuffer.buffer[0] = -1
}

func MouseMove(x int, y int) {
	if (s_MBuffer.status == 1 && 
		MouseGetBufNum(s_MBuffer.head, s_MBuffer.tail) < BUFSIZE - 1) {
		s_MBuffer.writing = true;
		s_MBuffer.buffer[s_MBuffer.tail] = 
			((x & 0xFFF) << 12) | 
			(y & 0xFFF) | 
			(PEN_MOVE << 24)
		if (s_MBuffer.tail == BUFSIZE - 1) {
			s_MBuffer.tail = 0
		} else {
			s_MBuffer.tail++
		}
		s_MBuffer.writing = false
	}	
}

func MouseLButtonDown(x int, y int) {
	if (MouseGetBufNum(s_MBuffer.head, s_MBuffer.tail) < BUFSIZE - 1) {
		s_MBuffer.status = 1
		s_MBuffer.writing = true
		s_MBuffer.buffer[s_MBuffer.tail] = 
			((x & 0xFFF) << 12) | 
			(y & 0xFFF) | 
			(PEN_DOWN << 24)
		if (s_MBuffer.tail == BUFSIZE - 1) {
			s_MBuffer.tail = 0
		} else {
			s_MBuffer.tail++
		}
		s_MBuffer.writing = false
	}	
}

func MouseLButtonUp(x int ,y int) {
	if (s_MBuffer.status == 1) {
		s_MBuffer.writing = true
		s_MBuffer.buffer[s_MBuffer.tail] = 
			((x & 0xFFF) << 12) | 
			(y & 0xFFF) | 
			(PEN_UP << 24)
		if (s_MBuffer.tail == BUFSIZE - 1) {
			s_MBuffer.tail = 0;
		} else {
			s_MBuffer.tail++
		}
		s_MBuffer.writing = false
		s_MBuffer.status = 0
	}	
}

func MouseGetMouseStatus() int {
    var status int
    if (MouseGetBufNum(s_MBuffer.head, s_MBuffer.tail) > 0) {
		/*
		MiscTrace("MouseGetBufNum: %d\n", 
			MouseGetBufNum(s_MBuffer.head, s_MBuffer.tail));
        */
		s_MBuffer.reading = true
        status = s_MBuffer.buffer[s_MBuffer.head];
		if (s_MBuffer.head == BUFSIZE - 1) {
			s_MBuffer.head = 0
		} else {
			s_MBuffer.head++
		}
		s_MBuffer.reading = false
        return status
    }
    return 0
}
