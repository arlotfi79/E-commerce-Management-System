from tkinter import *
from shoppingCart import showCartProducts
from orders import showOrderDetails

global firstName
global lastName
global email
global phoneNumber

def showNotifications():
    pass

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
    orders = {}    #id:date  TODO

    order_screen = Toplevel()
    order_screen.title("Orders")
    order_screen.geometry("400x600")

    Label(order_screen, text="Orders", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(order_screen, text="").pack()

    for item in orders.keys():
        Button(order_screen, text=orders[item], width=300, height=5, font=("Calibri", 13),
               command=lambda: showOrderDetails(item)).pack()



def submitAddress(address):
    pass
    #TODO save address


def addAddress():
    addAddres_screen = Toplevel()
    addAddres_screen.title("profile")
    addAddres_screen.geometry("400x300")

    Label(addAddres_screen, text="Add new Address", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(addAddres_screen, text="").pack()

    text = Text(addAddres_screen, width=50, height=10)
    text.pack()
    Label(addAddres_screen, text="").pack()
    address = text.get()
    Button(addAddres_screen, text="Add", width=30, height=2, bg="#0099d8", command=lambda: submitAddress(address)).pack()



def openProfile():
    profile_screen = Toplevel()
    profile_screen.title("profile")
    profile_screen.geometry("400x300")

    Label(profile_screen, text="profile", font = ("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(profile_screen, text="").pack()

    Button(profile_screen, text="Edit Info", width=30, height=2, bg="#0099d8", command=editInfo).pack()
    Button(profile_screen, text="My orders", width=30, height=2, bg="#0099d8", command=showOrders).pack()
    Button(profile_screen, text="Add address", width=30, height=2, bg="#0099d8", command=addAddress).pack()
    Button(profile_screen, text="Shopping Cart", width=30, height=2, bg="#0099d8", command=showCartProducts).pack()
    Button(profile_screen, text="Notifications", width=30, height=2, bg="#0099d8", command=showNotifications).pack()