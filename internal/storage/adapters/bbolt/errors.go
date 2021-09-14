package bbolt

type (
	InvalidBucketName struct {
		bucketName string
	}
)

func (e InvalidBucketName) Error() string {
	return "bucket '" + e.bucketName + "' does not exist"
}
