# Golang Interface

## Example interface: sort.Interface
    type Interface interface {
        Len() int
        Less(i, j int) bool
        Swap(i, j int)
    }
    func Sort(data Interface)

    //If there's a type myType that implement Len(), Less and Swap method, like this:
    func (a myType) Len() int {}
    func (a myType) Less(i, j int) bool {}
    func (a myType) Swap(i, j int) {}

    //Then we myType can be considered as "Interface" type, then we can use Sort func on it
    data := []myType{...}
    sort.Sort(data)

## Empty interface: 
    interface{}
    //All type implemented interface because it has zero method.

## Type assertion
    var myStr interface{} = "hello"
    v := myStr.(string)
    fmt.Println(v) // "hello"

    v2 := myStr.(bool)
    fmt.Println(v2) // panic

    // Generics:
    v := i.(T) // if i's type = T => v = value of i, otherwise panic
    // T == 'type' if used in switch

## Type switch
    switch v := i.(type) {
        case int:
            fmt.Println("Int: ",v)
        //...
    }

## map[string]interface{} or map[string]any:
    //Used to store data in its original state and data type. Example: JSON
