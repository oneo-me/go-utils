package xmldoc

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"oneo/smilepath"
	"os"
	"path/filepath"
	"strings"
)

// Doc 文档
type Doc struct {
	Header   string
	RootNode *Node
}

// NewDoc 初始化
func NewDoc() *Doc {
	doc := new(Doc)
	doc.Header = "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
	doc.RootNode = NewNode()
	doc.RootNode.Name = "XmlNode"
	return doc
}

// toString 字符串
func (doc *Doc) toString(beautify bool) string {
	str := ""
	if beautify {
		if doc.Header != "" {
			str += doc.Header + "\r\n"
		}
		str += doc.RootNode.toString(0, beautify)
	} else {
		str += doc.Header + doc.RootNode.toString(0, beautify)
	}
	return str
}

// String 字符串
func (doc *Doc) String(beautify bool) string {
	str := ""
	if beautify {
		if doc.Header != "" {
			str += doc.Header + "\r\n"
		}
		str += doc.RootNode.toString(0, beautify)
	} else {
		str += doc.Header + doc.RootNode.toString(0, beautify)
	}
	return str
}

// Save 保存到文件
func (doc *Doc) Save(file string, beautify bool) error {
	return doc.SaveContent(file, doc.toString(beautify))
}

// SaveContent 保存内容
func (doc *Doc) SaveContent(file, content string) error {
	dir := filepath.Dir(file)
	smilepath.Create(dir, false)
	writeFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer writeFile.Close()
	io.WriteString(writeFile, content)
	return nil
}

// LoadFile 从文件读取
func LoadFile(file string) *Doc {
	if !smilepath.FileExists(file) {
		return NewDoc()
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return Load(content)
}

// Load 从字符串加载
func Load(content []byte) *Doc {
	doc := NewDoc()
	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	var tree []*Node
	var currentNode *Node

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}
		switch element := token.(type) {
		case xml.ProcInst:
			doc.Header = "<?" + element.Target + " " + string(element.Inst) + "?>"
		case xml.StartElement:

			n := len(tree)
			if n == 0 {
				currentNode = NewNode()
				doc.RootNode = currentNode
			} else {
				currentNode = NewNode()
				currentNode.Parent = tree[n-1]
				tree[n-1].Add(currentNode)
			}

			currentNode.Name = element.Name.Local
			for _, attr := range element.Attr {
				currentNode.Attrs.Set(attr.Name.Local, attr.Value)
			}

			tree = append(tree, currentNode)
		case xml.EndElement:
			currentNode = nil
			tree = tree[:len(tree)-1]
		case xml.CharData:
			if str := strings.TrimSpace(string(element)); str != "" && []rune(str)[0] != 65279 {
				currentNode.Content = string(element)
			}
		}
	}

	return doc
}

func LoadFile(file string) (*Doc, error) {

}

func Unmarshal(xmlData []byte) (*Doc, error) {

}

func Marshal()
