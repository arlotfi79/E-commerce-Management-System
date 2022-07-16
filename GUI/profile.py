from tkinter import *
from tkinter import messagebox

import requests

from shoppingCart import showCartProducts
from orders import showOrderDetails

global firstName
global lastName
global email
global phoneNumber

global addAddres_screen


def showNotifications(account_token):
    notif_screen = Toplevel()
    notif_screen.title("Notification")
    notif_screen.geometry("400x550")
    Label(notif_screen, text="Massages", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(notif_screen, text="").pack()

    response = requests.get('http://localhost:8082/notif', headers={'Authorization': 'Bearer ' + account_token})
    if response.status_code == 200:
        notifs = response.json()
        for m in notifs:
            temp = m.json()
            Text(notif_screen, height=5, width=300, font=("Calibri", 13)).insert('end', temp["description"])


def showOrders(account_token):
    while True:
        response = requests.get('http://localhost:8082/order/all', headers={'Authorization': 'Bearer ' + account_token})
        if response.status_code == 200:
            listOforders = response.json()
            break

    order_screen = Toplevel()
    order_screen.title("Orders")
    order_screen.geometry("400x600")

    Label(order_screen, text="Orders", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(order_screen, text="").pack()

    for order in listOforders:
        temp = order.json()
        Button(order_screen, text=temp["description"] + " / " + temp["orderDate"], width=300, height=5, font=("Calibri", 13),
               command=lambda: showOrderDetails(temp, account_token)).pack()



def submitAddress(country, city, street, plaque, account_token):
    response = requests.post('http://localhost:8082/address/addNew', json = {"country": country,
            "city": city,
            "street": street,
            "plaque": plaque
    }, headers={'Authorization': 'Bearer ' + account_token})

    if response.status_code == 200:
        global addAddress_screen
        messagebox.showinfo("Add Address", "Address added Successfully")
        addAddress_screen.destroy()
        addAddress_screen.update()
    else:
        messagebox.showerror("err", "Please try again!")


def addAddress(account_token):
    global addAddress_screen
    addAddress_screen = Toplevel()
    addAddress_screen.title("profile")
    addAddress_screen.geometry("400x300")

    Label(addAddress_screen, text="Add new Address", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(addAddress_screen, text="").pack()

    global country
    global city
    global street
    global plaque

    country = StringVar()
    city = StringVar()
    street = StringVar()
    plaque = StringVar()

    # Set label for user's instruction
    Label(addAddress_screen, text="Please fill out the form", bg="#0099d8", width="300", height="2",
          font=("Calibri", 13)).pack()
    Label(addAddress_screen, text="").pack()

    # Set first name
    country_lable = Label(addAddress_screen, text="First Name * ", font=font)
    country_lable.pack()

    country_entry = Entry(addAddress_screen, textvariable=country, font=font, width=30)
    country_entry.pack()

    Label(addAddress_screen, text="").pack()
    # Set last name
    city_lable = Label(addAddress_screen, text="Last Name * ", font=font)
    city_lable.pack()

    city_entry = Entry(addAddress_screen, textvariable=city, font=font, width=30)
    city_entry.pack()

    Label(addAddress_screen, text="").pack()
    # Set street
    street_lable = Label(addAddress_screen, text="street * ", font=font)
    street_lable.pack()

    street_entry = Entry(addAddress_screen, textvariable=street, font=font, width=30)
    street_entry.pack()

    Label(addAddress_screen, text="").pack()
    # Set street
    plaque_lable = Label(addAddress_screen, text="street * ", font=font)
    plaque_lable.pack()

    plaque_entry = Entry(addAddress_screen, textvariable=street, font=font, width=30)
    plaque_entry.pack()

    Label(addAddress_screen, text="").pack()
    Button(addAddress_screen, text="Add", width=30, height=2, bg="#0099d8", command=lambda: submitAddress(country.get(), city.get(), street.get(), plaque.get(), account_token)).pack()


def showAllAddress(account_token):
    while True:
        response = requests.get('http://localhost:8082/address/getAddresses', headers={'Authorization': 'Bearer ' + account_token})
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
    Button(text="Add new Address", bg="#0099d8", height="2", width="30",font=("Calibri", 13), command=lambda: addAddress(account_token)).pack()


def removeFromWatchlist(productId, account_token):
    requests.post('http://localhost:8082/watchlist/remove', json = {"id": productId} ,headers={'Authorization': 'Bearer ' + account_token} )

def showProductDetails(product, account_token):
    details = product

    product_screen = Toplevel()
    product_screen.title("Product Details")
    product_screen.geometry("400x500")

    Label(product_screen, text="Product Details", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(product_screen, text="").pack()

    for info in details.keys():
        Label(product_screen, text = str(info) + " :", font=("Calibri", 13)).pack()
        Label(product_screen, text= str(details[info])).pack()
        Label(product_screen,text="").pack()


    Button(product_screen, text="Remove", width=30, height=2, bg="#0099d8",
           command=lambda: removeFromWatchlist(product["id"], account_token)).pack()



def showWhatchList(account_token):
    while True:
        response = requests.get('http://localhost:8082/product/watchlist', headers={'Authorization': 'Bearer ' + account_token})
        if response.status_code == 200:
            watchList = response.json()
            break

    global watchList_screen
    watchList_screen = Toplevel()
    watchList_screen.title("Shopping Cart")
    watchList_screen.geometry("400x600")

    Label(watchList_screen, text="Shopping Cart", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(watchList_screen, text="").pack()

    for p in watchList:
        Button(watchList_screen, text=p["name"] + p["color"], width=300, height=5, font=("Calibri", 13),
               command=lambda: showProductDetails(p, account_token)).pack()


def openProfile(token):
    profile_screen = Toplevel()
    profile_screen.title("profile")
    profile_screen.geometry("400x300")

    Label(profile_screen, text="profile", font = ("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(profile_screen, text="").pack()

    Button(profile_screen, text="My orders", width=30, height=2, bg="#0099d8", command=lambda: showOrders(token)).pack()
    Button(profile_screen, text="Add address", width=30, height=2, bg="#0099d8", command=lambda: showAllAddress(token)).pack()
    Button(profile_screen, text="Shopping Cart", width=30, height=2, bg="#0099d8", command=lambda: showCartProducts(token)).pack()
    Button(profile_screen, text="Notifications", width=30, height=2, bg="#0099d8", command=lambda: showNotifications(token)).pack()
    Button(profile_screen, text="Watch list", width=30, height=2, bg="#0099d8", command=lambda: showWhatchList(token)).pack()