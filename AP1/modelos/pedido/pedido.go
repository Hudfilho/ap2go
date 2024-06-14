package pedido

import (
	"AP1/modelos/produto"
)

type Pedido struct {
	ID         int                         `json:"id"`
	Delivery   bool                        `json:"delivery"`
	Produtos   []produto.ProdutoQuantidade `json:"produtos"`
	ValorTotal float32                     `json:"valorTotal"`
}
