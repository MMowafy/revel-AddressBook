package models

import (
	"fmt"
	"log"
	"github.com/gocql/gocql"
)

type AddressBookContact struct {
	PartitionNumber gocql.UUID
	ContactName string
	Email       string
	Nationality string
	Phone       string
	Address     string
}

var session, err = getDB()

func GetContacts(sortby string) []AddressBookContact{


	var contacts []AddressBookContact
	var contact  AddressBookContact

	if err==nil {
		if sortby =="contactname" || sortby =="phonenumber" {

			/*
			rows, err = session.Query("select distinct pk, contactname from addressbook inner join phonenumbers on pk=fk order by " + sortby)

			if err!=nil {
				fmt.Println(err.Error())
				*/
			iter := session.Query(`select partitioNumber,contactname from addressbook order by co` ,).Iter()
			for iter.Scan(&contact.PartitionNumber,&contact.ContactName) {
				fmt.Println("cassandra contact :", contact.PartitionNumber,&contact.ContactName)
				contacts =append(contacts,contact)
			}
			if err := iter.Close(); err != nil {
				log.Fatal(err)
			}
		} else {
			/*
				rows, err  = session.Query("select distinct pk, contactname from addressbook inner join phonenumbers on pk=fk")
				if err!=nil {
					fmt.Println(err.Error())
				}
			*/
			iter := session.Query(`select partitionNumber, contactname from addressbook`).Iter()
			for iter.Scan(&contact.PartitionNumber,&contact.ContactName) {
				fmt.Println("cassandra contact :", contact.PartitionNumber,&contact.ContactName)
				contacts =append(contacts,contact)
			}
			if err := iter.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("list of contacts to app = " , contacts)
	return  contacts


}

func AddNumber(contactname string,phone string) error  {
	result:=session.Query("insert into phonenumbers (contactname,phone_id,number) values (?,?,?)",contactname,gocql.TimeUUID(),phone)
	fmt.Println(result)
	var err error
	return err
}
func ViewDetails(contactname string, partitionnumber string)  ([]AddressBookContact,error){
	fmt.Println("in view details " , contactname, partitionnumber)
	var contacts []AddressBookContact
	var contact AddressBookContact
	iter := session.Query("select contactname,email,nationality,address from addressbook where partitionNumeber=? and contactname=?",partitionnumber,contactname).Iter()
	for iter.Scan(contact.PartitionNumber,&contact.ContactName) {
		fmt.Println("cassandra contact :", contact.PartitionNumber,&contact.ContactName)
		contacts =append(contacts,contact)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return contacts,err
}
func Delete(contactname string, number string,partitionnumber string) error  {
	fmt.Println("this is delete method ",number)
	var err error
	if number=="1" {
		session.Query("delete from addressbook where partitionNumber=? and contactname=?", partitionnumber,contactname)
	} else if number=="2" {
		fmt.Println("here is 2")
		 session.Query("delete from phonenumbers where contactname= ?", contactname)
	}
	return err
}
func AddContact(newContact AddressBookContact) (AddressBookContact,error) {
	uuid:=gocql.TimeUUID()
	var addContactToview AddressBookContact
	session.Query("insert into addressbook (partitionNumber,contactname,email,nationality,address) values (?,?,?,?,?) ",uuid, newContact.ContactName, newContact.Email, newContact.Nationality, newContact.Address)

	session.Query("insert into phonenumbers (contactname,phone_id,phonenumber) values (?,?,?)",
		newContact.ContactName,gocql.TimeUUID(),newContact.Phone)
	addContactToview = AddressBookContact{
		PartitionNumber:  uuid,
		ContactName: newContact.ContactName,
		Email:       newContact.Email,
		Nationality: newContact.Nationality,
		Phone:       newContact.Phone,
		Address:     newContact.Address,
	}
	return addContactToview,err
}
