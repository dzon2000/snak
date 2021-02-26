package net

import (
    "math/rand"
    "fmt"
)

var f = Sigmoid{}

type Net struct {
   Layers []Layer
}

func (n Net) Calc(input []float32) []float32 {
    var dd = input
    for _, l := range n.Layers {
        dd = l.Calc(dd)
    }
    return dd
}

func (n Net) String() string {
    return fmt.Sprintf("%v", n.Layers)
}

func NewNet(lenghts []int) Net {
    layers := make([]Layer, len(lenghts) - 1)
    for i := 1; i < len(lenghts); i++ {
        layers[i - 1] = newLayer(lenghts[i], lenghts[i - 1])
    }
    return Net{layers}
}

func newLayer(lenght, weights int) Layer {
   neurons := make([]Neuron, lenght)
   for i := 0; i < lenght; i++ {
       neurons[i] = newNeuron(weights)
   }
   return Layer{neurons}
}

func newNeuron(lenght int) Neuron {
    weights := make([]float32, lenght)
    for i := 0; i < lenght; i++ {
        weights[i] = rand.Float32()
    }
    return Neuron{
        weights,
        f,
    }
}
