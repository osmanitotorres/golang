package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID_USER int    `json:"id"`
	Email   string `json:"email"`
}

func main() {
	app := fiber.New()

	// Conectar ao banco de dados MySQL
	db, err := sql.Open("mysql", "osmanito:Lrfg@2024@tcp(127.0.0.1:3306)/banco2")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar a conexão com o banco de dados
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Rota para buscar dados da tabela 'users'
	app.Get("/users", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id_user, email FROM users")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID_USER, &user.Email); err != nil {
				return c.Status(500).SendString(err.Error())
			}
			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(users)
	})

	log.Fatal(app.Listen(":3000"))
}
