# Resumen de Variables y Constantes en Go

## Como se declaran?
Antes de empezar a realizar codigo es importante que tengas conocimiento la diferencia entre una variable y una comnstante, y como se declaran en el lenguaje de programación Go.

### Variables
Es un espacio de memoria que puede contener un valor guardado en si mismo, lo puedes modificar a lo largo del codigo, es decir que no es un valor fijo sino cambiable, se declaran de la siguiente manera:

```go
var nombreVariable tipoDato
```

#### Forma Corta de declarar una variable

Existe una forma de declarar de forma sencilla una variable, la condicion es que se le asigne un valor en ese momento para que se le asigne automaticamente el tipo de dato, se declaran de la siguiente manera:

```go
edad := 22 // Esta variable sera de tipo int
```

### Constantes
Es un espacio de memoria igual que la variable, nada mas que la unica diferencia es que el valor que contenga siempre va a ser fijo, se declaran de la siguiente manera:

```go
const nombreConstante tipoDato = valor
```

## Tipos de Datos
En Go existen varios tipos de datos, los cuales se pueden clasificar en:
- **Tipos de datos básicos:** int, float64, string, bool, etc.
- **Tipos de datos compuestos:** arrays, slices, maps, structs, etc.
- **Tipos de datos de referencia:** pointers, channels, etc.

> Para que entiendas uwu:
> - **int:** Es un tipo de dato que representa un número entero, es decir, sin decimales. Ejemplo: 1, 2, 3, etc.
> - **float64:** Es un tipo de dato que representa un número con decimales. Ejemplo: 3.14, 2.718, etc.
> - **string:** Es un tipo de dato que representa una cadena de caracteres. Ejemplo: "Hola", "Misael", etc.
> - **bool:** Es un tipo de dato que representa un valor booleano, es decir, verdadero o falso. Ejemplo: true, false.

## Ejemplos de Variables y Constantes
```go
	var nombre string = "Misael Hernandez"
    const pi float64 = 3.14
	fmt.Print(nombre) // Misael Hernandez
	fmt.Print(pi) // 3.14
```
> El **fmt** es un paquete de go que nos ayuda a imprimir(mostrar) en la consola de nuestra PC el resultado de una variable o operacion.

Nada mas para quedar claro, vamos a intentar modificar el valor de la variable

```go
    var nombre string = "Misael Hernandez"
    const pi float64 = 3.14
    fmt.Print(nombre) // Misael Hernandez
    fmt.Print(pi) // 3.14

    nombre = "Misael" // Modificando el valor de la variable
    fmt.Print(nombre) // Misael

    pi = 3.1416 // Modificando el valor de la constante
    fmt.Print(pi) // Esto generará un error
```

> Genera error en la constante ya que no se puede modificar el valor de una constante, mientras que la variable si se puede modificar su valor.
