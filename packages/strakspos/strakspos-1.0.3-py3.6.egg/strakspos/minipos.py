#!/usr/bin/env python3
# STRAKS POS - A self-hosted, STRAKS point-of-sale server
# Licence: Apache Licence
# Origin Author: Simon Volpert <simon@simonvolpert.com>
# Port Maintainer: STRAKS Developers
# Project page: https://github.com/straks/strakspos
# Consult the README file for usage instructions and other helpful hints

import os
import sys
from wsgiref.simple_server import make_server
import urllib.parse
import datetime
import qrcode
import io
import random
import base64
import threading

import strakspos.stak as stak # Local library file
import strakspos.sendmail as sendmail # Local library file



def home_screen():
	print('Welcome To STRAKS Point Of Sale')
	print('To begin please choose to add either STRAKS addresses or an extended public key (xpub)')
	
	choice_string = '[1] STRAKS Addresses\n[2] xpub\n'
	choice = input(choice_string)
	while choice != '1' and choice != '2':
		print('Invalid option')
		choice = input(choice_string)
	if choice == '1':
		inp = ''
		it = 1
		addresses = []


		print("Enter STRAKS addresses below, enter 'q' when you are done")
		while True:
			inp = input('address %s: '%it)
			if inp == 'q':
				break
			if not stak.validate_key(inp):
				print("Invalid address")
			else:
				addresses.append(inp)
				it +=1
			
	elif choice == '2':
		xpub = ''
		print("Enter extended public key from STRAKS electrum wallet")
		valid = False
		while not valid:
			res = input("xpub: ")
			if not stak.validate_key(res) or not res[:4] == 'xpub':
				print("Invalid xpub key")
			else:
				xpub = res
				valid = True

	print("\n-----------------------------------\n")

	path = os.path.join(os.path.expanduser('~'), '.strakspos')
	os.makedirs(path, exist_ok=True)
	print("Creating a strakspos configuration file in $HOME/.strakspos ")
	try:
		
		with open(os.path.join(lib_dir, 'strakspos.cfg.sample'), 'r') as f:
					lines = f.readlines()
					with open(os.path.join(path, 'strakspos.cfg'), 'w') as wf:
						for line in lines:
							wf.write(line)
							pass

						if choice == '1':
							print("Adding addresses into config")
							for addr in addresses:
								wf.write("address=%s\n"%addr)
						elif choice == '2':
							print("Adding xpub into config")
							wf.write("xpub=%s\n"%xpub)
						
	except Exception as e:
		print("Could not create straks.cfg file use sample at https://github.com/straks/strakspos/blob/master/strakspos/strakspos.cfg.sample")
		print("Got exception %s"%e)


usage = '''Usage: stakpos [DATA_DIRECTORY]
See the README file for more information.'''
config = {
	'addresses': [],
	'lock': {
		'@': None,
	},
	'cache': [],
}
binary_extensions = ['png', 'jpg', 'gif', 'woff']
mime_types = {
	'txt': 'text/plain',
	'css': 'text/css',
	'woff': 'application/font-woff',
	'js': 'text/javascript',
	'png': 'image/png',
	'jpg': 'image/jpeg',
	'gif': 'image/gif',
	'svg': 'image/svg+xml',
}
# Look for the directory containing the configuration files
lib_dir = os.path.dirname(os.path.abspath(__file__))
data_dir_locations = [
	os.path.join(os.path.expanduser('~'), '.strakspos'),
	os.path.join(os.path.expanduser('~'), '.config', 'strakspos'),
	lib_dir,
	os.getcwd()
]
if len(sys.argv) > 1:
	if sys.argv[1] == '-h' or sys.argv[1] == '--help':
		print(usage)
		sys.exit(0)
	else:
		data_dir_locations.insert(0, os.path.abspath(sys.argv[1]))
		if not os.path.isdir(data_dir_locations[0]):
			print('No such directory: ' + data_dir_locations[0])
for data_dir in data_dir_locations:
	try:
		os.chdir(data_dir)
	except (OSError, NotADirectoryError):
		continue
	if os.path.isfile('strakspos.cfg'):
		print('Using {dir} as data directory'.format(dir=data_dir))
		break
# Load the config file
try:
	with open('strakspos.cfg', 'r') as f:
		lines = f.readlines()

except (IOError, OSError, FileNotFoundError, PermissionError) as error:

	home_screen()

	try:
		path = os.path.join(os.path.expanduser('~'), '.strakspos')
		with open(path+'/strakspos.cfg', 'r') as f:
			lines = f.readlines()
	except:
		lines = []
		exit(98)


