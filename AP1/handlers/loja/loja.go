package loja

import (
	"AP1/processamento"
	"fmt"
	"net/http"
	"strconv"
)

var lojaAberta *bool

func AbrirLoja(w http.ResponseWriter, r *http.Request) {
	intervaloStr := r.FormValue("intervalo")
	if intervaloStr == "" {
		msg := "Parametro 'intervalo' faltando"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	intervalo, err := strconv.Atoi(intervaloStr)
	if err != nil {
		msg := "Erro ao converter 'intervalo'"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if lojaAberta != nil && *lojaAberta == true {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Loja ja foi aberta anteriormente")
		return
	}

	lojaAberta = new(bool)
	*lojaAberta = true
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Loja aberta com sucesso")

	go processamento.Processar(lojaAberta, intervalo)
}

func FecharLoja(w http.ResponseWriter, r *http.Request) {
	if lojaAberta == nil {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Loja nao foi aberta")
	} else if *lojaAberta == false {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Loja ja fechada")
	} else {
		*lojaAberta = false
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Loja fechada com sucesso")
	}
}
