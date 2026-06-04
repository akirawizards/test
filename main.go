// 意図的に脆弱な Go サンプル。
// 検知対象: CWE-78 / CWE-89、および AbemaTV Medium Block の選択済み Wiz SAST ルール:
// - SQL Injection in Go Database Queries Using database/sql Package
// - OS Command Injection
// - OS Command Injection via exec.Command
// SAST / Policy as Code の検知用。本番や実運用では使わないこと。
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	if os.Getenv("RUN_INTENTIONALLY_VULNERABLE_DEMO") != "1" {
		log.Println("intentionally vulnerable demo is disabled")
		return
	}

	db, err := sql.Open("postgres", os.Getenv("DEMO_DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", vulnerablePingHandler)
	mux.HandleFunc("/run", vulnerableExecHandler)
	mux.HandleFunc("/user", vulnerableUserLookupHandler(db))

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}

func vulnerablePingHandler(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")

	// VULNERABLE: targets "OS Command Injection" by concatenating user input into a shell command.
	output, err := exec.Command("sh", "-c", "ping -c 1 "+host).CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func vulnerableExecHandler(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("cmd")
	target := r.URL.Query().Get("target")

	// VULNERABLE: targets "OS Command Injection via exec.Command" with a user-controlled executable.
	output, err := exec.Command(command, target).CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func vulnerableUserLookupHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		// VULNERABLE: targets "SQL Injection in Go Database Queries Using database/sql Package".
		query := fmt.Sprintf("SELECT id, name FROM users WHERE name = '%s'", name)
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		fmt.Fprintln(w, "query completed")
	}
}
