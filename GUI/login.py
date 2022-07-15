from tkinter import *

import requests

global username_verify
global password_verify
global account_token

def login_verification(mainScreen):
    username = username_verify.get()
    password = password_verify.get()

    #TODO: login process
    response = requests.post('http://localhost:8082/signin', json={
        "username": username,
        "password": password
    })

    global account_token
    account_token = response.json()['access_token']


def login(mainScreen):
    login_screen = Toplevel()
    login_screen.title("Login")
    login_screen.geometry("400x300")

    Label(login_screen, text="Welcome back!", font = ("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(login_screen, text="").pack()

    font = ("Calibri", 11)

    global username_verify
    global password_verify

    username_verify = StringVar()
    password_verify = StringVar()

    Label(login_screen, text="Username ", font=font).pack()
    username_login_entry = Entry(login_screen, textvariable=username_verify, width=30, font=font)
    username_login_entry.pack()

    Label(login_screen, text="").pack()
    Label(login_screen, text="Password ", font=font).pack()

    password__login_entry = Entry(login_screen, textvariable=password_verify, show='*', width=30, font=font)
    password__login_entry.pack()

    Label(login_screen, text="").pack()
    Button(login_screen, text="Login", width=30, height=2, bg="#0099d8", command=lambda: login_verification(mainScreen)).pack()