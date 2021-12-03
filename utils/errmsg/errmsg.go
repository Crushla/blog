package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// 用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// 文章模块错误
	ERROR_ARTICLE_NOT_EXIST = 2001
	// 分类文章错误
	ERROR_CATENAME_USED      = 3001
	ERROR_CATENAME_NOT_EXIST = 3002
)

var Codemsg = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "用户名已经存在",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIST:     "用户不存在",
	ERROR_TOKEN_NOT_EXIST:    "token不存在",
	ERROR_TOKEN_RUNTIME:      "token已过期",
	ERROR_TOKEN_WRONG:        "token不正确",
	ERROR_TOKEN_TYPE_WRONG:   "token格式错误",
	ERROR_USER_NO_RIGHT:      "没有管理权限",
	ERROR_ARTICLE_NOT_EXIST:  "文章不存在",
	ERROR_CATENAME_USED:      "该分类已存在",
	ERROR_CATENAME_NOT_EXIST: "该分类不存在",
}

func GetErrMsg(code int) string {
	return Codemsg[code]
}
