import requests
import tkinter as tk
from tkinter import simpledialog

ROOT = tk.Tk()

ROOT.withdraw()
# the input dialog
msg = simpledialog.askstring(title="Hejsan",
                                  prompt="Meddelande till morfar: ")

requests.put(f'http://192.236.179.92:13001/?s={msg}')
