// generated by stringer -type CMType,TransactionStatus -output type_string.go; DO NOT EDIT

package pbchange

import "fmt"

const _CMType_name = "DEVICEPOLICYQUERYREPORTONTOLOGYNONE"

var _CMType_index = [...]uint8{0, 6, 12, 17, 23, 31, 35}

func (i CMType) String() string {
	if i < 0 || i >= CMType(len(_CMType_index)-1) {
		return fmt.Sprintf("CMType(%d)", i)
	}
	return _CMType_name[_CMType_index[i]:_CMType_index[i+1]]
}

const _TransactionStatus_name = "INITIALIZINGACTIVECOMPLETEFAILEDCLEANED"

var _TransactionStatus_index = [...]uint8{0, 12, 18, 26, 32, 39}

func (i TransactionStatus) String() string {
	if i < 0 || i >= TransactionStatus(len(_TransactionStatus_index)-1) {
		return fmt.Sprintf("TransactionStatus(%d)", i)
	}
	return _TransactionStatus_name[_TransactionStatus_index[i]:_TransactionStatus_index[i+1]]
}
