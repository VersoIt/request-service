package request

import _ "embed"

var (
	//go:embed sql/create_request.sql
	queryCreateRequest string

	//go:embed sql/get_handling_requests.sql
	queryGetHandlingRequestsForUpdate string

	//go:embed sql/get_request.sql
	queryGetRequest string
)
