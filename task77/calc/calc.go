package calc

import "fmt"

type Calculator struct {
	a float64
	b float64
}

func NewCalculator(a, b float64) Calculator {
	return Calculator{
		a: a,
		b: b,
	}
}

func (c *Calculator) Calculate(operation string) (float64, error) {
	switch operation {
	case "+":
		return c.add()
	case "-":
		return c.sub()
	case "*":
		return c.mult()
	case "/":
		return c.div()
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}

func (c *Calculator) add() (float64, error) {
	return c.a + c.b, nil
}

func (c *Calculator) sub() (float64, error) {
	return c.a - c.b, nil
}

func (c *Calculator) mult() (float64, error) {
	return c.a * c.b, nil
}

func (c *Calculator) div() (float64, error) {
	if c.b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return c.a / c.b, nil
}
