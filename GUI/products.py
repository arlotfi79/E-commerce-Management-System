from tkinter import *
from tkinter import messagebox

import requests

from review import showReviews

def addToCart(id, account_token):
    response = requests.post('http://localhost:8082/cart', json = {"productId": id, "productCount": 1}, headers={'Authorization': 'Bearer ' + account_token})
    if response.status_code == 200:
        messagebox.showinfo("Add To cart", "Added Successfully")


def showProductDetails(productDetails, categoryName, account_token):
    details = productDetails

    product_screen = Toplevel()
    product_screen.title("Product Details")
    product_screen.geometry("400x500")

    Label(product_screen, text="Product Details", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(product_screen, text="").pack()

    for info in details.keys():
        Label(product_screen, text = str(info) + " :", font=("Calibri", 13)).pack()
        Label(product_screen, text= str(details[info])).pack()
        Label(product_screen,text="").pack()


    Button(product_screen, text="Add to Cart", width=30, height=2, bg="#0099d8",
           command=lambda: addToCart(productDetails["id"], account_token)).pack()
    Label(product_screen, text="").pack()
    Button(product_screen, text="reviews", width=30, height=2, bg="#ffffff", command=lambda: showReviews(productDetails["id"], account_token)).pack()

    Label(product_screen, text="").pack()
    Label(product_screen, text="").pack()
    Label(product_screen, text="").pack()
    Label(product_screen, text="").pack()
    while True:
        response = requests.get('http://localhost:8082/product/byCategory', json={"name" : categoryName} , headers={'Authorization': 'Bearer ' + account_token})
        if response.status_code == 200:
            products = response.json()
            break

    for p in range(0, len(products), 2):
        name = products[p]["name"]
        Button(product_screen, text= str(name), width=300, height=5, font=("Calibri", 13), command= lambda: showProductDetails(p,categoryName, account_token)).pack()




def showAllProductsInCategory(categoryID, categoryName, account_token):
    while True:
        response = requests.get('http://localhost:8082/product/byCategory', json={"name" : categoryName} , headers={'Authorization': 'Bearer ' + account_token})
        if response.status_code == 200:
            products = response.json()
            break

    all_product_screen = Toplevel()
    all_product_screen.title(categoryName)
    all_product_screen.geometry("400x600")

    Label(all_product_screen, text=categoryName, font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(all_product_screen, text="").pack()

    for p in products:
        name = p["name"]
        Button(all_product_screen, text= str(name), width=300, height=5, font=("Calibri", 13), command= lambda: showProductDetails(p,categoryName, account_token)).pack()




