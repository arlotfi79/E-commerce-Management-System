from tkinter import *

def track(orderId):
    pass



def showOrderDetails(orderId):
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
           command=lambda: track(orderId)).pack()   ###TODO: track?