for line in lines:
	# Skip blank lines and comments
	if line.strip() == '' or line.startswith('#'):
		continue
	# Split to key and value pairs
	words = line.strip().split('=')
	key = words[0].strip()
	value = '='.join(words[1:]).strip()
	# Skip empty values
	if value == '':
		continue
	if key == 'address':
		config['addresses'].append(value)
	else:
		config[key] = value
# Read the auxillary address list, if present
try:
	with open('address.list', 'r') as f:
		lines = f.readlines()
except (IOError, OSError, FileNotFoundError, PermissionError):
	lines = []
for line in lines:
	_addr = line.strip()
	if _addr != '' and _addr not in config['addresses']:
		config['addresses'].append(_addr)

# Sanitize the config file
def cast_config_type(key, _type, default):
	try:
		config[key] = _type(config[key])
	except (KeyError, ValueError):
		config[key] = default

def clamp_config_value(_name, _min=None, _max=None, default=0):
	if _min is not None and config[_name] < _min:
		print('Invalid "{}" value, falling back to {}'.format(_name, default))
		config[_name] = default
	elif _max is not None and config[_name] > _max:
		print('Invalid "{}" value, falling back to {}'.format(_name, default))
		config[_name] = default

cast_config_type('taxrate', float, 0)
cast_config_type('port', int, 8080)
cast_config_type('propagation', int, 60)
clamp_config_value('propagation', 0, 100, 60)
cast_config_type('welcome_timeout', int, 120)
clamp_config_value('welcome_timeout', 0, None, 120)
cast_config_type('index', int, 0)
clamp_config_value('index', 0, None, 0)

def split_config_key(key, default):
	if key not in config:
		config[key] = [default]
	else:
		config[key] = config[key].split(',')

split_config_key('currencies', 'USD')
split_config_key('allowed_ips', '0.0.0.0')
# Prune meaningless values from configuration
config['allowed_ips'] = set(config['allowed_ips'])
if '127.0.0.1' in config['allowed_ips']:
	config['allowed_ips'].remove('127.0.0.1')

def pick_config_list(key, value_list):
	if key not in config:
		config[key] = value_list[0]
	else:
		if config[key] not in value_list:
			print("Invalid '{key}' value, falling back to '{default}'".format(key=key, default=value_list[0]))
			config[key] = value_list[0]

pick_config_list('unit', ['native', 'bits', 'cash', 'satoshi'])
pick_config_list('payment_return', ['request', 'welcome'])

if 'week_start' not in config or config['week_start'] == 'monday':
	config['week_start'] = 0
elif config['week_start'] == 'sunday':
	config['week_start'] = 1
else:
	print("Invalid 'week_start' value, falling back to 'monday'")
if 'label' not in config or config['label'] == '':
	config['label'] = 'STRAKS POS'
exchange_list = []
for e in stak.exchanges:
	exchange_list.append(e['name'])
if 'price_source' not in config or config['price_source'] not in exchange_list:
	print('Using default exchange rate source "{}"'.format(exchange_list[0]))
	config['price_source'] = exchange_list[0]
if 'custom_unit_satoshi' in config:
	config['custom_unit_satoshi'] = True if config['custom_unit_satoshi'].lower() in ['1', 'yes', 'on', 'true'] else False
config['auto_cents'] = True if 'auto_cents' in config and config['auto_cents'].lower() in ['1', 'yes', 'on', 'true'] else False
# Try to set up a custom block explorer
custom_explorer = None
try:
	custom_explorer = {
		'name': '.'.join(config['custom_explorer_url'].split('/')[2].split('.')[-2:]),
		'url': config['custom_explorer_url'],
		'tx_url': config['custom_tx_url'],
		'balance_key': config['custom_balance_key'],
		'confirmed_key': config['custom_confirmed_key'],
		'unconfirmed_key': config['custom_unconfirmed_key'],
		'last_tx_key': config['custom_last_tx_key'],
		'tx_time_key': config['custom_tx_time_key'],
		'tx_inputs_key': config['custom_tx_inputs_key'],
		'tx_in_double_spend_key': config['custom_tx_in_double_spend_key'],
		'tx_outputs_key': config['custom_tx_outputs_key'],
		'tx_out_value_key': config['custom_tx_out_value_key'],
		'tx_out_address_key': config['custom_tx_out_address_key'],
		'tx_double_spend_key': config['custom_tx_double_spend_key'],
		'tx_fee_key': config['custom_tx_fee_key'],
		'tx_size_key': config['custom_tx_size_key'],
		'unit_satoshi': config['custom_unit_satoshi'],
		'prefixes': config['custom_prefixes'],
	}
	for key in custom_explorer:
		if custom_explorer[key].lower() == 'none':
			custom_explorer[key] = None
	stak.explorers.insert(0, custom_explorer)
	custom_explorer = custom_explorer['name']
	print('Using custom explorer definition:', custom_explorer)
