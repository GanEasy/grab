package grab

import (
	"math/rand"
	"strings"
)

// 函数包

/*******https://github.com/astaxie/beego/blob/master/utils/slice.go start*********/

// InSliceIface checks given interface in interface slice.
func InSliceIface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// SliceUnique cleans repeated values in slice. 去重
func SliceUnique(slice []interface{}) (uniqueslice []interface{}) {
	for _, v := range slice {
		if !InSliceIface(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

// SliceShuffle shuffles a slice. 重组
func SliceShuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

/*******https://github.com/astaxie/beego/blob/master/utils/slice.go end*********/

// InArr 字符串被包含
func InArr(str string, arr []string) bool {
	for _, vv := range arr {
		if vv == str {
			return true
		}
	}
	return false
}

//Intersection 字符串交集
func Intersection(arr1 []string, arr2 []string) []string {
	for _, v := range arr1 {
		if !InArr(v, arr2) {
			arr2 = append(arr2, v)
		}
	}
	return arr2
}

//Union 字符串并集
func Union(arr1 []string, arr2 []string) (ret []string) {
	for _, v := range arr1 {
		if InArr(v, arr2) {
			ret = append(ret, v)
		}
	}
	return ret
}

//Nonion 字符串非集
func Nonion(arr1 []string, arr2 []string) (ret []string) {
	arr3 := Intersection(arr1, arr2)
	arr4 := Union(arr1, arr2)
	for _, v := range arr3 {
		if !InArr(v, arr4) {
			ret = append(ret, v)
		}
	}
	return ret
}

//Jaccard  杰卡德（Jaccard）相似系数
func Jaccard(arr1 []string, arr2 []string) float64 {
	return float64(len(Union(arr1, arr2))) / float64(len(Intersection(arr1, arr2)))
}

//JaccardMateGetURL  杰卡德（Jaccard）相似系数 匹配出目标url
/**
快速获取固定结构动态链接
url 为要验证参考的链接地址
demo1 有效的学习链接地址1，与url具有相同结构
demo2 有效的学习链接地址2，与url具有相同结构
to1   有效果的目标链接地址1， 将url变量替换到to1结构中 (to1为空时，保持原有结构)

todo 	t.Fatal(JaccardMateGetURL(`http://book.zongheng.com/book/book/658887.html`, `http://book.zongheng.com/book/769150.html`, `http://book.zongheng.com/book/316562.html`, `http://book.zongheng.com/showchapter/769150.html`))
相同值不同位置时抽风了。。
*/
func JaccardMateGetURL(url, demo1, demo2, to1 string) (string, bool) {
	// demo1,2的标签
	demotag1 := strings.Split(GetTag(demo1), ",")
	demotag2 := strings.Split(GetTag(demo2), ",")

	// url的标签
	urltag := strings.Split(GetTag(url), ",")

	if len(demotag1) != len(urltag) {
		return url, false
	}

	eduUn := Union(demotag1, demotag2)
	eduIn := Intersection(demotag1, demotag2)

	demoParamUn := Union(demotag1, Nonion(eduIn, eduUn))

	paramRp := Nonion(urltag, eduUn)

	if len(demoParamUn) != len(paramRp) {
		return url, false
	}
	if to1 != `` {
		to := to1
		for i, val := range demoParamUn {
			// t.Fatal(i, val)
			to = strings.Replace(to, val, paramRp[i], -1)
		}
		return to, true
	}
	return url, true

}
