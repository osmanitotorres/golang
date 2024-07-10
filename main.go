package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
}

type User struct {
	Id_user int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
}

func main() {

	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo de configuração: ", err)
		return
	}

	defer configFile.Close()

	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		log.Fatal("Erro ao fazer parse do arquivo de configuração:", err)
		return
	}

	user := config.DBUsername
	password := config.DBPassword
	host := config.DBHost
	port := config.DBPort
	database := config.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	app := fiber.New()

	// Conectar ao banco de dados MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar a conexão com o banco de dados
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//=====  END POINT 00002 =======   Rota para consultar usuários
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

	//=====  END POINT 00001 =======   Rota para inserir um novo usuário na tabela 'users'
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
			return c.Status(500).SendString(err.Error())
		}

		id, err := result.LastInsertId()

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		user.Id_user = int(id)

		return c.Status(201).JSON(user)
	})

	//===== END POINT 00003	====  Rota para receber arquivo

	app.Post("/upload", func(c *fiber.Ctx) error {
		// Recebe o arquivo
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Falha ao receber o arquivo")
		}

		// Caminho onde o arquivo será salvo
		filePath := fmt.Sprintf("./uploads/%s", file.Filename)

		// Salva o arquivo no diretório
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Falha ao salvar o arquivo")
		}

		return c.SendString("Arquivo enviado com sucesso")
	})

	// Cria o diretório uploads se não existir
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err = os.Mkdir("./uploads", 0755)
		if err != nil {
			log.Fatalf("Falha ao criar os diretórios de upload: %v", err)
		}
	}

	//========= Valida CNPJ =========
	app.Get("/validacnpj", func(c *fiber.Ctx) error {
		//cnpj := c.Query("cnpj")
		queryParams := c.Context().QueryArgs()
		params := make(map[string]string)
		// Itera sobre todos os parâmetros de consulta e adiciona ao mapa
		queryParams.VisitAll(func(key, value []byte) {
			params[string(key)] = string(value)
		})

		return c.JSON(fiber.Map{
			"queryParams": params,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
