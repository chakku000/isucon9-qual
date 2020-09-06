package main

import(
	"github.com/jmoiron/sqlx"
)


func get_transactions(items []int64) map[int64]TransactionEvidence {
	sql, params, err := sqlx.In("SELECT * FROM `transaction_evidences` WHERE `item_id` IN (?)", items)
	if err != nil{
		panic(err)
	}

	var transactions []TransactionEvidence
	err = dbx.Select(&transactions, sql, params...)
	if err != nil {
		panic(err)
	}

	ret := make(map[int64]TransactionEvidence)
	for i := range transactions{
		transaction := transactions[i]
		ret[transaction.ItemID] = transaction
	}
	return ret
}
