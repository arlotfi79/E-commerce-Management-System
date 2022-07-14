from tkinter import *
from products import showProductDetails, cart

global cart_screen


def order():
    global cart_screen
    global cart
    cart_screen.destroy()

    # TODO: save order

    cart = {}


def showCartProducts():
    global cart_screen
    cart_screen = Toplevel()
    cart_screen.title("Shopping Cart")
    cart_screen.geometry("400x600")

    Label(cart_screen, text="Shopping Cart", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(cart_screen, text="").pack()

    for item in cart.keys():
        Button(cart_screen, text= cart[item]["name"], width=300, height=5, font=("Calibri", 13), command= lambda: showProductDetails(item)).pack()


    Label(cart_screen, text="").pack()
    Button(text="Order", bg="#0099d8", height="2", width="30",font=("Calibri", 13), command=order).pack()
