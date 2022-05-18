package models

import (
	"fmt"

	_ "github.com/godror/godror"
	log "github.com/sirupsen/logrus"
)

func (r *Repo) FindReverse(csn string) (response interface{}, err error) {

	sqlComand := fmt.Sprintf("select key_part_sn from emesp.tp_assy_rec where serial_number ='%s'", csn)

	rows, err := r.db.Query(sqlComand)
	if err != nil {
		log.Debug("Query Database Failed!")
		return nil, err
	}
	defer rows.Close()

	var sz string
	result := make([]string, 0, 20)
	for rows.Next() {
		rows.Scan(&sz)
		result = append(result, sz)
	}
	result = append(result, csn)

	if len(result) != 0 {
		log.Debug("query success! return the result...")
		return result, nil
	} else {
		log.Debug("query failed! Target param not found...")
		return nil, nil
	}
}
