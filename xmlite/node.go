package xmldoc

// Node 节点
type Node struct {
	Name     string
	Attrs    *Attrs
	Content  string
	Parent   *Node
	Children []*Node
}

// NewNode 初始化
func NewNode() *Node {
	node := new(Node)
	node.Name = "Node"
	node.Attrs = NewAttrs()
	node.Content = ""
	node.Parent = nil
	node.Children = []*Node{}
	return node
}

// Add 添加
func (node *Node) Add(n *Node) {
	node.Children = append(node.Children, n)
}

// Each 遍历
func (node *Node) Each(includeChild bool, action func(*Node) bool) {
	for i := 0; i < len(node.Children); i++ {
		if action(node.Children[i]) {
			break
		}
		if includeChild {
			node.Children[i].Each(includeChild, action)
		}
	}
}

// Find 查找具备指定属性的节点
func (node *Node) Find(name, attr, value string, includeChild bool) *Node {
	var rn *Node
	node.Each(includeChild, func(n *Node) bool {
		if n.Name == name && n.Attrs.Get(attr) == value {
			rn = n
			return true
		}
		return false
	})
	if rn == nil {
		rn = NewNode()
		rn.Name = name
		rn.Attrs.Set(attr, value)
		node.Add(rn)
	}
	return rn
}

// FindName 查找指定名称的节点
func (node *Node) FindName(name string, includeChild bool) *Node {
	var rn *Node
	node.Each(includeChild, func(n *Node) bool {
		if n.Name == name {
			rn = n
			return true
		}
		return false
	})
	if rn == nil {
		rn = NewNode()
		rn.Name = name
		node.Add(rn)
	}
	return rn
}

// DeleteName 删除指定名称的节点
func (node *Node) DeleteName(name string, includeChild bool) {
	var ns []*Node
	node.Each(includeChild, func(n *Node) bool {
		if n.Name != name {
			ns = append(ns, n)
		}
		return false
	})
	node.Children = ns
}

// Delete 删除指定名称的节点
func (node *Node) Delete(name, attr, value string, includeChild bool) {
	var ns []*Node
	node.Each(false, func(n *Node) bool {
		if n.Name != name || n.Attrs.Get(attr) != value {
			n.Delete(name, attr, value, includeChild)
			ns = append(ns, n)
		}
		return false
	})
	node.Children = ns
}

func (node *Node) toString(tab int, beautify bool) string {
	str := ""
	if beautify {
		if node.Content == "" && len(node.Children) <= 0 {
			str += newStr("    ", tab) + "<" + node.Name + node.Attrs.String() + "/>" + "\r\n"
		} else {
			str += newStr("    ", tab) + "<" + node.Name + node.Attrs.String() + ">"
			if node.Content == "" {
				str += "\r\n"
			}
			str += node.Content
			node.Each(false, func(n *Node) bool {
				str += n.toString(tab+1, beautify)
				return false
			})
			if node.Content == "" && len(node.Children) > 0 {
				str += newStr("    ", tab)
			}
			str += "</" + node.Name + ">"
			if tab != 0 {
				str += "\r\n"
			}
		}
	} else {
		if node.Content == "" && len(node.Children) <= 0 {
			str += "<" + node.Name + node.Attrs.String() + "/>"
		} else {
			str += "<" + node.Name + node.Attrs.String() + ">"
			str += node.Content
			node.Each(false, func(n *Node) bool {
				str += n.toString(tab+1, beautify)
				return false
			})
			str += "</" + node.Name + ">"
		}
	}
	return str
}

// String 字符串
func (node *Node) String() string {
	return node.toString(0, true)
}

func newStr(str string, num int) string {
	rstr := ""
	for i := 0; i < num; i++ {
		rstr += str
	}
	return rstr
}
