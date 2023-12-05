package misc

import (
	"fmt"
	"sort"

	"github.com/rs/zerolog/log"
)

/*
Implement a banking system.
Level 1:
- Implement the ability to create an account.
  - "CREATE_ACCOUNT" should create an account in the system
	  - If an account already exists return "false", otherwise "true"
		signature: (timeStamp, accountId INT)
	- "DEPOSIT" should deposit the specified amount in an account
		signature: (timeStamp, accountId, amount INT)
	  - If account does not exist, return ""
		- If it does, return the update balance e.g. "100"
	- "PAY" should debit the specified amount from an account
		signature: (timeStamp, accountId, amount INT)
		- If account does not exist, return ""
		- If it does, return the update balance e.g. "100"

Level 2:
- Implement the ability to fetch top activity. Top activity here is
  the total of all the transactions for the account (deposit and pay).
	"TOP_ACTIVITY" Given n, print n accounts with total transactions in ascending order.
	- If two accounts have the same total transaction amount, sort by
	  account ID in descending order first
	- If less than n accounts are found, then print all with previous
	  logic
	signature: (timeStamp, n INT)
	The output is in string format as this "[101(500), 102(200), account3(200)]"

Level 3:
- Allow transfers: Amount is withdrawn from an existing account and added
   to another existing account
	 - The given amount of money should be withdrawn from source and held
	   until the transfer is accepted by the destination. If the transfer
		 expires, the money is put back to the source account.
	- retun empty if source and destination are same
	- return empty if source or destination does not exist
	- return empty if source does not have enough funds
	- transfer expires after 24h or (24 * 60 * 60 * 1000) = 86400000 milliseconds
	- a transfer expires at the next millisecond after the expiration period ends
	- a valid transfer return "transfer(ordinal number of the transfer)"
	- the transaction history for source and target accounts is only updated when
	  the transfer is accepted
	- transfers count towards the transactions of both source and destination
	signature: (timeStamp, sourceAccountId, dstAccountId, amount INT)

	Also implement, "ACCEPT_TRANSFER",
	- return "true" if trnafer was successfully accepted or "false"
	- return "false" if transfer with transferId does not exist or was already accepted
	  or expired
	- return "false" if the given accountId was not the target account for the transfer
*/

type Account struct {
	AccountID    int
	Balance      int
	LastUpdated  int
	Created      int
	Transactions []int
}

// Important Note on the reason for using *Account
// It boils down to the fact that when using a map
// as an embedded field (in this case we're assiging
// our map to the Accounts field) of the struct, we
// either:
// User non-ptr but first get, modify and re-assign
// (e.g. )
// or simply use pointer.
// https://stackoverflow.com/questions/17438253/accessing-struct-fields-inside-a-map-value-without-copying
// https://stackoverflow.com/questions/42605337/cannot-assign-to-struct-field-in-a-map
type Bank struct {
	Accounts   map[int]Account
	Transfers  map[string]Transfer
	TransferId int
}

type Transfer struct {
	Status string
	Query  []int
}

func NewAccount(timeStamp, accountId int) Account {
	return Account{
		AccountID:    accountId,
		Balance:      0,
		LastUpdated:  timeStamp,
		Created:      timeStamp,
		Transactions: []int{},
	}
}

func NewBank() Bank {
	log.Info().
		Msg("Creating new bank")
	return Bank{
		Accounts:   map[int]Account{},
		Transfers:  map[string]Transfer{},
		TransferId: 1,
	}
}

func (b *Bank) CreateAccount(timeStamp, accountId int) string {
	log.Info().
		Int("timeStamp", timeStamp).
		Int("accountId", accountId).
		Msg("Creating new bank account")
	if _, found := b.Accounts[accountId]; found {
		return "false"
	}
	acc := NewAccount(timeStamp, accountId)
	b.Accounts[accountId] = acc
	return "true"
}

func (b *Bank) Deposit(timeStamp, accountId, amount int) string {
	log.Info().
		Int("timeStamp", timeStamp).
		Int("accountId", accountId).
		Int("amount", amount).
		Msg("Depositing request")
	if _, found := b.Accounts[accountId]; !found {
		log.Info().
			Msg("Returning empty string")
		return ""
	}
	// To do this, use *Account in the Bank struct
	// b.Accounts[accountId].Balance += amount
	// But for now, we first get the map, modify it
	// and then re-assign it
	acc := b.Accounts[accountId]
	acc.Balance += amount
	acc.LastUpdated = timeStamp
	acc.Transactions = append(acc.Transactions, amount)
	b.Accounts[accountId] = acc
	log.Info().
		Int("Balance", b.Accounts[accountId].Balance).
		Msg("Returning balance")
	return fmt.Sprintf("%d", b.Accounts[accountId].Balance)
}

