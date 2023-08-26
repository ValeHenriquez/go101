# Go es rápido, sencillo y productivo

En términos generales, los lenguajes compilados se ejecutan mucho más rápido que los lenguajes interpretados, y Go no es una excepción.

Go es uno de los lenguajes de programación más rápidos, superando con facilidad a JavaScript, Python y Ruby en la mayoría de las pruebas de rendimiento.

Sin embargo, el código en Go no se ejecuta *tan* rápido como sus contrapartes compiladas en Rust y C. Dicho esto, *compila* mucho más rápido que estos dos, lo que hace que la experiencia del desarrollador sea muy productiva.

***

# Comparando la Velocidad de Go

Go es *generalmente* más rápido y liviano que los lenguajes interpretados o basados en máquinas virtuales, como:

* Python
* JavaScript
* PHP
* Ruby
* Java

Sin embargo, en términos de velocidad de ejecución, Go se queda atrás de algunos otros lenguajes compilados como:

* C
* C++
* Rust

Go es un poco más lento principalmente debido a su manejo automático de memoria, también conocido como "Go runtime". Una velocidad ligeramente menor es el precio que pagamos por la seguridad en la memoria y una sintaxis simple.

![comparación de velocidad](https://miro.medium.com/max/2020/1*nlpYI256BR71xMBWd1nlfg.png)

***

# Compilado vs Interpretado

Los programas compilados pueden ejecutarse sin acceso al código fuente original y sin acceso a un compilador.

Observa cómo esto es diferente de los lenguajes interpretados como Python y JavaScript. Con Python y JavaScript, el código se interpreta en [tiempo de ejecución](https://en.wikipedia.org/wiki/Runtime_(program_lifecycle_phase)) por un programa separado conocido como el "intérprete". Distribuir código para que los usuarios lo ejecuten puede ser complicado porque necesitan tener un intérprete instalado y acceso al código fuente original.

## Ejemplos de lenguajes compilados

* Go
* C
* C++
* Rust

## Ejemplos de lenguajes interpretados

* JavaScript
* Python
* Ruby

![compilado vs interpretado](https://i.imgur.com/ovHaWmS.jpg)

***

# Go es Fuertemente Tipado

Go tiene un tipado fuerte y estático, lo que significa que las variables solo pueden tener un único tipo. Una variable `string` como "hello world" no puede cambiarse a un `int`, como el número `3`.

Uno de los mayores beneficios de la tipificación fuerte es que los errores pueden detectarse en "tiempo de compilación". En otras palabras, los errores son más fáciles de identificar antes de tiempo porque se detectan cuando el código se compila antes de ejecutarse.

Contrasta esto con la mayoría de los lenguajes interpretados, donde los tipos de variables son dinámicos. La tipificación dinámica puede llevar a errores sutiles que son difíciles de detectar. Con los lenguajes interpretados, el código *debe* ejecutarse (a veces en producción si tienes mala suerte 😨) para detectar errores de sintaxis y tipo.

## Concatenación de cadenas

Dos cadenas pueden [concatenarse](https://en.wikipedia.org/wiki/Concatenation) con el operador `+`. Debido a que Go es fuertemente tipado, no te permitirá concatenar una variable de cadena con una variable numérica.

***

# Los programas en Go son amigables con la memoria

Los programas en Go son bastante livianos. Cada programa incluye una pequeña cantidad de código "extra" que se incluye en el binario ejecutable. A este código extra se le llama el [Go Runtime](https://go.dev/doc/faq#runtime). Uno de los propósitos del Go Runtime es limpiar la memoria no utilizada en tiempo de ejecución.

En otras palabras, el compilador de Go incluye una pequeña cantidad de lógica adicional en cada programa de Go para facilitar a los desarrolladores la escritura de código que sea eficiente en memoria.

## Comparación

Como regla general, los programas en Java usan *más* memoria que los programas comparables en Go porque Go no utiliza una máquina virtual completa para ejecutar sus programas, solo un pequeño runtime. El runtime de Go es lo suficientemente pequeño como para incluirlo directamente en el código de máquina compilado de cada programa en Go.

Como otra regla general, los programas en Rust y C++ usan ligeramente *menos* memoria que los programas en Go porque se le da más control al desarrollador para optimizar el uso de memoria del programa. El runtime de Go simplemente lo maneja automáticamente por nosotros.

## Uso de memoria en reposo

![uso de memoria en reposo](https://miro.medium.com/max/1400/1*Ggs-bJxobwZmrbfuoWGpFw.png)

En el gráfico de arriba, [Dexter Darwich compara el uso de memoria](https://medium.com/@dexterdarwich/comparison-between-java-go-and-rust-fdb21bd5fb7c) de tres programas *muy* simples escritos en Java, Go y Rust. Como puedes ver, Go y Rust utilizan *muy poca* memoria en comparación con Java.