except KeyError as error:
	if str(error) != "'custom_explorer_url'":
		print('Missing key in custom explorer definition:', error)

def write_config_file():
	# Open the current config file
	try:
		with open(os.path.join(data_dir, 'strakspos.cfg'), 'r') as f:
			lines = f.readlines()
	except (IOError, FileNotFoundError):
		# Or the config file sample
		with open(os.path.join(lib_dir, 'strakspos.cfg.sample'), 'r') as f:
			lines = f.readlines()
	# Iterate through the config
	for key in config:
		if key in ['addresses', 'lock', 'cache']:
			continue
		# Cast types into correct format
		if type(config[key]) in [list, set]:
			value = ','.join(config[key])
		elif key == 'week_start':
			value = {1: 'sunday', 0: 'monday'}[config[key]]
		else:
			value = config[key]
		if value is True or value is False:
			value = 'yes' if value else 'no'
		elif value is None:
			value = 'none'
		# Seek for first matching set key
		found = False
		for k in ['{}=', '#{}=']:
			if not found:
				k = k.format(key)
				for i, line in enumerate(lines):
					if line.startswith(k):
						found = True
						break
		# Replace the matching line
		if found:
			lines[i] = '{}={}\n'.format(key, value)
		# Or append to the end of the config file
		else:
			lines.append('{}={}\n'.format(key, value))
	# Disable removed keys
	for i, line in enumerate(lines):
		if line.startswith('sightings='):
			lines[i] = '## This option has been depreciated\n#' + line
	# Write the config file
	try:
		with open(os.path.join(data_dir, 'strakspos.cfg'), 'w') as f:
			f.writelines(lines)
	except (IOError, OSError, PermissionError) as error:
		print('Could not write config file: {}'.format(error))
	# Write the cached address list to address.list
	address_list = config['addresses'].copy()
	for entry in config['lock'].values():
		if type(entry) is dict:
			address_list.append(entry['address'])
	try:
		with open(os.path.join(data_dir, 'address.list'), 'w') as f:
			for address in address_list:
				f.write(address + '\n')
	except (IOError, OSError, PermissionError) as error:
		print('Could not write address.list: {}'.format(error))


# Simple logging wrapper
def logger(message):
	print('{} {}'.format(datetime.datetime.now().isoformat().split('.')[0].replace('T', ' '), message))


# Utility wrapper function
def load_file(filename):
	file_mode = 'rb' if filename.split('.')[-1] in binary_extensions else 'r'
	try:
		if 'html' in filename.split('.')[-1]:
			src = open(os.path.join(data_dir + '/resources', filename), file_mode)
		else:
			src = open(os.path.join(data_dir + '/assets', filename), file_mode)
	except (IOError, OSError, FileNotFoundError, PermissionError):
		try:
			src = open(os.path.join(lib_dir + '/resources', filename), file_mode)
		except (IOError, OSError, FileNotFoundError, PermissionError):
			src = open(os.path.join(lib_dir, filename), file_mode)
	data = src.read()
	src.close()
	return data


# Cast amount into preferred units
def format_amount(amount):
	if config['unit'] in ['bits', 'cash']:
		return stak.bits(float(amount))
	elif config['unit'] == 'satoshi':
		return str(int(float(amount) * 100000000))
	else:
		return amount


