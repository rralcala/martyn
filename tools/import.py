#!/usr/bin/env python3
import requests
import pprint
import csv

response = requests.get('http://localhost:8080/accounts')
accounts = {}

for item in response.json():
    accounts[item["description"].lower()] = item["id"]


response = requests.get('http://localhost:8080/cost-centers')
cost_centers = {}

for item in response.json():
    cost_centers[item["description"].lower()] = item["id"]


response = requests.get('http://localhost:8080/providers')
providers = {}

for item in response.json():
    providers[item["name"].lower()] = item["id"]



with open("Book1.csv") as f:
    spamreader = csv.reader(f, delimiter=',', quotechar='"')
    for row in spamreader:
        raw_date = row[0].split("/")
        if len(raw_date[2]) == 2:
            raw_date[2] = "20"+raw_date[2]
        date = f"{raw_date[2]}-{raw_date[0].zfill(2)}-{raw_date[1].zfill(2)}"
        
        provider = providers[row[2].lower()]
        description = row[3]
        account = accounts[row[4].lower()]
        cost_center = cost_centers[row[5].lower()]
        pos = row[6].strip().replace(',', '')
        if len(pos) > 0:
            amount = int(pos)
        else:
            amount = -int(row[7].strip().replace(',', ''))
   

        tx = {
            "date": date, 
            "provider": provider, 
            "Description": description, 
            "account": account, 
            "cost_center": cost_center,
            "amount": amount
        }

        response = requests.post('http://localhost:8080/transactions', json=tx)
        print(response)
        
