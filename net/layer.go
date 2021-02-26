package net

import "fmt"

type Layer struct {
    Neurons []Neuron
}

func (l Layer) Calc(input []float32) []float32 {
    res := make([]float32, len(l.Neurons))
    for i, n := range l.Neurons {
        res[i] = n.Calc(input)
    }
    return res
}

func (l Layer) String() string {
    return fmt.Sprintf("%v", l.Neurons)
}
