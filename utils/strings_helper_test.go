package utils_test

import (
	"testing"

	"github.com/codeh007/mtgolib/utils"
)

func Test_WildcardMatch(t *testing.T) {
	//能匹配
	if utils.WildcardMatch("t*e*s*t", "ttttteeeeeeeesttttttt") == false {
		t.Fail()
	}
	//不能匹配
	if utils.WildcardMatch("t*e*s*t", "tteeeeeettttttt") == true {
		t.Fail()
	}
	//能匹配
	if utils.WildcardMatch("http*://www.baidu.com/", "https://www.baidu.com/") == false {
		t.Fail()
	}
	//能匹配
	if utils.WildcardMatch("http*://www.baidu.com/*", "https://www.baidu.com/") == false {
		t.Fail()
	}
	if utils.WildcardMatch("http*://www.baidu.com/*/*", "https://www.baidu.com/") == true {
		t.Fail()
	}
	if utils.WildcardMatch("http*://www.baidu.com/*/*", "https://www.baidu.com/a/b") == false {
		t.Fail()
	}
	if utils.WildcardMatch("http*://www.baidu.com/*/*?id=3", "https://www.baidu.com/a/b?id=3") == false {
		t.Fail()
	}
}

func Split(r rune) bool {
	return r == ':' || r == '.'
}
func Test_CustomSplitStr(t *testing.T) {
	if utils.CustomSplitStr(`abc;ddd|eee|;;`)[2] != "eee" {
		t.Fail()
	}
	if utils.CustomSplitStr(`abc;ddd eee|;;`)[2] != "eee" {
		t.Fail()
	}
	if utils.CustomSplitStr("abc;ddd\teee|;;")[2] != "eee" {
		t.Fail()
	}
	if utils.CustomSplitStr("abc;ddd\neee|;;")[2] != "eee" {
		t.Fail()
	}
	if utils.CustomSplitStr(`abc;ddd
	eee|;;`)[2] != "eee" {
		t.Fail()
	}
}
