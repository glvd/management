package management

import (
	"testing"

	"github.com/godcong/go-trait"
	"github.com/goextension/log"
)

func init() {
	log.Register(trait.NewZapFileSugar("zap.log"))
	cfg := DefaultConfig()
	RegisterDatabase(MustDatabase(initSQLite3(cfg)))
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
	rows, e := _database.Table(SVideo{}).Rows(&SVideo{})
	if e != nil {
		t.Fatal(e)
	}

	for rows.Next() {
		var v SVideo
		if err := rows.Scan(&v); err != nil {
			t.Fatal(err)
		}
		log.Info(v)
	}
}
