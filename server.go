package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// struct buat ngeformat data user
type User struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

func main() {
	db = connectDB()
	defer db.Close()

	// 1. ENDPOINT GET / users ( ambil data dari database )
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// query ke database
		rows, err := db.Query("SELECT id, nama, email FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// loop hasil query dan masukin ke slice
		var users []User
		for rows.Next() {
			var u User
			err := rows.Scan(&u.ID, &u.Nama, &u.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	// 2. ENDPOINT POST /Users/add ( tambah ke database )
	http.HandleFunc("/users/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// baca json  dari bodi
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		//insert ke database
		err = db.QueryRow(
			"INSERT INTO users (nama, email) VALUES ($1, $2) RETURNING id",
			newUser.Nama, newUser.Email,
		).Scan(&newUser.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	})

	fmt.Println(" Server jalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
