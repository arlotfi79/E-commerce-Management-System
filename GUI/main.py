from tkinter import *
from register import register
from login import login
from productCategories import showCategories
from profile import openProfile

global main_screen

def main_account_screen():
    global main_screen
    main_screen = Tk()  # create a GUI window
    main_screen.title("E-commerce Management System")
    main_screen.geometry("500x350")

    Label(text="E-commerce Management System", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(text="").pack()
    Label(text="").pack()

    Button(text="Product categories", bg="#0099d8", height="2", width="30", command=showCategories, font=("Calibri", 13)).pack()
    Label(text="").pack()
    Label(text="").pack()
    Button(text="Profile", bg="#0099d8", height="2", width="30", command=openProfile, font=("Calibri", 13)).pack()

    main_screen.mainloop()



def logIn_screen():
    global main_screen
    main_screen = Tk()  # create a GUI window
    main_screen.geometry("400x250")  # set the configuration of GUI window
    main_screen.title("E-commerce Management System")  # set the title of GUI window

    # create a Form label
    Label(text="Please Log in first", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(text="").pack()

    # create Login Button
    Button(text="Login", height="2", width="30", command=login).pack()
    Label(text="").pack()

    # create a register button
    Button(text="Register", bg="#0099d8", height="2", width="30", command=register).pack()

    main_screen.mainloop()  # start the GUI


logIn_screen()  # call the main_account_screen() function
# main_account_screen()