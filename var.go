package main

import "fmt"
import "github.com/pquerna/ffjson/ffjson"
import "regexp"

// import "strconv"

var x, y int
var ( //这种只能出现在全局变量中，函数体内不支持
	a int
	b bool
)

//并行或同时赋值
var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
// g,h := 123,"hello"
/*会报错
# command-line-arguments
.\var.go:12:1: syntax error: non-declaration statement outside function body
*/

var i int64 = 123456
var m int = 1234454

func main() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h, i, m)
	//0 0 0 false 1 2 123 hello 123 hello

	str := `字数：1930
更新日期:2017-10-10 21:22:13`
	reg, _ := regexp.Compile("字数：([0-9]*)")
	c := reg.FindStringSubmatch(str)
	fmt.Printf("%v", c)
	return

	type (
		ChapterItem struct {
			Book_id    int
			Chapter_id int
			Content    string
			Name       string
			Price      int
			Vip        string
		}
		UserInfo struct {
			Nick_name string
		}
		Chapter struct {
			Data          ChapterItem
			Free          int
			Price_sub_all int
			User_info     UserInfo
		}
	)
	json := `{"status":{"code":0,"msg":"succ"},"free":1,"read_status":2,"price_sub_all":0,"user_info":{"uid":"58351112","nick_name":"\u5f69\u73b2\u5b89\u5a1c","sub_auto":0,"balance":0,"coupon_balance":0},"data":{"book_id":327976,"chapter_id":9980658,"name":"\u7b49\u7ea7\u5212\u5206","price":0,"price_extend":[],"content":"\u51e1\u6b66\u5883\uff1a\u6709\u5341\u4e2a\u9636\u6bb5\uff0c\u4ece\u4e00\u5230\u5341\u5206\u522b\u662f\u4e3a\u70bc\u6c14\u3001\u6dec\u4f53\u3001\u6b66\u4f53\u3001\u901a\u8109\u3001\u771f\u6c14\u3001\u795e\u8bc6\u3001\u771f\u7f61\u3001\u795e\u529b\u3001\u771f\u5f62\u3001\u5927\u5706\u6ee1\u5341\u4e2a\u5883\u754c\u3002\n\n\u771f\u6b66\u5883\uff1a\u4e5d\u4e2a\u5883\u754c\uff0c\u5c31\u662f\u4e00\u5230\u4e5d\u6bb5\u3002\n\n\u6781\u81f4\u4e09\u5927\u5883\uff1a\u7075\u6b66\u5883\u3001\u9b42\u6b66\u5883\u3001\u767e\u70bc\u5883\uff08\u90fd\u5206\u4e3a\u524d\u4e2d\u540e\u4e09\u671f\uff09\n\n\u6d85\u69c3\u5883\uff1a\n\n\u6d85\u69c3\u6709\u4e5d\u52ab\uff0c\u6e21\u8fc7\u4e5d\u52ab\uff0c\u5c31\u7b97\u5b8c\u5168\u6d85\u69c3\uff0c\u6210\u4e3a\u4ed9\u4eba\u3002\n\n\u6d85\u69c3\u4e4b\u540e\u5c31\u662f\u4ed9\u4eba\u7684\u5883\u754c\uff0c\u4e3b\u89d2\u8ddd\u79bb\u4ed9\u4eba\u5883\u754c\u8fd8\u6709\u4e00\u4e9b\u5e74\u5934\uff0c\u6682\u65f6\u4e0d\u5217\u4e3e\u3002\n\n\u4fee\u795e\u9053\u7684\u5883\u754c\uff1a\n\n\u795e\u9053\uff1a\n\n\u7b2c\u4e00\u5927\u9636\u6bb5\uff1a\u5a74\u513f\uff0c\u5e7c\u513f\uff0c\u6210\u957f\uff0c\u8715\u53d8\uff0c\u6210\u719f\uff0c\u5b8c\u7f8e\n\n\u7b2c\u4e8c\u5927\u9636\u6bb5\uff1a\u51fa\u795e\u5165\u5316\n\n\u4e39\u836f\uff1a\u51e1\u7ea7\u3001\u7075\u7ea7\u3001\u7384\u7ea7\u3001\u5730\u7ea7\u3001\u5929\u7ea7\u3001\u4ed9\u7ea7\uff08\u5206\u4e0b\u54c1\u3001\u4e2d\u54c1\u3001\u4e0a\u54c1\uff01\uff09\n\n\u70bc\u4e39\u5e08\u7684\u7b49\u7ea7\uff1a\n\n\u4e00\u7ea7\u70bc\u4e39\u5e08\u5230\u4e5d\u7ea7\u3001\u7136\u540e\u662f\u80fd\u70bc\u5236\u5730\u7ea7\u4e2d\u4e0b\u7684\u662f\u5927\u4e39\u5e08\uff0c\u4e39\u5b97\uff0c\u4e39\u738b\uff0c\n\n\u61c2\u5f97\u70bc\u4e39\u5c31\u662f\u4e00\u7ea7\u70bc\u4e39\u5e08\uff1a\n\n\u4e8c\u7ea7\uff1a\u9700\u8981\u61c2\u5f97\u70bc\u5236\u51e1\u7ea7\u4e2d\u4e0a\u54c1\u5404\u4e09\u79cd\n\n\u4e09\u7ea7\uff1a\u7075\u7ea7\u4e0b\u54c1\u4e09\u79cd\n\n\u56db\u7ea7\uff1a\u7075\u7ea7\u4e2d\u54c1\u4e09\u79cd\n\n\u4e94\u7ea7\uff1a\u7075\u7ea7\u4e0a\u54c1\u4e09\u79cd\n\n\u516d\u7ea7\uff1a\u7384\u7ea7\u4e0b\u54c1\u4e09\u79cd\n\n\u4e03\u7ea7\uff1a\u7384\u7ea7\u4e2d\u54c1\u4e09\u79cd\n\n\u516b\u7ea7\uff1a\u7384\u7ea7\u4e0a\u54c1\u4e09\u79cd\n\n\u4e5d\u7ea7\uff1a\u5730\u7ea7\u4e0b\u54c1\u4e09\u79cd\n\n\u5927\u4e39\u5e08\uff1a\u5730\u7ea7\u4e2d\u54c1\u4e09\u79cd\n\n\u4e39\u5b97\uff1a\u5730\u7ea7\u4e0a\u54c1\u4e09\u79cd\n\n\u4e39\u738b\uff1a\u53ea\u8981\u61c2\u5f97\u70bc\u5236\u5929\u7ea7\u4e39\u5c31\u662f\u4e39\u738b\u3002\n\n\u6682\u65f6\u5c31\u8fd9\u4e9b\uff0c\u4ee5\u540e\u4f1a\u6dfb\u52a0\u7684\u3002","vip":"N","platform_vip":[],"is_jingpin":"N","limit_free_tip":0,"created_at":1363361881000}}`
	var chapter Chapter
	err := ffjson.Unmarshal([]byte(json), &chapter)
	if err != nil {
		fmt.Printf("小说章节内容解析错误： %v\n", err)
		return
	}
	fmt.Println(chapter.Free)
	fmt.Printf("%v", chapter.Data)
	fmt.Printf("%v", chapter.User_info)
}
