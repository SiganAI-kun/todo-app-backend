package database

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IDbContext interface {
	WithContext(ctx context.Context) *DbContext
	Session() *gorm.DB
	BeginTransaction()
	Commit()
	Rollback()
	AutoMigrate(any ...any)
}

type DbContext struct {
	session     *gorm.DB
	transaction *gorm.DB
}

func NewDatabase(dsn string) *DbContext {
	dsn += "?parseTime=True"
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	session, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		panic("failed to connect database")
	}

	db := &DbContext{
		session: session,
	}

	return db
}

// Current は現在の Gorm データベースセッションを返します
func Current(ctx context.Context) *DbContext {
	value := ctx.Value("DB")
	if value == nil {
		fmt.Println("DB value is nil")
		return nil
	}

	dbContext, ok := value.(*DbContext)
	if !ok {
		fmt.Println("DB value is not of type *DbContext")
		return nil
	}

	return dbContext
}

// WithContext は指定されたコンテキストを使用して新しい Gorm データベースセッションを作成します
func (db *DbContext) WithContext(ctx context.Context) *DbContext {
	return &DbContext{
		session: db.session.WithContext(ctx),
	}
}

// Session は新しい Gorm データベースセッションを作成します
func (db *DbContext) Session() *gorm.DB {
	if db.transaction != nil {
		return db.transaction
	} else {
		return db.session
	}
}

// BeginTransaction はトランザクションを開始します
func (db *DbContext) BeginTransaction() {
	db.transaction = db.session.Begin()
}

// Commit はトランザクションをコミットします
func (db *DbContext) Commit() {
	if db.transaction != nil {
		db.transaction.Commit()
		db.transaction = nil
	} else {
		return
	}
}

// Rollback はトランザクションをロールバックします
func (db *DbContext) Rollback() {
	if db.transaction != nil {
		db.transaction.Rollback()
		db.transaction = nil
	} else {
		return
	}
}

// AutoMigrate は指定されたモデルをデータベースに自動マイグレーションします
func (db *DbContext) AutoMigrate(any ...any) {
	db.session.AutoMigrate(any...)
}
