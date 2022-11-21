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

### Rune
- A rune is an int32 value that is used for representing a single Unicode code point 
- Examples: 

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
- []string{"test"}
- make([]string, 4)
### Add new element
aSlice := []string{}
aSlice = append(aSlice, "test")
### Len, capacity
len(aSlice)
cap(aSlice)
### Iterate 
    for i in len(aSlice){
        //...
    }

