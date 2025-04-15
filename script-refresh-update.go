package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func fetchKSQLStream(url string, payload []byte) (*bufio.Scanner, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/vnd.ksql.v1+json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(resp.Body), nil
}

func extractStringsFromStream(scanner *bufio.Scanner) ([]string, error) {
	var results []string

	for scanner.Scan() {
		line := scanner.Bytes()

		var record []interface{}
		if err := json.Unmarshal(line, &record); err != nil {
			continue
		}

		if len(record) > 0 {
			if str, ok := record[0].(string); ok {
				results = append(results, str)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func refreshMaterializedView(dbname string) error {
	const (
		host     = "10.140.56.139"
		user     = "root"
		password = "va*8sa*8sa*s8as*alue"
		port     = 5432
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=db_%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("erro ao conectar no banco %s: %w", dbname, err)
	}
	defer db.Close()

	_, err = db.Exec("REFRESH MATERIALIZED VIEW smarket_busca_produto;")
	if err != nil {
		return fmt.Errorf("erro ao executar REFRESH no banco %s: %w", dbname, err)
	}

	fmt.Printf("Materialized view atualizada com sucesso no banco: %s\n", dbname)
	return nil
}

func main() {
	url := "http://localhost:8088/query-stream"
	payload := []byte(`{
		"sql": "SELECT * FROM unique_customers_beta_hw EMIT;",
		"streamsProperties": {}
	}`)

	scanner, err := fetchKSQLStream(url, payload)
	if err != nil {
		panic(err)
	}

	dbNames, err := extractStringsFromStream(scanner)
	if err != nil {
		panic(err)
	}

	err1 := refreshMaterializedView(dbNames[0])
	if err1 != nil {
		fmt.Println("Erro:", err1)
	}

	for _, dbname := range dbNames {
		err := refreshMaterializedView(dbname)
		if err != nil {
			fmt.Println("Erro:", err)
		}
	}

}
