package main

import (
	"database/sql"
	"errors"
)

const NotFound = 10000

//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
func Dao()  {
	row := 0
	row  = sql.ErrNoRows("")
	//需要Warp 这个error ，但是需要定义一个错误码，因为此处说明查询成功，但是只是无数据，这个无数据需要业务方判断使用需要对业务进行处理
	if row == 0 {
		return NotFound
	}
}

func MyError() {
	UserNotFound := errors. New("no rows")
}









