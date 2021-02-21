# go面试

## slice

### 切片追加

```go
package main

import "fmt"

func main() {
    s := make([]int, 10)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}
```

 运行结果:

![image-20210221233920582](http://cdn.cjpa.top/cdnimages/image-20210221233920582.png)

解释：

在初始化的时候已经添加了10个

## for range

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4}
    m := make(map[int]*int, 0)
    for k, v := range slice {
      fmt.Println(k, " ", v, " ", &v)
       m[k] = &v
    }
    for k, v := range m {
      fmt.Println(k, " ", v, " ", *v)
    }
}
```

运行结果:

![image-20210221233941227](http://cdn.cjpa.top/cdnimages/image-20210221233941227.png)

解释：

for range创建一个新的副本，并不是每次都都返回一个一模一样的地址

