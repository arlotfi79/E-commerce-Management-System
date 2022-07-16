from tkinter import *
from register import register
from login import login
from store_login import store_login
from store_register import register_store

def logIn_screen():
    main_screen = Tk()
    main_screen.geometry("400x500")
    main_screen.title("E-commerce Management System")


    Label(text="Costumer", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(text="").pack()

    Button(text="Login", height="2", width="30", command=lambda: login(main_screen)).pack()
    Label(text="").pack()

    Button(text="Register", bg="#0099d8", height="2", width="30", command=lambda: register()).pack()
    Label(text="").pack()
    Label(text="").pack()

    Label(text="Store", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(text="").pack()

    Button(text="Login", height="2", width="30", command=lambda: store_login(main_screen)).pack()
    Label(text="").pack()

    Button(text="Register", bg="#0099d8", height="2", width="30", command=lambda: register_store()).pack()

    main_screen.mainloop()


logIn_screen()