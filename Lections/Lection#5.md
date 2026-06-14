# Tipos de Datos Compuestos(Basicos)
En esta seccion vamos a estudiar lo que son los Tipos de datos Compuestos, que son aquellos que pueden **contener mas de un valor**, a diferencia de los tipos de datos basicos que solo pueden contener un valor. Los basicos que veremos aqui son los **arrays y los slices**

## Arrays
Un array es una secuencia de elementos del mismo tipo, con un tamaño fijo. Se declara de la siguiente manera:
```go
var nombre [tamaño]tipo
```
En donde:
- `nombre` es el nombre del array
- `tamaño` es el numero de elementos que el array puede contener
- `tipo` es el tipo de dato de los elementos del array(los que vimos anteriormente)

### Ejemplo de cuando se declararia un array
Supongamos que queremos almacenar las ventas de un dia, y sabemos que solo vamos a tener 5 ventas, entonces podemos declarar un array de la siguiente manera:
```go
var ventas_dia[5] int
```
> En este caso `ventas_dia` va a tener **5 elementos de tipo int(numeros enteros)**

### Como asignarle valores a un array
Para asignarle valores a un array, podemos hacerlo de la siguiente manera con el array que mostramos de ejemplo anteriormente:
```go
ventas_dia[0] = 100
ventas_dia[1] = 200
ventas_dia[2] = 300
ventas_dia[3] = 400
ventas_dia[4] = 500
```
Aqui podemos ver que:
- `ventas_dia` es el nombre del array a poner los valores
- `[numero de posicion]` es la posicion del elemento al que le queremos asignar el valor, y se empieza a contar desde 0
- `= valor` es el valor que queremos asignar a esa posicion del array

> Toma en cuenta que el numero de posicion de un array siempre se cuenta desde 0, es decir:

![image](/Assets/Array.png)

Aqui se puede observar que:
- **Elementos del Arreglo:** Son los valores que se encuentran dentro del array, en este caso son 5 elementos
- **Indice del Arreglo:** Es el numero que indica la posicion del elemento dentro del array, en este caso el indice va desde 0 hasta 4, ya que el array tiene 5 elementos.

### Como saber la longitud de un array
Para saber la longitud de un array, podemos usar la funcion `len()`, que nos devuelve el numero de elementos que tiene el array. Por ejemplo:
```go
fmt.Println(len(ventas_dia)) // Esto nos va a imprimir 5, ya que
```
> En este caso `len(ventas_dia)` nos devuelve 5, ya que el array `ventas_dia` tiene 5 elementos.

## Como saber que valores tiene un array
Para saber que valor tiene un array, podemos usar la funcion `fmt.Println()`, que nos permite imprimir el valor de un array. Por ejemplo:
```go
fmt.Println(ventas_dia) // Esto nos va a imprimir [100 200 300 400 500]
```
> Aqui solo sabemos todos los valores del array, pero desconocemos la posicion de cada valor. Aqui es donde vamos a trabajar **la segunda parte de los bucles**

## Slices
Un slice es una secuencia de elementos del mismo tipo, pero a diferencia de un array, un slice tiene un tamaño dinamico, es decir  que puede crecer o disminuir su tamaño dinamicamente. Se declara de la siguiente manera:
```go
var nombre []tipo
```
> Aqui podemos observar que se declara igual que un array con la diferencia de que no se especifica el tamaño. 

**Nota:** No explico la estructura del slice ya que es igual que la de array nada mas quitando el numero de elementos.

### Ejemplo de cuando se declararia un slice
Supongamos que queremos almacenar las calificaciones de los estudiantes de una clase, pero no sabemos cuantas calificaciones vamos a tener, entonces podemos declarar un slice de la siguiente manera:

```go
var listaCalificaciones []int
```
> En este caso `listaCalificaciones` va a tener un tamaño dinamico, es decir que puede crecer o disminuir su tamaño dinamicamente, y va a contener elementos de tipo int(numeros enteros)

### Como asignarle valores a un slice
Para asignarle valores a un slice, podemos usar la funcion `append()`, que nos permite agregar elementos a un slice. Por ejemplo:
```go
listaCalificaciones = append(listaCalificaciones,20)
listaCalificaciones = append(listaCalificaciones,30)
listaCalificaciones = append(listaCalificaciones,40)
```
Aqui podemos observar que:
- `listaCalificaciones` es el nombre del slice al que le queremos agregar los valores
- `append`: es la funcion que nos permite agregar elementos a un slice
- `(nombre del slice, valor)` es la sintaxis de la funcion append, donde el primer parametro es el nombre del slice al que le queremos agregar el valor, y el segundo parametro es el valor que queremos agregar al slice.

> Otra diferencia que podemos notar, es que no es necesario poner el numero de posicion para agregar un valor, **dinamicamente lo agrega**

### Como saber la longitud de un slice
Para saber la longitud de un slice, podemos usar la funcion `len()`, que nos devuelve el numero de elementos que tiene el slice. Por ejemplo:
```go
fmt.Println(len(listaCalificaciones)) // Esto nos va a imprimir 3, ya que el slice tiene 3 elementos
```
> En este caso `len(listaCalificaciones)` nos devuelve 3, ya que el slice `listaCalificaciones` tiene 3 elementos que agregamos anteriormente.

### Como saber que valores tiene un slice
Para saber que valor tiene un slice, podemos usar la funcion `fmt.Println()`, que nos permite imprimir el valor de un slice. Por ejemplo:
```go
fmt.Println(listaCalificaciones) // Esto nos va a imprimir [20 30 40]
```
> Aqui solo sabemos todos los valores del slice, pero desconocemos la posicion de cada valor. Aqui es donde vamos a trabajar **la segunda parte de los bucles**

## Modificar un valor de un array o slice
Para modificar un valor de un array o slice, podemos usar la sintaxis de asignacion, que es la misma que usamos para asignar valores a un array o slice.

### Modificar un valor de un array
```go
ventas_dia[0] = 100
ventas_dia[0] = 150 // Cambiara el valor de 100 a 150
```

### Modificar un valor de un slice
```go
listaCalificaciones = append(listaCalificaciones,20)
listaCalificaciones[0] = 25 // Cambiara el valor de 20 a 25
```

Ahora que ya hemos visto como funciona los arrays y los slices, es momento de que aprendamso sobre como podemos acceder a todos los valores de un array o slice con **bucles**. Esto lo veremos en la siguiente seccion.




