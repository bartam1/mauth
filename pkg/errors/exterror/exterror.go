package exterror

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnexpected     = ErrorType{"unexpected"}
	ErrorTypeRepository     = ErrorType{"repository"}
	ErrorTypeAuthorization  = ErrorType{"authorization"}
	ErrorTypeIncorrectInput = ErrorType{"incorrectinput"}
)

type ExtError struct {
	error     error
	slug      string
	errorType ErrorType
}

func (s ExtError) Error() string {
	return s.error.Error()
}
func (s ExtError) Slug() string {
	return s.slug
}

func (s ExtError) ErrorType() ErrorType {
	return s.errorType
}

func NewUnexpected(error error, slug string, log func()) ExtError {
	log()
	return ExtError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeUnexpected,
	}
}
func NewRepoSlug(err error, slug string, log func()) ExtError {
	log()
	return ExtError{
		error:     err,
		slug:      slug,
		errorType: ErrorTypeRepository,
	}
}
func NewRepo(err error, log func()) ExtError {
	log()
	return ExtError{
		error:     err,
		slug:      "",
		errorType: ErrorTypeRepository,
	}
}
func NewIncorrectInput(error error, slug string, log func()) ExtError {
	log()
	return ExtError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeIncorrectInput,
	}
}
func NewAuthorization(error error, slug string, log func()) ExtError {
	log()
	return ExtError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeAuthorization,
	}
}
