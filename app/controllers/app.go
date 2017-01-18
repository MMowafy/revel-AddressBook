package controllers

import (
	"github.com/revel/revel"
	"github.com/addressbook/app/models"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) AboutUs() revel.Result {
	return c.Render()
}
func (c App) ContactUs() revel.Result {
	return c.Render()
}

func (c App) Index() revel.Result {
	sortby:="pk"
	var Contacts  []models.AddressBookContact
	Contacts=models.GetContacts(sortby)
	return c.Render(Contacts)
}
func (c App) SortContacts() revel.Result {
	var sortby string
	c.Params.Bind(&sortby,"sortBy")
	fmt.Println("in sort contacts ====",sortby)
	var Contacts  []models.AddressBookContact
	Contacts=models.GetContacts(sortby)
	return c.RenderJson(Contacts)
}
func (c App) AddNumberToContact() revel.Result {
	var newnumber string
	var pk string
	c.Params.Bind(&newnumber,"newnumber")
	c.Params.Bind(&pk,"contactID")
	ok:=models.AddNumber(pk,newnumber)
	if ok!=nil {
		fmt.Println(ok.Error())
		c.RenderError(ok)
	}
	return c.Result
}
func (c App) ViewContactDetails() revel.Result  {
	var contactname string
	var partitionnumber string
	var ok error
	var ContactDetails  []models.AddressBookContact
	c.Params.Bind(&contactname,"contactname")
	c.Params.Bind(&partitionnumber,"partitionnumber")
	ContactDetails,ok=models.ViewDetails(contactname , partitionnumber )
	if ok!=nil {
		fmt.Println(ok.Error())
		c.RenderError(ok)
	}
	return c.RenderJson(ContactDetails)
}
func (c App) DeleteContact() revel.Result {
	var contactname string
	var partitionnumber string
	var number string
	var ok error
	c.Params.Bind(&contactname,"contactname")
	c.Params.Bind(&number,"number")
	c.Params.Bind(&partitionnumber,"partitionnumber")
	ok=models.Delete(contactname,number,partitionnumber)
	if ok!=nil {
		fmt.Println(ok.Error())
		c.RenderError(ok)
	}
	return c.Result
}
func (c App) AddNewContact() revel.Result  {
	var contact  models.AddressBookContact
	c.Params.Bind(&contact,"addressbookcontact")
	result,err:=models.AddContact(contact)
	if err!=nil{
		fmt.Println(err.Error())
		c.RenderError(err)
	}
	return c.RenderJson(result)
}

