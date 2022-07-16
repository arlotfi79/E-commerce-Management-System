from tkinter import *
from productCategories import showCategories
from profile import openProfile

def main_account_screen(mainScreen, token):
    main_screen = Tk()
    main_screen.title("E-commerce Management System")
    main_screen.geometry("500x350")

    Label(text="E-commerce Management System", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(text="").pack()
    Label(text="").pack()

    Button(text="Product categories", bg="#0099d8", height="2", width="30", command=lambda: showCategories(token), font=("Calibri", 13)).pack()
    Label(text="").pack()
    Label(text="").pack()
    Button(text="Profile", bg="#0099d8", height="2", width="30", command=lambda: openProfile(token), font=("Calibri", 13)).pack()

    main_screen.mainloop()