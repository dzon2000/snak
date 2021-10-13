package net

import "math"

type Func interface {
	Calc(x float32) float32
	Derivative(x float32) float32
}

type Sigmoid struct {
}

func (s Sigmoid) Calc(x float32) float32 {
	return 1 / (1 + float32(math.Exp(float64(-x))))
}

func (s Sigmoid) Derivative(x float32) float32 {
	return x * (1 - x)
}

type Identity struct {
}

func (i Identity) Calc(x float32) float32 {
	return x
}

func (i Identity) Derivative(x float32) float32 {
	return 0.
}
