package libs

import (
	"fmt"
	"strconv"
)

func getBit(bt byte, index int) (i int) {
	bt = (bt << index) >> 7
	, _ = strconv.Atoi(fmt.Sprintf("%b", bt))
	return i
}
