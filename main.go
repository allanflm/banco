package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuiza := contas.ContaPoupanca{}
	contaDaLuiza.Depositar(500)
	PagarBoleto(&contaDaLuiza, 350)
	fmt.Println(contaDaLuiza.ObterSaldo())

}
