from tkinter import *
from store_edit_info import editInfo
from store_add_product import addProduct

def main_account_screen(mainScreen):
    main_screen = Tk()
    main_screen.title("E-commerce Management System")
    main_screen.geometry("500x350")

    Label(text="E-commerce Management System", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(text="").pack()
    Label(text="").pack()

    Button(text="Add new Product", bg="#0099d8", height="2", width="30", font=("Calibri", 13), command=editInfo).pack()
    Label(text="").pack()
    Label(text="").pack()
    Button(text="Edit Info", bg="#0099d8", height="2", width="30", font=("Calibri", 13), command= addProduct).pack()

    main_screen.mainloop()