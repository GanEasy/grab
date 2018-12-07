package grab

import (
	"strings"
	"testing"
)

func Test_Jaccard(t *testing.T) {

	// struct Str {[]string}
	var str1 = []string{`A`, `b`, `c`, `d`}
	var str2 = []string{`A`, `B`, `c`, `E`, `F`}

	t.Fatal(Jaccard(str1, str2))
}

func Test_JaccardURL(t *testing.T) {

	urltag1 := strings.Split(GetTag(`http://book.zongheng.com/`), ",")
	urltag2 := strings.Split(GetTag(`http://book.zongheng.com/chapter/176024`), ",")

	// t.Fatal(GetTag(`https://blog.csdn.net/Jack_lyp2017/article/details/78741839`))

	// t.Fatal(len(urltag1))
	t.Fatal(Jaccard(urltag1, urltag2))
}

func Test_JaccardMateURL(t *testing.T) {

	urltag1 := strings.Split(GetTag(`http://book.zongheng.com/book/769150.html`), ",")
	urltag2 := strings.Split(GetTag(`http://book.zongheng.com/book/316562.html`), ",")

	urltag3 := strings.Split(GetTag(`http://book.zongheng.com/book/628887.html`), ",")

	eduUn := Union(urltag1, urltag2)
	eduIn := Intersection(urltag1, urltag2)

	paramUn := Union(urltag1, Nonion(eduIn, eduUn))

	paramRp := Nonion(urltag3, eduUn)

	to := `http://book.zongheng.com/showchapter/769150.html`
	for i, val := range paramUn {
		// t.Fatal(i, val)
		to = strings.Replace(`http://book.zongheng.com/showchapter/769150.html`, val, paramRp[i], -1)
	}

	// ut3 := Union(urltag3)

	t.Fatal(paramUn, eduUn, eduIn, to)
	// http://book.zongheng.com/book/769150.html
	//

	// t.Fatal(GetTag(`https://blog.csdn.net/Jack_lyp2017/article/details/78741839`))

	// t.Fatal(len(urltag1))
	t.Fatal(Jaccard(urltag1, urltag2), Jaccard(urltag1, urltag3))
}

func Test_JaccardMateGetURL(t *testing.T) {

	t.Fatal(JaccardMateGetURL(`http://book.zongheng.com/book/book/658887.html`, `http://book.zongheng.com/book/769150.html`, `http://book.zongheng.com/book/316562.html`, `http://book.zongheng.com/showchapter/769150.html`))
}
