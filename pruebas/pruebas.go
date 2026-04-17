package pruebas

import "fmt"

func ProcesarSensores(mapa  map[string]int) {

	for _, temp := range mapa {

		if temp>100{
			fmt.Println("Sobrecalentamiento")
		}else if temp<=0{
			fmt.Println("Temperatura baja")

		}else{
			fmt.Printf("Temperatura estable: [%d]\n",temp)
		}
	}

}	

func ProcesarTemperaturas(temperaturas []int) {

	for _, temp := range temperaturas {

		if temp>100{
			fmt.Println("Sobrecalentamiento")
		}else if temp<=0{
			fmt.Println("Temperatura baja")

		}else{
			fmt.Printf("Temperatura estable: [%d]\n",temp)
		}
	}

}	