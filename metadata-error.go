package errors

// WithMetadata annotates an error with metadata.
func WithMetadata(err error, meta any) error {
	if meta == nil {
		return err
	}
	return &metadataError{err, meta}
}

// Metadata returns the metadata associated with an error.
func Metadata(err error) any {
	if err == nil {
		return nil
	}

	me, ok := err.(iMetadataError)
	if !ok {
		return nil
	}

	return me.Metadata()
}

// TypedMetadata returns the metadata associated with an error.
// If the metadata is not of the specified type, ok is false.
func TypedMetadata[T any](err error) (ret T, ok bool) {
	meta := Metadata(err)
	if meta == nil {
		return
	}

	ret, ok = meta.(T)
	return
}

// TypedDeepMetadata returns the metadata of the specified type
// associated with an error or any of its wrappers.
func TypedDeepMetadata[T any](err error) (ret T, ok bool) {
	for {
		if err == nil {
			return
		}

		ret, ok = TypedMetadata[T](err)

		if ok {
			return
		}

		err = Unwrap(err)
	}
}

type iMetadataError interface {
	error
	Metadata() any
	Unwrap() error
}

type metadataError struct {
	error
	meta any
}

var _ iMetadataError = (*metadataError)(nil)

func (m *metadataError) Metadata() any {
	if m == nil {
		return nil
	}
	return m.meta
}

func (m *metadataError) Unwrap() error {
	if m == nil {
		return nil
	}
	return m.error
}
