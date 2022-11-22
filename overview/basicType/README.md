# Golang types

## Error Type
Custom error: 

- errors.News()
- fmt.Errorf()
## Numeric Type

## Byte
var a1 byte = 97 
fmt.Println(a1) 
// 97 

fmt.Printf("%c\n", a1)
// a 

## String, Character, Rune

String = Collection (a.k.a Slice) of bytes. 
1 byte to store 1 ASCII character. 
N bytes to store 1 Unicode character. 

### String

### Byte (uint8)
- A byte is an uint8 value that is used for representing a single ASCII character. 
- []byte{} is used for I/O performing. 
#### Examples: 
    var a byte = 97
    fmt.Printf("%T",a)     // uint8
    fmt.Printf("%c",a)    // a
    fmt.Println(a)    // 97

    b := make([]byte, 12)
	b = []byte("Byte slice €")
	fmt.Println("Byte slice as text: \n", string(b)) // byte slice to string

### Rune
- A rune is an int32 value that is used for representing a single Unicode code point 
#### Examples:
    r := '€'
    fmt.Printf("%T",r)
    // int32 
    fmt.Println(r)
    // 8364
    fmt.Printf("%c",r)
    // €

**Note:**
    r := `€`
    fmt.Printf("%T",r) 
    // string 

## Array
Examples:
    var a = [2]string{"Zero", "One"} 
    var S0 = a[0:1] // S0 = slice
	S0[0] = "S0" // Change to slice will affect array
    fmt.Println(a)
    //["S0","One","Two","Three"]

    //As the capacity of S0 changes, it is no longer connected to the same underlying array
    S0 = append(S0, "N1")
    S0 = append(S0, "N2")
    S0[0] = "S0_2"
    fmt.Println(a)
    //["S0","One","Two","Three"]

## Slice
- Any changes you make to a slice inside a function also affect the original slice 
### Create slice
    []string{"test"}
    make([]string, 4)
### Add new element
    aSlice := []string{}
    aSlice = append(aSlice, "test")
    // Use ... to append a slice
    aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
### Len, capacity
    len(aSlice)
    cap(aSlice)
### Iterate 
    for i in len(aSlice){
        //...
    }


