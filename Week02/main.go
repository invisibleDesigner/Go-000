package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// 题目：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

// 解决方案：应该wrap这个错误，并抛给上层，因为sql.ErrNoRows是一个sentinel,直接返回这个错误，没有任何堆栈信息，所以必须wrap

func dao() error {
	return sql.ErrNoRows
}

func service() error {
	err := dao()
	if err != nil {
		return errors.Wrap(err, "my service")
	}
	return nil
}

func main() {
	err := service()
	if err != nil {
		fmt.Printf("%+v", err)
	}
}

//dao:
//
//return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))
//
//
//biz:
//
//if errors.Is(err, code.NotFound} {
//
//}