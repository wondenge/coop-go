<p align="center">
<img src="src/assets/img/CoopSDK.png" alt="Co-op Bank SDK" title="Co-op Bank SDK" />
</p>
<h1 align="center">Co-operative Bank of Kenya Golang SDK</h3>

[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/wondenge/coop-go?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/wondenge/coop-go)](https://goreportcard.com/report/wondenge/coop-go)
[![License](https://img.shields.io/badge/license-Apache-blue.svg)]

# Overview

Coo-operative Bank of Kenya offers simple and RESTful APIs that allow one to quickly integrate and manage payments in web or mobile applications.

This Golang SDK provides an avenue to intergrate with the APIs directly inside your Go application to tap into the continuous innovation at the bank.

# Installation

```bash
$  go get github.com/wondenge/coop-go
```

# Usage

```go
package connectapi

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	connect "github.com/wondenge/coop-go/gen/connect"
)

```

<h1 align="left">Authentication API</h1>
An Access Token will be used to access, authorize and authenticate your interactions with the Co-op Bank’s Open Banking APIs. It will be generated using a Consumer Key and a Consumer Secret linked to your account. The Access Token will expire after a specified duration of time, after which regeneration of the Access Token will be required. To generate or re-generate an Access Token, your sign-in username and password will be required.

It is very important to keep your API credentials safe, as they can be used to access your account, make changes to your account and carry out transactions. It is also important to note that once you generate a new set of API credentials you must update all your applications with the new API details for your applications to continue working. Note that the Sandbox and Live environments have unique Consumer Key and Consumer Secret which are NOT be interchangeable.

<p align="left">
<img src="src/assets/img/AccountBalance.png" alt="AccountBalance API"  width="200" />
</p>
<h1 align="left">1. AccountBalance API</h1>
The Account Balance Enquiry API enables you to enquire about your own Co-operative Bank accounts' balance as at now for the specified account number.

<p align="left">
<img src="src/assets/img/AccountFullStatement.png" alt="AccountFullStatement API" width="200" />
</p>
<h2 align="left">2. AccountFullStatement API</h2>
The Account Full Statement Enquiry API enables you to enquire about your own Co-operative Bank accounts' full statement for the specified account number, start date and end date.

<p align="left">
<img src="src/assets/img/AccountMiniStatement.png" alt="AccountMiniStatement API" width="200" />
</p>
<h2 align="left">3. AccountMiniStatement API</h2>
The Account Mini Statement Enquiry API enables you to enquire about your own Co-operative Bank accounts' Mini statement for the specified account number.

<p align="left">
<img src="src/assets/img/AccountTransactions.png" alt="AccountTransactions API"  width="200" />
</p>
<h2 align="left">4. AccountTransactions API</h2>
The Account Transactions Enquiry API enables you to enquire about your own Co-operative Bank accounts' latest transactions for the specified account number and number of transactions to be returned.

<p align="left">
<img src="src/assets/img/AccountValidation.png" alt="AccountValidation API"  width="200" />
</p>
<h2 align="left">5. AccountValidation API</h2>
The Account Validation Enquiry API enables you to validate a Co-operative Bank account number.

<p align="left">
<img src="src/assets/img/ExchangeRate.png" alt="ExchangeRate API"  width="200" />
</p>
<h2 align="left"> 6. ExchangeRate API</h2>
The Exchange Rate Enquiry API enables you to enquire about the current SPOT exchange rate for the specified currencies.

<p align="left">
<img src="src/assets/img/IFTAccountToAccount.png" alt="IFTAccountToAccount API"  width="200" />
</p>
<h2 align="left">7. IFTAccountToAccount API</h2>
The Internal Funds Transfer Account to Account API enables you to transfer funds from your own Co-operative Bank account to other Co-operative Bank account(s).

<p align="left">
<img src="src/assets/img/INSSimulation.png" alt="INSSimulation API"  width="200" />
</p>
<h2 align="left">8. INSSimulation API</h2>
The INSSimulation API gives instant notifications or alerts on account activities(Debits,Credits) to the customer so that they can reflect this in their accounting backend.

<p align="left">
<img src="src/assets/img/PesaLinkSendToAccount.png" alt="PesaLinkSendToAccount API"  width="200" />
</p>
<h2 align="left">9. PesaLinkSendToAccount API</h2>
The PesaLink Send to Account Funds Transfer API enables you to transfer funds from your own Co-operative Bank account to Bank account(s) in IPSL participating banks.

<p align="left">
<img src="src/assets/img/PesaLinkSendToPhone.png" alt="PesaLinkSendToPhone API" width="200" />
</p>
<h2 align="left">10. PesaLinkSendToPhone API</h2>
The PesaLink Send to Phone Funds Transfer API enables you to transfer funds from your own Co-operative Bank account to a Phone Number(s) linked to a Bank account in an IPSL participating bank.

<p align="left">
<img src="src/assets/img/SendToM-Pesa.png" alt="SendToM-Pesa API"  width="200" />
</p>
<h2 align="left">11. SendToM-Pesa API</h2>
The Send to M-Pesa Funds Transfer API enables you to transfer funds from your own Co-operative Bank account to an M-Pesa account recipient.

<p align="left">
<img src="src/assets/img/TransactionStatus.png" alt="TransactionStatus API"  width="200" />
</p>
<h2 align="left">12. TransactionStatus API</h2>
The Transaction Status Enquiry API enables you to enquire about the status of a previously requested transaction for the specified transaction message reference.
