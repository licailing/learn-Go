package test

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var id = 12
	if str := Encode(id); str == "" {
		t.Error("加密失败")
	} else {
		t.Log("加密成功",str)
	}
}

func TestDecode(t *testing.T) {
	if id := Decode("p43h0ca2nlwg3h06");id==0 {
		t.Error("解密失败")
	} else {
		t.Log("解密成功",id)
	}
}

