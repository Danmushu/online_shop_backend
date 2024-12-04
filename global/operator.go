package global

import (
	"project/models"
	"project/utils/perms"
)

// todo

var OperatorList = []*models.Operator{
	{
		ID:   1,
		Usr:  "",
		Pwd:  "",
		Key:  "",
		Perm: perms.PermAll,
	},
	{
		ID:   2,
		Usr:  "",
		Pwd:  "",
		Key:  "",
		Perm: perms.PermAll,
	},
}
var OperatorEmailList = []string{}