# Create a payment request QR page
def create_invoice(parameters):
	logger('creating invoice')
	if 'currency' not in parameters:
		parameters['currency'] = config['currencies']
	currency = parameters['currency'][0]
	divider = 100 if config['auto_cents'] else 1
	fiat = float(parameters['amount'][0]) / divider
	if fiat <= 0.0:
		raise ValueError('Requested amount must be positive')
	# Check for address lock timeouts
	for k in list(config['lock']):
		if k != '@' and check_lock_timeout(k):
			logger('payment request {} timed out'.format(k))
			unlock_address(k)
	# Use the next available address
	if config['addresses'] == []:
		if 'xpub' in config:
			address = stak.generate_address(config['xpub'], config['index'])
			config['addresses'].append(address)
			config['index'] += 1
			write_config_file()
		else:
			return load_file('noaddrs.html')
	address = config['addresses'].pop(0)

	# Generate a lock tag
	tag = ''
	for i in range(7):
		tag += random.choice('0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ')
	# Lock the address
	lock_address(tag)
	request = config['lock'][tag]
	request['address'] = address
	# Get the exchange rate
	try:
		price = stak.get_price(currency, exchange=config['price_source'])
	except KeyboardInterrupt:
		raise
	except:
		print(sys.exc_info()[1])
		return load_file('timeout.html')
	# Calculate amount
	amount = stak.stak(fiat / price)
	if float(amount) > 20999950:
		raise ValueError('Requested amount is greater than logically possible')

	# Get current address state
	try:
		txid = stak.get_last_txid(address, explorer=custom_explorer)

	except KeyboardInterrupt:
		raise
	except Exception as e:
		unlock_address(tag)
		return load_file('timeout.html')
	request['seen_txids'] = [] if txid is None else [txid]
	request['amount'] = amount
	request['fiat'] = stak.fiat(fiat)
	request['currency'] = currency

	# Display invoice
	logger('new invoice {tag}: {amount} STAK ({fiat} {currency}) to {address}'.format(tag=tag, **request))
	label = urllib.parse.quote('%s ID:%s' % (config['label'], tag))
	data = 'straks:{addr}?amount={amt}&message={label}'.format(addr=address, amt=amount, label=label)
	image = qrcode.make(data, error_correction=qrcode.constants.ERROR_CORRECT_L)
	output = io.BytesIO()
	image.save(output)
	output = base64.b64encode(output.getvalue()).decode('UTF-8')
	filler = {
		'addr': address,
		'amt': format_amount(amount),
		'qr': output,
		'request': data,
		'fiat': stak.fiat(fiat),
		'cur': currency,
		'price': stak.fiat(price),
		'tag': tag,
		'return': config['payment_return'],
		'label': config['label'],
	}
	filler['token'] = 'STAK' if config['unit'] == 'native' else config['unit']
	page = load_file('invoice.html').format(**filler)
	return page


# API check if a payment was received
def check_payment(parameters):
	# Responses:
	# 0 - not yet received
	# 1 - payment detected (with txid)
	# 2 - payment request timed out
	# 3 - server connection error
	# 4 - client connection error
	# 5 - double spend detected (with txid)
	tag = parameters['id'][0]
	if tag not in config['lock']:
		return '2'
	# Update address lock
	if check_lock_timeout(tag):
		logger('payment request {} timed out'.format(tag))
		unlock_address(tag)
		return '2'
	lock_address(tag)
	# Check address state
	request = config['lock'][tag]
	address = request['address']
	amount = float(request['amount'])
	# No previously detected transaction
	if 'txid' not in request:
		try:
			txid = stak.get_last_txid(address)
		except KeyboardInterrupt:
			raise
		except:
			logger('Could not fetch address info: ' + str(sys.exc_info()[1]))
			return '3'
		# No new transactions
		if txid is None:
			return '0'
		# Previously seen transaction
		elif txid in request['seen_txids']:
			return '0'
		elif txid in config['cache']:
			return '0'
		# New transaction
		try:
			tx = stak.TxInfo(txid, explorer=stak.explorers[-1]['name'])
		except KeyboardInterrupt:
			raise
		except stak.TxNotFoundError:
			logger('Anomalous event: tx not found on reporting explorer')
			return '0' # TODO anomaly!
		except:
			logger('Could not fetch address info: ' + str(sys.exc_info()[1]))
			return '3'
		# Transaction is known-old
		if tx.time < request['ctime']:
			logger('Ignoring old tx {}'.format(txid))
			request['seen_txids'].append(txid)
			return '0'
		elif address not in tx.outputs:
			logger('Ignoring mis-addressed tx {}'.format(txid)) # TODO anomaly!
			request['seen_txids'].append(txid)
			return '0'
		# Wrong transaction amount
		elif tx.outputs[address] != amount:
			logger('Ignoring tx with wrong amount {}'.format(txid))
			request['seen_txids'].append(txid)
			return '0'
		# All checks passed make note of the transaction
		request['txid'] = txid
		logger('payment {} detected'.format(tag))
		# Propagation check needed, defer processing
		if config['propagation'] > 0:
			return '0'
		# Double spend check
		elif tx.double_spend:
			request['wait_confirm'] = True
			return '5 ' + txid
	# Previously detected transaction
	else:
		# Currently waiting for confirmation
		if 'wait_confirm' in request:
			try:
				tx = stak.TxInfo(txid)
			except KeyboardInterrupt:
				raise
			except:
				logger('Could not get transaction info', sys.exc_info()[1])
				return '3'
			if tx.confirmations == 0:
				if tx.double_spend:
					return '5 ' + txid
				else:
					return '0'
		# Not currently waiting for confirmation
		else:
			txid = request['txid']
			try:
				propagation, double_spend = stak.get_tx_propagation(txid, threshold=config['propagation'])
			except KeyboardInterrupt:
				raise
			except:
				logger('Could not get propagation information: ' + str(sys.exc_info()[1]))
				return '3'
			# Is double spend
			if double_spend:
				request['wait_confirm'] = True
				return '5 ' + txid
			# Low propagation
			elif propagation < config['propagation']:
				return '0'
	# Record the payment
	record_payment(tag)
	config['cache'].append(txid)
	unlock_address(tag)
	# Remove this address from future use if generated
	if 'xpub' in config:
		config['addresses'].remove(address)
		write_config_file()
	return '1 ' + txid


