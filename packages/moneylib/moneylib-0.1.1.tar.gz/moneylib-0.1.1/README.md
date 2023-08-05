# money
Python library for easy deal with money and currencies. The main point is that any operations with money will be processed with the amount in the particular currency.
You cannot use an amount without currency, because it could confuse you. You no longer have to watch so that the amounts and currencies move together.

## Installation
`pip install moneylib` - in your environment.

## How to use?

Firstly, you have to specify needed currencies. Where first argument is name, second argument is a currency code, third argument is a decimal places.
You also can define prefix and suffix for your currency. Them will be used for building of string representation of `Money`.

```python
from money import Currency, update_currencies

UAH = Currency('UAH', 'uah', 2, suffix=' uah')
USD = Currency('USD', 'usd', 2, prefix='$')
BTC = Currency('BTC', 'btc', 8, suffix=' btc')
update_currencies([UAH, USD, BTC])
```

Now, you can use defined variables or also you can use `money.Currency.get` method.
Notice, if you try to get undefined `Currency`, will be raised `money.exception.CurrencyDoesNotExist` exception.
It you have two or more defined currencies with the same codes, will be raised 'money.exceptions.MultipleCurrenciesFound' exception.
For example:
```python
from money import Currency
from money.exceptions import CurrencyDoesNotExist

try:
    LTC = Currency.get('ltc')
except CurrencyDoesNotExist as e:
    print(e)
    
# output: Currency with code "ltc" does not exist
```

Keep in mind, that `Money` is a class which contains integer representation of amount and `Currency`.
You can apply `str` function to get the string representation of `Money` with the defined suffix and prefix; 
`float` for getting the amount without currency; `int` for getting an amount in minimum units.

```python
a = Money(95, USD)

str(a)
# output: $95
float(a)
# output: 95.0
int(a)
# output: 9500

``` 

If you need to define `Money` using minimum units, you can use `int_to_money` function.

```python
Money.int_to_money(1000, USD)
# output: $10
```

All arithmetic operations are available. But keep in mind, that you can use amounts with the same currencies only.
Otherwise will be raised `TypeError`.
Division is available between `Money` and `int` or `float` (to divide some amount between 2 persons) and between `Money` and `Money` (how many money parts can be divided).

```python
Money(10, UAH) + Money(10, BTC)
# output: TypeError

a = Money(100, USD)
b = Money(45, USD)

a + b
# output: $145
a - b
# output: $45
a * 2
# output: $190
a // 2
# output: $47
a % 2
# output: $1
a // b
# output: 1.0
a % b
# output: 45.0
```

And comparison operations:

```python
a = Money(100, USD)
b = Money(45, USD)

a > b
# output: True
a < b
# output: False
a >= b
# output: True
a <= b
# output: False
a == b
# output: False
a != b
# output: True
a is b
# output: False
```

And one more example, which shows a real case of using this library.

```python
from money import Money, Currency

USD = Currency('USD', 'usd', 2, prefix='$')

# Let's imaging that you're a multimillionaire
balance = Money(5000000, USD)

print("Balance: {}".format(balance))

moneylib
amount = Money(1000000, USD)
# And you should pay a fee of 2%.
fee = amount * 0.02  # 2%

while True:
    print("Withdraw {} with fee {}".format(amount, fee))
    if balance >= amount + fee:
        balance -= amount + fee
        print("Balance: {}".format(balance))
    else:
        print("Balance too small")
        print("Balance: {}".format(balance))
        break
        
# output:
# Balance: $5000000
# Withdraw $1000000 with fee $20000
# Balance: $3980000
# Withdraw $1000000 with fee $20000
# Balance: $2960000
# Withdraw $1000000 with fee $20000
# Balance: $1940000
# Withdraw $1000000 with fee $20000
# Balance: $920000
# Withdraw $1000000 with fee $20000
# Balance too small
# Balance: $920000
```
