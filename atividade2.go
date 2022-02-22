package main

import "fmt"

func main() {
	fmt.Println("************ * Seja bem vindo ao sistema de reserva de Ticket * ************")

	var Ticket float32 = 10.50
	var name string
	var quant int
	var preco float32
	var confirma string = "s"
	var vetorNome []string
	var vetorQuant []int
	var vetorPreco []float32

	for confirma != "n" {
		fmt.Println("Qual o seu nome? ")
		fmt.Scanln(&name)
		vetorNome = append(vetorNome, name)

		fmt.Println("Quantos bilhetes você quer comprar? ")
		fmt.Scanln(&quant)
		vetorQuant = append(vetorQuant, quant)

		preco = float32(quant) * Ticket

		fmt.Println("Você confirma a compra de", quant, "bilheres no valor de: R$", preco, "? Digite (s/n)")
		fmt.Scanln(&confirma)

		if confirma == "s" {

			fmt.Println("Sua Compra dará o total de R$", preco)
			vetorPreco = append(vetorPreco, preco)

			fmt.Println("Deseja comprar mais? s/n")
			fmt.Scanln(&confirma)
			if confirma == "s" {
				continue
			} else {
				fmt.Println("nome:", name, "Quantidade de Ticket:", quant, "Preço:", preco)
			}
		}
	}
}
