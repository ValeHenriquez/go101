# Variables


### Tipos Básicos

Los tipos básicos de variables en Go son:


`uint`
* uint8   unsigned  8-bit integers (0 to 255)
* uint16  unsigned 16-bit integers (0 to 65535)
* uint32  unsigned 32-bit integers (0 to 4294967295)
* uint64  unsigned 64-bit integers (0 to 18446744073709551615)

`int`
* int8    signed  8-bit integers (-128 to 127)
* int16   signed 16-bit integers (-32768 to 32767)
* int32   signed 32-bit integers (-2147483648 to 2147483647)
* int64   signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

`byte` // alias for uint8

`rune` // alias for int32
       // represents a Unicode code point

`float32` `float64`

`string`


### Declarar una variable

Las variables se declaran utilizando la palabra clave `var`. Por ejemplo

```go
var pi float64 = 3.141
var name string 
var total uint = 12 
```

### Valores default

En Go, las variables tienen valores por defecto según su tipo. 

```go
var numInt int            // Valor por defecto: 0
var numFloat64 float64    // Valor por defecto: 0.0
var boolValue bool        // Valor por defecto: false
var stringValue string    // Valor por defecto: ""
```

### Declaración corta

Se puede utilizar la declaración de asignación corta `:=` en lugar de una declaración `var`. El operador `:=` infiere el tipo de la nueva variable en función del valor.

```go
var name string
```

```go
name := ""
```

### Declaración en la misma linea

```go
size, country := 20000000, "Chile"
```

```go
size := 20000000
country := "Chile"
```

### Conversión  de tipos

```go
metersFloat := 10.4

metersFloat := int(metersFloat)

```

### Constantes

Las constantes deben ser conocidas en tiempo de compilación

```go
const myAge = 30
```

# Formateo de Strings

