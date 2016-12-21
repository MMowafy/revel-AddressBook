package controllers

import (
	"github.com/revel/revel"
	"github.com/addressbook/app/models"
	"fmt"
)

type Login struct {
	*revel.Controller
}
func (lc Login) Login() revel.Result  {
	return lc.Render()
}
func (lc Login) TryToLogin() revel.Result  {
	var login models.UserLogin
	//login.Email=lc.Params.Get("email")
	//login.Password=lc.Params.Get("password")
	lc.Params.Bind(&login,"login")
	fmt.Println(login)
	ok:=models.VerifyUser(login)
	if ok {
		return lc.Redirect(App.Index)
	} else {
		return lc.RenderTemplate("login/login.html")
	}

}
