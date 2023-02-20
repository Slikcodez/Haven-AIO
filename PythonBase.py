import requests
import json
import random
import re
import uuid
import os
import tls_client
import sys
import time
from flask import Flask
import datetime
import logging
from bs4 import BeautifulSoup
from threading import Thread
import pytz
from random import randint
import globals
import os
from faker import Faker
import string
import datetime
from websocket import create_connection
import websockets
import os
import ctypes
os.environ['PYGAME_HIDE_SUPPORT_PROMPT'] = "hide"
import pygame
import base64
import io
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad,unpad
from pyee.base import EventEmitter
import subprocess
key = 'dskxuahzsnqavpew'
globals.version = "0.5.1-HEHEHEHAHAHAHA"
globals.var = None
globals.node = None
globals.carts = 0
globals.checkouts = 0
globals.declines = 0
globals.modVar = random.randint(0, 1)

globals.dropTime = 0

def random_string(length):
    letters_and_digits = string.ascii_letters + string.digits
    return ''.join(random.choice(letters_and_digits) for i in range(length))

def generate_string(min_length, max_length):
    length = random.randint(min_length, max_length)
    return random_string(length)


def runTls():
    command = f"{os.getcwd()}/Configs/Hibbett/Encryption/paymentEncryption2.exe"
    process = subprocess.Popen([command], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)

    # Keep the script running until it is manually stopped
    while True:
        pass


Thread(target=runTls).start()




def initalTimer():
    while True:
        if int(datetime.datetime.now().strftime("%H")) == globals.dropTime:
            globals.ee.emit('initialDrop')

Thread(target=initalTimer).start()
def dropTime(time):
    if time != None and globals.dropTime == 0:
        globals.dropTime = time
        return 'cloud'
    else:
        return globals.dropTime




globals.ee = EventEmitter()

globals.serverUrl = "http://38.102.8.15"


now = datetime.datetime.now()
fake = Faker()

