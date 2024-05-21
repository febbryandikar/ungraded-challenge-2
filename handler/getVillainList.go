package handler

import (
	"database/sql"
	"fmt"

	"ungraded-challenge-2/entity"
)

func GetVillainList(db *sql.DB) []entity.Villain {
	var villains []entity.Villain
	
		rows, err := db.Query("SELECT id, name, universe, image_url FROM villain")
		if err!= nil {
			fmt.Println(err.Error())
			return nil
		}
		defer rows.Close()
	
		for rows.Next() {
			var villain entity.Villain
			err := rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL)
			if err!= nil {
				fmt.Println(err.Error())
				return nil
			}
			villains = append(villains, villain)
		}
	
		return villains
}