package testgo

import (
	"fmt"
)

type data struct {
	val int
}

func (p_data *data) set(num int) {
	p_data.val = num
}

func (p_data *data) get() int {
	return p_data.val
}

func (p_data *data) show() {
	fmt.Printf("val=%d\n", p_data.val)
}

func main() {
	pData := &data{4}
	pData.show()
}
