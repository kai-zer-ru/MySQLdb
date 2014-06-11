# О библиотеке

Библиотека дл работы с [MySQL](http://www.mysql.com/) на языке [Go](http://golang.org/).

Для работы требуется [go-sql-driver](https://github.com/go-sql-driver/mysql)

Сейчас все ответы приходят в понятном и простом формате. Ответы от `GetOne` и  `GetArray` приходят в виде map[string]interface{}

Ответ от `GetOneField` приходит в виде interface{}, содержащего значение выбранного поля.

Доработка ещё в процессе, написаы основные функции.
## Автор

Kaizer666 - [http://vk.com/](http://vk.com/id_00000000000000000000000000)

## Установка

    go get github.com/kaizer666/MySQLdb
    
## Использование

<pre>

package main

import (
      "github.com/kaizer666/MySQLdb"
      "fmt"
      )

func main() {
    MyDB := MySQLdb.MySqlDB{
        Address:"localhost:3306",
        DbName:"MyDB", 
        User:"User", 
        Password:"PassWord",
        }
    MyDB.Connect()
    defer MyDB.Close()
    row,err := MyDB.GetOne("SELECT * FROM table1 WHERE name='kaizer666'")
    if err != nil {
        panic(err)
    }
    fmt.Println(row)
    // row = map[string]interface{}{
    // id: 1,name: kaizer666
    // }
    
    
    row,err := MyDB.GetArray("SELECT * FROM table1 WHERE name in ('kaizer666','kaizer',github')")
    if err != nil {
        panic(err)
    }
    fmt.Println(row)
    // row = []map[string]interface{}{
    // [id: 1,name: kaizer666]
    // [id: 2,name: kaizer]
    // [id: 3,name: github]
    // }
        
    
    row,err := MyDB.GetOneField("SELECT id FROM table1 WHERE name='kaizer666'","id)
    if err != nil {
        panic(err)
    }
    fmt.Println(row)
    // row = 1
}

</pre>



