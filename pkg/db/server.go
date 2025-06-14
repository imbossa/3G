package db

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DBOption -.
type DBOption struct {
	MaxIdleConns int
	MaxOpenConns int
	CONFIG       string
}

// DB holds the database connection.
type DB struct {
	URL    string
	Option DBOption
	Conn   *gorm.DB
}

// New creates a new DB instance and connects to the specified database using gorm.
func New(url string, opts ...Option) (*DB, error) {
	option := DBOption{
		MaxIdleConns: 5,
		MaxOpenConns: 10,
		CONFIG:       "",
	}
	for _, opt := range opts {
		opt(&option)
	}

	var typ, dsn string
	if idx := strings.Index(url, "://"); idx > 0 {
		typ = url[:idx]
		dsn = url[idx+3:]
	} else {
		typ = ""
		dsn = url
	}

	var dialector gorm.Dialector
	switch typ {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		if strings.HasPrefix(dsn, "postgres") {
			dialector = postgres.Open(dsn)
		} else if strings.HasPrefix(dsn, "sqlite") {
			dialector = sqlite.Open(dsn)
		} else {
			dialector = mysql.Open(dsn)
		}
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		return nil, fmt.Errorf("db - New - gorm.Open: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("db - New - db.DB(): %w", err)
	}

	sqlDB.SetMaxIdleConns(option.MaxIdleConns)
	sqlDB.SetMaxOpenConns(option.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if option.CONFIG != "" {
		log.Printf("db - New - custom config: %s", option.CONFIG)
	}

	return &DB{
		URL:    url,
		Option: option,
		Conn:   db,
	}, nil
}

// Close closes the database connection.
func (d *DB) Close() error {
	if d.Conn == nil {
		return nil
	}
	sqlDB, err := d.Conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
