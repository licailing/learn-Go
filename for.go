package main

import "fmt"

func main() {
	// var b int = 15
	// var a int
	// numbers := [6]int{1, 2, 3, 5} //数组

	// /* 3种格式
	// for init; condition;increment {}
	// for condition {}
	// for {} 与for(;;)一样
	// */
	// for a := 0; a < 10; a++ {
	// 	fmt.Printf("a 的值为：%d\n", a)
	// }

	// for a < b { //注意：a为0
	// 	a++
	// 	fmt.Printf("a 的值为：%d\n", a)
	// }

	// //for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
	// for i, x := range numbers {
	// 	fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	// }
	wechat_id := "gh_d7200fee047c"
	flag := false
	checkWechatId(wechat_id, flag)

}

func checkWechatId(wechat_id string, flag bool) {
	is_novel := false
	limit_num := 20
	novel_wechat_ids := []string{"gh_51ef65f02d63", "gh_800acd6d3b58", "gh_b7a8fed7abf1", "gh_d7200fee047c", "gh_77fd846b4b14", "gh_a32d8be2c6c0", "gh_4f1913ee1b61", "gh_72b7674de0d5", "gh_d3be7c1baa86"}
	for _, novel_wechat_id := range novel_wechat_ids {
		if wechat_id == novel_wechat_id {
			limit_num = 50
			is_novel = true
			break
		}
		fmt.Println("novel_wechat_id:", novel_wechat_id)
	}

	content := [2]string{
		"亲爱的%s你好，每日二维码推广只能新增20名好友，20名将不能再获得奖励，详情请联系客服MM",
		"亲爱的%s你好，你今日的二维码推广已达20人，不能再获得奖励，请明日再来获取奖励，详情请联系客服",
	}

	if is_novel {
		content = [2]string{
			"亲爱的%s你好，每日二维码推广只能新增50名好友，50名将不能再获得奖励，详情请联系客服MM",
			"亲爱的%s你好，你今日的二维码推广已达50人，不能再获得奖励，请明日再来获取奖励，详情请联系客服",
		}
	}

	fmt.Println("wechat_id=", wechat_id, "limit_num=", limit_num, "is_novel=", is_novel, "content=", content[0], "flag=", flag)
}

/**
a 的值为：0
a 的值为：1
a 的值为：2
a 的值为：3
a 的值为：4
a 的值为：5
a 的值为：6
a 的值为：7
a 的值为：8
a 的值为：9
a 的值为：1
a 的值为：2
a 的值为：3
a 的值为：4
a 的值为：5
a 的值为：6
a 的值为：7
a 的值为：8
a 的值为：9
a 的值为：10
a 的值为：11
a 的值为：12
a 的值为：13
a 的值为：14
a 的值为：15
第 0 位 x 的值 = 1
第 1 位 x 的值 = 2
第 2 位 x 的值 = 3
第 3 位 x 的值 = 5
第 4 位 x 的值 = 0
第 5 位 x 的值 = 0
*/
