package calc

import (
	"log"
)

type calculator struct {
}

func NewCalculator() calculator {
	return calculator{}
}
func (c *calculator) Calculator(a, b float64, optr string) float64 {
	switch optr {
	case "+":
		return c.add(a, b)
	case "-":
		return c.sub(a, b)
	case "*":
		return c.mult(a, b)
	case "/":
		return c.div(a, b)
	default:
		log.Println("Такой операции не существует или не предусмотренно разработчиком")
		log.Println("Such an operation does not exist or is not provided by the developer")
		return 0
	}
}
func (c *calculator) add(a, b float64) float64 {
	return a + b
}
func (c *calculator) sub(a, b float64) float64 {
	return a - b
}
func (c *calculator) mult(a, b float64) float64 {
	return a * b
}
func (c *calculator) div(a, b float64) float64 {
	if b == 0 {
		log.Println("ошибка, деление на 0")
		log.Println("error, division by 0")
		return 0
	} else {
		return a / b
	}
}
