// 只修改msgMap，不要添加其他代码

package outputmsg

import (
	"github.com/houyanzu/work-box/lib/output"
)

func init() {
	output.InitMsgMap(msgMap)
}

var msgMap = map[output.ErrorCode]map[string]string{
	0: {
		"zh": "ok",
		"en": "ok",
		"tw": "ok",
	},
	1: {
		"zh": "请重试",
		"en": "Please try again",
		"tw": "請重試",
	},
	3: {
		"zh": "登陆过期了，需要重新登录哟",
		"en": "Login Has Expired. Please Login Again",
		"tw": "登錄過期了，需要重新登錄喲",
	},
	6: {
		"zh": "错误",
		"en": "Error",
		"tw": "Error",
	},
}
