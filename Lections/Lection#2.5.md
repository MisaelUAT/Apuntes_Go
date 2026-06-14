# Operaciones Basicas en Go

## Operaciones de Texto(Strings)
- **Concatenacion:** Permite unir dos o mas variables de texto o parrafos. Puede ser de dos formas:

```go
string_concatenado:= "Hola como estas" + "Amigo mio?"
fmt.Println(string_concatenado)
```
> El resultado sera: "Hola como estasAmigo mio?"

```go
texto1:= "Hola"
texto2:= "como estas"
string_concatenado:= texto1 + " " + texto2
fmt.Println(string_concatenado)
```
> El resultado sera: "Hola como estas"

- **Longitud:** Permite obtener la longitud de una cadena de texto

```go
texto:="Misael"
longitud:= len(texto)
fmt.Println(longitud)
```
> El resultado sera la cantidad de caracteres que conforma el texto en este caso es de 6

## Operaciones con el Paquete "fmt"

- **fmt.priintln:** Permite imprimir en consola un mensaje o una variable

```go
fmt.Println("Hola Mundo")
```
> El resultado sera: "Hola Mundo"

- **fmt.Printf:** Permite imprimir en consola un mensaje o una variable con formato. Para agregar el valor de una variable o constante se debe de usar un **comodin**, existen los siguientes comodines:

  - `%s` string
  - `%d` int
  - `%f` float
  - `%t` bool
  - `\n` salto de línea

```go
nombre:= "Misael"
fmt.Printf("Hola %s", nombre)
```
> El resultado sera: "Hola Misael"

- **fmt.Scanln:** Permite leer un valor desde la consola y almacenarlo en una variable. Se usa el simbolo de & junto al nombre de la variable para almacenar el valor ingresado por el usuario en la variable

```go
var nombre string
fmt.Println("Ingresa tu nombre:")
fmt.Scanln(&nombre)
fmt.Println(nombre)
```
> El resultado sera el valor que el usuario ingrese en la consola

## Operaciones Aritmeticas con Go

- **Suma:** Permite sumar dos o mas numeros

```go
num1:= 5
num2:= 10
suma:= num1 + num2
fmt.Println(suma)
```
> El resultado sera: 15

- **Resta:** Permite restar dos o mas numeros

```go
num1:= 10
num2:= 5
resta:= num1 - num2
fmt.Println(resta)
```
> El resultado sera: 5

- **Multiplicacion:** Permite multiplicar dos o mas numeros

```go
num1:= 5
num2:= 10
multiplicacion:= num1 * num2
fmt.Println(multiplicacion)
```
> El resultado sera: 50

- **Division:** Permite dividir dos o mas numeros

```go
num1:= 10
num2:= 5
division:= num1 / num2
fmt.Println(division)
```
> El resultado sera: 2

- **Modulo:** Permite obtener el residuo de una division

```go
num1:= 10
num2:= 3
modulo:= num1 % num2
fmt.Println(modulo)
```
> El resultado sera: 1





