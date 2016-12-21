package models

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

type AddressBookContact struct {
	PK          int
	ContactName string
	Email       string
	Nationality string
	Phone       string
	Address     string
}

var db, err = getDB()
func GetContacts(sortby string) []AddressBookContact{

	var rows *sql.Rows
	var contacts []AddressBookContact

	if err==nil {
		if sortby =="contactname" || sortby =="phonenumber" {

			rows, err = db.Query("select distinct pk, contactname from addressbook inner join phonenumbers on pk=fk order by " + sortby)
			if err!=nil {
				fmt.Println(err.Error())
			}
		} else {
			rows, err  = db.Query("select distinct pk, contactname from addressbook inner join phonenumbers on pk=fk")
			if err!=nil {
				fmt.Println(err.Error())
			}
		}
		for rows.Next() {
			var contact AddressBookContact
			rows.Scan(&contact.PK,&contact.ContactName, )
			contacts = append(contacts, contact)
		}
		rows.Close()
	}
	return  contacts
}
func AddNumber(pk string,phone string) error  {
	_,err:=db.Exec("insert into phonenumbers (phone_id,fk,phonenumber) values (?,?,?)",nil,pk,phone)
	return err
}
func ViewDetails(pk string)  ([]AddressBookContact,error){
	var contacts []AddressBookContact
	rows, err  := db.Query("select phone_id, contactname,email,nationality,address,phonenumber from addressbook inner join phonenumbers on pk=fk where pk="+pk)
	if err != nil {
		return contacts,err
	}
	for rows.Next() {
		var contact AddressBookContact
		rows.Scan(&contact.PK,&contact.ContactName,&contact.Email,&contact.Nationality,&contact.Address, &contact.Phone)
		contacts = append(contacts, contact)
	}
	rows.Close()
	return contacts,err
}
func Delete(pk string, number string) error  {
	fmt.Println("this is delete method ",number)
	var err error
	if number=="1" {
		if _, err = db.Exec("delete from addressbook where pk= ?", pk); err != nil {
			return err
		}
	} else if number=="2" {
		fmt.Println("here is 2")
		var count int
		if _, err = db.Exec("delete from phonenumbers where phone_id= ?", pk); err != nil {
			return err
		}
		rows, err  := db.Query("select count(*) from addressbook inner join phonenumbers on pk=fk where pk="+pk)
		if err != nil {
			return err
		}
		for rows.Next() {
			rows.Scan(&count)
		}
		fmt.Println("number of records ==",count)


		//w.WriteHeader(http.StatusOK)
	}
	return err
}
func AddContact(newContact AddressBookContact) (AddressBookContact,error) {
	var addContactToview AddressBookContact
	resultOfInsertion, err := db.Exec("insert into addressbook (pk,contactname,email,nationality,address) values (?,?,?,?,?) ",
		nil, newContact.ContactName, newContact.Email, newContact.Nationality, newContact.Address)
	if err != nil {
		return addContactToview,err
	}
	primarykey, _ := resultOfInsertion.LastInsertId()
	_,err=db.Exec("insert into phonenumbers (phone_id,fk,phonenumber) values (?,?,?)",
		nil,primarykey,newContact.Phone)
	if err != nil {
		return addContactToview,err
	}

	addContactToview = AddressBookContact{
		PK:          int(primarykey),
		ContactName: newContact.ContactName,
		Email:       newContact.Email,
		Nationality: newContact.Nationality,
		Phone:       newContact.Phone,
		Address:     newContact.Address,
	}
	return addContactToview,err
}