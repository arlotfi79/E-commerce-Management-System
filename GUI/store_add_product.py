from tkinter import *
from tkinter import messagebox

import requests

global name
global color
global price
global weight
global quantity

global product_screen

def _addProduct(catId, account_token):
    response = requests.post('http://localhost:8082/product/addNewProduct', json={
        "name": name.get(),
        "color": color.get().Upper(),
        "price": float(price),
        "weight": float(weight),
        "quantity": int(quantity)
    }, headers={'Authorization': 'JWT ' + account_token})

    if response.status_code == 200:
        response = requests.post('http://localhost:8082/product/addNewProduct', json = {
        "product_id": 1,  #TODO product id?
        "category_id": catId
    }, headers={'Authorization': 'JWT ' + account_token})

        if response.status_code == 200:
            global product_screen
            messagebox.showinfo("add product", "Added Successfully")
            product_screen.destroy()
        else:
            messagebox.showerror("err", "Please try again!")



def addProductToCategory(catId, account_token):
    global product_screen

    product_screen = Toplevel()
    product_screen.title("Add Product")
    product_screen.geometry("400x550")

    font = ("Calibri", 11)


    # Set text variables
    global name
    global color
    global price
    global weight
    global quantity


    name = StringVar()
    color = StringVar()
    price = StringVar()
    weight = StringVar()
    quantity = StringVar()

    # Set label for user's instruction
    Label(product_screen, text="Please fill out the form", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(product_screen, text="").pack()

    # name
    firstName_lable = Label(product_screen, text="Product Name * ", font=font)
    firstName_lable.pack()

    firstName_entry = Entry(product_screen, textvariable=name, font=font, width= 30)
    firstName_entry.pack()

    Label(product_screen, text="").pack()
    # color
    lastName_lable = Label(product_screen, text="Color * ", font=font)
    lastName_lable.pack()

    lastName_entry = Entry(product_screen, textvariable=color, font=font, width= 30)
    lastName_entry.pack()

    Label(product_screen, text="").pack()
    # price
    email_lable = Label(product_screen, text="Price * ", font=font)
    email_lable.pack()

    email_entry = Entry(product_screen, textvariable=price, font=font, width= 30)
    email_entry.pack()

    Label(product_screen, text="").pack()
    # weight
    phoneNumber_lable = Label(product_screen, text="Weight * ", font=font)
    phoneNumber_lable.pack()

    phoneNumber_entry = Entry(product_screen, textvariable=weight, font=font, width= 30)
    phoneNumber_entry.pack()

    Label(product_screen, text="").pack()

    # quantity
    username_lable = Label(product_screen, text="Quantity * ", font=font)
    username_lable.pack()

    username_entry = Entry(product_screen, textvariable=quantity, font=font, width=30)
    username_entry.pack()

    Label(product_screen, text="").pack()

    Button(product_screen, text="ADD", bg="#0099d8", height="2", width="30", command=lambda: _addProduct(catId, account_token)).pack()

def addProduct(account_token):
    while True:
        response = requests.get('http://localhost:8082/category/all')
        if response.status_code == 200:
            categories = response.json()  # {id: name}
            break

    categories_screen = Toplevel()
    categories_screen.title("Product Categories")
    categories_screen.geometry("400x300")

    Label(categories_screen, text="Select Category", font = ("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(categories_screen, text="").pack()

    for c in categories.keys():
        Button(categories_screen, text=str(categories[c]), width=30, height=2, bg="#0099d8", command=lambda: addProductToCategory(c, account_token)).pack()