from tkinter import *
from shoppingCart import showCartProducts

def editInfo():
    pass

def showOrders():
    pass

def addAddress():
    pass


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
