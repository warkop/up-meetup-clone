package db

import (
	"log"

	"gorm.io/gorm"
)

type TransactionProto interface {
	// read
	Find(out interface{}, where ...interface{}) TransactionProto
	First(out interface{}, where ...interface{}) TransactionProto
	Take(out interface{}, where ...interface{}) TransactionProto
	Last(out interface{}, where ...interface{}) TransactionProto

	// write
	Create(value interface{}) TransactionProto
	Update(attrs ...interface{}) TransactionProto
	Delete(value interface{}, where ...interface{}) TransactionProto

	// condition
	Where(query interface{}, args ...interface{}) TransactionProto

	// support for raw query
	Raw(sql string, values ...interface{}) TransactionProto

	// scan result from raw sql query
	Scan(dest interface{}) TransactionProto

	Begin() TransactionProto
	Commit() TransactionProto
	Rollback() TransactionProto

	SetError(err error)
	GetError() error

	Stop()
}

type Transaction struct {
	Tx       *gorm.DB
	Instance *gorm.DB
	Error    error
	Stopped  bool
}

func InitializeTransaction(orm *gorm.DB, table string) TransactionProto {
	inst := orm.Table(table)

	if inst.Error != nil {
		log.Fatal(inst.Error)
	}

	return &Transaction{
		Instance: inst,
		Stopped:  false,
	}
}

func (tx *Transaction) Find(out interface{}, where ...interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.Find(out, where...)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) First(out interface{}, where ...interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.First(out, where...)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Take(out interface{}, where ...interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.Take(out, where...)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Last(out interface{}, where ...interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.Last(out, where...)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Create(value interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.Create(value)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Update(attrs ...interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.Updates(attrs)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Delete(value interface{}, where ...interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.Delete(value, where...)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Where(query interface{}, args ...interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.Where(query, args...)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Raw(sql string, values ...interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.Raw(sql, values...)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Scan(dest interface{}) TransactionProto {
	if tx.Stopped {
		return tx
	}

	t := tx.Tx.Scan(dest)

	if t.Error != nil {
		tx.SetError(t.Error)
		tx.Stop()
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Begin() TransactionProto {
	tx.Tx = tx.Instance.Begin()

	if tx.Tx.Error != nil {
		tx.SetError(tx.Tx.Error)
	} else {
		tx.SetError(nil)
	}

	tx.Stopped = false

	return tx
}

func (tx *Transaction) Commit() TransactionProto {
	t := tx.Tx.Commit()

	if t.Error != nil {
		tx.SetError(t.Error)
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) Rollback() TransactionProto {
	t := tx.Tx.Rollback()

	if t.Error != nil {
		tx.SetError(t.Error)
	}

	tx.Tx = t

	return tx
}

func (tx *Transaction) SetError(err error) {
	tx.Error = err
}

func (tx *Transaction) GetError() error {
	return tx.Error
}

func (tx *Transaction) Stop() {
	tx.Stopped = true
}
