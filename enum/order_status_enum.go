package enum
//支付状态
type OrderStatus int

const (
	Payed OrderStatus = 0
	UnPay OrderStatus = 1
	TimeOut OrderStatus = 2
)

func (p OrderStatus) String() string {
	switch p {
	case Payed:
		return "已付款"
	case UnPay:
		return "未付款"
	case TimeOut:
		return "超时"
	default:
		return "UNKNOWN"
	}
}