package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-collections/collections/set"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

type InfocEraDimensionsCacheModel struct {
	Id int `json:"id"`
	ProductId int `json:"product_id"`
	Concern string `json:"concern"`
	Dimension string `json:"dimension"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Values string `json:"values"`
	Merge int `json:"merge"`
}


type InfoValueDimension struct {
	Id int `json:"id"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Values []interface{} `json:"values"`
	Merge int `json:"merge"`
	Insert bool `json:"insert"`
}

type TaskJob struct {
	ProductId int `json:"product_id"`
	Concern string `json:"concern"`
	Dimension string `json:"dimension"`
}

type YAML struct {
	Mysql MySQLConfig `yaml:"global_db"`
	Path ProjectPath `yaml:"path"`
}

type ProjectPath struct {
	Project string `yaml:"project"`
	Log string `yaml:"log"`
	Data string `yaml:"data"`
}


type MySQLConfig struct {
	Name string `yaml:"name"`
	HOST string `yaml:"host"`
	USER string `yaml:"user"`
	PASSWORD string `yaml:"password"`
	PORT int `yaml:"port"`
}


func LoadConfig(ENVFILE string) *YAML {
	yamlFile, err := ioutil.ReadFile(ENVFILE)
	if err != nil {
		panic(err)
	}
	var yamlConfig YAML
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		panic(err)
	}
	return &yamlConfig
}


func ConDB(mysql *MySQLConfig) *sql.DB {
	//"root:@tcp(127.0.0.1:3306)/test?parseTime=true"
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		mysql.USER,
		mysql.PASSWORD,
		mysql.HOST,
		mysql.PORT,
		mysql.Name)
	SqlDB, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	err = SqlDB.Ping()
	if err != nil {
		panic(err)
	}
	SqlDB.SetMaxIdleConns(1)
	SqlDB.SetMaxOpenConns(10)
	return SqlDB
}


func is_exists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}


func QueryMaxId(SqlDB *sql.DB) (id int) {
	querySql := `SELECT
					 id
					FROM infoc_era_dimensions_cache
					WHERE merge != 0
					ORDER BY id DESC
					LIMIT 1`
	
	rows, err := SqlDB.Query(querySql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		return
	}
	return
}


func readIdFromFile(filePath string) (size int) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	size, err = strconv.Atoi(string(buf))
	if err != nil{
		return
	}
	return
}


func getMaxId(SqlDB *sql.DB, filePath string)(id int)  {
	if is_exists(filePath){
		id = readIdFromFile(filePath)
	}else {
		id = QueryMaxId(SqlDB)
		id -= 20
	}
	return
}


func RecodeMaxId(SqlDB *sql.DB, filePath string)  {
	id := QueryMaxId(SqlDB)
	fileHandle, err := os.Create(filePath)
	if err != nil{
		fmt.Println(err)
		return
	}
	_, err = fileHandle.WriteString(strconv.Itoa(id))
	if err != nil {
		fmt.Println(err)
	}
}

