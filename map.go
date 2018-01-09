package main

import "fmt"

func main() {
	/* map是一种无序的键值对集合（使用hash表实现） */
	/* var map_variable map[key_data_type]value_data_type */
	/* map_variable = make(map[key_data_type]value_data_type) */
	var countryCapitalMap map[string]string

	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japen"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	/*
		Capital of Japen is Tokyo
		Capital of India is New Delhi
		Capital of France is Paris
		Capital of Italy is Rome
	 */

	/* 查看元素是否在集合中存在 */
	captinal,ok := countryCapitalMap["United State"]

	if !ok {
		fmt.Print("Capital of United States is not present")
	} else {
		fmt.Println("Capital of United States is",captinal)
	}

	/*
		Capital of United States is not present
	 */

	 /* 删除集合的元素 */
	delete(countryCapitalMap,"France")
}
