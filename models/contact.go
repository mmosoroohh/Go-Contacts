package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	 u "github.com/mmosoroohh/go-contact/utils"
)

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"user_id"` //The user that this contact belongs to
}

/*
This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
 */

func (contact *Contact) Validate() (map[string] interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Please provide Contact name"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Please provide Phone number"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User provided is not recognised"), false
	}

	//All the required parameters are present
	return u.Message(true, "Success"), true
}

func (contact *Contact) Create() (map[string] interface{}) {

	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "Success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) (*Contact) {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) ([]*Contact) {

	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}