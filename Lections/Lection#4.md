# Bucles en Go(Primera Parte)
Los bucles son estructuras de control que permiten repetir un bloque de código varias veces. En Go, el bucle más común es el bucle `for`.

## Bucle `for`
Este bucle se compone de tres partes:
1. Siempre que se ponga palabra clave `for`, se inicializa una **variable contadora**
2. Se evalúa una **condición**
3. Se incrementa o decrementa la **variable contadora** dependiendo del uso que le demos al bucle.

> Tienes que tomar en cuenta que el bucle `for` en la parte de evaluar la condicion, se va a activar el bucle solo si se esta cumpliendo la condicion.

### Ejemplo de un bucle `for`:

#### Imprimir numeros del 1 al 10:
La estructura del for es la siguiente:

1. Despues de la palabra clave `for`, se inicializa la **variable contadora** `i:= 1`
2. Se evalua la condicion `i <= 10` o en representacion `1<=10`> para que el **bucle se active**
3. Indicamos que por cada vuelta que realize el for se va a **aumentar el valor** de `i` en 1, esto se representa con `i++`

```go
for i := 1; i <= 10; i++ {
	fmt.Println(i)
}
```
> Va a imprimir los numeros del 1 al 10, se va a detener hasta que `i` deje de ser **menor o igual a 10**

#### Imprimir numeros del 10 al 1:
En este caso, la estructura del for es la siguiente:
1. Despues de la palabra clave `for`, se inicializa la **variable contadora** `i:= 10`
2. Se evalua la condicion `i >= 1` o en representacion `10>=1`> para que el **bucle se active**
3. Indicamos que por cada vuelta que realize el for se va a **disminuir el valor** de `i` en 1, esto se representa con `i--`
```go
for i := 10; i >= 1; i-- {
    fmt.Println(i)
}
```
> Va a imprimir los numeros del 10 al 1, se va a detener hasta que `i` deje de ser **mayor o igual a 1**

#### Tabla de Multiplicacion entre dos numeros
En este caso, la estructura del for es la siguiente:
1. Despues de la palabra clave `for`, se inicializa la **variable contadora** `x:= 1`
2. Se evalua la condicion `x <= 10` o en representacion `1<=10`> para que el **bucle se active**
3.  Por cada vuelta que realize el for se va a **aumentar el valor** de `x` en 1, esto se representa con `x++`
```go
	var numero int
	fmt.Println("Ingrese un numero")
	fmt.Scanln(&numero)

	for x := 1; x <= 10; x++ {
		resultado := numero * x
		fmt.Printf("%d X %d  = %d \n", numero, x, resultado)
	}
```
Aqui podemos observar que dentro del `for`:

1. Se inicializa una variable  `resultado` que va a guardar el resultado de la multiplicacion entre `numero` y `x`
2. Se imprime el resultado de la multiplicacion con `fmt.Printf` y se muestra el numero ingresado, el valor de `x` y el resultado de la multiplicacion.
3. El bucle se va a repetir hasta que `x` deje de ser menor o **igual a 10**

## Nota del instructor

> La segunda parte de los bucles se va a posponer hasta que veamos los temas de **arrays** y **slices** para que se pueda entender mejor el uso de los bucles.