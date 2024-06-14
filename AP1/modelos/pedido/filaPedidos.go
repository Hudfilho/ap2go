package pedido

type FilaPedidos struct {
	arr []Pedido
}

var FilaP *FilaPedidos

func Iniciar() {
	FilaP = new(FilaPedidos)
	FilaP.arr = make([]Pedido, 0)
}

func (fila *FilaPedidos) FilaVazia() bool {
	return len(fila.arr) == 0
}

func (fila *FilaPedidos) OrdenarID() []Pedido {
	return fila.arr
}

func (fila *FilaPedidos) Incluir(tmp *Pedido) {
	fila.arr = append(fila.arr, *tmp)
}

func (fila *FilaPedidos) Expedir() bool {
	if len(fila.arr) > 0 {
		fila.arr = fila.arr[1:]
		return true
	}
	return false
}

func VerificarAlgoritimo(tipo string) int {
	switch tipo {
	case "bubblesort":
		return 1
	case "quicksort":
		return 2
	case "selectionsort":
		return 3
	default:
		return -1
	}
}

func (fila *FilaPedidos) OrdenarPreco(tipo int) []Pedido {
	switch tipo {
	case 1:
		return fila.bubbleSort()
	case 2:
		return fila.quickSort()
	case 3:
		return fila.selectionSort()
	}
	return nil
}

func (fila *FilaPedidos) bubbleSort() []Pedido {
	n := len(fila.arr)
	copia := make([]Pedido, n)
	copy(copia, fila.arr)

	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if copia[j].ValorTotal > copia[j+1].ValorTotal {
				copia[j], copia[j+1] = copia[j+1], copia[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}

	return copia
}

func (fila *FilaPedidos) quickSort() []Pedido {
	n := len(fila.arr)
	copia := make([]Pedido, n)
	copy(copia, fila.arr)

	quickSortHelper(copia, 0, n-1)

	return copia
}

func quickSortHelper(arr []Pedido, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSortHelper(arr, low, p-1)
		quickSortHelper(arr, p+1, high)
	}
}

func partition(arr []Pedido, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j].ValorTotal < pivot.ValorTotal {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func (fila *FilaPedidos) selectionSort() []Pedido {
	n := len(fila.arr)
	copia := make([]Pedido, n)
	copy(copia, fila.arr)

	for i := 0; i < n-1; i++ {
		minI := i
		for j := i + 1; j < n; j++ {
			if copia[j].ValorTotal < copia[minI].ValorTotal {
				minI = j
			}
		}
		copia[i], copia[minI] = copia[minI], copia[i]
	}

	return copia
}
