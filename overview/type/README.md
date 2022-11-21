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
    fmt.Printf("%T",a)
    // uint8
    fmt.Printf("%c",a)
    // a
    fmt.Println(a)
    // 97
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
    [4]string{"Zero", "One", "Two", "Three"} 
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