Go sigue la [tradición printf](https://cplusplus.com/reference/cstdio/printf/) del lenguaje C. 

* [fmt.Printf](https://pkg.go.dev/fmt#Printf) - Imprime una cadena formateada.
* [fmt.Sprintf()](https://pkg.go.dev/fmt#Sprintf) - Devuelve la cadena formateada.

#### [Documentación fmt Package](https://pkg.go.dev/fmt)

#### Ejemplos

```go
fmt.Println("Hola Mundo")

fmt.Printf("Tengo %v años", 10)

fmt.Printf("Tengo %v años", "muchos")

result := fmt.Sprintf("Tengo %v años", 10)

```

# Operadores 

```go
func main() {
	// Operadores Aritméticos (), *, /, %, +, -
	var a = 4 + 2*3
	fmt.Println(a)
	// Operadores de asignación: =, +=, -=, *=, /=, %=
	var b = 10
	b += 2
	fmt.Println(b)
	// declaración post-incremento y post-decremento: ++, --
	var c = 3
	c--
	fmt.Println(c)
	// Operadores Comparación: >, <, ==, !=, >=, <=
	fmt.Println(4 != 4)
	// Operadores Lógicos &&, ||
	var age = 70
	fmt.Println("adulto:", age >= 18 && age <= 60)
	fmt.Println("niño o viejo:", age < 18 || age > 60)
	// Unario: !
	fmt.Println(!(4 != 4))
}

```

# Condicionales

#### If-else

```go
func main() {
    emoji := "😋"

    if emoji == "🌵" {
		fmt.Println("es un cactus")
	} else if emoji == "😋" {
		fmt.Println("es una carita")
	} else {
		fmt.Printf("el emoji es: %s\n", emoji)
	}

	fmt.Println(emoji)
}
```


#### Switch

```go
func main() {
	emoji := "🍎"

	switch {
	case emoji == "🐈" || emoji == "🐕":
		fmt.Println("es un animal")
	case emoji == "🍌" || emoji == "🍎":
		fmt.Println("es una fruta")
	default:
		fmt.Println("no es un animal ni una fruta")
	}
}
```

# Ciclo `for`

En Go, solo existe una estructura de bucle, el bucle for, que puede ser utilizado de diversas formas.

#### Bucle for básico:

```go
func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

#### Ciclo for como "while":
Se puede usar el bucle for como un bucle "while" simplemente omitiendo las declaraciones de inicialización y post-incremento.

```go
func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

# Array

```go
func main() {
    // Declarar e inicializar un array de enteros
    var numeros [5]int
    numeros[0] = 1
    numeros[1] = 2
    numeros[2] = 3
    numeros[3] = 4
    numeros[4] = 5

    // Acceder e imprimir elementos del array
    fmt.Println(numeros[0]) // Imprime 1
    fmt.Println(numeros[2]) // Imprime 3

    // Declarar e inicializar un array de cadenas
    nombres := [3]string{"Alice", "Bob", "Charlie"}
    fmt.Println(nombres[1]) // Imprime "Bob"

    arr := [5]int{1, 2, 3, 4, 5}

    // Recorrer el array utilizando un bucle for
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

     // Recorrer el array utilizando un bucle for range
    for index, value := range arr {
        fmt.Printf("Índice: %d, Valor: %d\n", index, value)
    }
}
```

# Slice

A diferencia de los arrays, los slices tienen una longitud dinámica, lo que significa que pueden crecer o reducirse según sea necesario. Los slices también son más flexibles en términos de manipulación y manejo de datos.

```go
func main() {
    // Declarar e inicializar un slice de enteros
    numeros := []int{1, 2, 3, 4, 5}

    // Acceder e imprimir elementos del slice
    fmt.Println(numeros[0]) // Imprime 1
    fmt.Println(numeros[2]) // Imprime 3

    // Modificar elementos del slice
    numeros[1] = 10

    // Agregar elementos al slice (append)
    numeros = append(numeros, 6, 7, 8)

    // Crear un nuevo slice a partir de un subconjunto
    subSlice := numeros[2:5] // Incluye índices 2, 3, 4
    fmt.Println(subSlice)     // Imprime [3 6 7]

    // Crear un slice usando make
    otroSlice := make([]int, 3, 5) // Longitud 3, Capacidad 5

    // Copiar un slice
    copiaSlice := make([]int, len(numeros))
    copy(copiaSlice, numeros)

    // Eliminar elementos de un slice
    numeros = append(numeros[:2], numeros[3:]...)
    fmt.Println(numeros) // Imprime [1 10 6 7 8]
}
```
* append: Se utiliza para agregar uno o varios elementos a un slice.
* [:]: Se utiliza para crear subconjuntos de slices.
* make: Se utiliza para crear un nuevo slice con longitud y capacidad * especificadas.
* copy: Se utiliza para copiar elementos de un slice a otro.
* ...: Se utiliza para expandir un slice en una lista de argumentos para la función append.

# Funciones

```go
// Definición de una función que suma dos números
func suma(a, b int) int {
    return a + b
}

func dividir(a,b int) (int, err){
       if b != 0 {
              return a/b, nil
       }
       return 0, err
}

// Definición de una función que imprime un saludo
func saludar(nombre string) {
    fmt.Println("¡Hola,", nombre, "!")
}

func main() {
    // Llamada a la función suma y uso del resultado
    resultado := suma(5, 3)
    fmt.Println("Resultado:", resultado) // Imprime "Resultado: 8"

    // Llamada a la función saludar
    saludar("Alice") // Imprime "¡Hola, Alice !"
}
```

La convención para **exportar** funciones, variables, tipos y constantes a otros paquetes es bastante simple. Cualquier identificador (nombre de función, variable, tipo, etc.) que comience con una **letra mayúscula** se considera exportado y puede ser accedido desde otros paquetes

# Defer

La declaración defer se utiliza para posponer la ejecución de una función hasta que la función que lo contiene haya terminado. La función pospuesta se ejecutará justo antes de que la función envolvente finalice, ya sea por una declaración return, un panic u otro motivo de salida. Ej: Lectura de archivos.

### Ejemplo

```go
func main() {
	fmt.Println("Inicio de la función main")

	// La función dentro del defer se ejecutará al final de la función main
	defer fmt.Println("Defer 1: Esto se imprimirá al final")

	// Defer se ejecuta en el orden inverso en que se declaran
	defer fmt.Println("Defer 2: Esto se imprimirá antes que Defer 1")

	fmt.Println("Fin de la función main")
}
```


# Struct

Estructuras

```go
// Definición de un struct
type Persona struct {
    Nombre   string
    Edad     int
    Profesion string
}

func main() {
    // Crear una instancia de Persona
    persona1 := Persona{
        Nombre:   "Alice",
        Edad:     30,
        Profesion: "Ingeniera",
    }

    // Acceder a los campos del struct
    fmt.Println(persona1.Nombre)   // Imprime "Alice"
    fmt.Println(persona1.Edad)     // Imprime 30
    fmt.Println(persona1.Profesion) // Imprime "Ingeniera"

    // Modificar un campo del struct
    persona1.Edad = 31

    // Crear otra instancia de Persona
    persona2 := Persona{
        Nombre:   "Bob",
        Edad:     28,
        Profesion: "Desarrollador",
    }

    // Imprimir el struct completo
    fmt.Println(persona2) // Imprime "{Bob 28 Desarrollador}"
}
```

# Metodos

Funciones asociadas a los structs

```go
// Definición de un tipo de dato struct
type Rectangulo struct {
    Ancho  float64
    Altura float64
}

// Definición de un método asociado al tipo Rectangulo
func (r Rectangulo) Area() float64 {
    return r.Ancho * r.Altura
}

func main() {
    // Crear una instancia de Rectangulo
    rect := Rectangulo{Ancho: 10, Altura: 5}

    // Llamar al método Area en la instancia rect
    area := rect.Area()
    fmt.Println("Área:", area) // Imprime "Área: 50"
}
```

# Interfaces
Una interfaz especifica un contrato que otros tipos deben cumplir si desean considerarse implementaciones válidas de esa interfaz.

Se implementan implícitamente.

```go
// Definición de una interfaz
type Forma interface {
    Area() float64
    Perimetro() float64
}

// Definición de un tipo struct
type Rectangulo struct {
    Ancho  float64
    Altura float64
}

// Implementación de los métodos de la interfaz para Rectangulo
func (r Rectangulo) Area() float64 {
    return r.Ancho * r.Altura
}

func (r Rectangulo) Perimetro() float64 {
    return 2*r.Ancho + 2*r.Altura
}

func main() {
    // Crear una instancia de Rectangulo
    rect := Rectangulo{Ancho: 10, Altura: 5}

    // La instancia de Rectangulo cumple con la interfaz Forma
    var forma Forma
    forma = rect

    // Llamar a los métodos de la interfaz
    fmt.Println("Área:", forma.Area())         // Imprime "Área: 50"
    fmt.Println("Perímetro:", forma.Perimetro()) // Imprime "Perímetro: 30"
}
```
# Panic 
En términos simples, en Go, panic es una forma de detener bruscamente la ejecución del programa debido a un error crítico. Es como una señal de alarma que indica que algo inesperado y grave ha sucedido y que el programa no puede continuar su funcionamiento normal. Cuando se produce un panic, el programa se detiene y muestra información sobre el error
```go
package main

import "fmt"

func main() {
	division(10, 2)
	division(40, 3)
	division(12, 0)
	division(20, 4)
}

func division(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperándome del panic", r)
		}
	}()

	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("🤦‍♀️")
	}
}
```

# Errores
Los errores en Go son representados como valores que implementan la interfaz error, que consiste en un solo método llamado Error() que devuelve una cadena que describe el error.

```go
type error interface {
    Error() string
}
```
### Ejemplo
```go
package main

import (
    "fmt"
    "errors"
)

// Función que devuelve un error
func dividir(dividendo, divisor float64) (float64, error) {
    if divisor == 0 {
        return 0, errors.New("no se puede dividir por cero")
    }
    return dividendo / divisor, nil
}

func main() {
    resultado, err := dividir(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Resultado:", resultado) // Imprime "Resultado: 5"
    }

    resultado, err = dividir(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // Imprime "Error: No se puede dividir por cero"
    } else {
        fmt.Println("Resultado:", resultado)
    }
}
```

# Punteros

```go
func main() {
    valor := 42
    puntero := &valor // Obtener la dirección de memoria de la variable 'valor'

    fmt.Println("Valor:", valor)       // Imprime "Valor: 42"
    fmt.Println("Puntero:", puntero)   // Imprime la dirección de memoria
    fmt.Println("Valor en puntero:", *puntero) // Imprime el valor almacenado en la dirección apuntada por el puntero
}
```

## Funciones que reciben punteros

```go
func modifyValue(x *int) {
    *x = *x * 2
}

func main() {
    value := 5
    modifyValue(&value)
    fmt.Println(value) // Imprimirá 10, ya que el valor original fue modificado dentro de la función
}
```

# Packages

Los paquetes en Go agrupan funciones, tipos, variables y otros elementos relacionados en una sola unidad. Cada archivo fuente Go pertenece a un paquete y se declara al principio del archivo con la instrucción package.

### Ejemplo

**math.go:** 
```go
package math

func Sumar(a, b int) int {
    return a + b
}

func Restar(a, b int) int {
    return a - b
}
```

**main.go:**
```go
package math

func Sumar(a, b int) int {
    return a + b
}

func Restar(a, b int) int {
    return a - b
}
```

# Module

Los módulos son una característica para gestionar dependencias y versiones de manera más estructurada y eficiente. Al crear un módulo, se inicia el manejo de dependencias y versiones en tu proyecto.

Por ejemplo, si el proyecto se llama "myapp", se inicializa el módulo con el siguiente comando:


```cmd
go mod init myapp
```

El comando go mod init se ejecuta desde la línea de comandos en el directorio raíz de tu proyecto y toma un único argumento, que es el nombre del módulo. El argumento debe ser una ruta que identifique de manera única el módulo. El comando crea un archivo llamado go.mod en el directorio actual que contiene la información sobre el módulo y sus dependencias.