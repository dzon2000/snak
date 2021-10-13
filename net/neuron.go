package net

import "fmt"

type Neuron struct {
	Weights []float32
	Bias    float32
	Fn      Func
}

func (n Neuron) Calc(input []float32) float32 {
	res := n.Bias
	for i := range n.Weights {
		res += n.Weights[i] * input[i]
	}
	return n.Fn.Calc(res)
}

func (n Neuron) String() string {
	return fmt.Sprintf("%v", n.Weights)
}
