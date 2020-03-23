package lang

import "fmt"

func DemoMap() {

    /*
    map 变量声明格式：
    var variable_name map[key_type]value_type

    使用 make 函数 
    variable_name := make(map[key_type]value_type)
    */
    var m map[string]int        // 不初始化默认是 nil 
    m = make(map[string]int)    // 使用 make 函数定义
    m = map[string]int{"Go": 1, "PHP": 2}   // 也可以用数据直接初始化map
    fmt.Println(m)                          // map[Go:1 PHP:2]

    m["Java"] = 3
    fmt.Println(m)                          // map[Go:1 PHP:2 Java:3]
    m["C++"]++                              // 可以对一个原本不存在的key执行++
    fmt.Println(m)                          // map[Go:1 PHP:2 Java:3 C++:1]
    fmt.Println(m["C"])                     // 输出一个不存在的值，这里输出int的默认值：0，但是没有此key，此时map中依然没有键C

    // 判断键是否存在
    _, ok := m["C"]
    fmt.Println(ok) // false
    _, ok = m["Java"]
    fmt.Println(ok) // true

    // delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key
    delete(m, "Java")
    fmt.Println(m)                          // map[Go:1 PHP:2 C++:1]

    // map 的key需要支持相等判断
    // key 不能是 切片、函数、map等复杂类型
    fmt.Println(make(map[bool]int))
    fmt.Println(make(map[float32]int))
    fmt.Println(make(map[int]int))
    fmt.Println(make(map[[3]int]int))
    // fmt.Println(make(map[[]int]int))             // invalid map key type []int
    // fmt.Println(make(map[func(int, int)]int))    // invalid map key type func(int, int)
    // fmt.Println(make(map[map[string]int]int))

    // 值可以是 切片、函数、map等复杂类型
    fmt.Println(make(map [string] []int))
    fmt.Println(make(map [string] [][]int))
    fmt.Println(make(map [string] func(int, int) int))
    fmt.Println(make(map [string] map[string]int))


    var intvalue int = 30
    var int8value int8 = 127    // 数值大小超出类型所表示的范围，编译不通过，此处不能是128
    var fvalue float32 = 3.14
    var fvalue2 float64 = 3.14
    
    intvalue = int(int8value)
    fvalue = float32(intvalue)
    fvalue2 = float64(int8value)
    int8value = int8(intvalue)
    intvalue = int(fvalue)
    intvalue = int(fvalue2)
    // var bvalue bool = false
    // intvalue = int(bvalue)       // cannot convert bvalue (type bool) to type int
    // var i int = int(23.23233)    // constant 23.2323 truncated to integer
    var i23 float32 = 23.23233
    var i int = int(i23)            // OK
    // var s23 string = "23"
    // i = int(s23)                 // cannot convert s23 (type string) to type int
    fmt.Println(intvalue, int8value, fvalue, fvalue2, i)


    // Go 内置的 map 不是并发安全的，并发安全的 map 可以 使用标准包 sync 中的 map

    // map value类型如果是结构体，不能直接修改一个元素中的某个成员变量，必须通过整个元素赋值来修改
    book := Books{"Go", "somebody", "", 0}
    books := map[string]Books{"Go": book}
    // books["Go"].bookID = 3   // cannot assign to struct field books["Go"].bookID in map
    // (books["Go"]).bookID = 3 // 加上（）也不行，同样报上面的错误
    tmp := books["Go"]
    tmp.bookID = 3
    books["Go"] = tmp   // 整个元素赋值


}