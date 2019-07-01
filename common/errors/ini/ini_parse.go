package ini

type IniParseError struct {
	error_info string
}

func (e *IniParseError) Error() string {
	return e.error_info
}