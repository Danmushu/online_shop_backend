package perms

const (
	PermOP      = Perm(1 << 1)
	PermCode    = Perm(1 << 2) // 查询激活码
	PermAccount = Perm(1 << 3) // 账户封禁
	PermNews    = Perm(1 << 4) // 新闻
	PermNote    = Perm(1 << 5) // 笔记

	PermOrder     = Perm(1 << 7) // 订单
	PermProChange = Perm(1 << 8) //
	PermStats     = Perm(1 << 9) // 数据

	PermAll = Perm(1 << 10) // 全部

)
