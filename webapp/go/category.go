package main

import(
	"fmt"
)

// key : categoryID
var categoryMap map[int]*Category = make(map[int]*Category)


func constructCategoryMap(){
	categories := []Category{}
	err := dbx.Select(&categories, "SELECT * FROM `categories`")
	if err != nil{
		fmt.Println("Error in constructCategoryMap",err)
	}

	for i := range categories {
		categoryMap[categories[i].ID] = &categories[i]
	}

	for _, category := range categories {
		if category.ParentID == 0 {
			continue
		}
		categoryMap[category.ID].ParentCategoryName = categoryMap[category.ParentID].CategoryName
	}
}
