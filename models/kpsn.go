package models

import (
	"encoding/json"
	"fmt"

	_ "github.com/godror/godror"
	log "github.com/sirupsen/logrus"
)

type Kpsn struct {
	KPNO  string `json:"KEY_PART_NO"`
	PRE   string `json:"PRE_KPSN"`
	THIS  string `json:"THIS_KPSN"`
	DC    string `json:"Datecode"`
	LC    string `json:"Lotcode"`
	KPQTY string `json:"KP_QTY"`
	KPLOC string `json:"KP_LOCATION"`
}

func (r *Repo) Find(csn string) (response interface{}, err error) {

	sqlComand := fmt.Sprintf("select KEY_PART_NO,PRE_KPSN,THIS_KPSN,DATECODE,LOTCODE,KP_QTY,KP_LOCATION from EMESC.VP_ASSY_SUMMARY_SMT where serial_number =  '%s' and (key_part_no like '3B%%' OR key_part_no like '3V%%' OR key_part_no like '3G%%' OR key_part_no like '71%%') order by KEY_PART_NO", csn)

	rows, err := r.db.Query(sqlComand)
	if err != nil {
		log.Debug("Query Database Failed!")
		return nil, err
	}
	defer rows.Close()

	var kpsns []*Kpsn
	for rows.Next() {
		k := new(Kpsn)
		rows.Scan(&k.KPNO, &k.PRE, &k.THIS, &k.DC, &k.LC, &k.KPQTY, &k.KPLOC)
		kpsns = append(kpsns, k)
	}
	jsonStr, _ := json.Marshal(&kpsns)

	var result []Kpsn
	// var result []map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		log.Debug("Can't parse the json string")
		return nil, nil
	}

	if len(result) != 0 {
		log.Debug("query success! return the result...")
		return result, nil
	} else {
		log.Debug("query failed! Target param not found...")
		return nil, nil
	}
}
