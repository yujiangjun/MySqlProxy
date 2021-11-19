package main

import "mySqlProxy/config"

var (
	OK  = config.NewResponse(200, "ok") // 通用成功
	Err = config.NewResponse(500, "")   // 通用错误

	// ErrParam 服务级错误码
	ErrParam     = config.NewResponse(10001, "参数有误")
	ErrSignParam = config.NewResponse(10002, "签名参数有误")

	// ErrUserService 模块级错误码 - 用户模块
	ErrUserService = config.NewResponse(20100, "用户服务异常")
	ErrUserPhone   = config.NewResponse(20101, "用户手机号不合法")
	ErrUserCaptcha = config.NewResponse(20102, "用户验证码有误")

	// ErrOrderService 库存模块
	ErrOrderService = config.NewResponse(20200, "订单服务异常")
	ErrOrderOutTime = config.NewResponse(20201, "订单超时")
)