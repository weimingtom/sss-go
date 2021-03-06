package sssgo

type keyboard_buffer_struct struct {
	status int
	buffer [BUFSIZE]uint32
	head int
	tail int
	reading bool
	writing bool
}

var (
	s_KBuffer keyboard_buffer_struct
)

func KeyboardGetBufNum(head int, tail int) int {
	if (tail > head) {
		return tail - head
	} else if (tail < head) {
		return BUFSIZE - head + tail
	} else {
		return 0
	}
}

func KeyboardInit() {
    s_KBuffer.head = 0
    s_KBuffer.tail = 0
}

func KeyboardRelease() {
    s_KBuffer.head = 0
    s_KBuffer.tail = 1
    s_KBuffer.buffer[0] = uint32(-1 & 0xFFFFFFFF)
}

func KeyboardChar(key int) {
	if (KeyboardGetBufNum(s_KBuffer.head, s_KBuffer.tail) < BUFSIZE - 1) {
		s := 0
		s_KBuffer.writing = true
		s_KBuffer.buffer[s_KBuffer.tail] = uint32(key) | uint32(s << 8)
		if (s_KBuffer.tail == BUFSIZE - 1) {
			s_KBuffer.tail = 0
		} else {
			s_KBuffer.tail++
		}
		s_KBuffer.writing = false
	}
}

func KeyboardGetKeyboardStatus() uint32 {
    var status uint32
    if (KeyboardGetBufNum(s_KBuffer.head, s_KBuffer.tail) > 0) {
		/*
		MiscTrace("KeyboardGetBufNum: %d\n", 
		   KeyboardGetBufNum(s_KBuffer.head, s_KBuffer.tail));
		*/
        s_KBuffer.reading = true
        status = s_KBuffer.buffer[s_KBuffer.head]
		if (s_KBuffer.head == BUFSIZE - 1) {
			s_KBuffer.head = 0
		} else {
			s_KBuffer.head++
		}
		s_KBuffer.reading = false
        return status
    }
    return 0
}
