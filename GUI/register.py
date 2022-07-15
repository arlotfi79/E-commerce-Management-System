from tkinter import *

global firstName
global lastName
global email
global phoneNumber
global password

def register_user():
    fname = firstName.get()
    lname = lastName.get()
    mail = email.get()
    mobile = phoneNumber.get()
    pas = password.get()

    #TODO: register user





def register():
    register_screen = Toplevel()
    register_screen.title("Register")
    register_screen.geometry("400x550")

    font = ("Calibri", 11)


    # Set text variables
    global firstName
    global lastName
    global email
    global phoneNumber
    global password

    firstName = StringVar()
    lastName = StringVar()
    email = StringVar()
    phoneNumber = StringVar()
    password = StringVar()

    # Set label for user's instruction
    Label(register_screen, text="Please fill out the form", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(register_screen, text="").pack()

    # Set first name
    firstName_lable = Label(register_screen, text="First Name * ", font=font)
    firstName_lable.pack()

    firstName_entry = Entry(register_screen, textvariable=firstName, font=font, width= 30)
    firstName_entry.pack()

    Label(register_screen, text="").pack()
    # Set last name
    lastName_lable = Label(register_screen, text="Last Name * ", font=font)
    lastName_lable.pack()

    lastName_entry = Entry(register_screen, textvariable=lastName, font=font, width= 30)
    lastName_entry.pack()

    Label(register_screen, text="").pack()
    # Set last name
    email_lable = Label(register_screen, text="Email * ", font=font)
    email_lable.pack()

    email_entry = Entry(register_screen, textvariable=email, font=font, width= 30)
    email_entry.pack()

    Label(register_screen, text="").pack()
    # Set phone number
    phoneNumber_lable = Label(register_screen, text="Phone Number * ", font=font)
    phoneNumber_lable.pack()

    phoneNumber_entry = Entry(register_screen, textvariable=phoneNumber, font=font, width= 30)
    phoneNumber_entry.pack()

    Label(register_screen, text="").pack()
    # Set password label
    password_lable = Label(register_screen, text="Password * ", font=font)
    password_lable.pack()

    password_entry = Entry(register_screen, textvariable=password, show='*', font=font, width= 30)
    password_entry.pack()

    Label(register_screen, text="").pack()

    # Set register button
    Button(register_screen, text="Register", bg="#0099d8", height="2", width="30").pack()