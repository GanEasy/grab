package grab

import (
	"fmt"
)

//DefaultListReader 默认列表匹配器
type DefaultListReader struct {
}

// GetList 获取列表
func (r DefaultListReader) GetList() {
	fmt.Print(`a read`)
}

//DefaultInfoReader 默认详细页匹配器
type DefaultInfoReader struct {
}

// GetInfo 获取详细内容
func (r DefaultInfoReader) GetInfo() {
	fmt.Print(`a read`)
}

// // GetNextURL 获取详细内容
// func (r DefaultInfoReader) GetNextURL() string {
// 	return ``
// }
