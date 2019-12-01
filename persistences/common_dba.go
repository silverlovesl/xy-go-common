package persistences

import (
	"github.com/go-xorm/xorm"
)

// HandledDBError represents a error reason by code and reason string.
type HandledDBError struct {
	DBError error
}

// ThrowDBError は HandleDBError にラップした error を panic に渡す.
func ThrowDBError(err error) {
	panic(HandledDBError{
		DBError: err,
	})
}

// Error returns error reason.
func (e *HandledDBError) Error() string {
	return e.DBError.Error()
}

// CommonDBA is commonDBA of User.
type CommonDBA struct {
	engine *xorm.Engine
}

// GetCommonDBA return a common commonDBA instance.
func GetCommonDBA() *CommonDBA {
	return &CommonDBA{
		engine: GetEngine(),
	}
}

// RunInTransaction does statements in db transaction.
func (commonDBA *CommonDBA) RunInTransaction(fn func(session *xorm.Session) error) (result error) {
	session := commonDBA.engine.NewSession()
	defer func() {
		session.Close()
		err := recover()
		if err != nil {
			// HnadledDBError の場合は panic せずに error を返す.
			e, ok := err.(HandledDBError)
			if ok {
				// 戻り値をアンラップしたオリジナルの error に設定.
				result = e.DBError
			} else {
				panic(err)
			}
		}
	}()
	err := session.Begin()
	if err != nil {
		return err
	}

	err = fn(session)
	if err != nil {
		session.Rollback()
		return err
	}

	err = session.Commit()
	if err != nil {
		session.Rollback()
		return err
	}

	return nil
}

// Run does statements in db session without transaction.
func (commonDBA *CommonDBA) Run(fn func(session *xorm.Session) error) (result error) {
	session := commonDBA.engine.NewSession()
	defer func() {
		session.Close()
		err := recover()
		if err != nil {
			e, ok := err.(HandledDBError)
			if ok {
				// 戻り値をアンラップしたオリジナルの error に設定.
				result = e.DBError
			} else {
				panic(err)
			}
		}
	}()
	return fn(session)
}

// InsertMany Entity
func (commonDBA *CommonDBA) InsertMany(session *xorm.Session, records interface{}) int64 {
	count, err := session.InsertMulti(records)
	if err != nil {
		ThrowDBError(err)
	}
	return count
}

// InsertOne Entity
func (commonDBA *CommonDBA) InsertOne(session *xorm.Session, record interface{}) int64 {
	count, err := session.InsertOne(record)
	if err != nil {
		ThrowDBError(err)
	}
	return count
}

// UpdateByID 指定されたフィールドを更新する.
func (commonDBA *CommonDBA) UpdateByID(session *xorm.Session, id int, record interface{}) int64 {
	count, err := session.ID(id).AllCols().Update(record)
	if err != nil {
		ThrowDBError(err)
	}
	return count
}

// DeleteByID 指定されたフィールドを削除する.
func (commonDBA *CommonDBA) DeleteByID(session *xorm.Session, id int, record interface{}) int64 {
	count, err := session.ID(id).Delete(record)
	if err != nil {
		ThrowDBError(err)
	}
	return count
}
