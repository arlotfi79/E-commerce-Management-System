package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var states = []string{"AK", "AL", "AR", "AZ", "CA", "CO", "CT", "DC", "DE", "FL", "GA",
	"HI", "IA", "ID", "IL", "IN", "KS", "KY", "LA", "MA", "MD", "ME",
	"MI", "MN", "MO", "MS", "MT", "NC", "ND", "NE", "NH", "NJ", "NM",
	"NV", "NY", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX",
	"UT", "VA", "VT", "WA", "WI", "WV", "WY"}

type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	state       string
	business    bool
}

var contacts = make([]Contact, 0)

// Tview
var pages = tview.NewPages()
var contactText = tview.NewTextView()
var app = tview.NewApplication()
var form = tview.NewForm()
var contactsList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) to add a new contact \n(q) to quit")

func main() {
	contactsList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		setConcatText(&contacts[index])
	})

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(contactsList, 0, 1, true).
			AddItem(contactText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			form.Clear(true)
			addContactForm()
			pages.SwitchToPage("Add Contact")
		}
		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Add Contact", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func addContactList() {
	contactsList.Clear()
	for index, contact := range contacts {
		contactsList.AddItem(contact.firstName+" "+contact.lastName, " ", rune(49+index), nil)
	}
}

func addContactForm() *tview.Form {

	contact := Contact{}

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		contact.firstName = firstName
	})

	form.AddInputField("Last Name", "", 20, nil, func(lastName string) {
		contact.lastName = lastName
	})

	form.AddInputField("Email", "", 20, nil, func(email string) {
		contact.email = email
	})

	form.AddInputField("Phone", "", 20, nil, func(phone string) {
		contact.phoneNumber = phone
	})

	form.AddDropDown("State", states, 0, func(state string, index int) {
		contact.state = state
	})

	form.AddCheckbox("Business", false, func(business bool) {
		contact.business = business
	})

	form.AddButton("Save", func() {
		contacts = append(contacts, contact)
		addContactList()
		pages.SwitchToPage("Menu")
	})

	return form
}

func setConcatText(contact *Contact) {
	contactText.Clear()
	text := contact.firstName + " " + contact.lastName + "\n" + contact.email + "\n" + contact.phoneNumber
	contactText.SetText(text)
}
