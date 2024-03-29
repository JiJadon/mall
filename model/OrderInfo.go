package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderInfo struct { //订单商品表
	gorm.Model
	OrderNo      string        `json:"orderNo"form:"order_no" gorm:"unique"` //订单编号
	MemberId     uint          `json:"memberId"form:"member_id"`             //用户id
	AddressId    uint          `json:"addressId"form:"address_id"`           //地址ID
	PayMethod    uint          `json:"payMethod"form:"pay_method"`           //付款方式
	TotalCount   uint          `json:"totalCount"form:"total_count"`         //商品数量
	TotalPrice   float64       `json:"totalPrice"form:"total_price"`         //商品总价
	TransitPrice float64       `json:"transitPrice"form:"transit_price"`     //运费
	OrderStatus  uint          `json:"orderStatus"form:"order_status"`       //订单状态
	TradeNo      string        `json:"tradeNo"form:"trade_no" gorm:"unique"` //支付编号
	CommentTime  time.Time     `json:"commentTime"form:"comment_time"`       //评论时间
	OrderGoods   []*OrderGoods `json:"orderGoods"gorm:"-"`
}
