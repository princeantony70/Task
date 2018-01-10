package main

import (

        "net/http"
        "html/template"
        "database/sql"
    	_ "github.com/go-sql-driver/mysql"
    )


type Userdetail struct{
  Name string
  Age  int
  College string
  Qualification string
  Email string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "nfn"
    dbName := "goblog"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))



func New(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
//    selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
     selDB, err := db.Query("SELECT * FROM userdetail ORDER BY name DESC")
    if err != nil {
        panic(err.Error())
    }
    usr:= Userdetail{}
    res := []Userdetail{}
    for selDB.Next() {
    //    var id int
        var age int
        //var name, city string
        var name ,college,qualification,email string
        err = selDB.Scan(&name,&age,&college,&qualification,&email)
        if err != nil {
            panic(err.Error())
        }
        //emp.Id = id
        usr.Name =name
        //emp.Name = name
        //emp.City = city
        usr.Age = age
        usr.Email = email
        usr.College = college
        usr.Qualification = qualification

        res = append(res, usr)
        //res = append(res,usr)
    }
    tmpl.ExecuteTemplate(w, "New", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("name")

//    nId := r.URL.Query().Get("id")
  //  selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
      selDB, err := db.Query("SELECT * FROM userdetail WHERE name=?", nId)

    if err != nil {
        panic(err.Error())
    }
    usr := Userdetail{}
    for selDB.Next() {
  //      var id int
    //    var name, city string
                var age int
                var name ,college,qualification,email string
                err = selDB.Scan(&name,&age,&college,&qualification,&email)



        //err = selDB.Scan(&id, &name, &city)

        if err != nil {
            panic(err.Error())
        }
            usr.Name =name
            usr.Age = age
            usr.Email = email
            usr.College = college
            usr.Qualification = qualification

      //  emp.Id = id
        //emp.Name = name
        //emp.City = city
    }
    tmpl.ExecuteTemplate(w, "Show", usr)
    defer db.Close()
}

  func  main(){
  http.HandleFunc("/", New)
  http.HandleFunc("/show", Show)
  http.ListenAndServe(":8080", nil)


  }
