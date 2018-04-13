package main

import (
"fmt"
"net/http"
"io/ioutil"
"bytes"
"pholcus/common/goquery"
_ "github.com/go-sql-driver/mysql"
"regexp"
"time"
"strconv"
)

type (
	Yima struct {
		Code   string
		Name   string
		Option string
		Score  string
		Health string
		Xuere  string
		Qixu   string
		Hanshi string
		Xueyu  string
		Xuexu  string
		Ganyu  string
		Shire  string
	}
	YimaData struct {
		Data []*Yima
	}
)

func getUrlRespHtml(strUrl string, postDataStr string) string {
	var respHtml string = ""

	httpClient := &http.Client{}
	var httpReq *http.Request

	postDataBytes := []byte(postDataStr)
	postBytesReader := bytes.NewReader(postDataBytes)

	httpReq, _ = http.NewRequest("POST", strUrl, postBytesReader)

	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Add("Referer", "http://www.huofar.com/ymtest/")
	httpReq.Header.Add("Proxy-Connection","keep-alive")
	httpReq.Header.Add("Upgrade-Insecure-Requests","1")
	httpReq.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36")

	httpResp, err := httpClient.Do(httpReq)

	if err != nil {
		fmt.Printf("http get strUrl=%s\n response error=%s\n", strUrl, err.Error())
	}

	defer httpResp.Body.Close()

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		fmt.Printf("get response for strUrl=%s\n got error=%s\n", strUrl, errReadAll.Error())
	}
	respHtml = string(body)

	return respHtml
}

func save2db(data *YimaData) {
	db, err := ConnectDbForLocalPholcus()
	if err != nil {
		fmt.Printf("连接数据库失败：%v\n", err.Error())
		return
	}
	defer db.Close()

	sql := &bytes.Buffer{}
	args := []interface{}{}
	sql.WriteString("insert into ds_yima_style(code,name,option,score,health,xuere,qixu,hanshi,xueyu,xuexu,ganyu,shire)values")
	for i, c := range data.Data {
		if i > 0 {
			sql.WriteString(",")
		}
		sql.WriteString("(?,?,?,?,?,?,?,?,?,?,?,?)")
		args = append(args, c.Code, c.Name, c.Option, c.Score, c.Health, c.Xuere, c.Qixu, c.Hanshi, c.Xueyu, c.Xuexu, c.Ganyu, c.Shire)
	}

	fmt.Println(sql.String())
	_, err = db.Exec(sql.String(), args...)
	if err != nil {
		fmt.Printf("存入数据库失败：%v\n", err.Error())
		fmt.Println(err.Error())
	}
	fmt.Println("ok")
}

func main() {
	var yimaData YimaData
	baiduMainLoginUrl := "http://www.huofar.com/ymtest/trial/result?test=3&app=&version="
	reg, _ := regexp.Compile(":(.*?)%")
	reg1, _ := regexp.Compile("question\\[(\\d*)\\]=(0.5|1.5|1)")
	reg2, _ := regexp.Compile("[\\d\\.]*%")
	for i, postDataStr := range postDataArr {
		postDataStr = "uid=&test=0&from=&" + postDataStr
		html := &bytes.Buffer{}
		html_str := getUrlRespHtml(baiduMainLoginUrl, postDataStr)
		html.WriteString(html_str)
		doc, err := goquery.NewDocumentFromReader(html)
		if err != nil {
			fmt.Println("解析doc失败")
			break
		}
		t := doc.Find(".typeinfo h1").Text()
		if len(t) == 0 {
			fmt.Println("获取失败")
			return
		}
		s, _ := doc.Find("#rankPointer").Attr("style")
		score_arr := reg.FindStringSubmatch(s)
		score := "0"
		if len(score_arr) > 1 {
			score = score_arr[1]
		}
		score_i, _ := strconv.ParseFloat(score, 64)
		var rel_score float64
		if score_i > 0 {
			rel_score = Round(93-score_i, 2)
		}
		key_arr := reg1.FindStringSubmatch(postDataStr)
		key_name := ""
		if len(key_arr) > 1 {
			key_name = "[" + key_arr[1] + "]=" + key_arr[2]
		}
		hel_str := doc.Find(".rank-desc").Text()
		health_arr := reg2.FindStringSubmatch(hel_str)
		health := ""
		if len(health_arr) > 0 {
			health = health_arr[0]
		}
		t_s := doc.Find(".histogram")
		var xuere, qixu, hanshi, xueyu, xuexu, ganyu, shire string
		if t_s.Length() > 0 {
			xuere, _ = t_s.Find(".histo").Eq(0).Attr("v-score")
			qixu, _ = t_s.Find(".histo").Eq(1).Attr("v-score")
			hanshi, _ = t_s.Find(".histo").Eq(2).Attr("v-score")
			xueyu, _ = t_s.Find(".histo").Eq(3).Attr("v-score")
			xuexu, _ = t_s.Find(".histo").Eq(4).Attr("v-score")
			ganyu, _ = t_s.Find(".histo").Eq(5).Attr("v-score")
			shire, _ = t_s.Find(".histo").Eq(6).Attr("v-score")
		}
		content, _ := doc.Html()
		fmt.Println(i)
		yimaData.Data = append(yimaData.Data, &Yima{Health: health, Xuere: xuere, Qixu: qixu, Hanshi: hanshi, Xueyu: xueyu, Xuexu: xuexu, Ganyu: ganyu, Shire: shire, FormKey: key_name, RelScore: strconv.FormatFloat(rel_score, 'g', 12, 64), FormData: postDataStr, Content: content, Type: t, Score: score})
		time.Sleep(300 * time.Millisecond)
	}
	save2db(&yimaData)
}