# Write the details of the payment to a log file
def record_payment(tag):
	logger('payment {} received'.format(tag))
	request = config['lock'][tag]
	log_dir = os.path.join(data_dir, 'logs')
	if not os.path.isdir(log_dir):
		os.mkdir(log_dir)
	logfile = os.path.join(log_dir, datetime.date.today().isoformat() + '.log')
	with log_lock:
		with open(logfile, 'a') as log:
			log.write('{date}	{address}	{amount}	{fiat} {currency}	{tag}	{txid}\n'.format(date=datetime.datetime.now().isoformat(), tag=tag, **request))


# Lock an address to prevent concurrent access
def lock_address(tag):
	if tag not in config['lock']:
		config['lock'][tag] = {
			'ctime': datetime.datetime.now()
		}
	config['lock'][tag]['time'] = datetime.datetime.now()


# Free address lock
def unlock_address(tag):
	if tag in config['lock']:
		config['addresses'].append(config['lock'][tag]['address'])
		del(config['lock'][tag])


# Check address lock timeout
def check_lock_timeout(tag):
	if tag not in config['lock']:
		return True
	delta = datetime.datetime.now() - config['lock'][tag]['time']
	if delta >= datetime.timedelta(seconds=60):
		return True
	return False


# Parse a log file and add its contents to the table
def read_log_file(filename, plaintext=False, txids=False):
	if plaintext and txids:
		raise RuntimeError('read_log_file: the "plaintext" and "txids" parameters are incompatible')
	table = [] if txids else ''
	totals = {}
	token = 'STAK' if config['unit'] == 'native' else config['unit']
	for currency in config['currencies']:
		totals[currency] = 0
	try:
		logfile = open(os.path.join(data_dir, filename), 'r')
	except:
		if sys.exc_info()[0] not in [IOError, OSError, FileNotFoundError, PermissionError]:
			print(sys.exc_info()[1])
		return totals, table
	try:
		for line in logfile.readlines():
			line = line.strip().split('	')
			while len(line) < 6:
				line.append('')
			date, address, amount, fiat, tag, txid = line
			amount = format_amount(amount)
			fiat, currency = fiat.split(' ')
			totals[currency] += float(fiat)
			date = date.split('.')[0].replace('T', ' ')
			if txids:
				if txid != '':
					table.append(txid)
			elif plaintext:
				table += '{date}  {amt}  {fiat} {cur}  {tag}\n  Address: {addr}\n  TxID: {txid}\n'.format(date=date, addr=address, amt=str(amount).rjust(17 + len(token)), fiat=str(fiat).rjust(15), cur=currency, tag=tag, txid=txid)
			else:
				table += '''<tr class="%STYLE%">
<td><a id="toggle%ROW%" href="javascript:toggleRow(%ROW%);">+</a></td>
<td>{date}</td><td>{fiat} {cur}</td><td>{tag}</td>
</tr>
<tr class="%STYLE% expand" id="row%ROW%" style="display: none;">

<td colspan="4"><strong>Address:</strong> <a href="https://straks.info/en/address/{addr}" class="address" target="_blank">{addr}</a><br>
<strong>Amount:</strong> <span>{amt} {token}</span><br>
<p><strong>TxID:</strong><span class="txid" style="margin-left:5px;"><a style="word-wrap: break-word;" href="https://straks.info/en/transaction/{txid}" target="_blank">{txid}</a></span></p></td></tr>\n'''.format(date=date, amt=amount, fiat=fiat, cur=currency, tag=tag, token=token, addr=address, txid=txid)
	except:
		print('Log file is corrupted: {file} ({error})'.format(file=filename, error=sys.exc_info()[1]))
		msg = 'The log file for {file} is corrupted!'.format(file=filename.split('/')[1].split('.')[0])
		if txids:
			pass
		elif plaintext:
			table += msg
		else:
			table += '<tr class="%STYLE%"><td colspan="5" class="error">' + msg + '</td></tr>'
	logfile.close()
	return totals, table


