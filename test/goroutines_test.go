package test

import "testing"

func TestBuffer_goRoutines(t *testing.T) {
	Buffer_goRoutines()
}

func TestNoBuffer_goRoutines(t *testing.T) {
	NoBuffer_goRoutines()
}

/**
=== RUN   TestEncode
--- PASS: TestEncode (0.00s)
        encode_test.go:12: 加密成功 p43h0ca2nlwg3h06
=== RUN   TestDecode
--- PASS: TestDecode (0.00s)
        encode_test.go:20: 解密成功 12
=== RUN   TestBuffer_goRoutines
sent  0  to buffer_quit
recived  0  from buffer_quit
sent  9  to buffer_quit
recived  9  from buffer_quit
sent  6  to buffer_quit
recived  6  from buffer_quit
sent  7  to buffer_quit
recived  7  from buffer_quit
sent  8  to buffer_quit
recived  8  from buffer_quit
sent  1  to buffer_quit
sent  5  to buffer_quit
sent  2  to buffer_quit
sent  3  to buffer_quit
sent  4  to buffer_quit
recived  1  from buffer_quit
recived  5  from buffer_quit
recived  2  from buffer_quit
recived  3  from buffer_quit
recived  4  from buffer_quit
--- PASS: TestBuffer_goRoutines (0.01s)
=== RUN   TestNoBuffer_goRoutines
sent  2  to quit
recived  2  from quit
sent  6  to quit
recived  6  from quit
sent  7  to quit
sent  3  to quit
sent  8  to quit
sent  4  to quit
sent  0  to quit
sent  1  to quit
sent  9  to quit
sent  5  to quit
recived  7  from quit
recived  3  from quit
recived  8  from quit
recived  4  from quit
recived  0  from quit
recived  1  from quit
recived  9  from quit
recived  5  from quit
--- PASS: TestNoBuffer_goRoutines (0.00s)
PASS
ok      learn-Go/test   0.307s

 */