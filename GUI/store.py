from tkinter import *

def showStoreDetails(productId):
    details = {}  #TODO: get details by ID

    store_screen = Toplevel()
    store_screen.title("Store Details")
    store_screen.geometry("400x500")

    Label(store_screen, text="Store Details", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(store_screen, text="").pack()

    for info in details.keys():
        Label(store_screen, text = str(info) + " :", font=("Calibri", 13)).pack()
        Label(store_screen, text= str(details[info])).pack()
        Label(store_screen,text="").pack()