package service

import (
	"github.com/jeyog/seemeetcan-server/dataset"
)

type AuthService struct {
}

func (a *AuthService) loginJJInstar(studentId string, password string) dataset.XMain {
	var root dataset.XMain
	root.XMLns = "http://www.nexacroplatform.com/platform/dataset"
	var parms dataset.Parameters
	parms.Parameters = append(parms.Parameters, dataset.Parameter{Id: "fsp_action", Value: "JJLogin"})
	parms.Parameters = append(parms.Parameters, dataset.Parameter{Id: "fsp_cmd", Value: "login"})
	parms.Parameters = append(parms.Parameters, dataset.Parameter{Id: "GV_USER_ID", Value: "undefined"})
	parms.Parameters = append(parms.Parameters, dataset.Parameter{Id: "GV_IP_ADDRESS", Value: "undefined"})
	parms.Parameters = append(parms.Parameters, dataset.Parameter{Id: "fsp_logId", Value: "all"})
	parms.Parameters = append(parms.Parameters, dataset.Parameter{Id: "MAX_WRONG_COUNT", Value: "5"})
	root.Parameters = parms

	var ds dataset.Dataset
	ds.Id = "ds_cond"
	columnInfo := dataset.ColumnInfo{}
	columnInfo.Columns = make([]dataset.Column, 0)
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "SYSTEM_CODE", Type: "STRING", Size: "256"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "MEM_ID", Type: "STRING", Size: "10"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "MEM_PW", Type: "STRING", Size: "15"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "MEM_IP", Type: "STRING", Size: "20"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "ROLE_GUBUN1", Type: "STRING", Size: "256"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "ROLE_GUBUN2", Type: "STRING", Size: "256"})
	ds.ColumnInfo = columnInfo
	var rows dataset.Rows
	var row dataset.Row
	row.Cols = make([]dataset.Col, 0)
	row.Cols = append(row.Cols, dataset.Col{Id: "SYSTEM_CODE", Value: "INSTAR_WEB"})
	row.Cols = append(row.Cols, dataset.Col{Id: "MEM_ID", Value: studentId})
	row.Cols = append(row.Cols, dataset.Col{Id: "MEM_PW", Value: password})
	rows.Rows = make([]dataset.Row, 0)
	rows.Rows = append(rows.Rows, row)
	ds.Rows = rows
	root.Dataset = append(root.Dataset, ds)

	ds.Id = "fsp_sd_cmd"
	columnInfo.Columns = make([]dataset.Column, 0)
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "TX_NAME", Type: "STRING", Size: "100"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "TYPE", Type: "STRING", Size: "10"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "SQL_ID", Type: "STRING", Size: "200"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "KEY_SQL_ID", Type: "STRING", Size: "200"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "KEY_INCREMENT", Type: "INT", Size: "10"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "CALLBACK_SQL_ID", Type: "STRING", Size: "200"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "INSERT_SQL_ID", Type: "STRING", Size: "200"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "UPDATE_SQL_ID", Type: "STRING", Size: "200"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "DELETE_SQL_ID", Type: "STRING", Size: "200"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "SAVE_FLAG_COLUMN", Type: "STRING", Size: "200"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "USE_INPUT", Type: "STRING", Size: "1"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "USE_ORDER", Type: "STRING", Size: "1"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "KEY_ZERO_LEN", Type: "INT", Size: "10"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "BIZ_NAME", Type: "STRING", Size: "100"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "PAGE_NO", Type: "INT", Size: "10"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "PAGE_SIZE", Type: "INT", Size: "10"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "READ_ALL", Type: "STRING", Size: "1"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "EXEC_TYPE", Type: "STRING", Size: "2"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "EXEC", Type: "STRING", Size: "1"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "FAIL", Type: "STRING", Size: "1"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "FAIL_MSG", Type: "STRING", Size: "200"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "EXEC_CNT", Type: "INT", Size: "1"})
	columnInfo.Columns = append(columnInfo.Columns, dataset.Column{Id: "MSG", Type: "STRING", Size: "200"})
	ds.ColumnInfo = columnInfo
	row.Cols = make([]dataset.Col, 0)
	rows.Rows = make([]dataset.Row, 0)
	rows.Rows = append(rows.Rows, row)
	ds.Rows = rows
	root.Dataset = append(root.Dataset, ds)

	ds.Id = "gds_user"
	columnInfo.Columns = nil
	ds.ColumnInfo = columnInfo
	row.Cols = nil
	rows.Rows = nil
	ds.Rows = rows
	root.Dataset = append(root.Dataset, ds)
	return root
}
