# About

Library to work with the [MySQL] (http://www.mysql.com/) in language [Go] (http://golang.org/).

It requires [go-sql-driver] (https://github.com/go-sql-driver/mysql)

Now all the answers come in a clear and simple format. Answers from `GetOne` and `GetArray` come in the form of map [string] interface {}

Reply from `GetOneField` comes in the form of interface {}, contains the value of the selected field.

Refinement still in the process, write basic functions.
## Author

Kaizer666 - [http://vk.com/] (http://vk.com/id_00000000000000000000000000)

## Install

    go get github.com/kaizer666/MySQLdb
    
## Use

<pre>

package main

import (
      "github.com/kaizer666/MySQLdb"
      "fmt"
      )

func main () {
    MyDB: = MySQLdb.MySqlDB {
        Address: "localhost: 3306"
        DbName: "MyDB",
        User: "User",
        Password: "PassWord",
        }
    MyDB.Connect ()
    defer MyDB.Close ()
    row, err: = MyDB.GetOne ("SELECT * FROM table1 WHERE name = 'kaizer666'")
    if err! = nil {
        panic (err)
    }
    fmt.Println (row)
    // Row = map [string] interface {} {
    // Id: 1, name: kaizer666
    //}
    
    
    row, err: = MyDB.GetArray ("SELECT * FROM table1 WHERE name in ('kaizer666', 'kaizer', github ')")
    if err! = nil {
        panic (err)
    }
    fmt.Println (row)
    // Row = [] map [string] interface {} {
    // [Id: 1, name: kaizer666]
    // [Id: 2, name: kaizer]
    // [Id: 3, name: github]
    //}
        
    
    row, err: = MyDB.GetOneField ("SELECT id FROM table1 WHERE name = 'kaizer666'", "id)
    if err! = nil {
        panic (err)
    }
    fmt.Println (row)
    // Row = 1
}

</Pre>