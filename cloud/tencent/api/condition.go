package api

const (
	CondKeyCurrentTime = "qcs:current_time"
	CondKeyIP          = "qcs:ip"
	CondKeyResourceTag = "qcs:resource_tag"
	CondKeyRequestTag  = "qcs:request_tag"

	CondOpStringEQ    = "string_equal"
	CondOpStringNEQ   = "string_not_equal"
	CondOpStringEQIC  = "string_equal_ignore_case"
	CondOpStringNEQIC = "string_not_equal_ignore_case"

	CondOpNumericEQ   = "numeric_equal"
	CondOpNumericNEQ  = "numeric_not_equal"
	CondOpNumericLT   = "numeric_less_than"
	CondOpNumericLTEQ = "numeric_less_than_equal"
	CondOpNumericGT   = "numeric_greater_than"
	CondOpNumericGTEQ = "numeric_greater_than_equal"

	CondOpDateEQ   = "date_equal"
	CondOpDateNEQ  = "date_not_equal"
	CondOpDateLT   = "date_less_than"
	CondOpDateLTEQ = "date_less_than_equal"
	CondOpDateGT   = "date_greater_than"
	CondOpDateGTEQ = "date_greater_than_equal"

	CondOpBoolEQ = "bool_equal"

	CondOpBinaryEQ = "binary_equal"

	CondOpIpEQ  = "ip_equal"
	CondOpIpNEQ = "ip_not_equal"

	CondOpNullEQ = "null_equal"
)
