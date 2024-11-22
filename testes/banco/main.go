package main

/*
BANCO DE DADOS DO EXEMPLO
+-------+--------------+------+-----+---------+----------------+
| Field | Type         | Null | Key | Default | Extra          |
+-------+--------------+------+-----+---------+----------------+
| id    | int          | NO   | PRI | NULL    | auto_increment |
| name  | varchar(100) | YES  |     | NULL    |                |
| price | float        | YES  |     | NULL    |                |
+-------+--------------+------+-----+---------+----------------+
*/
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

func main() {
	// Conectar ao banco de dados
	dsn := "root:123456@tcp(localhost:3306)/teste"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close() // Garantir que a conexão será fechada no final

	// Testar a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Conexão falhou:", err)
	}
	fmt.Println("Conexão bem-sucedida!")

	inserirRegistro(db, "Bicicleta", 1100.00)
	inserirRegistro(db, "Farol", 100.00)
	inserirRegistro(db, "Motocicleta", 125000.00)
	inserirRegistro(db, "Capacete", 300.00)

	removerRegistro(db, 2)

	listarRegistros(db)

}

// Realiza o insert de um registro
func inserirRegistro(db *sql.DB, name string, price float32) {
	query := "INSERT INTO products (name, price) VALUES (?, ?)"
	result, err := db.Exec(query, name, price)
	if err != nil {
		log.Fatal("Erro ao inserir registro:", err)
	}

	id, _ := result.LastInsertId()
	fmt.Printf("Registro inserido com sucesso! ID: %d\n", id)
}

// Deletando registros
func removerRegistro(db *sql.DB, id int) {
	query := "DELETE FROM products WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("Erro ao remover registro:", err)
	}

	linhasAfetadas, _ := result.RowsAffected()
	fmt.Printf("Registro removido com sucesso! Linhas afetadas: %d\n", linhasAfetadas)
}

// Listando registros
func listarRegistros(db *sql.DB) {
	query := "SELECT id, name, price FROM products"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Erro ao listar registros:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price float32
		err := rows.Scan(&id, &name, &price)
		if err != nil {
			log.Fatal("Erro ao ler registro:", err)
		}
		fmt.Printf("ID: %d | Nome: %s | Preço: %v\n", id, name, price)
	}
}
