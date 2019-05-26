package xmldoc

// Attrs 属性集合
type Attrs struct {
	kvs []*Attr
}

// NewAttrs 初始化
func NewAttrs() *Attrs {
	attrs := new(Attrs)
	attrs.kvs = []*Attr{}
	return attrs
}

// Set 设置属性
func (attrs *Attrs) Set(name, value string) {
	isSet := false
	attrs.Each(func(attr *Attr) bool {
		if attr.Key == name {
			attr.Value = value
			isSet = true
			return true
		}
		return false
	})
	if !isSet {
		attr := new(Attr)
		attr.Key = name
		attr.Value = value
		attrs.kvs = append(attrs.kvs, attr)
	}
}

// Get 获取属性
func (attrs *Attrs) Get(name string) string {
	r := ""
	attrs.Each(func(attr *Attr) bool {
		if attr.Key == name {
			r = attr.Value
			return true
		}
		return false
	})
	return r
}

// Delete 删除属性
func (attrs *Attrs) Delete(name string) {
	var as []*Attr
	attrs.Each(func(attr *Attr) bool {
		if attr.Key != name {
			as = append(as, attr)
		}
		return false
	})
	attrs.kvs = as
}

// Clear 清理属性
func (attrs *Attrs) Clear() {
	attrs.kvs = []*Attr{}
}

// Each 遍历属性
func (attrs *Attrs) Each(action func(attr *Attr) bool) {
	for _, attr := range attrs.kvs {
		if action(attr) {
			break
		}
	}
}

// String 字符串
func (attrs *Attrs) String() string {
	str := ""
	attrs.Each(func(attr *Attr) bool {
		str += " " + attr.Key + "=\"" + attr.Value + "\""
		return false
	})
	return str
}
