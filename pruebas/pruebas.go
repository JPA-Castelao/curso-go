package pruebas

import "fmt"

func ProcesarSensores(lecturas []int) {

	for _, temp := range lecturas {

		if temp>100{
			fmt.Println("Sobrecalentamiento")
		}else if temp<=0{
			fmt.Println("Temperatura baja")

		}else{
			fmt.Printf("Temperatura estable: [%d]\n",temp)
		}
	}

}	