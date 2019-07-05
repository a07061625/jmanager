package utils

import (
	"github.com/axgle/mahonia"
)

//字符串编码转换
func JConvertStr(str string, oldEncoder string, newEncoder string) string {
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr := srcDecoder.ConvertString(str)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}