package main

import (
	
)

func main() {
	/*username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", datasourceName)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to db successful")
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50),
		age INT
	)`)

	if err != nil {
		log.Fatal(err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name= ? AND age= ?", "zzz", 40).Scan(&count)

	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		fmt.Println("Record already exists")
		return
	}
	//Insert some data into table
	_, err = db.Exec("INSERT INTO users (name,age) VALUE (?,?)", "zzz", 40)
	if err != nil {
		log.Fatal(err)
	}

	//Retrieve data from table
	rows, err := db.Query("SELECT id,name,age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int

		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d,Name: %s,Age: %d\n", id, name, age)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}*/

}
