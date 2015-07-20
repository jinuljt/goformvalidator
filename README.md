## go form validator ##

go form validator 其实只有一个函数，混合使用schema与validator来做golang的表单验证

[schema](https://github.com/gorilla/schema) 将 net/url.Values 转换为struct，

[go-validator](https://github.com/go-validator/validator) 校验数据

schema与go-validator的用法见他们各自的文档

简单用法：
```
type Person struct {
    Name string `validate:"min=5,max=10" schema:name"` //姓名长度 5-10个字节
    Age int `validate:"min=1,max=150" schema:age"` //年龄1-150岁
}

person := new(Person)
httpRequest.ParseForm()
if err := Validate(person, httpRequest.Form); err != nil {
    fmt.Printf("%v validate error due to %s", *person, err)
} else {
    fmt.Printf("%v validate success", *person)
}

```
