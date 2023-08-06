# -*- coding: utf-8 -*-
# @Time    : 2018/6/3 下午9:27
# @Author  : GuoXiaoMin
# @File    : AccountExample.py
# @Software: PyCharm
from sdk.src.account.Account import Account

# generate a new account
account = Account.new_account()

# export account
account_json = account.to_key(bytes("passphrase".encode()))
print(account_json)

# load account
account = Account.from_key(account_json, bytes("passphrase".encode()))
print(account.get_address_str())
print(account.get_private_key())
print(account.get_public_key())



