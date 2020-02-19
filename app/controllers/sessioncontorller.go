package controllers

import (
        "github.com/astaxie/beego"
)

// Controllerを登録
type SessionController struct {
    beego.Controller
}

func (this *SessionController) Home() {
    isAuthenticated := this.GetSession("authenticated")
    if isAuthenticated == nil || isAuthenticated == false {
        this.Ctx.WriteString("You need to get authentication to access the page")
        return
    }
    this.Ctx.ResponseWriter.WriteHeader(200)
    this.Ctx.WriteString("Home Page")
}

func (this *SessionController) Login() {
    // Login関数が呼び出されたらauthenticatedをTrueに設定する
    this.SetSession("authenticated", true)
    this.Ctx.ResponseWriter.WriteHeader(200)
    this.Ctx.WriteString("SUCCESS LOGIN")
}

func (this *SessionController) Logout() {
    // Logout関数が呼び出されたらauthenticatedをFalseに設定する
    this.SetSession("authenticated", false)
    this.Ctx.ResponseWriter.WriteHeader(200)
    this.Ctx.WriteString("SUCCESS LOGOUT")
}