# Display a log of recent transactions
def show_logs(parameters, plaintext=False):
	if 'date' not in parameters:
		date = datetime.date.today().isoformat()
	else:
		date = parameters['date'][0]
	# Process the current and calculate next and previous date
	days = []
	# Day scope
	if len(date) == 10:
		d = datetime.datetime.strptime(date, '%Y-%m-%d')
		delta = datetime.timedelta(1)
		next_date = (d + delta).date().isoformat()
		prev_date = (d - delta).date().isoformat()
		tag_s = 'W'
		scope_s = '%s-W%02d' % d.isocalendar()[0:2]
		tag_m = 'M'
		scope_m = '%s-%s' % (d.year, str(d.month).zfill(2))
		tag_l = 'Y'
		scope_l = str(d.year)
		days = [date]
	# Week scope
	elif len(date) == 8:
		# Convert ISO week to Python date
		_year = int(date[0:4])
		_week = int(date[6:8])
		ref_date = datetime.date(_year, 1, 4)
		ref_week, ref_day = ref_date.isocalendar()[1:3]
		d = (ref_date + datetime.timedelta(days=1-ref_day, weeks=_week-ref_week))
		# Calculate offsets
		delta = datetime.timedelta(7)
		next_date = '%s-W%02d' % (d + delta).isocalendar()[0:2]
		prev_date = '%s-W%02d' % (d - delta).isocalendar()[0:2]
		tag_s = 'D'
		scope_s = (d + datetime.timedelta(3)).isoformat()
		tag_m = 'M'
		scope_m = '%s-%s' % (d.year, str(d.month).zfill(2))
		tag_l = 'Y'
		scope_l = str(d.year)
		# Populate date list
		for i in range(7):
			days.append((d + datetime.timedelta(i - config['week_start'])).isoformat())
	# Month scope
	elif len(date) == 7:
		d = datetime.datetime.strptime(date, '%Y-%m')
		if d.month == 12:
			year, month = d.year + 1, 1
		else:
			year, month = d.year, d.month + 1
		next_date = '%s-%s' % (year, str(month).zfill(2))
		if d.month == 1:
			year, month = d.year - 1, 12
		else:
			year, month = d.year, d.month - 1
		prev_date = '%s-%s' % (year, str(month).zfill(2))
		tag_s = 'D'
		scope_s = '%s-%s-15' % (d.year, str(d.month).zfill(2))
		tag_m = 'W'
		scope_m = '%s-W%02d' % (d + datetime.timedelta(15)).isocalendar()[0:2]
		tag_l = 'Y'
		scope_l = str(d.year)
		# Populate date list
		_date = datetime.date(d.year, d.month, 1)
		while _date.month == d.month:
			days.append(_date.isoformat())
			_date += datetime.timedelta(1)
	# Year scope
	elif len(date) == 4:
		d = datetime.datetime.strptime(date, '%Y')
		next_date = str(d.year + 1)
		prev_date = str(d.year - 1)
		tag_s = 'D'
		scope_s = '%s-06-15' % d.year
		tag_m = 'W'
		scope_m = '%s-W26' % d.year
		tag_l = 'M'
		scope_l = '%s-06' % d.year
		# Populate date list
		_date = datetime.date(d.year, 1, 1)
		while _date.year == d.year:
			days.append(_date.isoformat())
			_date += datetime.timedelta(1)
	else:
		raise ValueError
	# Create a transaction table and calculate totals
	if plaintext:
		page = '===== Summary for {date} ====='.format(date=date)
	else:
		page = load_file('logs.html')
	table = ''
	table_head = ''
	table_foot = ''
	summary = ''
	totals = {}
	# Compile transaction table and calculate date totals
	for _date in days:
		_totals, _table = read_log_file(os.path.join('logs', _date + '.log'), plaintext)
		table += _table
		for k in _totals.keys():
			if k in totals:
				totals[k] += _totals[k]
			else:
				totals[k] = _totals[k]
	for sign in totals.keys():
		if totals[sign] != 0:
			if plaintext:
				chunk = '{fiat} {cur}'
			else:
				chunk = '<span>{fiat} {cur}</span>'
			summary += chunk.format(fiat=stak.fiat(totals[sign]), cur=sign)
	# Format and return the logs page
	if table != '':
		if plaintext:
			token = '' if config['unit'] == 'native' else config['unit']
			table_head = '=== Date & Time ===||=== Amount ({token}) ===||== Amount (fiat) ==||== ID =='.format(token=token)
		else:
			table_head = '<table class="table" style="table-layout: fixed;"> <thead> <tr> <th style="width:5%;">#</th> <th style="width:45%;">Date</th> <th>Amount</th> <th>ID</th> </tr> </thead>\n'
			
			table_foot = '</table>\n'
	else:
		if plaintext:
			summary = 'No transactions.'
		else:
			summary = '<span>No transactions.</span>'
	if plaintext:
		return '\n'.join([page, summary, table_head, table])
	else:
		# Load print header and footer
		try:
			header = load_file('log_header.html')
		except (IOError, OSError, FileNotFoundError, PermissionError):
			header = ''
		try:
			footer = load_file('log_footer.html')
		except (IOError, OSError, FileNotFoundError, PermissionError):
			footer = ''
		row_count = 1
		style = 'odd'
		while '%ROW%' in table:
			table = table.replace('%ROW%', str(row_count), 3)
			table = table.replace('%STYLE%', style, 2)
			style = 'even' if style == 'odd' else 'odd'
			row_count += 1
		table = table_head + table + table_foot
		# Pack the above variables into a filler dict
		label = config['label']
		params = {}
		_names = ['date', 'prev_date', 'tag_s', 'scope_s', 'tag_m', 'scope_m', 'tag_l', 'scope_l', 'next_date', 'header', 'summary', 'table', 'footer', 'label']
		for n in _names:
			params[n] = locals()[n]
		return page.format(**params)


