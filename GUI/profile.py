from tkinter import *
from tkinter import messagebox

import requests

from shoppingCart import showCartProducts
from orders import showOrderDetails
from login import account_token

global firstName
global lastName
global email
global phoneNumber

global addAddres_screen

def showNotifications():
    pass
    #TODO

def submitChanges():
    fname = firstName.get()
    lname = lastName.get()
    mail = email.get()
    mobile = phoneNumber.get()

    #TODO: edit info

def editInfo():
    edit_screen = Toplevel()
    edit_screen.title("Register")
    edit_screen.geometry("400x550")

    font = ("Calibri", 11)

    # Set text variables
    global firstName
    global lastName
    global email
    global phoneNumber

    firstName = StringVar()
    lastName = StringVar()
    email = StringVar()
    phoneNumber = StringVar()

    # Set label for user's instruction
    Label(edit_screen, text="Personal Info", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(edit_screen, text="").pack()

    # Set first name
    firstName_lable = Label(edit_screen, text="First Name * ", font=font)
    firstName_lable.pack()

    firstName_entry = Entry(edit_screen, textvariable=firstName, font=font, width=30)
    firstName_entry.pack()

    Label(edit_screen, text="").pack()
    # Set last name
    lastName_lable = Label(edit_screen, text="Last Name * ", font=font)
    lastName_lable.pack()

    lastName_entry = Entry(edit_screen, textvariable=lastName, font=font, width=30)
    lastName_entry.pack()

    Label(edit_screen, text="").pack()
    # Set last name
    email_lable = Label(edit_screen, text="Email * ", font=font)
    email_lable.pack()

    email_entry = Entry(edit_screen, textvariable=email, font=font, width=30)
    email_entry.pack()

    Label(edit_screen, text="").pack()
    # Set phone number
    phoneNumber_lable = Label(edit_screen, text="Phone Number * ", font=font)
    phoneNumber_lable.pack()

    phoneNumber_entry = Entry(edit_screen, textvariable=phoneNumber, font=font, width=30)
    phoneNumber_entry.pack()

    Label(edit_screen, text="").pack()
    # Set password label
    password_lable = Label(edit_screen, text="Password * ", font=font)
    password_lable.pack()

    Label(edit_screen, text="").pack()

    # Set register button
    Button(edit_screen, text="Submit", bg="#0099d8", height="2", width="30",
           command=submitChanges).pack()


def showOrders():
    while True:
        response = requests.get('http://localhost:8082/', headers={'Authorization': 'JWT ' + account_token}) #TODO Complete
        if response.status_code == 200:
            orders = response.json()  #{id: name} #TODO che sheklie?
            break

    order_screen = Toplevel()
    order_screen.title("Orders")
    order_screen.geometry("400x600")

    Label(order_screen, text="Orders", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(order_screen, text="").pack()

    for item in orders.keys():
        Button(order_screen, text=orders[item], width=300, height=5, font=("Calibri", 13),
               command=lambda: showOrderDetails(item)).pack()



def submitAddress(country, city, street, plaque):
    response = requests.post('http://localhost:8082/address/addNew', json = {"country": country,
            "city": city,
            "street": street,
            "plaque": plaque
    }, headers={'Authorization': 'JWT ' + account_token})

    if response.status_code == 200:
        global addAddres_screen
        messagebox.showinfo("Add Address", "Address added Successfully")
        addAddres_screen.destroy()
    else:
        messagebox.showerror("err", "Please try again!")


def addAddress():
    global addAddres_screen
    addAddres_screen = Toplevel()
    addAddres_screen.title("profile")
    addAddres_screen.geometry("400x300")

    Label(addAddres_screen, text="Add new Address", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(addAddres_screen, text="").pack()

    global country
    global city
    global street
    global plaque

    country = StringVar()
    city = StringVar()
    street = StringVar()
    plaque = StringVar()

    # Set label for user's instruction
    Label(addAddres_screen, text="Please fill out the form", bg="#0099d8", width="300", height="2",
          font=("Calibri", 13)).pack()
    Label(addAddres_screen, text="").pack()

    # Set first name
    country_lable = Label(addAddres_screen, text="First Name * ", font=font)
    country_lable.pack()

    country_entry = Entry(addAddres_screen, textvariable=country, font=font, width=30)
    country_entry.pack()

    Label(addAddres_screen, text="").pack()
    # Set last name
    city_lable = Label(addAddres_screen, text="Last Name * ", font=font)
    city_lable.pack()

    city_entry = Entry(addAddres_screen, textvariable=city, font=font, width=30)
    city_entry.pack()

    Label(addAddres_screen, text="").pack()
    # Set street
    street_lable = Label(addAddres_screen, text="street * ", font=font)
    street_lable.pack()

    street_entry = Entry(addAddres_screen, textvariable=street, font=font, width=30)
    street_entry.pack()

    Label(addAddres_screen, text="").pack()
    # Set street
    plaque_lable = Label(addAddres_screen, text="street * ", font=font)
    plaque_lable.pack()

    plaque_entry = Entry(addAddres_screen, textvariable=street, font=font, width=30)
    plaque_entry.pack()

    Label(addAddres_screen, text="").pack()
    Button(addAddres_screen, text="Add", width=30, height=2, bg="#0099d8", command=lambda: submitAddress(country.get(), city.get(), street.get(), plaque.get())).pack()


def showAllAddress():
    while True:
        response = requests.get('http://localhost:8082/address/getAddresses', headers={'Authorization': 'JWT ' + account_token})
        if response.status_code == 200:
            addresses = response.json()  #{id: text}
            break

    address_screen = Toplevel()
    address_screen.title("Address")
    address_screen.geometry("400x300")

    Label(address_screen, text="Addresses", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(address_screen, text="").pack()

    for m in addresses.keys():
        Text(address_screen, height=5, width=300, font=("Calibri", 13)).insert('end', addresses[m])

    Label(address_screen, text="").pack()
    Button(text="Add new Address", bg="#0099d8", height="2", width="30",font=("Calibri", 13), command=addAddress).pack()



def openProfile():
    profile_screen = Toplevel()
    profile_screen.title("profile")
    profile_screen.geometry("400x300")

    Label(profile_screen, text="profile", font = ("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(profile_screen, text="").pack()

    Button(profile_screen, text="Edit Info", width=30, height=2, bg="#0099d8", command=editInfo).pack()
    Button(profile_screen, text="My orders", width=30, height=2, bg="#0099d8", command=showOrders).pack()
    Button(profile_screen, text="Add address", width=30, height=2, bg="#0099d8", command=showAllAddress).pack()
    Button(profile_screen, text="Shopping Cart", width=30, height=2, bg="#0099d8", command=showCartProducts).pack()
    Button(profile_screen, text="Notifications", width=30, height=2, bg="#0099d8", command=showNotifications).pack()