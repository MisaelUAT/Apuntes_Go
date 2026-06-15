# Introduccion a Go

En este curso aprenderemos los conceptos basicos de Go, un lenguaje de programacion creado por Google en el 2009. Go es un lenguaje de programacion de codigo abierto que hace facil construir software simple, confiable y eficiente.

## ¿Que es Go?
Go es un lenguaje de programacion de codigo abierto que hace facil construir software simple, confiable y eficiente. Go es un lenguaje de programacion de codigo abierto que hace facil construir software simple, confiable y eficiente.

## Caracteristicas de Go
- Go es un lenguaje de programacion de codigo abierto
- Go es un lenguaje compilado
- Su sintaxis es sencilla de leer
- Permite crear programas de alto rendimiento

## Comandos Basicos de Go

- go run: Permite ejecutar un programa en Go

```go
go run
```
> Para ejecutar un programa en Go se debe usar el comando go run seguido del nombre del archivo con extension .g

o tambien:

```go
go run <archivo>.go
```

- go build: Permite compilar un programa en Go

```go
go build
```
> Compilar un programa significa que se creara un archivo ejecutable
- go fmt: Permite formatear un programa en Go

```go
go fmt
```
> Formatear el programa significa que el codigo se vera mas limpio y ordenado

## Estructura Basica de un lenguaje en Go

Normalmente como en todos los programas, suelen tener una estructura basica al inicar un proyecto del mismo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola Mundo")
}
```

Explicacion de la estructura:

- package main: Indica que el archivo es un ejecutable

- import "fmt": Importa la libreria fmt que permite imprimir en consola

- func main(): Es la funcion principal del programa

- fmt.Println("Hola Mundo"): Imprime en consola el mensaje "Hola Mundo"

**El resultado sera:**
![Hola Mundo](/Assets/go_run.png)

## Cadenas de Texto

En Go, las cadenas de texto se definen con comillas dobles. Por ejemplo:

```go
palabra:= "Hola Mundo";
```
> En este caso la cadena de texto esta almacenando un string llamado "Hola Mundo".

- Si quieres dar un salto de renglon en una cadena de texto, puedes hacerlo con el caracter \n. Por ejemplo:

```go
palabra:= "Hola \n Mundo";
```
