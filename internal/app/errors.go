package app

type (
	InvalidStorageType struct {
		t string
	}
)

func (e InvalidStorageType) Error() string {
	return "storage type '" + e.t + "' does not exist"
}

func InvalidStorageTypeError(t string) error {
	return InvalidStorageType{t}
}
