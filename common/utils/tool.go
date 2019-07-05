package utils

import (
	"github.com/axgle/mahonia"
	"github.com/json-iterator/go"
)

var jsonIterator  = jsoniter.ConfigCompatibleWithStandardLibrary

//字符串编码转换
func JConvertStr(str string, oldEncoder string, newEncoder string) string {
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr := srcDecoder.ConvertString(str)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

//JSON编码
func JJsonEncode(obj interface{}) ([]byte, error) {
	return jsonIterator.Marshal(obj)
}

//JSON解码
func JJsonDecode(jsonBlob []byte, obj interface{}) error {
	return jsonIterator.Unmarshal(jsonBlob, &obj)
}
