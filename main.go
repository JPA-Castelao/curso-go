package main

import (
	"curso-go/pruebas"
	"errors"
	"fmt"
)
var nivelPaquete =6
const Pi =3.14

const (
    A=1
    B=2
)


func Sapo(){
    fmt.Println("Sapooo")
}
func addicionar(x,y int) int{
    return x+y
}

func dividir(x,y int)(int,error){

    if y <=0{
        return 0, errors.New("No puede ser 0")
    }

    return x/y,nil
}

func main() {
    // fmt.Println("Hello, Gopher!")

    // fmt.Println(mathutil.Add(12,5))
    // fmt.Println(mathutil.Resta(12,5))
    // fmt.Println(math.Sqrt(5))

    // var x int = 5
    // fmt.Println(x)

    // var a,b = 3,3
    // fmt.Println(a+b)

    // y:=4
    // fmt.Println(y)

    // nombre,edad:="Jerson",26

    // fmt.Println(nombre,edad)

    // fmt.Println(mathutil.Sapo)
    // fmt.Println(Pi)

    // fmt.Println(A+B)

    // suma:=addicionar(34,55)
    // Sapo()

    // //la funcion div devuelve 
    // div,err:=dividir(1,0)
    // if err!=nil{
    //     fmt.Println(err)
    // }
    // fmt.Println(div)

  


    // fmt.Println(suma)



    // //alcance limitado
    // if y:=10;y>5{
    //     fmt.Println("y es mayor de 5")
    // }

    // //if else chainings
    // if w:=1;w>0{
    //         fmt.Println("w es mayor que 0")
    // }else if w==0{
    //         fmt.Println("w es igual a")
    // }else{
    //         fmt.Println(" w es negativo")
    // }


    // //for
    // for i:=1;i<=5;i++{
    //     fmt.Println(i)
    // }

    // nums:=[]int{1,2,3}
    // for _,num:= range nums{
    //     fmt.Println(num)
    // }

    // //return


    lecturas:=[]int{50,46,43,-24,123,0,34,187}

    sensores:=map[string]int{
        

    }

    pruebas.ProcesarSensores(lecturas)

}