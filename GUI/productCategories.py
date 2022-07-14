from tkinter import *
from products import showAllProductsInCategory

def showCategories():
    categories = {}   #{id: name}    TODO: get all categories

    categories_screen = Toplevel()
    categories_screen.title("Product Categories")
    categories_screen.geometry("400x300")

    Label(categories_screen, text="All Product Categories", font = ("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(categories_screen, text="").pack()

    for c in categories.keys():
        Button(categories_screen, text=str(categories[c]), width=30, height=2, bg="#0099d8", command=lambda: showAllProductsInCategory(c, categories[c])).pack()