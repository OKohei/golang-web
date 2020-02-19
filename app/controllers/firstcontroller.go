package controllers

import "github.com/astaxie/beego"

// Controllerを登録
type FirstController struct {
    beego.Controller
}

// Employeeの構造体を作成する
type Student struct {
    ID        int    `json:"id"`
    FirstName string `json:"firstName"`
    LastName  string `json:lastName`
}

// Studentタイプを作成
type Students []Student

// Employeeのスライスを作成
var students []Student

func init() {
    // Studentデータを作成
    students = Students {
        Student{ID: 1, FirstName: "Harry", LastName: "Potter"},
        Student{ID: 2, FirstName: "Ron", LastName: "Weasley"},
        Student{ID: 3, FirstName: "Hermione", LastName: "Granger"},
    }
}

func (this *FirstController) GetStudents() {
    // 戻り値のヘッダーを200に設定
    this.Ctx.ResponseWriter.WriteHeader(200)
    // jsonデータで返却
    this.Data["json"] = students
    this.ServeJSON()
}

func (this*FirstController) Dashboard() {
    this.Data["students"] = students
    this.TplName = "dashboard.tpl"
}
