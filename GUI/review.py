from tkinter import *

global addMassage_Screen
global massage_Screen
global addreview_Screen
global review_screen

def _addReview(subjectInput, textInput):
    global addreview_Screen
    global review_screen

    subject = subjectInput.get()
    text = textInput.get()

    #TODO: add review

    addreview_Screen.destroy()
    review_screen.destroy()


def addReview(productId):
    addreview_Screen = Toplevel()
    addreview_Screen.title("Add new Review")
    addreview_Screen.geometry("300x250")


    subject = Entry(addreview_Screen,width=50)
    subject.pack()

    Label(addreview_Screen, text="").pack()

    text = Text(addreview_Screen, width=50, height=10)
    text.pack()

    Label(addreview_Screen, text="").pack()
    Button(text="Add", bg="#0099d8", height="2", width="30",font=("Calibri", 13), command= lambda:_addReview(subject, text)).pack()



def _addMassage(textInput):
    global addMassage_Screen
    global massage_Screen
    text = textInput.get()

    #TODO: add massage

    addMassage_Screen.destroy()
    massage_Screen.destroy()


def addMassage():
    global addMassage_Screen

    addMassage_Screen = Toplevel()
    addMassage_Screen.title("Add new Massage")
    addMassage_Screen.geometry("300x250")

    text = Text(addMassage_Screen, width=50, height=10)
    text.pack()

    Label(addMassage_Screen, text="").pack()
    Button(text="Add", bg="#0099d8", height="2", width="30",font=("Calibri", 13), command= lambda:_addMassage(text)).pack()


def upVote(reviewId):
    pass

def downVote(reviewId):
    pass

def showMassages(reviewId):
    global massage_Screen

    massages = {}  #{id: text} TODO: get massages

    massage_Screen = Toplevel()
    massage_Screen.title("Massages")
    massage_Screen.geometry("400x600")

    Label(massage_Screen, text="Massages", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(massage_Screen, text="").pack()

    for m in massages.keys():
        Text(massage_Screen, height=5, width=300, font=("Calibri", 13)).insert('end', massages[m])

    Label(massage_Screen, text="").pack()
    Button(text="Add massage", bg="#0099d8", height="2", width="30",font=("Calibri", 13), command=addMassage).pack()
    Button(text="UpVote", bg="#018600", height="2", width="30", font=("Calibri", 13), command=lambda: upVote(reviewId)).pack()
    Button(text="DownVote", bg="#c10100", height="2", width="30", font=("Calibri", 13), command=lambda: downVote(reviewId)).pack()



def showReviews(productID):
    reviews = {}  #{id: subject}  TODO: get rewiews

    review_screen = Toplevel()
    review_screen.title("Reviews")
    review_screen.geometry("400x600")

    Label(review_screen, text="Reviews", font=("Calibri", 13), bg="#0099d8", width="300", height="2").pack()
    Label(review_screen, text="").pack()

    for r in reviews.keys():
        Button(review_screen, text= reviews[r], width=300, height=5, font=("Calibri", 13), command= lambda: showMassages(r)).pack()


    Label(review_screen, text="").pack()
    Button(text="Add review", bg="#0099d8", height="2", width="30",font=("Calibri", 13), command=lambda: addReview(productID)).pack()