package test

import (
	"strconv"
	"strings"
	"fmt"
	"bytes"
)

/**
 * 加密
 * @param unknown_type $str
 * @return string
 */
func Encode(id int) string {
	var spool = [16]string{"kpo0zt1h6a", "m047gl8foj", "wg3h0ca2nl", "j6oyw8kclh", "tods62c4fi", "c9f5aeb62v", "kim7pltyaw", "pkneclgasz", "3kf46cxi2h", "qo83ezx524", "ivbondxc47", "gh6jxkdwbt", "s8pw43ukhf", "09xzmjru63", "wbyn6hai3d", "q7v9264bjl"}
	len_mask := "d6a78fe94c352b1"
	str := strconv.Itoa(id)
	if len(str) == 0 {
		return ""
	}
	total_len := 16
	str_len := len(str)
	if str_len+1 > total_len {
		return ""
	}
	var tmp bytes.Buffer
	ii := 0

	for i := 0; i < str_len; i++ {
		ii = i
		str_byte, _ := strconv.Atoi(string(str[i]))
		tmp.WriteByte(spool[i][str_byte])
	}
	ii += 1

	str_byte, _ := strconv.Atoi(string(str[ii-1]))
	last_char := str_byte

	for i := ii; i < total_len-1; i++ {
		tmp.WriteByte(spool[last_char][i%10])
	}
	tmp.WriteByte(len_mask[str_len-1])
	return tmp.String()
}

func Decode(str string) int64 {
	var spool = [16]string{"kpo0zt1h6a", "m047gl8foj", "wg3h0ca2nl", "j6oyw8kclh", "tods62c4fi", "c9f5aeb62v", "kim7pltyaw", "pkneclgasz", "3kf46cxi2h", "qo83ezx524", "ivbondxc47", "gh6jxkdwbt", "s8pw43ukhf", "09xzmjru63", "wbyn6hai3d", "q7v9264bjl"}
	len_mask := "d6a78fe94c352b1"
	str_len := len(str)
	if str_len == 0 {
		return 0
	}
	total_len := 16
	if str_len != total_len {
		return 0
	}

	var tmp bytes.Buffer
	count := strings.Index(len_mask, str[15:16]) + 1
	for i := 0; i < count; i++ {
		pos := strings.Index(spool[i], str[i:i+1])
		tmp.WriteString(strconv.Itoa(pos))
	}

	if s, err := strconv.ParseInt(tmp.String(), 10, 64); err == nil {
		return s
	} else {
		fmt.Println("xxx")
		return 0
	}

}