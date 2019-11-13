package management

import (
	"testing"

	"github.com/godcong/go-trait"
	"github.com/goextension/log"
)

func init() {
	log.Register(trait.NewZapSugar())
	cfg := DefaultConfig()
	RegisterDatabase(MustDatabase(MakeDBInstance(cfg)))
	e := SyncTable()
	if e != nil {
		panic(e)
	}
}

// TestInsertOrUpdate ...
func TestInsertOrUpdate(t *testing.T) {
	i, e := InsertOrUpdate(&Video{})
	if e != nil {
		t.Fatal(e)
	}
	if i == 0 {
		t.Failed()
	}
}

// TestFindAll ...
func TestFindAll(t *testing.T) {
	var v SVideo
	rows, e := _database.Table(v.TableName()).Rows(&v)
	if e != nil {
		t.Fatal(e)
	}
	count := 0
	for rows.Next() {
		if err := rows.Scan(&v); err != nil {
			t.Fatal(err)
		}
		count++
		log.Info(v)
	}
	log.Info("total:", count)
}
