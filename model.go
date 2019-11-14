package management

import (
	"fmt"
	"net/url"
	"reflect"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/goextension/log"
	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
	"github.com/xormsharp/xorm"
)

const mysqlStatement = "%s:%s@tcp(%s)/%s?loc=%s&charset=%s&parseTime=true"
const createDatabase = "CREATE DATABASE `%s` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_bin'"

// Model ...
type model struct {
	ID        string     `xorm:"id pk"`
	CreatedAt time.Time  `xorm:"created_at created"`
	UpdatedAt time.Time  `xorm:"updated_at updated"`
	DeletedAt *time.Time `xorm:"deleted_at deleted"`
	Version   int        `xorm:"version"`
}

// Model ...
type Model struct {
	model `xorm:"extends"`
}

// DBConfig ...
type DBConfig struct {
	ShowSQL      bool   `json:"show_sql"`
	ShowExecTime bool   `json:"show_exec_time"`
	UseCache     bool   `json:"use_cache"`
	Create       bool   `json:"create"`
	DBType       string `json:"db_type"`
	Addr         string `json:"addr"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Schema       string `json:"schema"`
	Charset      string `json:"charset"`
	Prefix       string `json:"prefix"`
	Location     string `json:"location"`
}

// BeforeInsert ...
type BeforeInsert interface {
	BeforeInsert()
}

// IModel ...
type IModel interface {
	Table() interface{}
	ID() string
	SetID(string)
	Version() int
	SetVersion(int)
}

// ISync ...
type ISync interface {
	Sync() error
}

// FindResult ...
type FindResult func(rows *xorm.Rows) error

// ConfigOptions ...
type ConfigOptions func(config *DBConfig)

var (
	_              = mysql.Config{}
	_              = sqlite3.Error{}
	_database      *xorm.Engine
	_databaseTable map[string]ISync
)

// ShowSQLOptions ...
func ShowSQLOptions(b bool) ConfigOptions {
	return func(config *DBConfig) {
		config.ShowSQL = b
	}
}

// UseCacheOptions ...
func UseCacheOptions(b bool) ConfigOptions {
	return func(config *DBConfig) {
		config.UseCache = b
	}
}

// SchemaOption ...
func SchemaOption(s string) ConfigOptions {
	return func(config *DBConfig) {
		config.Schema = s
	}
}

// LoginOption ...
func LoginOption(addr, user, pass string) ConfigOptions {
	return func(config *DBConfig) {
		config.Addr = addr
		config.Username = user
		config.Password = pass
	}
}

// MakeDBInstance ...
func MakeDBInstance(config DBConfig) (engine *xorm.Engine, e error) {
	switch config.DBType {
	case "mysql":
		engine, e = initSQLite3(config)
	case "sqlite3":
		engine, e = initSQLite3(config)
	}
	if e != nil {
		return nil, e
	}
	if config.ShowSQL {
		engine.ShowSQL()
	}
	if config.ShowExecTime {
		engine.ShowExecTime()
	}
	return
}

// InitMySQL ...
func initMySQL(ops ...ConfigOptions) (*xorm.Engine, error) {
	config := &DBConfig{
		ShowSQL:  true,
		UseCache: true,
		DBType:   "mysql",
		Addr:     "localhost",
		Username: "root",
		Password: "111111",
		Schema:   "glvd",
		Location: url.QueryEscape("Asia/Shanghai"),
		Charset:  "utf8mb4",
		Prefix:   "",
	}
	for _, op := range ops {
		op(config)
	}
	dbEngine, e := xorm.NewEngine(config.DBType, config.dbSource())
	if e != nil {
		return nil, e
	}
	defer dbEngine.Close()
	sql := fmt.Sprintf(createDatabase, config.Schema)

	_, e = dbEngine.DB().Exec(sql)
	if e == nil {
		log.Infow("create database", "database", config.Schema)
	}
	engine, e := xorm.NewEngine(config.DBType, config.source())
	if e != nil {
		return nil, e
	}

	return engine, nil
}

// Source ...
func (d *DBConfig) source() string {
	return fmt.Sprintf(mysqlStatement,
		d.Username, d.Password, d.Addr, d.Schema, d.Location, d.Charset)
}

func (d *DBConfig) dbSource() string {
	return fmt.Sprintf(mysqlStatement,
		d.Username, d.Password, d.Addr, "", d.Location, d.Charset)
}

// ID ...
func (m Model) ID() string {
	return m.model.ID
}

// SetID ...
func (m *Model) SetID(id string) {
	m.model.ID = id
}

// Version ...
func (m Model) Version() int {
	return m.model.Version
}

// SetVersion ...
func (m *Model) SetVersion(v int) {
	m.model.Version = v
}

// BeforeInsert ...
func (m *Model) BeforeInsert() {
	if m.model.ID == "" {
		m.model.ID = UUID().String()
	}
}

func liteSource(name string) string {
	return fmt.Sprintf("file:%s?cache=shared&mode=rwc&_journal_mode=WAL", name)
}

func initSQLite3(config DBConfig) (*xorm.Engine, error) {
	engine, e := xorm.NewEngine(config.DBType, liteSource(config.Schema+".db"))
	if e != nil {
		return nil, e
	}

	return engine, nil
}

// MustDatabase ...
func MustDatabase(engine *xorm.Engine, err error) *xorm.Engine {
	if err != nil {
		panic(err)
	}
	return engine
}

// RegisterDatabase ...
func RegisterDatabase(engine *xorm.Engine) {
	if _database == nil {
		_database = engine
	}
}

// registerTable ...
func registerTable(m ISync) {
	if _databaseTable == nil {
		_databaseTable = make(map[string]ISync)
	}
	_databaseTable[reflect.TypeOf(m).Name()] = m
}

// SyncTable ...
func SyncTable() (e error) {
	for _, v := range _databaseTable {
		if err := v.Sync(); err != nil {
			return err
		}
	}
	return nil
}

// InsertOrUpdate ...
func InsertOrUpdate(m IModel) (i int64, e error) {
	i, e = _database.InsertOne(m)
	if e != nil {
		return 0, e
	}
	return i, e
}

// FindAll ...
func FindAll(model IModel, f FindResult, limit int, start ...int) (e error) {
	table := _database.Table(model.Table())
	if limit == 0 {
		return nil
	} else if limit > 0 {
		table = table.Limit(limit, start...)
	}
	rows, e := table.Rows(model)
	if e != nil {
		return e
	}

	for rows.Next() {
		if err := f(rows); err != nil {
			return err
		}
	}

	return nil
}

// MustSession ...
func MustSession(session *xorm.Session) *xorm.Session {
	if session == nil {
		return _database.Where("")
	}
	return session
}

// IsExist ...
func IsExist(m IModel) bool {
	i, e := _database.
		Where("id = ?", m.ID()).
		Count(m.Table())
	if e != nil || i <= 0 {
		return false
	}
	return true
}

// UUID ...
func UUID() uuid.UUID {
	return uuid.Must(uuid.NewUUID())
}

// MustString  must string
func MustString(val, src string) string {
	if val != "" {
		return val
	}
	return src
}

// CheckDatabase ...
func CheckDatabase() bool {
	if _database == nil {
		return false
	}
	return _database.Ping() == nil
}
