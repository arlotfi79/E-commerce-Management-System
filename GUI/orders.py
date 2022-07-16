from tkinter import *

global addMassage_Screen

def _addMassage(textInput, account_token):
    global addMassage_Screen
    text = textInput.get()
     #TODO add massage

    addMassage_Screen.destroy()
    addMassage_Screen.update()


def addMassage(account_token):
    global addMassage_Screen
    addMassage_Screen = Toplevel()
    addMassage_Screen.title("Add new Massage")
    addMassage_Screen.geometry("300x250")

    text = Text(addMassage_Screen, width=50, height=10)
    text.pack()

    Label(addMassage_Screen, text="").pack()
    Button(text="Add", bg="#0099d8", height="2", width="30",font=("Calibri", 13), command= lambda:_addMassage(text, account_token)).pack()



def track(orderId, account_token):
    allMassages = {} #TODO

    track_screen = Toplevel()
    track_screen.title("Order Tracking")
    track_screen.geometry("400x500")

    for id in allMassages:
        Text(track_screen, height=5, width=300, font=("Calibri", 13)).insert('end', allMassages[id])

    Label(track_screen, text="").pack()
    Button(track_screen, text="Add new massage", width=30, height=2, bg="#0099d8",
           command= lambda: addMassage(account_token)).pack()


def showOrderDetails(orderId, account_token):
    details = {}  #TODO: get order details

    product_screen = Toplevel()
    product_screen.title("Order Details")
    product_screen.geometry("400x500")

    Label(product_screen, text="Order Details", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(product_screen, text="").pack()

    for info in details.keys():
        Label(product_screen, text = str(info) + " :", font=("Calibri", 13)).pack()
        Label(product_screen, text= str(details[info])).pack()
        Label(product_screen,text="").pack()


    Button(product_screen, text="Track", width=30, height=2, bg="#0099d8",
           command=lambda: track(orderId, account_token)).pack()