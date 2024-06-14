package processamento

import (
	"AP1/handlers/pedidos"
	"fmt"
	"time"
)

func Processar(lojaPonteiro *bool, intervalo int) {
	for {
		time.Sleep(time.Duration(intervalo) * time.Second)
		if *lojaPonteiro {
			if pedidos.ExpedirPedido() {
				logComTempo("Pedido expedido")
			}
		} else {
			return
		}
	}
}

func logComTempo(msg string) {
	tempo := time.Now()
	tempoF := tempo.Format("02/01/2006 15:04:05")
	fmt.Printf("%s %s\n", tempoF, msg)
}
