package produto

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type ArvoreProduto struct { // TODO
	raiz    *Node
	tamanho int //diminuir
}

type Node struct {
	Data Produto
	Esq  *Node
	Dir  *Node
}

const (
	menor    = -1
	maior    = 1
	igual    = 0
	esquerda = 5
	direita  = 6
)

var ArvoreProdutos *ArvoreProduto

func (arvore *ArvoreProduto) ArvoreVazia() bool { return arvore.raiz == nil }

func Iniciar() {
	ArvoreProdutos = new(ArvoreProduto)
	ArvoreProdutos.raiz = nil
	ArvoreProdutos.tamanho = 0
}

// verifica se o input tem caracteres especiais
func validarNome(nome string) bool {
	for _, char := range nome {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

// retorna um slice ordenado da arvore
var copiaPtrTmp *[]Produto //
func (arvore *ArvoreProduto) GetArvore() []Produto {

	if arvore.ArvoreVazia() {
		return nil
	}

	res := make([]Produto, 0, arvore.tamanho)
	copiaPtrTmp = &res

	getArvoreProdutosCont(arvore.raiz.Esq)
	res = append(res, arvore.raiz.Data)
	getArvoreProdutosCont(arvore.raiz.Dir)

	copiaPtrTmp = nil
	return res
}

func getArvoreProdutosCont(node *Node) {
	if node == nil {
		return
	}
	getArvoreProdutosCont(node.Esq)
	*copiaPtrTmp = append(*copiaPtrTmp, node.Data)
	getArvoreProdutosCont(node.Dir)
}

func (arvore *ArvoreProduto) AcharProduto(val string) *Node {
	val = strings.ToLower(val)
	atual := arvore.raiz
	for atual != nil {
		res := strings.Compare(strings.ToLower(atual.Data.Nome), val)
		if res == igual {
			return atual
		}
		if res == maior {
			atual = atual.Esq
		} else {
			atual = atual.Dir
		}
	}
	return nil
}

// Acha o pai com base no nome
func (arvore *ArvoreProduto) acharPai(val string) (*Node, int, bool) {
	val = strings.ToLower(val)
	atual := arvore.raiz
	var anterior *Node
	anterior = nil
	existe := false
	dir := 0

	for atual != nil {
		res := strings.Compare(strings.ToLower(atual.Data.Nome), val)
		if res == igual {
			existe = true
			break
		}
		anterior = atual
		if res == menor {
			atual = atual.Dir
			dir = direita
		} else {
			atual = atual.Esq
			dir = esquerda
		}
	}

	return anterior, dir, existe
}

// Adiciona um produto
func (arvore *ArvoreProduto) AdicionarProduto(adicionar *Produto) error {

	if !validarNome(adicionar.Nome) {
		return errors.New("nome inválido pois possui caracteres especiais")
	}

	if arvore.ArvoreVazia() {
		arvore.raiz = &Node{*adicionar, nil, nil}
		arvore.tamanho++
		return nil
	}

	pai, dir, existe := arvore.acharPai(adicionar.Nome)

	if existe {
		msg := fmt.Sprintf("produto de nome '%s' ja existe", adicionar.Nome)
		return errors.New(msg)
	}

	if dir == esquerda {
		pai.Esq = &Node{*adicionar, nil, nil}
	} else {
		pai.Dir = &Node{*adicionar, nil, nil}
	}

	arvore.tamanho++
	return nil
}

// Remove o produto pelo nome
func (arvore *ArvoreProduto) RemoverProduto(nome string) error {
	if arvore.ArvoreVazia() {
		return errors.New("árvore está vazia")
	}

	pai, dir, existe := arvore.acharPai(nome)
	if !existe {
		msg := fmt.Sprintf("produto de nome '%s' não encontrado", nome)
		return errors.New(msg)
	}

	prod := pai.Esq
	if dir == direita {
		prod = pai.Dir
	}

	if prod == nil {
		return errors.New("produto não encontrado")
	}

	numFilhos := 0

	if prod.Dir != nil {
		numFilhos++
	}
	if prod.Esq != nil {
		numFilhos++
	}

	if numFilhos == 0 {
		if dir == direita {
			pai.Dir = nil
		} else {
			pai.Esq = nil
		}
		return nil
	}

	if numFilhos == 1 {
		neto := prod.Dir
		if neto == nil {
			neto = prod.Esq
		}

		if dir == direita {
			pai.Dir = neto
		} else {
			pai.Esq = neto
		}
		return nil
	}

	tmp := prod
	var tmpPai *Node = nil
	if dir == direita {
		for tmp.Esq != nil {
			tmpPai = tmp
			tmp = tmp.Esq
		}
		pai.Dir = tmp
		tmpPai.Esq = nil
	} else {
		for tmp.Dir != nil {
			tmpPai = tmp
			tmp = tmp.Dir
		}
		pai.Esq = tmp
		tmpPai.Dir = nil
	}

	return nil
}

func (arvore *ArvoreProduto) RemoverProduto2(nome string) error {
	if arvore.ArvoreVazia() {
		return errors.New("árvore está vazia")
	}

	if arvore.raiz.Data.Nome == nome {
		//TODO
	}

	pai, dir, existe := arvore.acharPai(nome)

	if !existe {
		return errors.New("Produto não existe")
	}

	prod := pai.Esq
	var prodPtr **Node
	prodPtr = &pai.Esq
	if dir == direita {
		prod = pai.Dir
		prodPtr = &pai.Dir
	}

	numFilhos := 0

	if prod.Esq != nil {
		numFilhos++
	}
	if prod.Dir != nil {
		numFilhos++
	}

	if numFilhos == 0 {
		*prodPtr = nil
		prod = nil
		arvore.tamanho--
		return nil
	}

	if numFilhos == 1 {
		tmp := prod.Esq
		if tmp == nil {
			tmp = prod.Dir
		}
		*prodPtr = tmp
		prod = nil
		arvore.tamanho--
		return nil
	}

	if dir == esquerda {
		tmp := prod.Dir
		//tmpPai := prod
		for tmp.Dir != nil {
			//tmpPai = tmp
			tmp = tmp.Dir
		}

	}

	arvore.tamanho--
	return nil
}
