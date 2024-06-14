package metricas

import (
	"AP1/modelos/metrica"
	"encoding/json"
	"net/http"
	"time"
)

var metricas metrica.MetricasSistema
var tempoComeco time.Time

func GetMetricas(w http.ResponseWriter, r *http.Request) {
	pedidosTotal := float32(metricas.PedidosEmAndamento) + float32(metricas.PedidosEncerrados)
	if pedidosTotal != 0 {
		metricas.TicketMedio = metricas.FaturamentoTotal / pedidosTotal
	} else {
		metricas.TicketMedio = 0
	}
	metricas.TempoTotal = int(time.Since(tempoComeco).Seconds())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metricas)
}

func IncProdutosCadastrados(i int)  { metricas.ProdutosCadastrados += i }
func IncPedidosEmAndamento(i int)   { metricas.PedidosEmAndamento += i }
func SomFaturamentoTotal(i float32) { metricas.FaturamentoTotal += i }

func SomPedidosEncerrados() {
	metricas.PedidosEncerrados++
	metricas.PedidosEmAndamento--
}

func SetTempoComeco() { tempoComeco = time.Now() }
