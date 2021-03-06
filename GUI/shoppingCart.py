from tkinter import *
from products import showProductDetails

global cart_screen


def order(account_token):
    global cart_screen
    cart_screen.destroy()
    cart_screen.update()

    #TODO: save order



def showCartProducts(account_token):
    cart = {} #TODO

    global cart_screen
    cart_screen = Toplevel()
    cart_screen.title("Shopping Cart")
    cart_screen.geometry("400x600")

    Label(cart_screen, text="Shopping Cart", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(cart_screen, text="").pack()

    for item in cart.keys():
        Button(cart_screen, text= cart[item]["name"], width=300, height=5, font=("Calibri", 13), command= lambda: showProductDetails(item)).pack()


    Label(cart_screen, text="").pack()
    Button(text="Order", bg="#0099d8", height="2", width="30",font=("Calibri", 13), command=lambda: order(account_token)).pack()
