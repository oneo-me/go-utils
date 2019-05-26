package fastxml

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// LoadFile 加载 Xml 文件
func LoadFile(file string, v interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err == nil {
		return Unmarshal(data, v)
	}
	return err
}

// Unmarshal 加载 Xml 字符串
func Unmarshal(xmlData []byte, v interface{}) error {
	return xml.Unmarshal(xmlData, v)
}

// Marshal Xml 转字符串
func Marshal(v interface{}, beautify bool) ([]byte, error) {
	var data []byte
	var err error
	if beautify {
		data, err = xml.MarshalIndent(v, "", "    ")
	} else {
		data, err = xml.Marshal(v)
	}
	if err == nil {
		return data, nil
	}
	return nil, err
}

// Save 保存 Xml
func Save(file string, v interface{}, beautify bool) error {
	data, err := Marshal(v, beautify)
	if err == nil {
		err = ioutil.WriteFile(file, data, os.ModePerm)
	}
	return err
}