# Serve a static file or return a 404
def serve_static_file(request):
	status = '200 OK'
	headers = [('Content-type', 'text/html; charset=UTF-8')]
	# Handle specific content-types
	extension = request.split('.')[-1]
	if extension in mime_types:
		headers = [('Content-Type', mime_types[extension])]
	# Try to load the requested file
	try:
		page = load_file(os.path.join('assets', request))
	except:
		headers = [('Content-type', 'text/html; charset=UTF-8')]
		status = '404 Not Found'
		page = load_file('error.html').format(err=status)
		print(sys.exc_info()[1])
	return status, headers, page


def send_email(parameters, config):
	# Unset previous status
	config['lock']['@'] = None
	# Assemble message parts
	if 'date' not in parameters:
		date = datetime.date.today().isoformat()
	else:
		date = parameters['date'][0]
	listing = show_logs(parameters, plaintext=True)
	subject = '[STRAKS POS] Transaction listing for {}'.format(date)
	# Send and store exit status
	config['lock']['@'] = sendmail.send(config, config['email'], subject, listing)


# Main webapp function
def minipos(environ, start_response):
	headers = [('Content-type', 'text/html; charset=UTF-8')]
	status = '200 OK'
	page = ''
	filler = ()
	if 'HTTP_X_REAL_IP' in environ:
		ip_addr = environ['HTTP_X_REAL_IP']
	else:
		ip_addr = environ['REMOTE_ADDR']
	subnet = '.'.join(ip_addr.split('.')[0:3]) + '.0'
	if ip_addr != '127.0.0.1' and '0.0.0.0' not in config['allowed_ips'] and ip_addr not in config['allowed_ips'] and subnet not in config['allowed_ips']:
		status = '403 Not Allowed'
		page = load_file('error.html').format(err=status)
		start_response(status, headers)
		return [bytes(page, 'UTF-8')]
	request = environ['PATH_INFO'].lstrip('/').split('/')[-1]
	parameters = urllib.parse.parse_qs(environ['QUERY_STRING'])
	# Handle specific app pages
	if request == 'invoice':
		try:
			page = create_invoice(parameters)
		except ValueError:
			status = '303 See Other\nLocation: request'
			page = 'Redirecting...'
		except:
			if sys.exc_info()[0] is KeyError:
				print('Missing required GET argument: ' + str(sys.exc_info()[1]))
			else:
				print(sys.exc_info()[1])
			status = '400 Bad Request'
			page = load_file('error.html').format(err=status)
	elif request == 'check':
		if 'id' not in parameters:
			status = '400 Bad Request'
			page = load_file('error.html').format(err=status)
		else:
			tag = parameters['id'][0]
		# Welcome page JavaScript check
		if tag == '0':
			page = '2'
		# Email sending check
		elif tag == '@':
			sent = config['lock']['@']
			page = '-1' if sent is None else '1' if sent else '2'
		# Payment received check
		else:
			page = check_payment(parameters)
		headers = [('Content-type', 'text/plain')]
	elif request == 'cancel':
		try:
			tag = parameters['id'][0]
			if tag:
				logger('payment {} cancelled'.format(tag))
				unlock_address(tag)
		except:
			if sys.exc_info()[0] is KeyError:
				print('Missing required GET argument: ' + str(sys.exc_info()[1]))
			else:
				print(sys.exc_info()[1])
		status = '303 See Other\nLocation: {}'.format(config['payment_return'])
		page = 'Redirecting...'
	elif request == 'logs':
		try:
			page = show_logs(parameters)
		except:
			print(sys.exc_info()[1])
			status = '400 Bad Request'
			page = load_file('error.html').format(err=status)
	elif request == 'email':
		headers = [('Content-type', 'text/plain')]
		if 'email' not in config:
			print('Sendmail failed: Email address is not set')
			page = '0'
		else:
			# Send an email in a subthread
			subthread = threading.Thread(target=send_email, args=(parameters, config))
			subthread.setDaemon(True)
			subthread.start()
			page = '-1'
	elif request == 'welcome':
		try:
			footer = load_file('welcome_footer.html')
		except (IOError, OSError):
			footer = ''
		page = load_file('welcome.html').format(label=config['label'], welcome_footer=footer)
	elif request == 'request':
		if config['taxrate'] < 0:
			tax = 'Discount'
		else:
			tax = 'Tax'
		filler = {
			'currencies': repr(config['currencies']),
			'timeout': config['welcome_timeout'],
			'cur': config['currencies'][0],
			'tax': tax,
			'taxrate': config['taxrate'],
			'label': config['label'],
			'centkey': '00' if config['auto_cents'] else '.',
		}
		filler['currency_disabled'] = 'disabled' if len(config['currencies']) == 1 else ''
		filler['tax_disabled'] = 'disabled' if config['taxrate'] == 0 else ''
		page = load_file('request.html').format(**filler)
	# Redirect blank request to main page
	elif request == '':
		try:
			footer = load_file('welcome_footer.html')
		except (IOError, OSError):
			footer = ''
		page = load_file('welcome.html').format(label=config['label'], welcome_footer=footer)
	
	# Load non-generated files from disk
	if page == '':
		status, headers, page = serve_static_file(request)
	# Serve the page
	start_response(status, headers)
	if type(page) is bytes:
		return [page]
	return [bytes(page, 'UTF-8')]


