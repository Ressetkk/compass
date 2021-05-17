// Code generated by "stringer -type ErrorType"; DO NOT EDIT.

package apperrors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InternalError-10]
	_ = x[UnknownError-11]
	_ = x[NotFound-20]
	_ = x[NotUnique-21]
	_ = x[InvalidData-22]
	_ = x[InsufficientScopes-23]
	_ = x[TenantRequired-24]
	_ = x[TenantNotFound-25]
	_ = x[Unauthorized-26]
	_ = x[InvalidOperation-27]
	_ = x[OperationTimeout-28]
	_ = x[EmptyData-29]
	_ = x[InconsistentData-30]
	_ = x[NotUniqueName-31]
	_ = x[ConcurrentOperation-32]
	_ = x[InvalidStatusCondition-33]
	_ = x[CannotUpdateObjectInManyBundles-34]
}

const (
	_ErrorType_name_0 = "InternalErrorUnknownError"
	_ErrorType_name_1 = "NotFoundNotUniqueInvalidDataInsufficientScopesTenantRequiredTenantNotFoundUnauthorizedInvalidOperationOperationTimeoutEmptyDataInconsistentDataNotUniqueNameConcurrentOperationInvalidStatusConditionCannotUpdateObjectInManyBundles"
)

var (
	_ErrorType_index_0 = [...]uint8{0, 13, 25}
	_ErrorType_index_1 = [...]uint8{0, 8, 17, 28, 46, 60, 74, 86, 102, 118, 127, 143, 156, 175, 197, 228}
)

func (i ErrorType) String() string {
	switch {
	case 10 <= i && i <= 11:
		i -= 10
		return _ErrorType_name_0[_ErrorType_index_0[i]:_ErrorType_index_0[i+1]]
	case 20 <= i && i <= 34:
		i -= 20
		return _ErrorType_name_1[_ErrorType_index_1[i]:_ErrorType_index_1[i+1]]
	default:
		return "ErrorType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
