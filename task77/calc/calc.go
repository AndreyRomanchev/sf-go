package calc

import "fmt"

type calculator struct {
	a float64
	b float64
}

func NewCalculator(a, b float64) calculator {
	return calculator{
		a: a,
		b: b,
	}
}

func (c *calculator) Calculate(operation string) (float64, error) {
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

func (c *calculator) add() (float64, error) {
	return c.a + c.b, nil
}

func (c *calculator) sub() (float64, error) {
	return c.a - c.b, nil
}

func (c *calculator) mult() (float64, error) {
	return c.a * c.b, nil
}

func (c *calculator) div() (float64, error) {
	if c.b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return c.a / c.b, nil
}
