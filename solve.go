package main

import (
	"fmt"
	"snak/net"
)

func main() {
	//var neurons1 = []net.Neuron{net.Neuron{22}, net.Neuron{11}}
	//var neurons2 = []net.Neuron{net.Neuron{14}, net.Neuron{88}}
	//var layers = []net.Layer{net.Layer{neurons1}, net.Layer{neurons2}}
	//net := net.Net{layers}
	//fmt.Println(net.Calc())
	net := net.NewNet([]int{2, 5, 1})
	fmt.Println(net.Calc([]float32{0., 1.}))
	fmt.Println(net)
}
