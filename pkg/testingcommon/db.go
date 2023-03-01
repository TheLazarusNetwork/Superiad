package testingcommon

import (
	"fmt"

	"github.com/MyriadFlow/superiad/pkg/store"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Referred from https://medium.com/@jarifibrahim/using-gorm-hooks-to-clean-up-test-fixtures-in-golang-99b0fcb04354
func DeleteCreatedEntities() func() {
	db := store.DB
	type entity struct {
		table   string
		keyname string
		key     interface{}
	}
	var entries []entity
	hookName := "cleanupHook"

	db.Callback().Create().After("gorm:create").Register(hookName, func(tx *gorm.DB) {
		statement := tx.Statement.Statement
		schema := tx.Statement.Schema

		if len(schema.PrimaryFields) == 0 {
			return
		}

		primaryFieldKey := ""
		primaryFieldIndex := 0
		for i, v := range schema.Fields {
			if v.PrimaryKey {
				primaryFieldKey = v.DBName
				primaryFieldIndex = i
				break
			}
		}

		primaryFieldValue := statement.Vars[primaryFieldIndex]

		fmt.Printf("Inserted entities of %s with %s=%v\n", statement.Table, primaryFieldKey, primaryFieldValue)

		entries = append(entries, entity{table: statement.Table, keyname: primaryFieldKey, key: primaryFieldValue})
	})
	return func() {
		// Remove the hook once we're done
		defer db.Callback().Create().Remove(hookName)

		for i := len(entries) - 1; i >= 0; i-- {
			entry := entries[i]

			var deleteValue string
			switch v := entry.key.(type) {
			case int:
				deleteValue = fmt.Sprint(v)

			case string:
				deleteValue = fmt.Sprintf("'%v'", v)

			default:
				log.WithFields(log.Fields{
					"err": "not Implemented",
				}).Fatal("not implemented")
			}

			q := fmt.Sprintf(`DELETE FROM %v WHERE %v=%v`, entry.table, entry.keyname, deleteValue)
			if err := db.Exec(q).Error; err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Warnf("failed to exec query: %s, in clean up hook, %s", q, err)
			}
		}

	}
}
