# Variables en Go
Las variables son un espacio de memoria que se utiliza para almacenar un valor.

## Variables Globales

La estructura de una variable en Go es la siguiente:

```go
var nombre_variable tipo_dato
```

Un ejemplo seria:

```go
var nombre string
```
## Variables Locales

Otro ejemplo de declarar su variable seria con el operador de asignacion pero debe de tener un valor definidio:

```go
nombre:= "Juan"
```

## Constantes

Las constantes son variables que no pueden cambiar su valor. Se declaran de la siguiente manera:

```go
const nombre string = "Juan"
```
> En este caso la constante nombre almacenara un string llamado "Juan" y no puedes ser cambiado pero si se puede calcular su valor.

- Puedes declarar varias constantes en una sola:

```go
	const(
		num_1= 10
		num_2= 30
		Texto_1= "Holis"
		condicion= true
	)
```
> En este caso se declaran 4 constantes en una sola.

Ejemplo:
```go
const num_3= num_1 + num_2
```

## Tipos de Datos Basicos

En Go, los tipos de datos son los siguientes:

### String

Las cadenas de texto se definen con comillas dobles. Por ejemplo:

```go
palabra:= "Hola Mundo";
```

### Int

Los numeros enteros se definen con el tipo de dato int. Por ejemplo:

```go
numero:= 10;
```

### Float

Los numeros decimales se definen con el tipo de dato float. Por ejemplo:

```go
decimal:= 10.5;
```

### Boolean

Los valores booleanos se definen con el tipo de dato bool. Por ejemplo:

```go
booleano:= true;
```

> Nota: Existen mas tipos de Valores y clasificaciones, puedes consultarlo en la documentacion oficial de Go.