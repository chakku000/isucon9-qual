package main

import(
	"github.com/jmoiron/sqlx"
	"log"
)


// return map[key of user_id] = UserSimple
func gen_sellerMap(items []Item) map[int64]UserSimple {
	ids := make([]int64, 0)
	for _, item := range items {
		ids = append(ids, item.SellerID, item.BuyerID)
	}
	sql, params , err := sqlx.In("SELECT * FROM `users` WHERE `id` IN (?)", ids)
	if err != nil {
		log.Println(err)
		return nil
	}

	var users []User
	err = dbx.Select(&users, sql, params...)
	if err != nil {
		log.Println(err)
		return nil
	}

	ret := make(map[int64]UserSimple)
	for i := range users {
		user := users[i]
		ret[user.ID] = UserSimple{
			ID : user.ID,
			AccountName: user.AccountName,
			NumSellItems: user.NumSellItems,
		}
	}
	return ret
}
