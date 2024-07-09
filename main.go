package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id_user int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
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
		rows, err := db.Query("SELECT id_user, email, name FROM users")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.Id_user, &user.Email, &user.Name); err != nil {
				return c.Status(500).SendString(err.Error())
			}
			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(users)
	})

	// Rota para inserir um novo usuário na tabela 'users'
	app.Post("/users", func(c *fiber.Ctx) error {
		var user User

		// Parse do corpo da requisição
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		if user.Name == "" || user.Email == "" {
			return c.Status(400).SendString("Nome e Email não podem estar vazios!")
		}

		// Inserir um novo registro na tabela 'users'
		result, err := db.Exec("INSERT INTO users (name, email) VALUES (?,?)", user.Name, user.Email)
		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		id, err := result.LastInsertId()

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		user.Id_user = int(id)

		return c.Status(201).JSON(user)
	})

	log.Fatal(app.Listen(":3000"))
}
