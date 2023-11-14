package model

import (
	"time"

	"github.com/google/uuid"
)

// Definición de la estructura Order (Orden)
type Order struct {
	OrderID     uint64     `json:"order_id"`     // Identificador único de la orden (etiqueta para serialización JSON)
	CustomerID  uuid.UUID  `json:"customer_id"`  // Identificador único del cliente (etiqueta para serialización JSON)
	LineItem    []LineItem `json:"line_items"`   // Lista de productos en la orden (etiqueta para serialización JSON)
	CreatedAt   *time.Time `json:"created_at"`   // Fecha de creación de la orden (etiqueta para serialización JSON)
	ShippedAt   *time.Time `json:"shipped_at"`   // Fecha de envío de la orden (etiqueta para serialización JSON)
	CompletedAt *time.Time `json:"completed_at"` // Fecha de completado de la orden (etiqueta para serialización JSON)
}

// Definición de la estructura LineItem (Ítem de Línea)
type LineItem struct {
	ItemID   uuid.UUID `json:"item_id"`  // Identificador único del ítem (etiqueta para serialización JSON)
	Quantity uint      `json:"quantity"` // Cantidad del ítem en la orden (etiqueta para serialización JSON)
	Price    uint      `json:"price"`    // Precio del ítem (etiqueta para serialización JSON)
}
