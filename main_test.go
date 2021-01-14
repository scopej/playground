package main

import (
	"context"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	ctx := context.Background()
	//	ctx, _ = context.WithTimeout(ctx, time.Nanosecond)

	ctx, cancelFunc := context.WithCancel(ctx)
	cancelFunc()
	var rowCount int
	tableName := "fake_table"
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Table(tableName).Select("count(1)").Scan(&rowCount).Error
	})

	t.Error(err)
}
