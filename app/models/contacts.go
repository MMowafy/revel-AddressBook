package models

import (
	"fmt"
	"log"
	"github.com/gocql/gocql"
	"strconv"
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
			iter := session.Query(`select contact_id,contactname from addressbook ` ,).Iter()
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
			iter := session.Query(`select contact_id, contactname from addressbook`).Iter()
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
	fmt.Println("hello from addnumber ", contactname , phone)
	var email string
	var nationality string
	var address string
	var err error
	err =session.Query(`select email,nationality,address from phones_by_user where contactname=? limit 1`,contactname).Scan(&email,&nationality,&address)
	if err!=nil {
		return err
	}
	phoneIDuuid:=gocql.TimeUUID()
	batch:=gocql.NewBatch(gocql.LoggedBatch)

	batch.Query(`insert into phonenumbers (contactname,phone_id,number) values (?,?,?)`,contactname,phoneIDuuid,phone)

	batch.Query("insert into phones_by_user (contactname,phone_id,email,nationality,address,number) values (?,?,?,?,?,?) ",
		contactname,phoneIDuuid,email,nationality,address,phone)
	fmt.Println(batch)
	err=session.ExecuteBatch(batch)
	return err
}
func ViewDetails(contactname string)  ([]AddressBookContact,error){
	fmt.Println("in view details " , contactname)
	var contacts []AddressBookContact
	var contact AddressBookContact
	counter:=0
	iter := session.Query("select phone_id,contactname,number,email,nationality,address from phones_by_user where contactname=?",contactname).Iter()
	for iter.Scan(contact.PartitionNumber,&contact.ContactName,contact.Phone,&contact.Email,&contact.Nationality,&contact.Address) {
		counter++
		fmt.Println("counter" , strconv.Itoa(counter))
		fmt.Println("cassandra contact :", contact.PartitionNumber,&contact.ContactName)
		contacts =append(contacts,contact)
	}
	err := iter.Close()
	return contacts,err
}
func Delete(contactname string, number string,partitionnumber string) error  {
	fmt.Println("this is delete method ",number)
	batch:=gocql.NewBatch(gocql.LoggedBatch)
	var err error
	if number=="1" {
		batch.Query("delete from addressbook where contact_id=? and contactname=?", partitionnumber,contactname)

		batch.Query("delete from phonenumbers where contactname=?",contactname)

		batch.Query("delete from phones_by_user where  contactname=?",contactname)
		fmt.Println(batch)
		err=session.ExecuteBatch(batch)

	} else if number=="2" {
		fmt.Println("here is 2")
		batch.Query("delete from phonenumbers where contactname=? and phone_id=?",contactname)
		batch.Query("delete from phones_by_user where  contactname=? and phone_id=?",contactname)
		fmt.Println(batch)
		err=session.ExecuteBatch(batch)
	}
	return err
}
func AddContact(newContact AddressBookContact) (AddressBookContact,error) {
	contactIDuuid:=gocql.TimeUUID()
	phoneIDuuid:=gocql.TimeUUID()
	var addContactToview AddressBookContact
	batch:=gocql.NewBatch(gocql.LoggedBatch)
	batch.Query("insert into addressbook (contact_id,contactname,email,nationality,address) values (?,?,?,?,?) ",
		contactIDuuid, newContact.ContactName, newContact.Email, newContact.Nationality, newContact.Address)

	batch.Query("insert into phonenumbers (contactname,phone_id,number) values (?,?,?)",
		newContact.ContactName,phoneIDuuid,newContact.Phone)

	batch.Query("insert into phones_by_user (contactname,phone_id,email,nationality,address,number) values (?,?,?,?,?,?) ",
		 newContact.ContactName,phoneIDuuid, newContact.Email, newContact.Nationality, newContact.Address,newContact.Phone)
	fmt.Println(batch)
	err:=session.ExecuteBatch(batch)
	if err!=nil {
		fmt.Println(err)
		return addContactToview,err
	}
	/*
	err:=session.Query("insert into addressbook (partitionNumber,contactname,email,nationality,address) values (?,?,?,?,?) ",contactIDuuid, newContact.ContactName, newContact.Email, newContact.Nationality, newContact.Address).Exec()
	if err !=nil {
		return addContactToview,err
	}
	err=session.Query("insert into phonenumbers (contactname,phone_id,number) values (?,?,?)",
		newContact.ContactName,gocql.TimeUUID(),newContact.Phone).Exec()
	if err !=nil {
		return addContactToview,err
	}
	*/
	addContactToview = AddressBookContact{
		PartitionNumber:  contactIDuuid,
		ContactName: newContact.ContactName,
		Email:       newContact.Email,
		Nationality: newContact.Nationality,
		Phone:       newContact.Phone,
		Address:     newContact.Address,
	}
	return addContactToview,err
}
