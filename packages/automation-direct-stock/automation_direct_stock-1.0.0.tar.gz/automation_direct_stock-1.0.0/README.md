# automation_direct_stock
**I am in no way associated with automationdirect.com**

A simple [python module](https://github.com/PhilipConte/automation_direct_stock) to easily check if the product at a given url on automationdirect.com is in stock

## Installation
**Requires Python 3 and pip**

From the PyPI:
```bash
pip install [--user] automation_direct_stock
```

From GitHub:
```bash
# Clone this repository
git clone https://github.com/PhilipConte/automation_direct_stock.git

# Go into the repository
cd automation_direct_stock

# install the module using pip
pip install [--user] .
```

## Usage
Use from the commandline or a python script (make sure to put the url in quotes)
```bash
#commandline
adstock "url_to_check"

#python script
from automation_direct_stock import isInStock
isInStock("url_to_check")
```
## Possible Return Values
```bash
#if in stock
'In Stock'

#if out of stock
'Out Of Stock'

#if not found or url invalid
'error'

#(commandline only) if invalid number of arguments passed
'bad arguments'
```