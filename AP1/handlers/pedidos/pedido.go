package pedidos

import (
	"AP1/handlers/metricas"
	"AP1/modelos/pedido"
	"AP1/modelos/produto"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

var idPedido int = 1

func Iniciar() {
	pedido.Iniciar()
}

func IncluirPedido(w http.ResponseWriter, r *http.Request) {
	var tmp pedido.Pedido
	//Decodificação do json
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Criar mensagem de resposta
	resposta := fmt.Sprintf("Pedido criado com sucesso (ID = %d)", idPedido)
	if tmp.ID != 0 {
		resposta += "\n[Aviso: ID fornecido no json desconsiderado]"
	}
	if tmp.ValorTotal != 0 {
		resposta += "\n[Aviso: Valor fornecido no json desconsiderado]"
	}

	tmp.ValorTotal = 0

	//Verifica se todos os produtos do pedido são validos
	arvore := produto.ArvoreProdutos
	for _, i := range tmp.Produtos {
		prodNode := arvore.AcharProduto(i.ProdutoNome) //TODO checar campo quantidade
		if prodNode == nil {
			msg := fmt.Sprintf("produto com nome '%s' não existe", i.ProdutoNome)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		//Incrementa o valor total
		prod := prodNode.Data
		tmp.ValorTotal += prod.Valor * float32(i.Quantidade)
	}

	tmp.ID = idPedido

	//Inclue a taxa de entrega
	if tmp.Delivery {
		tmp.ValorTotal += 10
	}

	//Inclui o desconto de 10% caso o valor exceda 100
	if tmp.ValorTotal > 100 {
		tmp.ValorTotal *= 0.9
	}

	//arredonda para 2 casas decimais
	tmp.ValorTotal = float32(math.Round(float64((tmp.ValorTotal * 100))) / 100)

	//Adiciona o pedido a fila
	fila := pedido.FilaP
	fila.Incluir(&tmp)

	//Resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, resposta)

	idPedido++ //Incrementar o contador de id para o próximo pedido

	metricas.IncPedidosEmAndamento(1)
	metricas.SomFaturamentoTotal(tmp.ValorTotal)
}

func ExpedirPedido() bool {
	fila := pedido.FilaP
	if fila.Expedir() {
		metricas.SomPedidosEncerrados()
		return true
	}
	return false
}

func ListarPedidos(w http.ResponseWriter, r *http.Request) {
	fila := pedido.FilaP
	if fila.FilaVazia() {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, "Não ha pedidos em andamento")
		return
	}

	algoritimo := r.FormValue("tipo")

	if algoritimo == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fila.OrdenarID())
		return
	}

	algoritimoIndex := pedido.VerificarAlgoritimo(algoritimo)
	if algoritimoIndex == -1 {
		http.Error(w, "Algoritimo invalido", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fila.OrdenarPreco(algoritimoIndex))
}
