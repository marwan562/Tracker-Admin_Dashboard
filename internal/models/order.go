package models

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

var (
	OrderStatuses = []string{"Order placed", "Preparing", "Baking", "Ready for delivery", "Delivered", "Cancelled"}

	PizzaTypes = []string{
		"pepperoni",
		"margherita",
		"supreme",
		"veggie",
	}

	PizzaSizes = []string{
		"small",
		"medium",
		"large",
		"extra large",
	}
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct {
	ID              string `gorm:"primaryKey;size:14;autoIncrement" json:"id"`
	CustomerName    string `gorm:"size:100;not null" json:"customer_name"`
	CustomerEmail   string `gorm:"size:100;not null" json:"customer_email"`
	CustomerPhone   string `gorm:"size:15;not null" json:"customer_phone"`
	CustomerAddress string `gorm:"size:255;not null" json:"customer_address"`
	CustomerCity    string `gorm:"size:100;not null" json:"customer_city"`
	CustomerCountry string `gorm:"size:100;not null" json:"customer_country"`
	Status          string `gorm:"size:100;not null" json:"status"`
	Items           []OrderItem
	CreatedAt       time.Time `gorm:"size:100;not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"size:100;not null" json:"updated_at"`
}

type OrderItem struct {
	ID           string    `gorm:"primaryKey;size:14;autoIncrement" json:"id"`
	OrderID      int       `gorm:"size:14;not null" json:"order_id"`
	Type         string    `gorm:"size:100;not null" json:"type"`
	Size         string    `gorm:"size:100;not null" json:"size"`
	Quantity     int       `gorm:"size:100;not null" json:"quantity"`
	Price        float64   `gorm:"size:100;not null" json:"price"`
	Instructions string    `gorm:"size:255" json:"instructions"`
	CreatedAt    time.Time `gorm:"size:100;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"size:100;not null" json:"updated_at"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = shortid.MustGenerate()
	}
	return nil
}

func (o *OrderModel) CreateOrder(order *Order) error {
	return o.DB.Create(order).Error
}

func (o *OrderModel) GetOrder(id string) (*Order, error) {
	var order Order
	err := o.DB.Preload("Items").First(&order, "id = ?", id).Error
	return &order, err
}