func (b *Bank) Pay(timeStamp, accountId, amount int) string {
	log.Info().
		Int("timeStamp", timeStamp).
		Int("accountId", accountId).
		Int("amount", amount).
		Msg("Pay request")
	if _, found := b.Accounts[accountId]; !found {
		log.Info().
			Msg("Returning empty string")
		return ""
	}
	if b.Accounts[accountId].Balance < amount {
		log.Info().
			Msg("Returning empty string")
		return ""
	}
	acc := b.Accounts[accountId]
	acc.Balance -= amount
	acc.LastUpdated = timeStamp
	acc.Transactions = append(acc.Transactions, amount)
	b.Accounts[accountId] = acc
	log.Info().
		Int("Balance", b.Accounts[accountId].Balance).
		Msg("Returning balance")
	return fmt.Sprintf("%d", b.Accounts[accountId].Balance)
}

func (b *Bank) TopActivity(timeStamp, n int) string {
	// Gather accounts in a slice
	accounts := []Account{}
	for _, acc := range b.Accounts {
		accounts = append(accounts, acc)
	}

	// Sort the slice based on total transactions
	sort.Slice(accounts, func(i, j int) bool {
		// This is left for knowledge as it would have been the way
		// if a separate slice was maintained with only keys of map
		// But here we can just sort (in-place) our account slice
		// and use it for printing purpose further down
		//return b.Accounts[accounts[i]].TotalTransactions() < b.Accounts[accounts[j]].TotalTransactions()

		// Here the custom sorting "func" helps in a multi sort
		if accounts[i].TotalTransactions() == accounts[j].TotalTransactions() {
			return accounts[i].AccountID > accounts[j].AccountID
		} else {
			return accounts[i].TotalTransactions() < accounts[j].TotalTransactions()
		}
	})

	if n > len(accounts) {
		return FormatAccountForPrint(accounts[:])
	} else {
		return FormatAccountForPrint(accounts[:n])
	}
}

func (b *Bank) Transfer(timeStamp, sourceAccountId, targetAccountId, amount int) string {
	if sourceAccountId == targetAccountId {
		return ""
	}

	_, sourceFound := b.Accounts[sourceAccountId]
	_, targetFound := b.Accounts[targetAccountId]
	if !(sourceFound && targetFound) {
		return ""
	}

	if sourceAcc := b.Accounts[sourceAccountId]; sourceAcc.Balance < amount {
		return ""
	}

	txId := fmt.Sprintf("transfer%d", b.TransferId)
	b.TransferId++

	tx := Transfer{
		Status: "pending",
		Query:  []int{timeStamp, sourceAccountId, targetAccountId, amount},
	}

	b.Transfers[txId] = tx
	return txId
}

func (b *Bank) AcceptTransfer(timeStamp, accountId int, transferId string) string {
	var tx Transfer
	var found bool
	if tx, found = b.Transfers[transferId]; !found {
		return "false"
	}

	if tx.Status == "accepted" {
		return "false"
	}

	if timeStamp-tx.Query[0] > 86400000 {
		tx.Status = "expired"
		return "false"
	}

	// The account accepting this transfer is not the intended (target)
	// account
	if tx.Query[2] != accountId {
		return "false"
	}

	b.Pay(tx.Query[0], tx.Query[1], tx.Query[3])   // Deduct from source, ts should be when tx request created
	b.Deposit(timeStamp, tx.Query[2], tx.Query[3]) // Deposit to target, ts should be when tx request accepted

	tx.Status = "accepted"
	b.Transfers[transferId] = tx

	return "true"
}

func FormatAccountForPrint(accounts []Account) (accountDetails string) {
	accountDetails = "["
	for i := range accounts {
		accountDetails += fmt.Sprintf("%d(%d)", accounts[i].AccountID, accounts[i].TotalTransactions())
		if i < len(accounts)-1 {
			accountDetails += ", "
		}
	}
	accountDetails += "]"
	return
}

func (b *Bank) OperationParser(operation string, parameters []int) (result string) {
	log.Info().
		Str("operation", operation).
		Msg("Bank request")
	switch operation {
	case "CREATE_ACCOUNT":
		result = b.CreateAccount(parameters[0], parameters[1])
	case "DEPOSIT":
		result = b.Deposit(parameters[0], parameters[1], parameters[2])
	case "PAY":
		result = b.Pay(parameters[0], parameters[1], parameters[2])
	case "TOP_ACTIVITY":
		result = b.TopActivity(parameters[0], parameters[1])
	case "TRANSFER":
		result = b.Transfer(parameters[0], parameters[1], parameters[2], parameters[3])
	case "ACCEPT_TRANSFER":
		txId := fmt.Sprintf("transfer%d", parameters[2])
		result = b.AcceptTransfer(parameters[0], parameters[1], txId)
	}
	return
}

// This is a non-pointer reciever because the
// top activity sorter complains
func (a Account) TotalTransactions() (total int) {
	for _, transaction := range a.Transactions {
		total += transaction
	}
	return
}
