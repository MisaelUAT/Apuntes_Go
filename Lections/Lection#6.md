# Bucles en Go(Segunda Parte)

En la primera parte de los bucles, vimos como funcionan los bucles `for` y como se pueden usar para repetir un bloque de código varias veces. En esta segunda parte, vamos a ver como podemos usar los bucles `for` para recorrer **arrays y slices.**

## Para que nos sirve el recorrido de arrays y slices con bucles `for`
El recorrido de arrays y slices con bucles `for` nos sirve para poder acceder a cada elemento de un array o slice, y poder realizar alguna operacion con ese elemento, como por ejemplo imprimirlo, modificarlo, etc.

## Recorrido de arrays con bucles `for` 
Para recorrer un array con un bucle `for`, podemos usar la siguiente estructura:

```go
for i := 0; i <= len(nombre_del_array) - 1; i++ {
    // Aqui va el bloque de codigo que queremos repetir para cada elemento del array
}
```

Aqui podemos observar que:
- `i := 0` es la inicializacion de la variable contadora, que va empezar desde 0, ya que el indice de un array empieza desde 0
- `i <= len(nombre_del_array) - 1` es la condicion para que el bucle se active, en este caso el bucle se va a activar mientras `i` sea **menor o igual a la longitud del array menos 1**, 

### Ejemplos de recorridos de arrays con bucle `for`

#### Ejercicio 
Pidele al usuario que ingrese 5 numeros de las cuales vas a sacar:

1. Cual es el primer y ultimo valor de la secuencia?
2. Cual es el numero mayor y menor de la secuencia?
3. La suma total de los elementos de la secuencia
4. El promedio de la secuencia

**La solucion es la siguiente:**
```go
var numerosUsuario[5] int
var numero_sumaTotal int = 0
var numero_Mayor int
var numero_Menor int

for x:= 0; x <=len(numerosUsuario) - 1; x++{
	fmt.Printf("Ingre el Valor %d",x)
	fmt.Scanln(&numerosUsuario[x])

	numero_sumaTotal = numero_sumaTotal + numerosUsuario[x]

	if x == 0{
	numero_Mayor= numerosUsuario[0]
    numero_Menor = numerosUsuario[0]
	}

	if numerosUsuario[x] > numero_Mayor{
		numero_Mayor = numerosUsuario[x]
	} 
	
	if numerosUsuario[x] < numero_Menor{
		numero_Menor = numerosUsuario[x]
	}
}

valor_final:= len(numerosUsuario) - 1

fmt.Printf("El Primer Valor es %d \n el Ultimo vaolor es %d \n el numero total de elementos es de %d",numerosUsuario[0],numerosUsuario[valor_final],len(numerosUsuario))
fmt.Printf(" \n La suma de los elementos es de %d \n El promedio es de %d", numero_sumaTotal, numero_sumaTotal/len(numerosUsuario))
fmt.Printf("El numero mayor es de %d  El numero menor es de %d",numero_Mayor,numero_Menor)
```

Desglosando este codigo, podemos observar que:

**1. Se declara las variables que se van a usar entre ellas:**

* `numero_usuario`: Array a usar
* `numero_sumaTotal`: Variable a almacenar la suma de los elementos
* `numero_Mayor y numero_Menor`: Variables que almacenaran el numero mayor y menor de la secuencia

**2. Armamos nuestro for para que recorra nuestro array**
- for x:= 0; **Inicializamos en la Posicion 0**
- x <=len(numerosUsuario) - 1; **Indicamos que si la variable `x` sigue siendo menor o igual a 4**

**3. Pedimos que el usuario ingrese los numeros y los guarde en en nuestro array con:** `fmt.Scanln(&numerosUsuario[x])` donde `x` es la posicion que se encuentra el for

**4. Se pone un condicional para que cuando `x` sea igual a 0 inicialize las variables y les asigne el valor que se encuentra en la posicion 0**
```go
if x == 0 {
	numero_Mayor= numerosUsuario[0]
    numero_Menor = numerosUsuario[0]
	}
```

**5. Se ponen los condicionales para evaluar si el valor  de la posicion`x` del array es mayor o menor que el valor que se encuentra en las variables `numero_Mayor y numero_Menor`**
```go
if numerosUsuario[x] > numero_Mayor{
		numero_Mayor = numerosUsuario[x]
	} 
	
	if numerosUsuario[x] < numero_Menor{
		numero_Menor = numerosUsuario[x]
	}
```
> Si se cumple la condicion  dependiendo del `if` el valor que tiene en ese momento el `array en posicion de x` se guardara en la variable que se encuentra en el if, en este caso el `numero_Mayor y numero_Menor`

**6. Ya se imprimen los valores obtenidos**
```go
valor_final:= len(numerosUsuario) - 1

fmt.Printf("El Primer Valor es %d \n el Ultimo vaolor es %d \n el numero total de elementos es de %d",numerosUsuario[0],numerosUsuario[valor_final],len(numerosUsuario))
fmt.Printf(" \n La suma de los elementos es de %d \n El promedio es de %d", numero_sumaTotal, numero_sumaTotal/len(numerosUsuario))
fmt.Printf("El numero mayor es de %d  El numero menor es de %d",numero_Mayor,numero_Menor)
```

**Nota:** Si quieres usar variables y mostrarlas fuera del for, es necesario que si o si los declares antes de iniciar el bucle

## Recorrido de un slice 






