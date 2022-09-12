package main

import (
    "fmt"
    "github.com/beego/beego/v2/client/orm"
    _ "github.com/go-sql-driver/mysql"
)

// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}

func init() {
    orm.RegisterDataBase("default", "mysql", "root:admin@tcp(127.0.0.1:3306)/beego?charset=utf8&loc=Local")
    orm.RegisterModel(new(User))
    orm.RunSyncdb("default", false, true)
}

func main() {
    o := orm.NewOrm()
    user := User{Name: "slene"}

    // insert
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    // update
    user.Name = "astaxie"
    num, err := o.Update(&user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
    // read one
    u := User{Id: user.Id}
    err = o.Read(&u)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