func FindTask(SqlDB *sql.DB, id int) (result []TaskJob, err error) {
	querySql := `SELECT
					 product_id, concern, dimension, count(1) as n
					FROM infoc_era_dimensions_cache
					WHERE id >=? AND merge = 0
					GROUP BY product_id, concern, dimension`
	
	rows, err := SqlDB.Query(querySql, id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next(){
		b := TaskJob{}
		var t int
		err = rows.Scan(&b.ProductId, &b.Concern, &b.Dimension, &t)
		if err != nil {
			return
		}
		result = append(result, b)
	}
	return
}

func mergeCache(yamlConfig *YAML)  {
	SqlDB := ConDB(&yamlConfig.Mysql)
	defer SqlDB.Close()
	MaxIdFilePath := path.Join(yamlConfig.Path.Data, "MaxId.txt")
	id := getMaxId(SqlDB, MaxIdFilePath)

	taskJob, err := FindTask(SqlDB, id)
	RecodeMaxId(SqlDB, MaxIdFilePath)
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	
	for _, task := range taskJob{
		wg.Add(1)
		go dealTask(task, SqlDB, &wg)
	}
	wg.Wait()
}

func dealTask(job TaskJob, db *sql.DB, wg *sync.WaitGroup)  {
	defer wg.Done()
	querySql := "SELECT id, start_date, end_date, `values`, `merge` FROM infoc_era_dimensions_cache WHERE " +
		"product_id=? AND concern=? " + "AND dimension=? ORDER BY start_date, end_date"
	var rows *sql.Rows
	rows, err := db.Query(querySql, job.ProductId, job.Concern, job.Dimension)
	if err != nil {
		return
	}
	defer rows.Close()
	var result []InfoValueDimension
	for rows.Next() {
		b := InfoValueDimension{}
		var t string
		err := rows.Scan(&b.Id, &b.StartDate, &b.EndDate, &t, &b.Merge)
		if err != nil {
			return
		}
		err = json.Unmarshal([]byte(t), &b.Values)
		if err != nil {
			return
		}
		result = append(result, b)
	}
	
	var tmp InfoValueDimension
	var deleteInfo []int
	var AddInfo []InfocEraDimensionsCacheModel
	for index , b := range result{
		if index == 0{
			tmp = b
		}else {
			if tmp.StartDate == b.StartDate && tmp.EndDate <= b.EndDate  && tmp.EndDate >= b.StartDate {
				deleteInfo = append(deleteInfo, tmp.Id)
				tmp = b
			}else if tmp.StartDate < b.StartDate && tmp.EndDate >= b.EndDate {
				deleteInfo = append(deleteInfo, b.Id)
			}else if tmp.StartDate <= b.StartDate && tmp.EndDate >= b.StartDate &&
				tmp.EndDate <= b.EndDate {
				deleteInfo = append(deleteInfo, b.Id)
				if !tmp.Insert{
					deleteInfo = append(deleteInfo, tmp.Id)
				}
				
				tmp = InfoValueDimension{
					StartDate: tmp.StartDate,
					EndDate: b.EndDate,
					Values: mergeValue(tmp.Values, b.Values),
					Merge: 1,
					Insert:true,
				}
			}else {
				AddCache(AddInfo, &tmp, &job)
				tmp = b
			}
		}
	}
	AddInfo = AddCache(AddInfo, &tmp, &job)
	updateCache(deleteInfo, AddInfo, db)
}


func AddCache(
	AddInfo []InfocEraDimensionsCacheModel,
	tmp *InfoValueDimension,
	job *TaskJob) ([]InfocEraDimensionsCacheModel) {
	if tmp.Insert {
		lang, err := json.Marshal(tmp.Values)
		if err != nil {
			return AddInfo
		}
		
		AddInfo = append(AddInfo, InfocEraDimensionsCacheModel{
			ProductId: job.ProductId,
			Concern: job.Concern,
			Dimension: job.Dimension,
			StartDate: tmp.StartDate,
			EndDate: tmp.EndDate,
			Values: string(lang),
			Merge: 1,
		})
	}
	return AddInfo
}


func mergeValue(a []interface{}, b []interface{}) (d []interface{}) {
	a1 := set.New(a...)
	b1 := set.New(b...)
	c := a1.Union(b1)
	c.Do(func(i interface{}) {
		d = append(d, i)
	})
	return
}


func updateCache(deleteInfo []int, AddInfo []InfocEraDimensionsCacheModel, db *sql.DB)  {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}else if err != nil {
			_ = tx.Rollback()
		}else {
			err = tx.Commit()
		}
	}()
	
	if len(deleteInfo) > 0{
		stringDeleteInfo := []string{}
		for _, ele := range deleteInfo {
			stringDeleteInfo = append(stringDeleteInfo, strconv.Itoa(ele))
		}
		
		deleteSql := "DELETE FROM infoc_era_dimensions_cache WHERE id in (" +
			strings.Join(stringDeleteInfo, ",") + ")"
		_, err := tx.Exec(deleteSql)
		if err != nil {
			return
		}
	}
	
	if len(AddInfo) > 0 {
		InsertSql := "INSERT INTO infoc_era_dimensions_cache (product_id, concern, dimension, " +
			"start_date, end_date, `values`, `merge`) VALUES "
		info := AddInfo[0]
		info.StartDate = info.StartDate[0:10]
		info.EndDate = info.EndDate[0:10]
		tmp := ("(" +
			strconv.Itoa(info.ProductId) + "," +
			"'" + info.Concern + "'"+ "," +
			"'" + info.Dimension + "'" + "," +
			"'" + info.StartDate + "'" + "," +
			"'" + info.EndDate + "'" + "," +
			"'" + info.Values + "'" + "," +
			strconv.Itoa(info.Merge) +
			")")
		InsertSql = InsertSql +  tmp
		
		for index, info := range AddInfo {
			if index == 0{
				continue
			}
			info.StartDate = info.StartDate[0:10]
			info.EndDate = info.EndDate[0:10]
			tmp := ("(" +
				strconv.Itoa(info.ProductId) + "," +
				"'" + info.Concern + "'" + "," +
				"'" + info.Dimension + "'" + "," +
				"'" + info.StartDate + "'" + "," +
				"'" + info.EndDate + "'" + "," +
				"'" + info.Values + "'" + "," +
				strconv.Itoa(info.Merge) +
				")")
			InsertSql = InsertSql + "," + tmp
		}
		_, err := tx.Exec(InsertSql)
		if err != nil {
			return
		}
	}
}

func main() {
	// ENVFILE := "/Users/zhou/Documents/workspace/infocera_conf.yaml"
	var ENVFILE string
	
	flag.StringVar(&ENVFILE, "envfile", "nothing","infocera config file")
	flag.Parse()
	if ENVFILE == "nothing"{
		panic(errors.New("配置文件必须配置"))
	}
	yamlConfig := LoadConfig(ENVFILE)
	LockFile := path.Join(yamlConfig.Path.Data, "merge_cache.lock")
	lock, e := os.Create(LockFile)
	if e != nil {
		log.Printf("main: 创建文件锁失败: %s", e)
		os.Exit(-1)
	}
	defer os.Remove(LockFile)
	defer lock.Close()
	
	// 尝试独占文件锁
	e = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if e != nil {
		log.Printf("main: 独占文件锁失败: %s", e)
		os.Exit(-1)
	}
	defer syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
	mergeCache(yamlConfig)
}
