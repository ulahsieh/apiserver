package models

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
	log "github.com/sirupsen/logrus"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Find(csn string) (response interface{}, err error) {

	sqlComand := fmt.Sprintf("SELECT DISTINCT SERIAL_NUMBER,PRE_KPSN,THIS_KPSN,DATECODE,LOTCODE from EMESC.VP_ASSY_SUMMARY_SMT where serial_number = '%s'", csn)

	rows, err := r.db.Query(sqlComand)
	if err != nil {
		log.Debug("Query Database Failed!")
		return nil, err
	}
	defer rows.Close()

	var sn, pre, this, dc, lc string
	var result []map[string]interface{}
	var eachrow map[string]interface{}
	eachrow = make(map[string]interface{})
	for rows.Next() {
		rows.Scan(&sn, &pre, &this, &dc, &lc)
		eachrow["SERIAL_NUMBER"] = sn
		eachrow["PRE_KPSN"] = pre
		eachrow["THIS_KPSN"] = this
		eachrow["DATECODE"] = dc
		eachrow["LOTCODE"] = lc
		result = append(result, eachrow)
	}

	if len(result) != 0 {
		log.Debug("query success! return the result...")
		return result, nil
	} else {
		log.Debug("query failed! Target param not found...")
		return nil, nil
	}
}
