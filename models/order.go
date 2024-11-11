package models

type Order struct {
	Orderid         uint        `json:"order_id" gorm:"primaryKey;column:orderid"`
	Status          string      `json:"status"`
	Items           []OrderItem `json:"items" gorm:"foreignKey:orderid;references:orderid"`
	Userid          int         `json:"userId"`
	Totalamount     float32     `json:"totalamount"`
	Shippingaddress string      `json:"shippingAddress"`
}

type OrderItem struct {
	Orderitemid uint    `json:"orderitemid,omitempty" gorm:"primaryKey;column:orderitemid"`
	Orderid     uint    `json:"order_id"`
	Productid   uint    `json:"productID"`
	Variantid   uint    `json:"variant_id"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (Order) TableName() string {
	return "orders"
}
func (OrderItem) TableName() string {
	return "orderitems"
}

type OrderRequest struct {
	UserId          int         `json:"userId"`
	Items           []OrderItem `json:"items"`
	ShippingAddress string      `json:"shippingAddress"`
}
