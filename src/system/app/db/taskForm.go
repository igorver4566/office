package db

import (
	"database/sql"
)

type TaskFormItems struct {
	Manager   []DevStruct `json:"manager"`
	Owner     []DevStruct `json:"owner"`
	Developer []DevStruct `json:"developer"`
}

type DevStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func selectByName(table string, d *sql.DB, where string) []DevStruct {
	var items []DevStruct
	var item DevStruct
	rows, err := d.Query("SELECT id, name FROM " + table + " " + where)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return items
}

func GetFormItems() interface{} {
	var items TaskFormItems
	d := Init()
	defer d.Close()
	items.Manager = selectByName("users", d, "WHERE role_id IN (1, 2)")
	items.Owner = selectByName("owner", d, "")
	items.Developer = selectByName("users", d, "WHERE role_id = 3")
	return items
}
