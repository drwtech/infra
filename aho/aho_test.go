package aho

import (
	"math/rand"
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func randSentence() string {
	keywords := []string{"血", "暴", "惨", "残忍", "用于", "定位", "有限", "字符串", "集合", "的", "元素", "在", "输入", "文本", "中", "它", "同时", "匹配", "所有", "算法", "复杂度", "长度", "加上", "搜索"}

	var sentence string
	for i := 0; i < 5; i++ {
		keyword := keywords[rand.Intn(len(keywords))]
		sentence += keyword
		sentence += " "
	}
	return sentence
}

// 场景：适用于敏感词检测，比如一个文本是否包含指定的敏感词（多个）
func TestAHO(t *testing.T) {
	dict := []string{"血", "暴", "惨", "残忍"}
	aho := NewAHO(dict)

	assert(t, aho.Contains("这太残忍了") == true)
}

func TestNormal(t *testing.T) {
	dict := []string{"血", "暴", "惨", "残忍"}
	normal := NewNormal(dict)
	assert(t, normal.Contains("这太残忍了") == true)
}

// BenchmarkAHO
// BenchmarkAHO-12    	 2324970	       479.0 ns/op
func BenchmarkAHO(b *testing.B) {
	dict := []string{"血", "暴", "惨", "残忍"}
	aho := NewAHO(dict)
	for i := 0; i < b.N; i++ {
		sent := randSentence()
		aho.Contains(sent)
	}
}

// BenchmarkNormal
// BenchmarkNormal-12    	 2717008	       443.0 ns/op
func BenchmarkNormal(b *testing.B) {
	dict := []string{"血", "暴", "惨", "残忍"}
	aho := NewNormal(dict)
	for i := 0; i < b.N; i++ {
		sent := randSentence()
		aho.Contains(sent)
	}
}
