package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	// "golang.org/x/tools/go/cfg"
)

//这个的查询都是单条sql语句，没有设计事务（当然也不需要上下文）。

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// 多行查询
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	// 在其他C语言里面经常使用sprintf来拼接字符串，当然这个也可以使用fmt.Sprintf
	// Caution: Don’t use string formatting functions such as fmt.Sprintf
	// to assemble an SQL statement! You could introduce an SQL injection risk.
	rows, err := db.Query("select * from album where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	// Err的接口文档写的很好了不注释了
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
	}
	return albums, nil
}

// 读取一条记录
func albumByID(id int64) (Album, error) {
	var album Album

	row := db.QueryRow("select * from album where id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumByID %d: no such album", id)
		}
		return album, fmt.Errorf("albumByID %d: %v", id, err)
	}
	return album, nil
}

// add data
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("insert into album (title,artist,price) values (?,?,?)",
		alb.Title, alb.Artist, alb.Price)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func main() {
	conf := mysql.Config{
		User:   "qzyDB",
		Passwd: "helloQzy",
		Net:    "tcp",
		Addr:   "120.24.178.74:3306",
		DBName: "recordings",
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	fmt.Println("begin select")
	// 有 bug 、、、
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found %d's \n%v", len(albums), albums)

	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	// add album
	albID, err := addAlbum(Album{
		Title:  "The modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v", albID)
}
