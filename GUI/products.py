from tkinter import *
from review import showReviews

def addToCart(id, details):
    pass #TODO

def showProductDetails(productId):
    details = {}  #TODO: get details by ID

    product_screen = Toplevel()
    product_screen.title("Product Details")
    product_screen.geometry("400x500")

    Label(product_screen, text="Product Details", bg="#0099d8", width="300", height="2", font=("Calibri", 13)).pack()
    Label(product_screen, text="").pack()

    for info in details.keys():
        Label(product_screen, text = str(info) + " :", font=("Calibri", 13)).pack()
        Label(product_screen, text= str(details[info])).pack()
        Label(product_screen,text="").pack()


    Button(product_screen, text="Add to Cart", width=30, height=2, bg="#0099d8",
           command=lambda: addToCart(productId, details)).pack()
    Label(product_screen, text="").pack()
    Button(product_screen, text="reviews", width=30, height=2, bg="#ffffff", command=showReviews).pack()



def showAllProductsInCategory(categoryID, categoryName):
    products = {} #{id: name} TODO: get all products in a category

    all_product_screen = Toplevel()
    all_product_screen.title(categoryName)
    all_product_screen.geometry("400x600")

    Label(all_product_screen, text=categoryName, font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(all_product_screen, text="").pack()

    for p in products.keys():
        Button(all_product_screen, text= str(products[p]), width=300, height=5, font=("Calibri", 13), command= lambda: showProductDetails(p)).pack()
