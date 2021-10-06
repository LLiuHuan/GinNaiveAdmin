package request

import "GinNaiveAdmin/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
