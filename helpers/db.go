package helpers

type NotExist struct{}

func (e NotExist) Error() string {
	return "Data not exist"
}

func CheckRowsAffected(rows int64) error {
	if rows == 0 {
		return NotExist{}
	}

	return nil
}
