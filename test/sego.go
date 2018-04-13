package main

import (
	"fmt"
	"github.com/huichen/sego"
)

func main() {
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("E:/go_workspace/src/learn-Go/test/dictionary.txt")

	// 分词
	text := []byte("任何一个用户可以上传原创内☺容的网站、App、直播平台等■载体，就是▲UGC产品。因为※内容是网民自主上传，所以存在着各种各样的运营风险。")
	segments := segmenter.Segment(text)

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	//fmt.Println(sego.SegmentsToString(segments, false))
	//搜索模式
	//fmt.Println(sego.SegmentsToString(segments, true))

	fmt.Println(sego.SegmentsToSlice(segments,true))
}
