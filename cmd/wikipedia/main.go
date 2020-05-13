package wikipedia

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/donng/teemo/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
)


type Event struct {
	Time string
	Desc string
}

func Sync() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://zh.wikipedia.org", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("accept-language", "zh-CN")
	resp, _ := client.Do(request)
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	// daily news
	todayNews := doc.Find("#column-itn li").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})

	var todayHistory []Event
	// today in history, using slice due to that map is out of order
	timeSli := doc.Find("#column-otd dt").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	descSli := doc.Find("#column-otd dd").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	for i := 0; i < len(timeSli); i++ {
		todayHistory = append(todayHistory, Event{
			Time: timeSli[i],
			Desc: descSli[i],
		})
	}

	date := time.Now().Format("20060102")
	newsJson, _ := json.Marshal(todayNews)
	historyJson, _ := json.Marshal(todayHistory)
	fmt.Println(strings.Count(string(historyJson), ""))

	dbConf := setting.Setting.Mysql
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbConf.User, dbConf.Password, dbConf.DBName))
	if err != nil {
		panic(fmt.Sprintf("mysql connect error, err: %s", err))
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO wikipedia(date, news, history) VALUES (?, ?, ?)")
	if err != nil {
		panic(fmt.Sprintf("prepare error, err: %s", err))
	}
	defer stmt.Close()

	_, err = stmt.Exec(date, newsJson, historyJson)
	//todo 错误打印日志
	if err != nil {
		panic(fmt.Sprintf("insert error, err: %s", err))
	}
}
