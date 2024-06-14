package produtos

import (
	"AP1/handlers/metricas"
	"AP1/modelos/produto"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

var idProduto int = 1

func Iniciar() { produto.Iniciar() }

func CriarProduto(w http.ResponseWriter, r *http.Request) {
	var tmp produto.Produto
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err != nil {
		http.Error(w, "Erro em decodificar json", http.StatusBadRequest)
		return
	}

	//Verificar se o json tem todas as informações
	if tmp.Descricao == "" || tmp.Nome == "" || tmp.Valor == 0 {
		http.Error(w, "Json sem informações necessárias", http.StatusBadRequest)
		return
	}

	resposta := fmt.Sprintf("Produto criado com sucesso (ID = %d)", idProduto)
	if tmp.ID != 0 {
		resposta += "\n[Aviso: ID fornecido no json desconsiderado]"
	}

	//arredonda para 2 casas decimais
	tmp.Valor = float32(math.Round(float64((tmp.Valor * 100))) / 100)

	arvore := produto.ArvoreProdutos
	tmp.ID = idProduto
	err = arvore.AdicionarProduto(&tmp)

	if err != nil {
		http.Error(w, "ERRO: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, resposta)

	idProduto++
	metricas.IncProdutosCadastrados(1)
}

// Fornece os dados de um produto com base no nome
func ListarProduto(w http.ResponseWriter, r *http.Request) {
	nomeProd := r.FormValue("nome")
	if nomeProd == "" {
		http.Error(w, "Nome do produto não fornecido", http.StatusBadRequest)
		return
	}

	arvore := produto.ArvoreProdutos
	produtoNode := arvore.AcharProduto(nomeProd)
	if produtoNode == nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtoNode.Data)
}

func RemoverProduto(w http.ResponseWriter, r *http.Request) {
	//Obter nome pelo form
	nomeProd := r.FormValue("nome")

	//Caso o nome seja valido, remover o produto
	arvore := produto.ArvoreProdutos
	err := arvore.RemoverProduto(nomeProd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	metricas.IncProdutosCadastrados(-1)

	//Mandar respota
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Produto removido com sucesso"))
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {
	arvore := produto.ArvoreProdutos
	if arvore.ArvoreVazia() {
		http.Error(w, "Não ha produtos cadastrados", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arvore.GetArvore())
}
