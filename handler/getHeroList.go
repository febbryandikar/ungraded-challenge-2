package handler

import (
	"database/sql"
	"fmt"

	"ungraded-challenge-2/entity"
)

func GetHeroList(db *sql.DB) []entity.Hero {
	var heroes []entity.Hero

	rows, err := db.Query("SELECT id, name, universe, skill, image_url FROM hero")
	if err!= nil {
        fmt.Println(err.Error())
        return nil
    }
	defer rows.Close()

	for rows.Next() {
		var hero entity.Hero
        err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL)
        if err!= nil {
            fmt.Println(err.Error())
            return nil
        }
        heroes = append(heroes, hero)
	}

	return heroes
}