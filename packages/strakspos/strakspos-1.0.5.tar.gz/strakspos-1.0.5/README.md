![STRAKS-POS](strakspos/assets/repo_logo.png)

____________________________________

```
  Licence: Apache Licence
  Origin Author: Simon Volpert (https://github.com/simon-v/minipos/)
  Port Maintainer: STRAKS Developers (freeman@straks.tech)
  Language: Python
  Homepage: https://straks.tech/
```


STRAKS POS is a simple self-hosted point-of-sale server, intended for use by small merchants and brick-and-mortar stores, that can be operated from any device with a relatively modern web browser.

With STRAKS POS you can begin receiving STAK payments without exposing your funds to a third party, or even to your own cashiers at any point in the process, by simply giving it a list of STRAKS addresses or an extended public key to work with.

### Installation

To run STRAKS POS, you will need an Internet-connected computer with Python 3 and pip installed. Optionally git aswell.

The easiest way to install is using pip, note depending on your version of pip you may need to run pip3 instead of pip:
```
#Windows & Mac
pip install strakspos

#Linux
sudo pip3 install strakspos --install-option="--install-scripts=/usr/local/bin"
```

You can also install directly from the source code:

Using git
```
git clone https://github.com/straks/strakspos.git
cd strakspos
python install.py
```

If you do not have git installed use the following option
Using .zip:
```
download: https://github.com/straks/strakspos/archive/master.zip
unzip: strakspos.zip
cd strakspos-master
python install.py
```

### Setup

Once installed the server can be run in one of three ways. The first is a command the installiation will make avaiable in your terminal or command prompt named `strakspos`. The second is to run the package using python3 from the command line `python3 -m strakspos.main`. The third is to run the strakspos.py file `python3 strakspos.py` that is avaiable in the root directory of the source code downloaded from github or extracted from the .zip file. 

The first time you run `strakspos` it will prompt you to enter either STRAKS addresses or an extended public key:

```
Welcome To STRAKS Point Of Sale
To begin please choose to add either STRAKS addresses or an extended public key (xpub)
[1] STRAKS Addresses
[2] xpub
```
Choose either option. For the addresses copy and paste the STRAKS addresses you want to include in the system, press q when you are done. For the extended public key generate it from the STRAKS electrum wallet 'Wallet->Information->Master Public key' and insert it into the prompt. STRAKS POS will then create a configuration file `strakspos.cfg` in the $HOME/.strakspos folder. It will use this as the defualt location for its data directory. 

Additionally if you would like to supply a large amount of static STRAKS addresses, create an `address.list` file in the data directory that contains a newline seperated list of STRAKS addresses.

By default strakspos will look in the following directories in order for a strakspos.cfg file, if one is found it will use that as the default data directory.
* `$HOME/.strakspos`
* `$HOME/.config/strakspos`
* The current directory .

Additionally a data directory can be supplied as the first argument:

`strakspos /path/to/directory'`


Below are some of the other configuration options you may consider changing before starting the server:
* `currencies` - The fiat currencies you accept
* `taxrate` - The tax rate that can be optionally applied to payments
* `port` - The network port to host the point of sale server on (default: 8080)
* `label` - The label to display on the POS
* `allowed_ips` - The IP addresses that are allowed to access the site (default: 0.0.0.0 [all])


Finally, run the `strakspos`, and take note of the computer's IP address. You will use it to connect to the STRAKS POS server.

To obtain your computer's IP address use the following commands:
```
ifconfig / ipaddr - Mac/Linux
ipconfig - Windows
```

If you wish to run the STRAKS POS as a background task on Mac or Linux use the following, and take note of the process ID that is displayed.
`strakspos &`


A systemd service file has also been provided in the source code for use on linux machines.

### Usage

Navigate to the server's address and port from any device with a relatively modern browser.

In the request creation page, enter the amount you wish to charge. You can use the percent button to apply a sales tax or crypto discount to the amount you have entered. Press the green check mark button.

Have your buyer scan the resulting QR code with their STRAKS wallet and authorize the transaction. Alternatively, you can click/tap on the QR code to copy the request URI to clipboard and send it to your buyer in a text message.

Wait for the transaction to be detected by the system. Press the `Finish` button and hand the buyer his purchase (in any order).

To review your sales, use the built-in log browser, accessible from the triple bar button. You can view a daily, weekly, monthly and yearly summary, print it, or email it to yourself at the configured email address.


### Customization

If you would like to have a custom header and footer on your welcome and log pages, create files in your data directory named  `welcome_footer.html`, `log_header.html` and `log_footer.html` with the desired HTML content.

Any file placed in the data directory overrides its counterpart in the library directory. Images and other files that you want to be directly accessible to the web browser should be placed in the `assets` subdirectory.

### Setting up with a web server

STRAKS POS doesn't need a web server to run, as it provides its own. It also doesn't care what the URL used to access it looks like, and will happily serve its content on a bare IP address, a subdomain or a directory. However, if you wish it to be accessible on port 80, you will need to set up a reverse proxy. An example Nginx configuration, is provided below

```
server {
       server_name strakspos;


       location /strakspos {
        proxy_pass http://127.0.0.1:8080/;
        proxy_cache off;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass_header Server;
       }
}

```
