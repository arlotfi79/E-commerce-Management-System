from tkinter import *
import requests
from datetime import datetime
from tkinter import messagebox

global firstName
global lastName
global email
global phoneNumber
global username
global password
global bDate
global gender
global register_screen


def register_new_store():
    fname = firstName.get()
    lname = lastName.get()
    mail = email.get()
    mobile = phoneNumber.get()
    uname = username.get()
    pas = password.get()
    birthDate = bDate.get()
    gndr = gender.get()

    response = requests.post('http://localhost:8082/signup', json={
        "username": uname,
        "email": mail,
        "name": fname,
        "lastName": lname,
        "password": pas,
        "gender": gndr.upper(),
        "birthDate": birthDate + "T00:00:00.419Z",
        "joinDate": datetime.now().isoformat() + "Z",
        "phoneNumber": mobile
    })
    if response.status_code == 200:
        messagebox.showinfo("registration", "Register Successfully")
        global register_screen
        register_screen.destroy()
    else:
        messagebox.showerror("err", "Please try again!")


def register_store():
    register_screen = Toplevel()
    register_screen.title("Register")
    register_screen.geometry("400x550")

    font = ("Calibri", 11)

    # Set text variables
    global firstName
    global lastName
    global email
    global phoneNumber
    global username
    global password
    global bDate
    global gender

    firstName = StringVar()
    lastName = StringVar()
    email = StringVar()
    phoneNumber = StringVar()
    username = StringVar()
    password = StringVar()
    bDate = StringVar()
    gender = StringVar()

    # Set label for user's instruction
    Label(register_screen, text="Please fill out the form", bg="#0099d8", width="300", height="2",
          font=("Calibri", 13)).pack()
    Label(register_screen, text="").pack()

    # Set first name
    firstName_lable = Label(register_screen, text="First Name * ", font=font)
    firstName_lable.pack()

    firstName_entry = Entry(register_screen, textvariable=firstName, font=font, width=30)
    firstName_entry.pack()

    Label(register_screen, text="").pack()
    # Set last name
    lastName_lable = Label(register_screen, text="Last Name * ", font=font)
    lastName_lable.pack()

    lastName_entry = Entry(register_screen, textvariable=lastName, font=font, width=30)
    lastName_entry.pack()

    Label(register_screen, text="").pack()
    # Set last name
    email_lable = Label(register_screen, text="Email * ", font=font)
    email_lable.pack()

    email_entry = Entry(register_screen, textvariable=email, font=font, width=30)
    email_entry.pack()

    Label(register_screen, text="").pack()
    # Set phone number
    phoneNumber_lable = Label(register_screen, text="Phone Number * ", font=font)
    phoneNumber_lable.pack()

    phoneNumber_entry = Entry(register_screen, textvariable=phoneNumber, font=font, width=30)
    phoneNumber_entry.pack()

    Label(register_screen, text="").pack()
    # Set username
    username_lable = Label(register_screen, text="Username * ", font=font)
    username_lable.pack()

    username_entry = Entry(register_screen, textvariable=username, font=font, width=30)
    username_entry.pack()

    Label(register_screen, text="").pack()

    # Set password
    password_lable = Label(register_screen, text="Password * ", font=font)
    password_lable.pack()

    password_entry = Entry(register_screen, textvariable=password, show='*', font=font, width=30)
    password_entry.pack()

    Label(register_screen, text="").pack()

    # Set Bdate
    bDate_lable = Label(register_screen, text="Birth Date* (2013-10-21) ", font=font)
    bDate_lable.pack()

    bDate_entry = Entry(register_screen, textvariable=bDate, font=font, width=30)
    bDate_entry.pack()

    Label(register_screen, text="").pack()

    # Set gender
    gender_lable = Label(register_screen, text="Gender* (male/ female/ other) ", font=font)
    gender_lable.pack()

    gender_entry = Entry(register_screen, textvariable=gender, font=font, width=30)
    gender_entry.pack()

    Label(register_screen, text="").pack()

    # Set register button
    Button(register_screen, text="Register", bg="#0099d8", height="2", width="30",
           command=register_new_store).pack()
