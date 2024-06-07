# Estructuras de Control

## Introducción

Una estructura de decision es una estructura de control que permite tomar decisiones basadas en una condición. En Go, las estructuras de decisión son `if`, `else if` y `else`. En esta leccion aprenderemos sobre su uso, y el como se ajusta con los operadores de comparacion y logicos.

## Estructura de Control If

La estructura de control `if` se utiliza para tomar decisiones basadas en una condición. Por ejemplo:

```go
// Variable a evaluar
edad:= 16;
// Estructura de control
if edad >= 18{
    fmt.Println("Eres mayor de edad");
}
```
> En este sencillo programa se evalua si la variable `edad` es mayor o igual a 18. Si es asi, se imprime en pantalla "Eres mayor de edad". 

### Else

La setencia `if-else` se utiliza para tomar poder agregarle a nuestro if, mas opciones para evaluar. Por ejemplo:

```go
// Variable a evaluar
edad:= 16;
// Estructura de control if
if edad >= 18{
    fmt.Println("Eres mayor de edad");
}
// La setencia else
else{
    fmt.Println("Eres menor de edad");
}
```

> Cuando es else se ejecuta si la condicion del if no se cumple con ninugna de las condiciones anteriores.

### Else If

La setencia `else if` se utiliza para agregar mas condiciones a nuestro if. Por ejemplo:

```go
edad := 15

if edad <= 17 {
	fmt.Println("Es un niño")
} else if edad >= 18 {
	fmt.Println("Es un adulto")
}
```
> En este caso, si la edad es menor o igual a 17, se imprime "Es un niño". Si no, se evalua si la edad es mayor o igual a 18, en cuyo caso se imprime "Es un adulto". Se que te estaras preguntando que significan los simbolos `<=` y `>=`, estos son operadores de comparacion, los cuales veremos en el siguiente apartado.

## Operadores de Comparación

Los operadores de comparación se utilizan para comparar dos valores. En Go, los operadores de comparación son:

- Operador Mayor que `>`: Se utiliza para verificar si un valor es mayor que otro.

Ejemplo:

```go
a := 10
b := 20

if a > b {
        fmt.Println("a es mayor que b")
 }
 ```
> En este caso, la condición `a > b` es falsa porque 10 no es mayor a 20, por lo que el mensaje no se imprime en pantalla. 

- Operador Menor que `<`: Se utiliza para verificar si un valor es menor que otro.

Ejemplo:

```go
a := 10
b := 20

if a < b {
        fmt.Println("a es menor que b")
 }
 ```
> En este caso, la condición `a < b` es verdadera porque 10 es menor a 20, por lo que el mensaje se imprime en pantalla.

- Operador Mayor o igual que `>=`: Se utiliza para verificar si un valor es mayor o igual que otro.

Ejemplo:

```go
a := 10
b := 20

if a >= b {
        fmt.Println("a es mayor o igual que b")
 }
 ```

> En este caso, la condición `a >= b` es falsa porque 10 no es mayor o igual a 20, por lo que el mensaje no se imprime en pantalla.

- Operador Menor o igual que `<=`: Se utiliza para verificar si un valor es menor o igual que otro.

Ejemplo:
```go
a := 10
b := 20

if a <= b {
        fmt.Println("a es menor o igual que b")
 }
 ```
> En este caso, la condición `a <= b` es verdadera porque 10 es menor a 20 pero no es igual, por lo que el mensaje se imprime en pantalla. Asi que aqui puede cumplirse si la variable es igual que la otra o que sea menor que la otra.

- Operador Igual que `==`: Se utiliza para verificar si un valor es igual a otro.

Ejemplo:

```go
a := 10
b := 20

if a == b {
        fmt.Println("a es igual que b")
 }
 ```
 > En este caso, la condición `a == b` es falsa porque 10 no es igual a 20, por lo que el mensaje no se imprime en pantalla.

- Operador Diferente que `!=`: Se utiliza para verificar si un valor es diferente a otro.

Ejemplo:

```go
a := 10
b := 20

if a != b {
        fmt.Println("a es diferente que b")
 }
 ```
 > En este caso, la condición `a != b` es verdadera porque 10 es diferente a 20, por lo que el mensaje se imprime en pantalla.

 ## Operadores Lógicos

Los operadores lógicos se utilizan para combinar dos o más condiciones dentro de una misma setencia. En Go, los operadores lógicos son:

- Operador AND `&&`: Se utiliza para verificar si dos o más condiciones son verdaderas.

Ejemplo:

```go

a := 10
b := 20

if a > 5 && b > 10 {
        fmt.Println("a es mayor que 5 y b es mayor que 10")
 }
 ```
 > En este caso, la condición `a > 5 && b > 10` es verdadera porque 10 es mayor que 5 y 20 es mayor que 10, por lo que el mensaje se imprime en pantalla.


- Operador OR `||`: Se utiliza para verificar si al menos una de las condiciones es verdadera.

Ejemplo:

```go

a := 10
b := 20

if a > 5 || b > 10 {
        fmt.Println("a es mayor que 5 o b es mayor que 10")
 }
 ```
 > En este caso, la condición `a > 5 || b > 10` es verdadera porque 10 es mayor que 5, por lo que el mensaje se imprime en pantalla.


- Operador NOT `!`: Se utiliza para negar una condición.

Ejemplo:

```go

a := 10
b := 20

if !(a > 5) {
        fmt.Println("a no es mayor que 5")
 }
 ```
 > En este caso, la condición `!(a > 5)` es falsa porque 10 es mayor que 5, por lo que el mensaje no se imprime en pantalla. En pocas palabras el operador `!` niega la condicion que se le pase aunque sea verdadera.

 