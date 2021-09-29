package attrs

import (
	"math"
	"strings"
)

type Attr struct {
	name string
	val  int
}

func HasAttr(a []Attr, s string) bool {
	var tmp bool
	for _, v := range a {
		if v.name == s {
			tmp = true
		}
	}
	return tmp
}

func RemoveAttr(a []Attr, s string) []Attr {
	tmp := a
	for _, v := range tmp {
		if v.name == s {
			v.name = ""
			v.val = 0
		}
	}
	return tmp
}

func UpdateAttr(a []Attr, b Attr) []Attr {
	tmp := a
	for _, v := range tmp {
		if v.name == b.name {
			v.val = b.val
		}
	}
	return tmp
}

func UpdateMultiAttr(a []Attr, b []Attr) []Attr {
	tmp := a
	for _, v1 := range b {
		if v1.val == NonAttrVal {
			tmp = RemoveAttr(tmp, v1.name)
		} else {
			for _, v := range tmp {
				if v.name == v1.name {
					v.val = v1.val
				}
			}
		}
	}
	return tmp
}

func AttrsFromString(s string) []Attr {
	mystr := strings.Split(s, " ")
	var a []Attr
	for _, v := range mystr {
		var b Attr
		b.name = v
	}
	return a
}

var NonAttrVal = math.MinInt32
