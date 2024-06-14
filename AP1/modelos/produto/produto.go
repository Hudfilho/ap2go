package produto

type Produto struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Valor     float32 `json:"valor"`
}

type ProdutoQuantidade struct {
	ProdutoNome string `json:"produto"`
	Quantidade  int32  `json:"quantidade"`
}
