import requests
import tkinter as tk
from tkinter import scrolledtext

root = tk.Tk()
root.title('test_loop')
root.geometry('750x625')
root.configure(background='ivory3')
textw = scrolledtext.ScrolledText(root, width=40, height=25)
textw.grid(column=0, row=1, sticky='nsew')
textw.config(background='light grey', foreground='black', font='arial 20 bold', wrap='word', relief='sunken', bd=5)


def show_msg():
    txt = requests.get('http://127.0.0.1:13001/').text
    textw.delete('1.0', 'end')
    textw.insert('end', txt)
    root.after(2000, lambda: show_msg())

tk.Button(root, text='', bg='light grey', width=15, font='arial 12', relief='raised', bd=5,
          command=show_msg).grid(row=0, column=0)

root.mainloop()