package perms

type Perm int

func (perm1 Perm) Has(perm2 Perm) bool {
	if perm1&perm2 == 0 {
		return false
	}
	return true
}

func (perm1 Perm) HasNo(perm2 Perm) bool {
	if perm1.Has(perm2) {
		return false
	}
	return true
}

func (perm1 Perm) In(perm2 Perm) bool {
	if perm2&perm1 == 0 {
		return false
	}
	return true
}

func (perm1 Perm) NotIn(perm2 Perm) bool {
	if perm1.In(perm2) {
		return false
	}
	return true
}

func (perm1 Perm) Add(perm2 Perm) Perm {
	return perm1 | perm2
}
