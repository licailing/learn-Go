package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

type (
	ApiRequest struct {
		Url      string
		PostDict map[string]string
	}
)

func getUrlRespHtml(strUrl string, postDict map[string]string) string {
	var respHtml string = ""

	httpClient := &http.Client{}
	var httpReq *http.Request
	if nil == postDict {
		httpReq, _ = http.NewRequest("GET", strUrl, nil)
	} else {
		postValues := url.Values{}
		for postKey, PostValue := range postDict {
			postValues.Set(postKey, PostValue)
		}
		postDataStr := postValues.Encode()
		postDataBytes := []byte(postDataStr)
		postBytesReader := bytes.NewReader(postDataBytes)
		httpReq, _ = http.NewRequest("POST", strUrl, postBytesReader)
		httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

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

func pay(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/pay"

	api.PostDict["words"] = "最近牙痛，怎么办？"
	api.PostDict["amount"] = "1"
	api.PostDict["fee"] = "0.01"
	api.PostDict["num"] = "2"
	api.PostDict["anony"] = "0"
	api.PostDict["square"] = "0"
}

func share(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/share"

	api.PostDict["id"] = "1"
	api.PostDict["page"] = "pages/pack-detail"
}

func show(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/show"

	api.PostDict["id"] = "1"
	fmt.Sprintf("url:%s,dict:%v", api.Url, api.PostDict)
}

func get_solutions(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/get_solutions"

	api.PostDict["id"] = "51"
	api.PostDict["page"] = "1"
}

func add_solution(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/add_solution"

	api.PostDict["id"] = "2"
	api.PostDict["voice"] = ""
	api.PostDict["len"] = ""
	api.PostDict["content"] = "吃点消炎药，好的快点"
}

func adopt(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/adopt"

	api.PostDict["id"] = "2"
	api.PostDict["soulID"] = "2"
}

func user_data_for_send(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/user_data_for_send"
}

func send_records(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/send_records"

	api.PostDict["page"] = "1"
}

func user_data_for_get(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/user_data_for_get"
}

func get_records(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/get_records"

	api.PostDict["page"] = "1"
}

func get_square_base_data(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/get_square_base_data"
}

func get_square_data(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/get_square_data"

	api.PostDict["page"] = "1"
}

func save_temp_id(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/save_temp_id"

	api.PostDict["formId"] = "12334"
}

func get_my_balance(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/get_my_balance"
}

func withdraw(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/withdraw"

	api.PostDict["amount"] = "1"
	api.PostDict["fee"] = "0.01"
}

func detectSensWords(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/detectSensWords"

	api.PostDict["content"] = "你妹的，大爷，怎么不去死"
}

func filterSensWords(api *ApiRequest) {
	api.Url = "https://m.tdamm.com/weapp/api/trouble/filterSensWords"

	api.PostDict["content"] = "你妹的，大爷，怎么不去死"
}

func Request(api *ApiRequest) {
	token := "73f0c30a77ce7e987b5a6af1dafcf9dc6140b3b4544bba15df5641b51223fbfe0061653af2c99ae3113d4bad17de7f52cc62f87f3b41af0fee444fa908674874xVccOZkeZriW3y3nCCZR0kg1T6ved6QfNNV%252FTm6x%252FhWQYx6XKVRQWJY3QOxbGYGUm%252B9Q%252FpXF07jJIfnxi6eCL4CoUNEIv%252FX2d0QleybazgPFkhifoSJnw0guvpWn8%252FiXrlRJ7GO%252F3stD7VV15NEKQg%253D%253D"
	api.Url = api.Url + "?token=" + token
	html := getUrlRespHtml(api.Url, api.PostDict)
	fmt.Println(html)
}

func main() {
	var api ApiRequest
	api.PostDict = make(map[string]string)
	//pay(&api)
	//share(&api)
	//show(&api)
	get_solutions(&api)
	//add_solution(&api)
	//adopt(&api)
	//user_data_for_send(&api)
	//send_records(&api)
	//user_data_for_get(&api)
	//get_records(&api)
	//get_square_base_data(&api)
	//get_square_data(&api)
	//get_my_balance(&api)
	//withdraw(&api)
	//detectSensWords(&api)
	//filterSensWords(&api)
	Request(&api)
}
