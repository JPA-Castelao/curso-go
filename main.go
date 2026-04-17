package main

import (
	"errors"
	"fmt"
	"time"
)
var nivelPaquete =6
const Pi =3.14

const (
    A=1
    B=2
)

const MasaRobot =72.4

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

    //arrays
    // nums:=[]int{1,2,3}
    // for _,num:= range nums{
    //     fmt.Println(num)
    // }

    // //return


//     lecturas:=[]int{50,46,43,-24,123,0,34,187}

//     sensores:=map[string]int{
//     "Brazo Izquierdo": 32,
//     "Motor Trasero":  110,
//     "Cámara":         -5,


//     pruebas.ProcesarSensores(sensores)
//     pruebas.ProcesarTemperaturas(lecturas)


 

// fmt.Printf("\nMasa en Marte: %.2f",mathutil.CalcularPeso(MasaRobot,mathutil.GravedadMarte))
// fmt.Printf("\nMasa en Tierra: %.2f",mathutil.CalcularPeso(MasaRobot,mathutil.GravedadTierra))

// //switch
// letra:="a"
// switch letra {
// case "a","e","i","o","u":
//     fmt.Println("Es una vocal")
//     default:
//     fmt.Println("Consonante")
// }

// num:=0
// switch num {
// case 0:
// case 1:
//         fmt.Println("Cero-1")
// default:
//     fmt.Println("Otro número")

// }




    
//  var fruta string="manzana"
// fruta  = "Manzana golden";
// switch{
// case strings.Contains(strings.ToLower(fruta),"golden"):
//      fmt.Println("Es una manzana")
// case strings.Contains(strings.ToLower(fruta),"reineta"):
   
//      fmt.Println("Es una manzana")
   
//     default:
//     fmt.Println("No es una manzana")
//     }




// x:=2

// switch x{
// case 1:
//     fmt.Println("One")
// case 2:
//     fmt.Println("Two")
//     break 
// case 3:
//     fmt.Println("Three")    
// }
// fmt.Println("Switch completo")


//continue se salta una iteracion
// for i:=0;i<5;i++{
//     if i==2 {
//         continue
//     }
//         fmt.Println(i)
// }   



// m:=map[string]int{"a":1,"b":2,"c":3}
// coleccion:=[]string{"manzana","pera","uva"}

// for key,value:= range m{
//     fmt.Printf("Clave: %s, Valor: %d\n",key,value)
// }
// for v:=range coleccion{
//     fmt.Printf("Valor: %s\n",v)

// }


// sensoreCamara1:=map[string]int{"Brazo Izquierdo": 32, "Motor Trasero": 110, "Cámara": -5}

// for sensor,valor:= range sensoreCamara1{

// fmt.Printf("\n%s: %d",sensor,valor)



// }

//esto es un ejemplo de defer, se ejecuta al final de la funcion, incluso si hay un return antes
// func deferir(){
//     defer fmt.Println("Esto va al final")
//     fmt.Println("Esto va al principio")
// }


// type, crea tipos personalizados

// cantidad :=Litro(58)

// fmt.Printf("\n Litros: %.2f",cantidad)
// fmt.Printf("\n Galones: %.2f",cantidad.ToGalones())

// type Litro float64
// const LitrosPorGalon = 3.78541
// func (l Litro) ToGalones() float64{
//     return float64(l)/LitrosPorGalon
// }


//struct, es una coleccion de campos, cada campo tiene un nombre y un tipo
// p:=Persona{Nombre:"Jerson",Edad:26}


// fmt.Print(p)

// type  AR float64

// r:=Rectangulo{Width:5,Height:10}

// ar:=AR(r.AreaRectangulo())
// fmt.Println(ar)

// func ( r Rectangulo) AreaRectangulo() float64{

//     return r.Width * r.Height
// }

// type Persona struct {
//     Nombre string
//     Edad int 
// }


// type Rectangulo struct {
//     Width float64
//     Height float64

// }


//go routines




printNumbers(4)

go printNumbers(5)
go printNumbers(5)

time.Sleep(time.Second)

ch :=make(chan int)

    go func(){
        ch<- 42
    }()
    val:=<-ch
    fmt.Println(val)


}
func printNumbers(n int){

    for i:=0;i<=n;i++{
        fmt.Println(i)
        time.Sleep(100*time.Millisecond)
    }
    
}



