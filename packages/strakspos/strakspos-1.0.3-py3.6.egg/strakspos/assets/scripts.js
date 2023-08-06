/*
MiniPOS web front-end common page scripts
Author: Simon Volpert <simon@simonvolpert.com>
Project page: https://github.com/simon-v/minipos/
This program is free software, released under the Apache License, Version 2.0. See the LICENSE file for more information
*/

// HTTP request wrapper
function loadHTTP(url, callback) {
	var xmlhttp = new XMLHttpRequest();
	xmlhttp.onreadystatechange = function() {
		if ( xmlhttp.readyState == 4 && callback ) {
			if ( xmlhttp.status == 200 ) {
				callback(xmlhttp.responseText.trim());
			}
			else {
				callback("4");
			}
		}
	};
	xmlhttp.open("GET", url, true);
	xmlhttp.send();
}

// Check if JavaScript is enabled
function jsCheck() {
	loadHTTP('check?id=0', jsCheckReady);
}
function jsCheckReady(code) {
	if ( code == 2 ) {
		document.getElementById("begin").style.display = "block";
		document.getElementById("noscript").style.display = "none";
	}
}

// Form button controls
function textAppend(value) {
	document.getElementById("amountbox").value += value;
	validateInput();
	return_timer = 0;
}
function textRemove() {
	var field = document.getElementById("amountbox"),
	value = field.value;
	field.value = value.substring(0, value.length - 1);
	return_timer = 0;
}
function addCents() {
	if ( document.getElementById("cents").value == "00" ) {
		textAppend(0);
		textAppend(0); // To allow validation to trigger
	}
	else {
		var field = document.getElementById("amountbox"),
		value = field.value;
		if ( !~ value.indexOf(".") ) {
			field.value = value + ".";
		}
		validateInput();
	}
	return_timer = 0;
}

// Form input validation
function validateInput() {
	var field = document.getElementById("amountbox"),
	value = field.value;
	// Decimal number leading zero
	if ( value[0] == "." ) {
		field.value = "0" + value;
	}
	// Two decimal places
	else if ( ~ value.indexOf(".") && value.length - value.indexOf(".") > 3 ) {
		textRemove();
	}
	// Leading zeroes
	else if ( value.length > 1 && value[0] == "0" && value[1] != "." ) {
		field.value = value.substring(1, value.length);
	}
	// Numbers only
	else if ( !~ "1234567890.".indexOf(value[value.length - 1]) ) {
		field.value = value.substring(0, value.length - 1);
	}
}

// Add the tax to the entered amount
function addTax() {
	var field = document.getElementById("amountbox"),
	percents = document.getElementById("percents").innerHTML,
	tax = percents.substring(0, percents.length - 1) / 100.0 + 1;
	if ( document.getElementById("cents").value == "." ) {
		var decimals = 2;
	}
	else {
		var decimals = 0;
	}
	field.value = (field.value * tax).toFixed(decimals);
	if ( field.value.substring(field.value.length - 3, field.value.length) == ".00" ) {
		field.value = field.value.substring(0, field.value.length - 3);
	}
	return_timer = 0;
}

// Cycle currency sign
var cur = 1;
function cycleCurrency() {
	var button = document.getElementById("currency");
	button.value = currencies[cur++ % currencies.length];
	document.getElementById("currencybox").value = button.value;
	return_timer = 0;
}

// Turn cancel button into confirm button
function showConfirmButton(response) {
	// Split response to code and txid
	if ( response.length > 1 ) {
		var code = response[0];
		var txid = response.substring(2, response.length);
	}
	else {
		var code = response;
	}	

	switch ( code ) {
		// Waiting for payment
		case "0":
			setTimeout(checkPayment, 2000);
			break;
		// Payment detected
		case "1":
			document.getElementById("cancel").style.display = "none";
			document.getElementById("finish").style.display = "inline";
			swal_it("Payment received", "<a href='https://straks.info/en/transaction/" + txid + "/' target='_blank'>See it on the blockchain</a>", "success");
			break;
		// Payment request timed out
		case "2":
			swal_it("Timed out", "The payment request has timed out.", "error");
			break;
		// Server connection error
		case "3":
			swal_it("Connection error", "Server connection error.", "error");
			setTimeout(checkPayment, 2000);
			break;
		// Client connection error
		case "4":
			swal_it("Connection error", "Client connection error.", "error");
			setTimeout(checkPayment, 2000);
			break;
		// Double spend detected
		case "5":
			swal_it("Double spend", "Double spend detected: " + txid, "error");
			setTimeout(checkPayment, 30000);
		}
}

