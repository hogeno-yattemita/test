package main

import (
	// "database/sql"
	// "context"

	"go_docker/handler"
	"net/http"
	// "go_docker/user/repository/mysql"
	// "github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
)

func main() {
	r := handler.NewRouter()
	http.ListenAndServe(":8080", r)

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/userdb1")
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	defer db.Close()

		// err = db.Ping()
		// if err != nil {
		// 	panic(err.Error())
		// }

		// rdb := redis.NewClient(&redis.Options{
		// 	Addr:     "redis:6379",
		// })

		// ctx := context.Background()
		// err = rdb.Set(ctx, "key", "value", 0).Err()
		// if err != nil {
		// 	panic(err)
		// }

		// userResolver := model.NewUserResolver(db)
		
		// users, err := userResolver.Users(ctx, struct{ Offset int }{Offset: 0})
		// if err != nil {
		// 	panic(err)
		// }
		// log.Println(users)


		// rows, err := db.Query("SELECT * FROM test")
		// if err != nil {
		// 	panic(err.Error())
		// }

		// defer rows.Close()

		// for rows.Next() {
		// 	var id int
		// 	var name string
		// 	err = rows.Scan(&id, &name)
		// 	if err != nil {
		// 		panic(err.Error())
		// 	}
		// }

	// 	userRepository := mysql.NewMysqlUserRepository(db)

	// 	user, err := userRepository.GetById(context.Background(), 1)


	// 	c.JSON(200, gin.H{
	// 		"message": user,
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

