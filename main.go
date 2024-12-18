package main

import (
	"fmt"
	"math/rand"
)

func main() {
	
	fmt.Println("*** Juego Adivinar Número Aleatorio ***")
	jugar()

}

func jugar(){
	numAleatorio := rand.Intn(100)
	var numIngresado int 
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingrese un número (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Felicitaciones, adivinaste el número!")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El número a adivinar es mayor")
		} else if numIngresado > numAleatorio {
			fmt.Println("El número a adivinar es menor")
		}
	}
	fmt.Println("Se acabaron los intentos, el número era: ", numAleatorio)
	jugarNuevamente()
}	

func jugarNuevamente(){
	var eleccion string
	fmt.Print("Quieres jugar nuevamente (s/n)?: ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar ... hasta pronto")
	default:
		fmt.Println("Elección inválida \n Inténtalo nuevamente")
		jugarNuevamente()
	}
}