// Deprecated
// Display an informational popup dialog
function displayPopup(text, fade) {
	var popupBox = document.getElementById("popup");
	document.getElementById("popup_text").innerHTML = text;
	popupBox.style.display = "block";
	if ( fade ) {
		popupBox.style.animation = "fader 1.5s linear 1";
		popupBox.style.WebkitAnimation = "fader 1.5s linear 1";
		setTimeout(dismissPopup, 1501);
	}
}
function dismissPopup() {
	var popupBox = document.getElementById("popup");
	document.getElementById("popup_text").innerHTML = "";
	popupBox.style.display = "none";
	popupBox.style.animation = "initial";
	popupBox.style.WebkitAnimation = "initial";
}

// Check whether or not payment was made
function checkPayment() {
	if (document.getElementById("finish").style.display == "none") {
		loadHTTP("check?" + request_string, showConfirmButton);
	}
}

// Open the log page
function openLogs(date) {
	if ( date )  {
		window.open("logs?date=" + date, "_self", true);
	}
	else {
		window.open("logs", "_self", true);
	}
}

// Email the contents of a log page
function sendEmail(date) {
	if ( date ) {
		loadHTTP("email?date=" + date, checkEmailSent);
	}
	else {
		loadHTTP("email", checkEmailSent);
	}
}
function emailStatus() {
	loadHTTP("check?id=@", checkEmailSent);
}
function checkEmailSent(response) {
	var button = document.getElementById("email");
	switch ( response ) {
		case "0":
			swal_it('Email not configured', "Please configure email.", "error");
			break;
		case "1":
			button.value = "Email";
			swal_it("Email sent.", "An email has been sent.", "success");
			break;
		case "2":
			button.value = "Email";
			swal_it("Failed", "Email sending failed. See the server log for more information.", "error");
			break;
		// Replace the email button with a spinner
		case "-1":
			switch ( button.value ) {
				case "Email":
				case "·····":
					button.value = "·    ";
					break;
				case "·    ":
					button.value = "··   ";
					break;
				case "··   ":
					button.value = "···  ";
					break;
				case "···  ":
					button.value = "···· ";
					break;
				case "···· ":
					button.value = "·····";
			}
			setTimeout(emailStatus, 100);
	}
}

// Automatically return to the welcome page after a timeout
var return_timer = 0;
function returnTimer() {
	if ( welcome_timeout > 0 ) {
		return_timer++;
		if ( return_timer == welcome_timeout ) {
			window.open("welcome", "_self");
		}
		else {
			setTimeout(returnTimer, 1000);
		}
	}
}

// Copy-to-clipboard
function copy() {
	var field = document.getElementById("copy");

	// Only process if there is no current popup
	if ( document.getElementById("popup").style.display == "none" ) {
		try {
			field.style.display = "block";
			field.select();
			document.execCommand("copy");
			field.style.display = "none";
		}
		finally {
			//swal_it("Copied to clipboard.", true);
		}
	}
}

// Switch log viewer collapsed row display
function toggleRow(row) {
	var table_row = document.getElementById("row" + row),
	toggle = document.getElementById("toggle" + row);
	if ( ( table_row.computedStyle && table_row.computedStyle.display == "none" ) || getComputedStyle(table_row, null).display == "none" ) {
		table_row.style.display = "table-row";
		toggle.innerHTML = "&ndash;";
	}
	else {
		table_row.style.display = "none";
		toggle.innerHTML = "+";
	}
}
