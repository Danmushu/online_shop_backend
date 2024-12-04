package global

type MAC string

func (mac MAC) In(li []string) bool {
	for i := 0; i < len(li); i++ {
		if string(mac) == li[i] {
			return true
		}
	}
	return false
}

func (mac MAC) NotIn(li []string) bool {
	return !mac.In(li)
}
