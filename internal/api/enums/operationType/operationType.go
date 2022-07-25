package operationType

// Type is the variable type for this enumerator.
type Type string

const (
	// Delete indicates the operation is to delete.
	Delete Type = `delete`

	// Head indicates the operation is a head operation.
	Head Type = `head`

	// Get indicates the operation is to get some data.
	Get Type = `get`

	// Options indicates the operation is an options operation.
	Options Type = `options`

	// Patch indicates the operation is a patch operation.
	Patch Type = `patch`

	// Post indicates the operation is to post new data.
	Post Type = `post`

	// Put indicates the operation is to put to data.
	Put Type = `put`

	// Trace indicates the operation is a trace operation.
	Trace Type = `trace`
)

// All gets the list of all enumerator values in this enumerator.
func All() []Type {
	return []Type{
		Delete,
		Head,
		Get,
		Options,
		Patch,
		Post,
		Put,
		Trace,
	}
}
