# Funciones en Go

## Definicion de funciones

Las funciones en Go son bloques de codigo reutilizables que agrupan una tarea especifica. Sirven para organizar mejor el programa, evitar repetir codigo y hacer que el proyecto sea mas facil de leer, mantener y depurar.

En otras palabras, una funcion recibe datos, realiza una accion y, si es necesario, devuelve un resultado.

## Para que sirve una funcion en Go

Las funciones sirven para:

- Reutilizar codigo sin escribir lo mismo varias veces.
- Dividir un programa grande en partes mas pequenas y ordenadas.
- Hacer que el codigo sea mas facil de entender.
- Facilitar pruebas y correccion de errores.

## Estructura basica de una funcion

La estructura general de una funcion en Go es la siguiente:

```go
func nombreFuncion() {
	// codigo que realiza una tarea
}
```

Aqui:

- `func` indica que estamos creando una funcion.
- `nombreFuncion` es el nombre que le damos a la funcion.
- Los parentesis `()` pueden contener parametros.
- Las llaves `{}` contienen el codigo que ejecutara la funcion.

### Como mandar a llamar una funcion basica

La estructura que se usa para poder usar una **funcion sin parametros:**

```go
func main(){
    nombreFuncion()
}

func nombreFuncion() {
	// codigo que realiza una tarea
}
```

> Siempre se llama las funciones apartir de la funcion principal `main` pero puedes tambien usarlas en **funciones que tu realices tambien**

### Ejemplos de una funcion basica

#### Para mostrar un mensaje

La estructura para realizar esta funcion es la siguiente:

```go
func main(){
//Llamada a la funcion Saludar()
Saludar()
}

//Funcion basica a usar
func Saludar(){
//Codigo a ejecutar cuando se mande a llamar
    fmt.Println("Hola,como estas?")
}
```

> Al momento de que la ejecutes, va a mostrar por consola el mensaje que esta en la funcion **Hola,como estas?**

#### Imprimir el resultado de una Operacion

```go
func main(){
Sumar() // Imprimira en pantalla 9-
}

// Estructura de la funcion Sumar
func Sumar(){
resultado:= 40 + 50
fmt.Println(resultado)
}
```

> Aqui en esta funcion solo lo que hace es imprimir el resultado de una operacion, pero tambien en las `funciones` puedes pedirle al usuario `valores` con el que trabajar en ellas.

## Parametros en una funcion

Los parametros son los datos que una funcion recibe para poder trabajar con ellos. Se escriben dentro de los parentesis de la funcion y cada parametro tiene un nombre y un tipo de dato.

La estructura general es la siguiente:

```go
func nombre(param1 tipoDato, param2 tipoDato) {
	// codigo que usa los parametros
}
```

Aqui:

- `param1` y `param2` son los datos que recibira la funcion.
- `tipoDato` indica el tipo de dato, por ejemplo `int`, `string` o `float64`.
- Dentro de la funcion puedes usar esos valores para hacer operaciones, mostrar mensajes o tomar decisiones.

### Ejemplos para trabajar con Parametros

#### Suma de dos numeros

En este ejemplo le pediremos al usuario que nos proporcione dos numeros de los cuales de ellos vamos a **sumar y imprimir el resultado**.

```go
func main(){
// El usuario dentro del llamado nos brinda 2 valores para poder llamar la funcion
Sumar(10,10) // Imprimira en pantalla 20
}

// Estructura de la funcion Sumar
// num1 y num2 son los parametros que pediremos al usuario
func Sumar(num1 int, num2 int){
// resultado almacena los valores que nos brinda el usuario
resultado:= num1 + num2
fmt.Println(resultado)
}
```

> A diferencia del ejemplo anterior de suma en el que nosotros establecimos los valores y solo imprimia el resultado, aqui estamos trabajando con los **valores que nos proporcione el usuario**

#### Validar Edad

En este ejemplo le pediremos al usuario que nos brinde una edad para poder validarla

```go
func main(){
// El usuario dentro del llamado nos brinda 1 valor para la funcion
validarEdad(22) // Imprimira "Es Mayor de Edad"
}

// Estructura de la funcion Sumar
// edad es el parametro que pediremos al usuario
func validarEdad(edad int){
//Evaluacion del valor que le pedimos al usuario
if edad >= 18{
    fmt.Println("Es mayor de Edad")
} else if edad <=17{
    fmt.Println("Es menor de Edad")
}
}
```

## Retorno de valores en una funcion

Cuando una funcion devuelve un valor, significa que ademas de hacer una tarea, nos regresa un resultado para poder usarlo despues en otra parte del programa.

La estructura general es la siguiente:

```go
func nombre(param1 tipoDato, param2 tipoDato) tipoRetorno {
    // codigo que procesa los datos
    return valor
}
```

Aqui:

- `tipoRetorno` indica el tipo de dato que va a regresar la funcion, por ejemplo `int`, `string` o `bool`.
- `return` es la palabra que se usa para devolver el valor.
- El valor que se **devuelve** debe ser del **mismo tipo que se declaro en la funcion.**

### Ejemplos de retorno de valores en una funcion

#### Imprimir una cadena Concatenada

En este ejemplo le pediremos al usuario dos cadenas de texto para imprimirlas en un solo mensaje

```go
func main() {
    // Llamamos a la funcion y guardamos el resultado en una variable
    resultado := concatenarCadenas("Hola", " Mundo")

    // Imprimimos el valor que regreso la funcion
    fmt.Println(resultado)
}

// Funcion que recibe dos cadenas y regresa una sola cadena unida
func concatenarCadenas(cadena1 string, cadena2 string) string {
    // Unimos las dos cadenas en una sola variable
    textoConcatenado := cadena1 + cadena2

    // Retornamos el resultado final
    return textoConcatenado
}
```
> Imprimira: **Hola Mundo**

#### Comparar numeros
```go
func main() {
    // Llamamos a la funcion y mostramos el numero mayor
    fmt.Println(mayorNumero(20, 10))
}

// Funcion que recibe dos numeros y devuelve el mayor
func mayorNumero(n1 int, n2 int) int {
    // Si n1 es mayor, se regresa n1
    if n1 > n2 {
        return n1
    }

    // Si no, se regresa n2
    return n2
}
```
> Imprimira el **numero 20**
