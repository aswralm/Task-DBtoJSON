

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/user/DIGITBUSH"
	
// )

// func main() {
// 	db, err := sql.Open("mysql", "root:aswralm@/db_miniecomerse")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	var result string
// 	err = db.QueryRow("SELECT 'Connected to MySQL server successfully!'").Scan(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(result)
// }


package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/user/DIGITBUSH/model/repository"
)

type Output struct {
	Orders   []model.Orders   `json:"Orders"`
	Products []model.Product `json:"Products"`
	Reviews  []model.Reviews  `json:"Reviews"`
}

func main() {
	db, err := sql.Open("mysql", "root:aswralm@/db_miniecomerse")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	orderList, err := model.GetListOrders(db)
	if err != nil {
		log.Fatal(err)
	}

	productList, err := model.GetListProduct(db)
	if err != nil {
		log.Fatal(err)
	}

	reviewList, err := model.GetListReview(db)
	if err != nil {
		log.Fatal(err)
	}

	output := Output{
		Orders:   orderList,
		Products: productList,
		Reviews:  reviewList,
	}

	dataJSON, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dataJSON))
}

