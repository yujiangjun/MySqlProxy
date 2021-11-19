package const1

var (
	OK  = NewResponse(200, "ok") // 通用成功
	Err = NewResponse(500, "")   // 通用错误

	// ErrParam 服务级错误码
	ErrParam     = NewResponse(10001, "参数有误")
	ErrSignParam = NewResponse(10002, "签名参数有误")

	// ErrUserService 模块级错误码 - 用户模块
	ErrUserService = NewResponse(20100, "用户服务异常")
	ErrUserPhone   = NewResponse(20101, "用户手机号不合法")
	ErrUserCaptcha = NewResponse(20102, "用户验证码有误")

	// ErrOrderService 库存模块
	ErrOrderService = NewResponse(20200, "订单服务异常")
	ErrOrderOutTime = NewResponse(20201, "订单超时")
)