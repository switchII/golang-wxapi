package search

import (
	"github.com/huichen/sego"
)

// 分词处理
var segmenter *sego.Segmenter

// 加载词典
func getSegmenter() *sego.Segmenter {
	if segmenter == nil {
		segmenter = &sego.Segmenter{}
		segmenter.LoadDictionary("assets/dictionary.txt")
	}
	return segmenter
}

// 得到分词
func Keyword(word string) []string {
	text := []byte(word)
	s := getSegmenter().Segment(text)
	output := sego.SegmentsToSlice(s, false)
	return output
}
