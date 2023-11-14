package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/FrancoRutiliano/orders-api/model"
	"github.com/redis/go-redis/v9"
)

// RedisRepo representa un repositorio que interactúa con Redis.
type RedisRepo struct {
	Client *redis.Client // Cliente de Redis
}

// orderIDKey devuelve la clave para la orden en Redis basada en su ID.
func orderIDKey(id uint64) string {
	return fmt.Sprintf("Order:%d", id)
}

// Insert inserta una orden en el repositorio de Redis.
func (r *RedisRepo) Insert(ctx context.Context, order model.Order) error {
	// Convierte la orden a formato JSON.
	data, err := json.Marshal(order)

	if err != nil {
		return fmt.Errorf("error al codificar la orden: %w", err)
	}

	// Genera la clave de Redis para la orden.
	key := orderIDKey(order.OrderID)

	// Establece los datos en Redis con SetNX (solo si la clave no existe).
	res := r.Client.SetNX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		return fmt.Errorf("error al establecer: %w", err)
	}

	return nil
}

// ErrorNotExist es un error que indica que la orden no existe.
var ErrorNotExist = errors.New("la orden no existe")

// FindById busca una orden en el repositorio de Redis por su ID.
func (r *RedisRepo) FindById(ctx context.Context, id uint64) (model.Order, error) {
	// Genera la clave de Redis para la orden.
	key := orderIDKey(id)

	// Obtiene el valor asociado a la clave en Redis.
	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// Si no se encuentra la clave en Redis, se devuelve el error ErrorNotExist.
		return model.Order{}, ErrorNotExist
	} else if err != nil {
		// Si ocurre un error diferente al no encontrar la clave, se devuelve un error explicativo.
		return model.Order{}, fmt.Errorf("obtener orden: %w", err)
	}

	// Declara una variable para almacenar la orden decodificada.
	var order model.Order

	// Decodifica el valor JSON recuperado en la estructura de orden.
	err = json.Unmarshal([]byte(value), &order)
	if err != nil {
		// Si hay un error al decodificar el JSON, se devuelve un error explicativo.
		return model.Order{}, fmt.Errorf("error al decodificar el JSON de la orden: %w", err)
	}

	// Devuelve la orden encontrada sin errores.
	return order, nil
}
