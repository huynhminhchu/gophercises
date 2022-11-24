# Golang Composite types

## Map
### Declare:
    aMap := map[string]int{} 
    aMap["test"] = 1 

    m := map[string]int {
        "key1":1,
        "key2":-2
    }

    anotherMap := make(map[string]string) 

    // Note:
    var aMap = map[string]int
    aMap["chu"] = "dep" // Error
    // We have to allocate this nil map before assigning, like this:
    aMap = make(map[string]int)


### Search: 
    //If ok = true then v exists 
    v, ok := aMap[k]  
### Iterate: 
    for i, v := range aMap { 
        fmt.Println(i,v) 
    } 
## Struct
### Declare:
    type Entry struct {
        Name string
        Surname string
        Tel string
    }
### Create new instances:
    pS := new(Entry)
    pS2 := Entry{Name: "chu", Surname: "minh", Tel: "01234567"}

