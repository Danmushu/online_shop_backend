package global

type MAC string

// In 方法检查当前MAC实例是否在提供的字符串切片li中
func (mac MAC) In(li []string) bool {
	for i := 0; i < len(li); i++ {
		if string(mac) == li[i] {
			return true
		}
	}
	return false
}

// NotIn 方法检查当前MAC实例是否不在提供的字符串切片li中
func (mac MAC) NotIn(li []string) bool {
	return !mac.In(li)
}
