package main

import(
	"github.com/jmoiron/sqlx"
)

func getShipmentMap(transactions []int64) map[int64]Shipping{
	sql, params, err := sqlx.In("SELECT * FROM `shippings` WHERE `transaction_evidence_id` IN (?)", transactions)
	if err != nil {
		panic(err)
	}

	var shippings []Shipping
	err = dbx.Select(&shippings, sql, params...)
	if err != nil {
		panic(err)
	}

	ret := make(map[int64]Shipping)
	for i := range shippings{
		ret[shippings[i].TransactionEvidenceID] = shippings[i]
	}
	return ret
}
