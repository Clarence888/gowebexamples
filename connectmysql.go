package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

//`database / sql`的 便捷软件包，用于查询各种SQL数据库 go附带
/*
调用 db.Query 执行 SQL 语句, 此方法会返回一个 Rows 作为查询的结果
通过 rows.Next() 迭代查询数据.
通过 rows.Scan() 读取每一行的值
调用 db.Close() 关闭查询*/
//MySQL数据库 go get -u github.com/go-sql-driver/mysql

func main() {

	//连接到MySQL数据库a
	//并不会立即建立一个数据库的网络连接, 也不会对数据库链接参数的合法性做检验, 它仅仅是初始化一个sql.DB对象. 当真正进行第一次数据库查询操作时, 此时才会真正建立网络连接;
	//返回的sql.DB对象是协程并发安全的. sql.DB的设计就是用来作为长连接使用的。不要频繁Open, Close
	db, err := sql.Open("mysql", "xxxx:xxxx@(xxxx:3306)/superstar?parseTime=true")
	if err == nil {

		//查找有没有user表

		createQuery := `
    CREATE TABLE if not exists users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
		_, err := db.Exec(createQuery)
		if err == nil {
			fmt.Println("创建成功")
		}

		//插入一条记录
		{
			username := "wwwww"
			password := "dddddddddan"
			createdAt := time.Now()
			insertSql := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`

			result, err := db.Exec(insertSql, username, password, createdAt)
			if err == nil {
				userID, err9 := result.LastInsertId()
				fmt.Println(userID, err9)
			} else {
				fmt.Printf("insert data error: %v\n", err)
				return
			}
		}
		//日常练习的时候 可以用{}来做每个作用域的区分 这样就不用注释代码了
		{
			//预处理插入
			stmt, _ := db.Prepare(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`)
			defer stmt.Close()
			ret, err := stmt.Exec("cadasdasdas", "xyssadsd", time.Now())
			fmt.Println(err)
			newInsertId, err2 := ret.LastInsertId()
			fmt.Println(newInsertId, err2)
		}

		//查询

		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)
		// Query the database and scan the values into out variables. Don't forget to check for errors.
		//单行查询 scan
		query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
		err1 := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
		if err1 == nil {
			fmt.Println(username)
		}

		rows, err2 := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
		defer rows.Close()
		/*
			每次db.Query操作后, 都建议调用rows.Close(). 因为 db.Query() 会从数据库连接池中获取一个连接, 这个底层连接在结果集(rows)未关闭前会被标记为处于繁忙状态。当遍历读到最后一条记录时，会发生一个内部EOF错误，自动调用rows.Close(),但如果提前退出循环，rows不会关闭，连接不会回到连接池中，连接也不会关闭, 则此连接会一直被占用. 因此通常我们使用 defer rows.Close() 来确保数据库连接可以正确放回到连接池中; 不过阅读源码发现rows.Close()操作是幂等操作，即一个幂等操作的特点是其任意多次执行所产生的影响均与一次执行的影响相同, 所以即便对已关闭的rows再执行close()也没关系.
		*/

		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		if err2 == nil {
			var users []user
			for rows.Next() {
				var u user
				err3 := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
				if err3 == nil {
					users = append(users, u)
				}
			}
			err4 := rows.Err() // check err
			if err4 == nil {
				fmt.Println(users)
			}
		}

		//删除
		//_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1) // check err
		if err != nil {
			log.Fatal(err)
		}
	}
	//err1 := db.Ping()
}

//延伸
/*
.预编译语句(Prepared Statement)
预编译语句(PreparedStatement)提供了诸多好处, 因此我们在开发中尽量使用它. 下面列出了使用预编译语句所提供的功能:

PreparedStatement 可以实现自定义参数的查询
PreparedStatement 通常来说, 比手动拼接字符串 SQL 语句高效.
PreparedStatement 可以防止SQL注入攻击
一般用Prepared Statements和Exec()完成INSERT, UPDATE, DELETE操作。

stmt, _ := dbw.Db.Prepare(`INSERT INTO user (name, age) VALUES (?, ?)`)
	defer stmt.Close()

	ret, err := stmt.Exec("xys", 23)
*/