def runBot():

    def titleStats():
        while True:
            ctypes.windll.kernel32.SetConsoleTitleW(f"Hibbett Haven - {str(globals.version)} | Carts {str(globals.carts)} | Checkouts {str(globals.checkouts)} | Declines {str(globals.declines)}                                                                                               {datetime.datetime.now().strftime('%b %d %I:%M %p')}")
            time.sleep(1)

    def decrypt(enc):
        enc = enc.encode("utf-8")
        enc = base64.b64decode(enc)
        cipher = AES.new(key.encode('utf-8'), AES.MODE_ECB)
        encVal = unpad(cipher.decrypt(enc),16)
        encVal= encVal.decode("utf-8", "ignore")
        return encVal

    def playSound(song):

        try:

            res = requests.post(f"{globals.serverUrl}:4567/api/v1/getsong",
                json={
                    "song": song
                }
            )

            pygame.mixer.init()
            sound_file_data = base64.b64decode(res.json()['data'])
            sound_file = io.BytesIO(sound_file_data)
            sound = pygame.mixer.Sound(sound_file)
            ch = sound.play()
            while ch.get_busy():
                pygame.time.wait(20)

        except:
            pass

    def inputb(message):
        globals.modVar += 1

        if globals.modVar % 2 == 0:
            return input(f"\033[1m{message}\033[0m")
        else:
            return input(f"\033[1m{message}\033[0m")

    def printb(message):
        globals.modVar += 1

        if globals.modVar % 2 == 0:
            print(f"\033[31m{message}\033[31m")
        else:
            print(f"\033[31m{message}\033[31m")

    def checkNode(license):

        while True:
        
            try:

                res = requests.post(f"{globals.serverUrl}:4567/api/v1/youlikemen",
                    headers={},
                    json={
                        "checkVal": "get",
                        "license": license
                    }
                )

                if res.status_code == 200:
                    if res.json()['node'] != "regen susness":
                        if int(res.json()['node']) == globals.node:
                            time.sleep(5)
                            continue
                        else:
                            os._exit(0)
                    else:
                        nodeRes = newAuth.nodeFunc()

                        if nodeRes == True:
                            time.sleep(5)
                            continue
                        else:
                            print("Sorry, there was a problem with UID verification, contact devs")
                            os._exit(0)
                else:
                    os._exit(0)

            except:
                continue

    def authLoop(key):

        while True:

            try:
                res = requests.post(f"{globals.serverUrl}:4567/auth",
                    headers = {
                            "Content-Type": "application/json",
                            "Authorization": "Bearer pk_cqzp7m2w4zsl5jlx233swz3uj7prgkmc5lmzg205"
                        },
                    json = {
                        'license': key,
                        'hwid': ':'.join(re.findall('..', '%012x' % uuid.getnode())),
                        "type": "check"
                    }
                )
                    
                if res.status_code == 200:
                    time.sleep(5)
                    continue
                else:
                    os._exit(0)
            except:
                continue

    def log(message, thread):
        message = message.upper()
        printb(f"Account {thread} [{str(datetime.datetime.now().strftime('%H:%M:%S'))}]: {message}")




    

    class HibbettTask():

        def __init__(self, mode, threadVal, uaBool):
            self.mode = mode
            self.thread = threadVal
            self.session = None
            self.paymentId = None
            self.token = None
            self.userAgent = f"Hibbett | CG/6.3.0 (com.hibbett.hibbett-sports; build:{str(randint(1,15000))}; iOS 16.0.2)"
            self.paymentType = None
            self.sessionId = None
            self.sessionEX = None
            self.customerId = None
            self.email = None
            self.password = None
            self.four = None
            self.checkedOut = False
            self.genningPx = False
            self.session = requests.session()
            self.proxy = None
            self.initialSku = None
            self.detected = False
            self.sku = None
            if uaBool != "random":
                self.userAgentVal = uaBool
            else:
                self.userAgentVal = random.choice(['2','3'])

        def run_task(self):
            self.userAgent = f"Hibbett | CG/6.3.0 (com.hibbett.hibbett-sports; build:{str(randint(1,15000))}; iOS 16.0.2)"
            if self.genningPx == False:
                Thread(target=self.loopPxToken).start()
                self.genningPx = True
            self.token = None
            self.getSession()

        def getSession(self):
            try:
                with open("./Configs/proxies.txt", "r") as proxies:
                    proxy = random.choice(proxies.read().split("\n"))

                    if "hhproxies" in proxy:
                        proxy = proxy.replace("hhproxies", "")
                        proxy = decrypt(proxy)

                        with open("./Configs/key.json", "r") as keyFile:
                            keyFile = json.loads(keyFile.read())
                        
                        if keyFile["key"] in proxy:
                            proxy = proxy.replace("splitval", "")
                            proxy = proxy.replace(keyFile["key"], "")
                        else:
                            print("Stop using stolen proxies bozo, next time you will be banned")
                            time.sleep(2)
                            os._exit(0)

                    IP = proxy.split(":")[0]
                    Port = proxy.split(":")[1]
                    User = proxy.split(":")[2]
                    Pass = proxy.split(":")[3]
                    proxies = f"http://{str(User)}:{str(Pass)}@{str(IP)}:{str(Port)}"


                self.proxy = proxies
                self.getRealSession()

            except Exception as e:
                print(e)
                log("Error getting proxy", self.thread)
                self.getSession()
    
        def getRealSession(self):
            try:
                sessionVar = requests.session()
                self.session = sessionVar
                self.getAccInfo()
            except Exception as e:
                printb(e)
                log("Error getting session", self.thread)
                self.getRealSession()

        def getAccInfo(self):
            
            with open("./Configs/Hibbett/accounts.txt") as account:
                e = account.read().split("\n")
                for item in e:
                    if item == "":
                        e.remove(item)

                splitaroony = e[int(self.thread) - 1].split(":")
                self.email = splitaroony[0]
                self.password = splitaroony[1]
                self.four = splitaroony[2]
                self.cvv_value = splitaroony[3]

                self.getCustomerId()
        
        def hour(self):
            hour_val = int(datetime.datetime.now().strftime("%M"))
            return hour_val

        def loopPxToken(self):
            time.sleep(3)
            if self.proxy:
                time.sleep(3)
                res = requests.post(f"{globals.serverUrl}:4567/pxGen", headers={
                            'Connection': 'keep-alive',
                            'Content-Type': 'application/json; charset=utf-8',
                            'version': '6.2.1',
                            'platform': 'ios',
                            "User-Agent": self.userAgent,
                            'Accept': '*/*',
                            'Accept-Encoding': 'br;q=1.0, gzip;q=0.9, deflate;q=0.8',
                            #another method for px is to rempove X-PX-ORIGINAL-TOKEN and change X-PX-AUTHORIZATION to 2, Expirmental is changing X-PX-AUTHORIZATION to 4 and remove X-PX-ORIGINAL-TOKEN
                        }, json={'proxy': self.proxy})
                if len(res.text) > 100:
                    self.token = res.text
                time.sleep(180)
            else:
                self.loopPxToken()

        def getCustomerId(self):
            log("LOGGING IN", self.thread)
            try:
                self.returnToken()
                res = self.session.post("http://127.0.0.1:8082",
                    headers={
                        'poptls-url': "https://hibbett-mobileapi.prolific.io/users/login",
                        'poptls-proxy': self.proxy,
                        'Accept': '*/*',
                        'Accept-Encoding': 'br;q=1.0, gzip;q=0.9, deflate;q=0.8',
                        'Accept-Language': 'en-US;q=1.0',
                        'Connection': 'keep-alive',
                        'Content-Type': 'application/json; charset=utf-8',
                        'Host': 'hibbett-mobileapi.prolific.io',
                        'platform': 'ios',
                        'version': '6.3.0',
                        'x-api-key': '0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0',
                        'X-PX-AUTHORIZATION': '2',
                        #'X-PX-ORIGINAL-TOKEN': f'{self.userAgentVal}:{self.returnToken()}',
                        "User-Agent": self.userAgent
                    },
                    json={
                        "login": self.email,
                        "password": self.password
                    }
                )

                if res.status_code == 401:
                    print(res.text)
                    log("INVALID ACCOUNT CREDENTIALS", self.thread)
                elif res.status_code == 403:
                    log("PX BLOCK", self.thread)
                    
                    self.run_task()
                elif res.status_code == 200:
                    log("LOGGED IN", self.thread)

                    self.sessionEX = self.hour()
                    self.sessionId = res.json()['sessionId']
                    self.customerId = res.json()['customerId']
                    try:
                        if self.paymentId == None:
                            self.getPaymentId()
                        else:
                            if self.thread == 1:
                                if self.mode == "1":
                                    self.monitorMonitor()
                                else:
                                    self.cloudMonitor()
                            else:
                                self.waitingForOneTap()
                    except Exception as e:
                        pass
                else:
                    log(f"UNKNOWN ERROR {str(res.status_code)}", self.thread)
                    time.sleep(60)
                    self.getSession()
            except Exception as e:
                print(e)
                log("ERROR LOGGING IN, RETRYING...", self.thread)
                time.sleep(15)
                self.getSession()

        def getPaymentId(self):
            try:
                res = self.session.get(f"http://127.0.0.1:8082",
                    headers = {
                        'poptls-url': f"https://hibbett-mobileapi.prolific.io/users/{self.customerId}/payment_methods",
                        'poptls-proxy': self.proxy,
                        'Accept': '*/*',
                        'Accept-Encoding': 'br;q=1.0, gzip;q=0.9, deflate;q=0.8',
                        'Accept-Language': 'en-US;q=1.0',
                        "Authorization": f"Bearer {self.sessionId}",
                        'Connection': 'keep-alive',
                        'Content-Type': 'application/json; charset=utf-8',
                        'Host': 'hibbett-mobileapi.prolific.io',
                        'platform': 'ios',
                        'version': '6.2.1',
                        'x-api-key': '0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0',
                        'X-PX-AUTHORIZATION': self.userAgentVal,
                        "User-Agent": self.userAgent
                    }
                )
            except:
                log('error getting payment id', self.thread)
            if res.status_code != 200:
                self.run_task()
            else:
                for payment in res.json():
                    if payment['paymentObject']['number'] == self.four:
                        self.paymentType = payment['paymentObject']['cardType'] 
                        self.paymentId = payment['id']
                        log("GOT PAYMENT ID", self.thread)
                        if dropTime(None) != 0:
                            self.initialSku = random.choice(open("./configs/Hibbett/InitialSkus.txt").read().split("\n"))
                            @globals.ee.on('initialDrop')
                            def trigger():
                                print("triggering")
                                self.initateOneTap()
                        elif self.thread == 1:
                            if self.mode == "1":
                                self.monitorMonitor()
                            else:
                                self.cloudMonitor()
                        else:
                            self.waitingForOneTap()
                            
        def cloudMonitor(self):

            printb("Monitor: CONNECTED TO CLOUD MONITOR (ALPHA)")

            while self.detected == False:

                if self.sessionEX != self.hour() + 1:

                    ws = create_connection("ws://38.102.8.15:8001")

                    while True:
                        msg = ws.recv()
                        Thread(target=self.checkinput,args=(json.loads(msg),)).start()

                else:
                    self.getCustomerId()

        def checkinput(self, jsondata):

            with open("./Configs/Hibbett/skus.txt", "r") as getSkus:
                skus = getSkus.read().split("\n")

            with open("./Configs/Hibbett/sizes.txt", "r") as getSizes:
                sizes = getSizes.read().split("\n")
                
            if jsondata['Sku'].upper() in str(skus).upper() and jsondata['Size'] in str(sizes):
                self.sku = jsondata['Sku']
                globals.var = jsondata['Varient']
                globals.ee.emit('startFunc')
                self.size = "Joe Biden"
                self.detected = True
                Thread(target=self.clearMonitor).start()
                self.initateOneTap()
            else:
                pass
        
        def clearMonitor(self):
            time.sleep(15)
            self.detected = False
            globals.var = None

        def returnToken(self):
            return generate_string(150,300)
            if self.token:
                return self.token
            else:
                return "".join(random.choices(string.ascii_uppercase + string.digits, k=randint(100,150)))

        def monitorMonitor(self):

            printb("Monitor: INITIALIZED LOCAL MONITOR (RESTOCK)")

            while True and self.detected == False:

                with open("./Configs/Hibbett/skus.txt", "r") as getSkus:
                    skus = getSkus.read().split("\n")

                now1 = datetime.datetime.now(pytz.timezone("EST"))
                minute = now1.strftime("%M")

                if self.sessionEX != self.hour() + 1:

                    if minute == "15" or minute == "45" or minute == "30" and self.checkedOut == False:

                        self.checkForThread = 0

                        def runThreadLoop(sku):
                            self.monitorSku(sku)

                        for sku in skus:
                            Thread(target=runThreadLoop, args = (sku, )).start()

                        time.sleep(1)
                        continue

                    else:
                        time.sleep(1)
                        continue
                
                else:
                    print("TATE")
                    self.getCustomerId()

        def monitorSku(self, skuVal):

            skuVal = skuVal.upper()

            try:

                response = self.session.get(f'https://hibbett-mobileapi.prolific.io/ecommerce/products/{skuVal}', 
                    headers={
                        'Host': 'hibbett-mobileapi.prolific.io',
                        'X-PX-AUTHORIZATION': self.userAgentVal,
                        'X-PX-ORIGINAL-TOKEN': f'{self.userAgentVal}:{self.returnToken}',
                        'Accept': '*/*',
                        'version': '6.0.0',
                        'Accept-Language': 'en-US;q=1.0, es-US;q=0.9',
                        'platform': 'ios',
                        'User-Agent': self.userAgent,
                        'Connection': 'keep-alive',
                        'Content-Type': 'application/json; charset=utf-8',
                    }, proxy=self.proxy
                )
                
                if "captcha" not in response.text:
                    try:
                        totalSkus = response.json()['skus']
                        for sku in totalSkus:
                            if sku['isAvailable'] != False:

                                with open("./Configs/Hibbett/sizes.txt", "r") as getSizes:
                                    sizes = getSizes.read().split("\n")

                                if sku['size'] in str(sizes):
                                    self.sku = skuVal
                                    self.detected = True
                                    self.size = sku['size']
                                    globals.var = sku['id']
                                    self.initateOneTap()
                    except:
                        pass
            except:
                pass
        
        def waitingForOneTap(self):
            log("WAITING FOR MONITOR PING", self.thread)


            '''while globals.var == None:
                time.sleep(.1)
                continue
            '''

            @globals.ee.on('startFunc')
            def startTask():
                self.initateOneTap()
        
        def initateOneTap(self):
            
            #Starting the experimential thread task
            #Thread(target=startFull, args=(self,)).start()
            # ---------------------------------------------

            if self.initialSku:
                var = self.initialSku
            else:
                var = globals.var
            try:
                
                res = self.session.post(f"http://127.0.0.1:8082",
                    headers = {
                        'poptls-url': f"https://hibbett-mobileapi.prolific.io/ecommerce/cart/one_tap?cardSecurityCode={self.cvv_value}",
                        'poptls-proxy': self.proxy,
                        'Connection': 'keep-alive',
                        "Authorization": f"Bearer {self.sessionId}",
                        'Host': 'hibbett-mobileapi.prolific.io',
                        'Accept-Language': 'en-US;q=1.0',
                        'Content-Type': 'application/json; charset=utf-8',
                        'version': '6.2.1',
                        'platform': 'ios',
                        "User-Agent": self.userAgent,
                        'X-PX-AUTHORIZATION': self.userAgentVal,
                        'X-PX-ORIGINAL-TOKEN': f'{self.userAgentVal}:{"".join(random.choices(string.ascii_uppercase + string.digits, k=randint(100,150)))}',
                        'Accept': '*/*',
                        'x-api-key': '0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0',
                        'Accept-Encoding': 'br;q=1.0, gzip;q=0.9, deflate;q=0.8',
                        #another method for px is to rempove X-PX-ORIGINAL-TOKEN and change X-PX-AUTHORIZATION to 2, Expirmental is changing X-PX-AUTHORIZATION to 4 and remove X-PX-ORIGINAL-TOKEN
                    },
                    json={
                        "preferredBillingAddressId" : "main",
                        "preferredShippingAddressId" : "main",
                        "cartItems" : [
                            {
                            "quantity" : 1,
                            "personalizations" : [

                            ],
                            "product" : {
                                "id" : var
                            },
                            "customerId" : self.customerId,
                            "sku" : {
                                "id" : var
                            }
                            }
                        ],
                        "customerId" : self.customerId,
                        "preferredPaymentMethodId": self.paymentId,
                    }
                )
                if res.status_code == 200:
                    log("POSTED 1/3", self.thread)
                    globals.carts += 1
                    self.sessionId = res.json()['bmSessionToken']
                    self.cartId = res.json()['id']
                    self.submitPayment()
                elif res.status_code == 404:
                    log("ITEM OOS... RETRYING", self.thread)
                    if self.thread == 1:
                        self.cloudMonitor()
                    else:
                        self.waitingForOneTap()
                elif res.status_code == 400 and "basket" in res.text:
                    log("ITEM OOS", self.thread)
                    if self.thread == 1:
                        self.cloudMonitor()
                    else:
                        self.waitingForOneTap()
                elif res.status_code == 403:
                    log("PX BLOCK 1/3", self.thread)
                    self.getSession()
                else:
                    log(f"UNKNOWN BLOCK 1/3 {str(res.status_code) + res.text}", self.thread)
                    if self.thread == 1:
                        self.cloudMonitor()
                    else:
                        self.waitingForOneTap()
            except Exception as e:
                print(e)
                log("ERROR AT 1/3", self.thread)
                self.run_task()

        def submitPayment(self):
        
            try:
                res = self.session.put(f"http://127.0.0.1:8082",
                headers = {
                    'poptls-url': f"https://hibbett-mobileapi.prolific.io/ecommerce/cart/{self.cartId}/customer?",
                    'poptls-proxy': self.proxy,
                    'Accept': '*/*',
                    'Accept-Encoding': 'br;q=1.0, gzip;q=0.9, deflate;q=0.8',
                    'Accept-Language': 'en-US;q=1.0',
                    "Authorization": f"Bearer {self.sessionId}",
                    'Connection': 'keep-alive',
                    'Content-Type': 'application/json; charset=utf-8',
                    'Host': 'hibbett-mobileapi.prolific.io',
                    'platform': 'ios',
                    'version': '6.2.1',
                    'x-api-key': '0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0',
                    'X-PX-AUTHORIZATION': self.userAgentVal,
                    'X-PX-ORIGINAL-TOKEN': f'{self.userAgentVal}:{"".join(random.choices(string.ascii_uppercase + string.digits, k=randint(100,150)))}',
                    "User-Agent": self.userAgent
                },
                json={"email": self.email}, 
                )
            except:
                log("error at 2/3", self.thread)
                self.run_task()
            if res.status_code != 200:
                if self.thread == 1:
                    self.cloudMonitor()
                else:
                    self.waitingForOneTap()
            else:

                log("POSTED 2/3", self.thread)
                
                try:
                    res = self.session.post(f"http://127.0.0.1:8082",
                    headers = {
                        'poptls-url': f"https://hibbett-mobileapi.prolific.io/ecommerce/cart/{self.cartId}/place_order?cardSecurityCode={self.cvv_value}&customer={self.customerId}&phone=&oneTapCheckout=true&firstName=&optIn=false",
                        'poptls-proxy': self.proxy,
                        'Accept': '*/*',
                        'Accept-Encoding': 'br;q=1.0, gzip;q=0.9, deflate;q=0.8',
                        'Accept-Language': 'en-US;q=1.0',
                        "Authorization": f"Bearer {self.sessionId}",
                        'Connection': 'keep-alive',
                        'Content-Type': 'application/json; charset=utf-8',
                        'Host': 'hibbett-mobileapi.prolific.io',
                        'platform': 'ios',
                        'version': '6.2.1',
                        'x-api-key': '0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0',
                        'X-PX-AUTHORIZATION': self.data.userAgentVal,
                        'X-PX-ORIGINAL-TOKEN': f'{self.userAgentVal}:{"".join(random.choices(string.ascii_uppercase + string.digits, k=randint(100,150)))}',
                        "User-Agent": self.userAgent
                    }, 
                    )
                except:
                    log("error at 1/3", self.thread)
                    self.run_task()

                if "decline" in res.text:
                    log("CARD DECLINED", self.thread)
                    self.sendWebhook("Card Declined  😭", 13238272)
                    globals.declines += 1
                    self.checkedOut = True
                elif res.status_code == 200:
                    log("CHECKED OUT", self.thread)
                    self.sendWebhook("Successful Checkout  🥳", 2082611)
                    Thread(target=playSound, args=("checkout",)).start()
                    globals.checkouts += 1
                    self.checkedOut = True
                elif "sold out" in res.text:
                    log("OOS", self.thread)
                    self.sendWebhook("OOS While Processing 😕", 16291095)
                    globals.var = None
                    self.detected = False
                    if self.thread == 1:
                        if self.mode == "1":
                            self.monitorMonitor()
                        else:
                            self.cloudMonitor()
                    else:
                        self.waitingForOneTap()
                elif res.status_code == 403: 
                    log("PX BLOCKED", self.thread)
                    self.detected == False
                    globals.var = None
                    self.run_task()
                elif "security code" in res.text:
                    log("INVALID CVV", self.thread)
                    self.sendWebhook("Card Declined  😭", 13238272)
                    globals.declines += 1
                    self.checkedOut = True
                else:
                    log(f"OOS AT PROCCESSING {res.text}", self.thread)
                    globals.var = None
                    self.detected - False
                    if self.thread == 1:
                        if self.mode == "1":
                            self.monitorMonitor()
                        else:
                            self.cloudMonitor()
                    else:
                        self.waitingForOneTap()
            
        def sendWebhook(self, message, color):

            with open("./Configs/key.json", "r") as settings:

                settings = json.loads(settings.read())
                hook = settings["webhook"]

                requests.post(hook, 
                    json={
                        "content": None,
                        "embeds": [
                            {
                                "title": message,
                                "color": color,
                                "fields": [
                                    {
                                        "name": "SKU",
                                        "value": self.sku
                                    },
                                    {
                                        "name": "Profile",
                                        "value": "||" + self.email + "||"
                                    }
                                ]
                            }
                        ],
                        "username": "HibbettHaven",
                        "avatar_url": "https://i.ibb.co/Vmzt71J/Hibbett-Haven.png",
                        "attachments": []
                    }
                )

                requests.post('https://discord.com/api/webhooks/1041190976036819065/JNyP99DCpedegtqUHx6RKpDgO7yQVonSW3ydu20M7_qa42yVANYdOFLADQpxcqEv3Mej',
                    json={
                        "content": None,
                        "embeds": [
                            {
                                "title": message,
                                "color": color,
                                "fields": [
                                    {
                                        "name": "SKU",
                                        "value": self.sku
                                    }
                                ]
                            }
                        ],
                        "username": "HibbettHaven",
                        "avatar_url": "https://i.ibb.co/Vmzt71J/Hibbett-Haven.png",
                        "attachments": []
                    }
                )
            
    class HibbetGen():
        
        def __init__(self, thread, jsonData, uaBool):
            self.thread = thread
            self.profile = jsonData
            self.userAgent = f"Hibbett | CG/6.3.0 (com.hibbett.hibbett-sports; build:{str(randint(1,15000))}; iOS 16.0.2)"
            self.userAgentVal = uaBool
            self.abr = {
                "Alabama": "AL",
                "Alaska": "AK",
                "Arizona": "AZ",
                "Arkansas": "AR",
                "California": "CA",
                "Colorado": "CO",
                "Connecticut": "CT",
                "Delaware": "DE",
                "Florida": "FL",
                "Georgia": "GA",
                "Hawaii": "HI",
                "Idaho": "ID",
                "Illinois": "IL",
                "Indiana": "IN",
                "Iowa": "IA",
                "Kansas": "KS",
                "Kentucky": "KY",
                "Louisiana": "LA",
                "Maine": "ME",
                "Maryland": "MD",
                "Massachusetts": "MA",
                "Michigan": "MI",
                "Minnesota": "MN",
                "Mississippi": "MS",
                "Missouri": "MO",
                "Montana": "MT",
                "Nebraska": "NE",
                "Nevada": "NV",
                "New Hampshire": "NH",
                "New Jersey": "NJ",
                "New Mexico": "NM",
                "New York": "NY",
                "North Carolina": "NC",
                "North Dakota": "ND",
                "Ohio": "OH",
                "Oklahoma": "OK",
                "Oregon": "OR",
                "Pennsylvania": "PA",
                "Rhode Island": "RI",
                "South Carolina": "SC",
                "South Dakota": "SD",
                "Tennessee": "TN",
                "Texas": "TX",
                "Utah": "UT",
                "Vermont": "VT",
                "Virginia": "VA",
                "Washington": "WA",
                "West Virginia": "WV",
                "Wisconsin": "WI",
                "Wyoming": "WY",
                "District of Columbia": "DC",
                "American Samoa": "AS",
                "Guam": "GU",
                "Northern Mariana Islands": "MP",
                "Puerto Rico": "PR",
                "United States Minor Outlying Islands": "UM",
                "U.S. Virgin Islands": "VI",
            }

        def run_task(self):
            self.getSession()

        def getSession(self):
            try:
                with open("./Configs/proxies.txt", "r") as proxies:
                    proxy = random.choice(proxies.read().split("\n"))

                    if "hhproxies" in proxy:
                        proxy = proxy.replace("hhproxies", "")
                        proxy = decrypt(proxy)

                        with open("./Configs/key.json", "r") as keyFile:
                            keyFile = json.loads(keyFile.read())
                        
                        if keyFile["key"] in proxy:
                            proxy = proxy.replace("splitval", "")
                            proxy = proxy.replace(keyFile["key"], "")
                        else:
                            print("Stop using stolen proxies bozo, next time you will be banned")
                            time.sleep(2)
                            os._exit(0)

                    IP = proxy.split(":")[0]
                    Port = proxy.split(":")[1]
                    User = proxy.split(":")[2]
                    Pass = proxy.split(":")[3]
                    proxies = f"http://{str(User)}:{str(Pass)}@{str(IP)}:{str(Port)}"

                self.proxy = proxies
                self.getRealSession()

            except Exception as e:
                print(e)
                log("Error getting proxy", self.thread)
                self.getSession()
    
        def getRealSession(self):
            try:
                sessionVar = tls_client.Session(client_identifier="safari_ios_15_0")
                self.session = sessionVar
                self.generateAccount()
            except Exception as e:
                print(e)
                log("Error getting session", self.thread)
                self.getRealSession()

        def generateAccount(self):

            self.password = ''.join(random.choice(string.ascii_lowercase) for i in range(10)) + "!A2"

            res = self.session.post("https://hibbett-mobileapi.prolific.io/users/register?",
                headers={
                    "Accept": "/",
                    "Accept-Encoding": "br;q=1.0, gzip;q=0.9, deflate;q=0.8",
                    "Accept-Language": "en-US;q=1.0",
                    "Connection": "keep-alive",
                    "Content-Type": "application/json; charset=utf-8",
                    "Host": "hibbett-mobileapi.prolific.io",
                    "platform": "ios",
                    "User-Agent": self.userAgent,
                    "version": "6.2.1",
                    "x-api-key": "0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0",
                    "X-PX-AUTHORIZATION": "2"
                },
                json={
                    "phone" : "206" + str(random.randint(1000000,9999999)),
                    "password" : self.password,
                    "firstName" : self.profile['billingAddress']['name'].split(" ")[0],
                    "login" : self.profile['billingAddress']['email'],
                    "subscribeToEmail" : False,
                    "agreeToTerms" : True,
                    "email" : self.profile['billingAddress']['email'],
                    "lastName" : self.profile['billingAddress']['name'].split(" ")[1]
                },
                proxy=self.proxy
            )

            if "sessionId" in res.text:
                log("Initialized Account", self.thread)

                self.sessionId = res.json()['sessionId']
                self.customerId = res.json()['customerId']

                self.addAddress()
            elif "Please enter valid phone number" in res.text:
                log("SMS Error, Retrying...", self.thread)
                time.sleep(1)
                self.generateAccount()
            elif "Your email is invalid" in res.text:
                log("Email Address Invalid", self.thread)
            elif "captcha" in res.text:
                log("PerimiterX Block", self.thread)
                time.sleep(15)
                self.run_task()
            else:
                log("Error Generating Account", self.thread)

        def addAddress(self):
            
            res = self.session.post(f"https://hibbett-mobileapi.prolific.io/users/{self.customerId}/addresses",
                headers={
                    "Accept": "*/*",
                    "Accept-Encoding": "br;q=1.0, gzip;q=0.9, deflate;q=0.8",
                    "Accept-Language": "en-US;q=1.0",
                    "Connection": "keep-alive",
                    "Content-Type": "application/json; charset=utf-8",
                    "Authorization": f"Bearer {self.sessionId}",
                    "Host": "hibbett-mobileapi.prolific.io",
                    "platform": "ios",
                    "User-Agent": self.userAgent,
                    "version": "6.2.1",
                    "x-api-key": "0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0",
                    "X-PX-AUTHORIZATION": "2"
                },
                json={
                    "phone" : "206" + str(random.randint(1000000,9999999)),
                    "city" : self.profile['billingAddress']['city'],
                    "country" : "US",
                    "id" : "main",
                    "firstName" : self.profile['billingAddress']['name'].split(" ")[0],
                    "address1" : self.profile['billingAddress']['line1'],
                    "isPrimary" : True,
                    "zip" : self.profile['billingAddress']['postCode'],
                    "address2" : self.profile['billingAddress']['line2'],
                    "state" : self.abr[self.profile['billingAddress']['state']],
                    "lastName" : self.profile['billingAddress']['name'].split(" ")[1]
                },
                proxy=self.proxy
            )

            if res.status_code == 200:
                log("Added Address", self.thread)
                self.genNonce()
            elif "captcha" in res.text:
                log("PerimiterX Block", self.thread)
                time.sleep(15)
                self.addAddress()
            else:
                log("Error Adding Address", self.thread)

        def genNonce(self):

            res = self.session.get(f"https://hibbett-mobileapi.prolific.io/users/radial/nonce",
                headers={
                    "Accept": "*/*",
                    "Accept-Encoding": "br;q=1.0, gzip;q=0.9, deflate;q=0.8",
                    "Accept-Language": "en-US;q=1.0",
                    "Connection": "keep-alive",
                    "Content-Type": "application/json; charset=utf-8",
                    "Authorization": f"Bearer {self.sessionId}",
                    "Host": "hibbett-mobileapi.prolific.io",
                    "platform": "ios",
                    "User-Agent": self.userAgent,
                    "version": "6.2.1",
                    "x-api-key": "0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0",
                    "X-PX-AUTHORIZATION": "2"
                },
                proxy=self.proxy
            )

            if res.status_code == 200:
                self.nonce = res.json()['nonce']
                log("GOT NONCE", self.thread)
                self.getPan()
            elif res.status_code == 403:
                log("PX BLOCK, RETRYING", self.thread)
                time.sleep(15)
                self.genNonce()
            else:
                log("ERROR GETTING NONCE, RETRYING", self.thread)
                time.sleep(15)
                self.genNonce()

        def getPan(self):

            res = self.session.post(f"https://hostedpayments.radial.com/hosted-payments/pan/tokenize?access_token={self.nonce}",
                headers={
                    "Accept": "*/*",
                    "Accept-Encoding": "br;q=1.0, gzip;q=0.9, deflate;q=0.8",
                    "Accept-Language": "en-US;q=1.0",
                    "Connection": "keep-alive",
                    "Content-Type": "application/json; charset=utf-8",
                    "Authorization": f"Bearer {self.sessionId}",
                    "Host": "hibbett-mobileapi.prolific.io",
                    "platform": "ios",
                    "User-Agent": self.userAgent,
                    "version": "6.2.1",
                    "x-api-key": "0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0",
                    "X-PX-AUTHORIZATION": "2"
                },
                json={
                    "paymentAccountNumber": self.profile['paymentDetails']['cardNumber']
                },
                proxy=self.proxy
            )

            
            if res.status_code == 200:
                log("GOT PAN", self.thread)
                self.accountToken = res.json()['account_token']
                self.postPan()
            elif res.status_code == 403:
                log("PX BLOCK, RETRYING", self.thread)
                time.sleep(15)
                self.getPan()
            else:
                log("ERROR GETTING PAN, RETRYING", self.thread)
                time.sleep(15)
                self.getPan()

        def postPan(self):

            res = self.session.post(f"https://hostedpayments.radial.com/hosted-payments/encrypt/pancsc?access_token={self.nonce}",
                headers={
                    "Accept": "*/*",
                    "Accept-Encoding": "br;q=1.0, gzip;q=0.9, deflate;q=0.8",
                    "Accept-Language": "en-US;q=1.0",
                    "Connection": "keep-alive",
                    "Content-Type": "application/json; charset=utf-8",
                    "Authorization": f"Bearer {self.sessionId}",
                    "Host": "hibbett-mobileapi.prolific.io",
                    "platform": "ios",
                    "User-Agent": self.userAgent,
                    "version": "6.2.1",
                    "x-api-key": "0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0",
                    "X-PX-AUTHORIZATION": "2"
                },
                json={
                    "cardSecurityCode": self.profile['paymentDetails']['cardCvv'],
                    "paymentAccountNumber": self.profile['paymentDetails']['cardNumber']
                },
                proxy=self.proxy
            )

            if res.status_code == 200:
                log("POSTED PAN", self.thread)
                self.encryptedCard = res.json()['encryptedCardSecurityCode']
                self.addPayment()
            elif res.status_code == 403:
                log("PX BLOCK, RETRYING", self.thread)
                time.sleep(15)
                self.postPan()
            else:
                log("ERROR POSTING PAN, RETRYING", self.thread)
                time.sleep(15)
                self.postPan()
        
        def addPayment(self):

            if self.profile['paymentDetails']['cardType'] == "MasterCard":
                self.profileType = "Master Card"
            else:
                self.profileType = self.profile['paymentDetails']['cardType']

            res = self.session.post(f"https://hibbett-mobileapi.prolific.io/users/{self.customerId}/payment_methods?addressId=main",
                headers={
                    "Accept": "*/*",
                    "Accept-Encoding": "br;q=1.0, gzip;q=0.9, deflate;q=0.8",
                    "Accept-Language": "en-US;q=1.0",
                    "Connection": "keep-alive",
                    "Content-Type": "application/json; charset=utf-8",
                    "Authorization": f"Bearer {self.sessionId}",
                    "Host": "hibbett-mobileapi.prolific.io",
                    "platform": "ios",
                    "User-Agent": self.userAgent,
                    "version": "6.2.1",
                    "x-api-key": "0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0",
                    "X-PX-AUTHORIZATION": "2"
                },
                json={
                    "type":"CREDIT_CARD",
                    "paymentObject": {
                        "cardType": self.profileType,
                        "number": self.profile['paymentDetails']['cardNumber'],
                        "nameOnCard": self.profile['paymentDetails']['nameOnCard'],
                        "expirationMonth": int(self.profile['paymentDetails']['cardExpMonth'].replace("0","")),
                        "creditCardToken": self.accountToken,
                        "encryptedCVNValue": self.encryptedCard,
                        "nameOnCard": self.profile['paymentDetails']['nameOnCard'],
                        "expirationYear": int(self.profile['paymentDetails']['cardExpYear'])
                    }
                },
                proxy=self.proxy
            )

            if res.status_code == 200:
                log("Account Created Successfully", self.thread)
                f = open("./Configs/Hibbett/AccountGen/generatedAccounts.txt", "a")
                f.write(f"{self.profile['billingAddress']['email']}:{self.password}:{self.profile['paymentDetails']['cardNumber'][-4:]}:{self.profile['paymentDetails']['cardCvv']}\n")
                f.close()
            else:
                print(res.text)
                print(res.status_code)
                log("ERROR ADDING PAYMENT, RETRYING", self.thread)
                time.sleep(15)
                self.addPayment()

    def runHibbet(mode, threads, uaVal):
        new_class = HibbettTask(mode, threads, uaVal)
        new_class.run_task()
        
    def generateHibbett(jsonData, threads, uaVal):
        new_class = HibbetGen(threads, jsonData, uaVal)
        new_class.run_task()

    class AuthContext:
        
        def __init__(self, license):
            self.license = license

        def authFunc(self):

            returnVal = False

            try:
                res = requests.post(f"{globals.serverUrl}:4567/auth",
                    headers = {
                            "Content-Type": "application/json",
                            "Authorization": "Bearer pk_cqzp7m2w4zsl5jlx233swz3uj7prgkmc5lmzg205"
                        },
                    json = {
                        'license': self.license,
                        'hwid': ':'.join(re.findall('..', '%012x' % uuid.getnode())),
                        "type": "init"
                    }
                )
                
                if res.status_code == 200:
                    returnVal = True
                else:
                    printb("Error authenticating CLI")
            except:
                printb("Error authenticating CLI")

            return returnVal

        def nodeFunc(self):

            returnVal = False 
            globals.node = random.randint(10000, 99999)

            try:
                res = requests.post(f"{globals.serverUrl}:4567/api/v1/youlikemen",
                    headers = {},
                    json={
                        "checkVal": "add",
                        "license": self.license,
                        "node": globals.node
                    }
                )

                if res.status_code == 200:
                    returnVal = True
                    
            except:
                pass

            return returnVal

        def launchCLI(self):

            os.system('cls')

            printb(" _   _                       ")
            printb("| | | | __ ___   _____ _ __  ")
            printb("| |_| |/ _` \ \ / / _ \ '_ \ ")
            printb("|  _  | (_| |\ V /  __/ | | |")
            printb("|_| |_|\__,_| \_/ \___|_| |_|")
        
            printb("")
            printb(f"VERSION: {globals.version}")
            printb("")
            printb("1. TASKS")
            printb("2. PROFILES")
            printb("3. PROXIES")
            printb("")
            # Thread(target=playSound, args=("heheheha",)).start()
            inputbRes = inputb("...")

            if inputbRes == "1":

                printb("1. GENERAL RESTOCK LOCAL")
                printb("2. ALPHA RESTOCK CLOUD")
                printb("3. ACCOUNT CREATION")
                printb("")
                
                mode = inputb("Please enter mode...")

                if mode == "1" or mode == "2":

                    with open("./Configs/Hibbett/accounts.txt", "r") as tskCt:
                        taskCount = tskCt.read().split("\n")

                    if len(taskCount) <= 100:
                        threadVal = 1

                        printb("")
                        resInputVal = str(input("What Mode? 1, xx2xx, xx3xx, xxEXPxx, or INITIAL..."))
                        
                        uaValParam = None
                        if resInputVal == "1":
                            uaValParam = '2'
                        elif resInputVal == "2":
                            uaValParam = '3'
                        elif resInputVal == "3":
                            uaValParam = '4'
                        elif str(resInputVal.lower()) == 'i':
                            dropTime(int(input("What hour is the drop?")))
                            uaValParam = 'random'
                        else:
                            uaValParam = 'random'

                        try:
                            app = Flask(__name__)

                            @app.route("/")
                            def hello():
                                return "Hello World!"
                            logging.getLogger('werkzeug').disabled = True
                            cli = sys.modules['flask.cli']
                            cli.show_server_banner = lambda *x: None
                            
                            def runflask():
                                if __name__ == "__main__":
                                    
                                    app.run(debug=True, use_reloader=False, port=12347)
                            Thread(target=runflask).start()
                        except:
                            print("More than one instance open bozo")
                            os._exit(0)
                        
                            
                        for i in range(len(taskCount)):
                            if mode == "1":
                                Thread(target=runHibbet, args=("1", threadVal, uaValParam,)).start()
                            elif mode == "2":
                                Thread(target=runHibbet, args=("2", threadVal, uaValParam,)).start()

                            threadVal += 1
                    else:
                        printb("Task count must be less than 10")
                        time.sleep(2)
                        newAuth.launchCLI()
                else:
                    with open("./Configs/Hibbett/AccountGen/profiles.json", "r") as tskCt:
                        taskCount = json.loads(tskCt.read())

                    uaValParam = None

                    printb("")
                    resInputVal = inputb("Would you like to run Safe or Experimental? (S/E) ")

                    if resInputVal[0] == "s" or resInputVal[0] == "S":
                        uaValParam = False
                    else:
                        uaValParam = True

                    threadVal = 1
                    for profile in taskCount:
                        Thread(target=generateHibbett, args=(profile,threadVal,uaValParam,)).start()
                        threadVal += 1

            elif inputbRes == "4":
                with open("./Configs/proxies.txt", "r") as proxies:
                    proxy = proxies.read().split("\n")
                    if len(proxy) == 1:
                        printb(f"{len(proxy)} proxy loaded")
                    else:
                        printb(f"{len(proxy)} proxies loaded")
                time.sleep(2)
                newAuth.launchCLI()
            else:
                printb("Check config file")
                time.sleep(2)
                newAuth.launchCLI()

    with open("./Configs/key.json", "r") as checkForLicense:
        checkForLicense = json.loads(checkForLicense.read())
        if checkForLicense["key"] == "":
            print("Welcome to Hibbett Haven!")
            print("")
            licenseKey = inputb("Enter your license key...")
        else:
            licenseKey = checkForLicense["key"]

    printb("Validating...")
    newAuth = AuthContext(licenseKey)
    getRes = newAuth.authFunc()

    if getRes == True:
        with open("./Configs/key.json", "r") as keyFile:
            keyFile = json.loads(keyFile.read())
            keyFile["key"] = licenseKey
            with open("./Configs/key.json", "w") as keyFileWrite:
                keyFileWrite.write(json.dumps(keyFile))
                keyFileWrite.close()
                Thread(target=titleStats).start()
                newAuth.launchCLI()


try:
    Thread(target=runBot).start()
except Exception as e:
    print(e)




while True:
    time.sleep(100000)


