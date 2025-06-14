//go:build migrate

package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func getDialector(dsn string) (gorm.Dialector, error) {
	if strings.HasPrefix(dsn, "mysql://") || strings.HasPrefix(dsn, "mysql:") {
		trim := strings.TrimPrefix(dsn, "mysql://")
		trim = strings.TrimPrefix(trim, "mysql:")
		return mysql.Open(trim), nil
	}
	if strings.HasPrefix(dsn, "postgres://") || strings.HasPrefix(dsn, "postgresql://") {
		return postgres.Open(dsn), nil
	}
	if strings.HasPrefix(dsn, "sqlite://") {
		path := strings.TrimPrefix(dsn, "sqlite://")
		if strings.HasPrefix(path, "/") {
			return sqlite.Open(path), nil
		}
		return nil, fmt.Errorf("sqlite file path must be absolute")
	}
	if strings.HasSuffix(dsn, ".db") || strings.HasSuffix(dsn, ".sqlite") || strings.HasSuffix(dsn, ".sqlite3") {
		return sqlite.Open(dsn), nil
	}
	return nil, fmt.Errorf("unknown or unsupported database dialect in DSN: %s", dsn)
}

func RunMigrate() {
	_ = godotenv.Load()

	databaseURL := ""
	flag.StringVar(&databaseURL, "database", "", "database url")
	flag.Parse()

	if databaseURL == "" {
		databaseURL = os.Getenv("DB_URL")
	}
	if databaseURL == "" {
		log.Fatalf("migrate: database url must be set by -database or env DB_URL")
	}

	var (
		attempts = _defaultAttempts
		err      error
		db       *gorm.DB
	)

	dialector, err := getDialector(databaseURL)
	if err != nil {
		log.Fatalf("migrate: DSN parse error: %v", err)
	}

	for attempts > 0 {
		db, err = gorm.Open(dialector, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			sqlDB, err2 := db.DB()
			if err2 == nil {
				if pingErr := sqlDB.Ping(); pingErr == nil {
					break
				}
			}
		}

		log.Printf("Migrate: DB connect failed, attempts left: %d, error: %v", attempts, err)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: DB connect error: %v", err)
	}

	log.Printf("Migrate: DB connect success (no AutoMigrate, only connection test)")
}

func init() {
	if len(os.Args) > 0 && (strings.HasSuffix(os.Args[0], ".test") || strings.Contains(os.Args[0], "test")) {
		return // 跳过测试
	}
	RunMigrate()
}
