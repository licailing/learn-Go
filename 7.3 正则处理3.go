package main

import (
	"fmt"
	"regexp"
	"strconv"
	"math"
	"crypto/md5"
	"io"
	"encoding/hex"
	"bytes"
	"github.com/pquerna/ffjson/ffjson"
	"go.tdamm.com/module/wechat"
	"encoding/json"
	"net/url"
	"math/rand"
)

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	src := []byte(`
		call hello alice
		hello bob
		call hello eve
		`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
	/*
		hello('alice')
		hello('eve')
	*/

	str := "question[117]=0&question[144]=0&question[145]=0&question[146]=0&question[148]=0&question[149]=0&question[150]=0&question[151]=0&question[152]=0&question[187]=0&question[154]=0&question[186]=0&question[156]=0&question[157]=0&question[164]=0&question[188]=0&question[196]=0&question[162]=0&question[155]=0&question[189]=0&question[177]=0&question[191]=0&question[161]=0&question[165]=0&question[190]=0&question[171]=0&question[178]=0&question[192]=0&question[200]=0&question[201]=0&question[202]=0&question[203]=0&question[204]=0&question[205]=0&question[206]=0&question[207]=0&question[208]=0&question[209]=0&question[212]=0&question[210]=0&question[211]=0&question[185]=0&question[183]=0"
	reg, _ := regexp.Compile("\\[\\d*\\]=(?:0.5|1.5|1)")
	str_arr := reg.FindAllString(str,-1)
	name := &bytes.Buffer{}
	for i, v := range str_arr{
		if i > 0 {
			name.WriteString("&")
		}
		name.WriteString(v)
	}
	fmt.Println(name.String())
	s := "89.8"
	s_i, _ := strconv.ParseFloat(s,64)
	score := 93 - s_i
	fmt.Println(s_i, score)

	fmt.Println(Round(score,2))

	str1 := `{"uid":2168838,"options":{"145":"0.5","146":"0.5","148":"0.5","149":"0.5","150":"0.5","151":"0.5","152":"0.5","187":"0.5","154":"0.5","186":"0.5","156":"0.5","157":"0.5","164":"0.5","188":"0.5","196":"0.5","162":"0.5","155":"0.5","189":"0.5","177":"0.5","191":"0.5","161":"0.5","165":"0.5","190":"0.5","171":"0.5","178":"0.5","192":"0.5","200":"0.5","201":"0.5","202":"0.5","203":"0.5","204":"0.5","205":"0.5","206":"0.5","207":"0.5","208":"0.5","209":"0.5","212":"0.5","210":"0.5","211":"0.5"}}`
	type Args struct {
		Uid     int
		Options map[string]string
	}

	var data Args
	ffjson.Unmarshal([]byte(str1), &data)
	options := []string{"117", "144", "145", "146", "148", "149", "150", "151", "152", "187", "154", "186", "156", "157", "164", "188", "196", "162", "155", "189", "177", "191", "161", "165", "190", "171", "178", "192", "200", "201", "202", "203", "204", "205", "206", "207", "208", "209", "212", "210", "211", "185", "183"};
	postValues := url.Values{}
	postValues.Set("uid", "")
	postValues.Set("test", "0")
	postValues.Set("from", "")
	for _, op := range options {
		value := "0"
		if key, ok := data.Options[op]; ok {
			value = key
		}
		postValues.Set("question["+op+"]", value)
	}
	postDataStr := postValues.Encode()
	fmt.Println(postDataStr)
	if key,ok := data.Options["111"];ok {
		fmt.Println(key)
	}
	fmt.Println(Md5(""))

	tempdata := map[string]map[string]string{"first": {"value": "", "color": "#000"}, "keyword1": {"value": "史上最准月经测试", "color": "#000"}, "keyword2": {"value": "黎彩玲", "color": "#000"}, "keyword3": {"value": "健康型", "color": "#000"}, "remark": {"value": "\n\n点击查看测试结果详情>>", "color": "#000"}}
	template := map[string]interface{}{"touser": "open_id", "template_id": "{temp_id:TestShape}", "url": wechat.Forward("gh_b5836bab3d35", "http://m.tdamm.com/doctor/index/ym_show?id=1"), "data": tempdata}
	bytes, _ := json.Marshal(template)
	body := string(bytes)
	fmt.Println(body)
	var rps map[string]interface{}
	_ = json.Unmarshal([]byte(body), &rps)
	a_data := rps["data"].(map[string]interface{})
	a_keywords2 := a_data["keyword2"].(map[string]interface{})
	nickname := a_keywords2["value"].(string)
	a_keyword3 := a_data["keyword3"].(map[string]interface{})
	style_name := a_keyword3["value"].(string)
	openid := rps["touser"].(string)
	url_str:= rps["url"].(string)
	u,_ := url.Parse(url_str)
	url_query,_ := url.ParseQuery(u.RawQuery)
	wechat_id := url_query["wechat_id"][0]
	to := url_query["to"][0]
	u1,_ := url.Parse(to)
	url1_query,_ := url.ParseQuery(u1.RawQuery)
	id := url1_query["id"][0]
	fmt.Println(nickname,style_name,openid,wechat_id,id)

}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}