# Populate txid cache from recent log entries
_today = datetime.datetime.today().strftime('logs/%Y-%m-%d.log')
config['cache'] += read_log_file(_today, txids=True)[1]
_yesterday = (datetime.datetime.today() - datetime.timedelta(1)).strftime('logs/%Y-%m-%d.log')
config['cache'] += read_log_file(_yesterday, txids=True)[1]


# Start the web server


def main():
	# Make sure xpub works
	if 'xpub' in config:
		print("Using xpub: %s"%config['xpub'])
		try:
			if not stak.validate_key(config['xpub']):
				print('xpub is invalid, address generation unavailable')
				del(config['xpub'])
		except ImportError:
			print('pycoin is not installed, address generation unavailable')
			del(config['xpub'])
	for addr in config['addresses']:
		if addr.startswith('xp'):
			print('Discarding extended key from address list')
			config['addresses'].remove(addr)
		elif 'xpub' in config and not stak.validate_key(addr):
			print('Discarding invalid address {}'.format(addr))
			config['addresses'].remove(addr)
	if config['addresses'] == [] and 'xpub' not in config:
		print('No receiving addresses available. Please add some receiving addresses or an extended public key to your config file.')
		sys.exit(2)
	httpd = make_server('', config['port'], minipos)
	print('Serving strakspos on port {}...'.format(config['port']))
	try:
		httpd.serve_forever()
	except KeyboardInterrupt:
		print('Server stopped.')

random.seed()
log_lock = threading.Lock()

main()