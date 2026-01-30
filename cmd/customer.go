package main

import (
	"log/slog"
	"net/http"
	"pizza-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderFormData struct {
	PizzaTypes []string `form:"pizza_types"`
	PizzaSizes []string `form:"pizza_sizes"`
}

type OrderRequest struct {
	Name         string   `form:"name" binding:"required, min=2, max=100"`
	Phone        string   `form:"phone" binding:"required, min=10, max=15"`
	Address      string   `form:"address" binding:"required, min=2, max=100"`
	Sizes        []string `form:"sizes" binding:"required, min=1, dive,valid_pizza_size"`
	PizzaTypes   []string `form:"pizza_types" binding:"required, min=1, dive,valid_pizza_type"`
	Instructions []string `form:"instructions" binding:"required, min=2, max=100"`
}

func (h *Handler) ServeNewOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", OrderFormData{
		PizzaTypes: models.PizzaTypes,
		PizzaSizes: models.PizzaSizes,
	})
}

func (h *Handler) HandleNewOrderPost(c *gin.Context) {
	var form OrderRequest

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems := make([]models.OrderItem, len(form.Sizes))
	for i := range form.Sizes {
		orderItems[i] = models.OrderItem{
			Size:         form.Sizes[i],
			Type:         form.PizzaTypes[i],
			Instructions: form.Instructions[i],
		}
	}

	order := &models.Order{
		CustomerName:    form.Name,
		CustomerPhone:   form.Phone,
		CustomerAddress: form.Address,
		Status:          "Order placed",
		Items:           orderItems,
	}

	if err := h.orders.CreateOrder(order); err != nil {
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Order created successfully", "order", order)
	c.String(http.StatusOK, "Order created successfully")
	c.Redirect(http.StatusSeeOther, "/customer/"+order.ID)
}

func (h *Handler) ServeCustomer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "Order ID is required")
		return
	}
	order, err := h.orders.GetOrder(id)
	if err != nil {
		slog.Error("Failed to get order", "error", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "customer.tmpl", gin.H{
		"Order": order,
	})
}