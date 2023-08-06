[![Build Status](https://travis-ci.org/debonzi/smartflags-client.svg?branch=master)](https://travis-ci.org/debonzi/smartflags-client)
[![PyPI](https://img.shields.io/pypi/v/smartflags-client.svg)](https://github.com/debonzi/smartflags-client)
[![PyPI](https://img.shields.io/pypi/pyversions/smartflags-client.svg)](https://github.com/debonzi/smartflags-client)
[![Coverage Status](https://coveralls.io/repos/github/debonzi/smartflags-client/badge.svg)](https://coveralls.io/github/debonzi/smartflags-client)

# Smart Flags Client (Beta)

## About
Simple to use client to Smartflags System (https://smartflag.herokuapp.com)

```
# -*- encoding: utf-8 -*-
import time
import logging

import smartflags

formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
ch = logging.StreamHandler()
ch.setFormatter(formatter)

# Change to DEBUG to see response time.
smartflags.logger.setLevel(logging.INFO)
smartflags.logger.addHandler(ch)

smartflags.setup('{{environment.key}}')


@smartflags.enabled('send_new_letter')
def print_chat_enabled():
    print("Hello!! I am a chat bot ;)")


if __name__ == '__main__':
    while True:
        try:
            print_chat_enabled()
            time.sleep(0.5)
        except KeyboardInterrupt:
            print('Bye!!')
            break
```
