package errors

import "errors"

var (
	ErrConnectDatabase = errors.New("FAILED_CONNECT_TO_DATABASE")

	ErrMigrateDatabase = errors.New("FAILED_MIGRATE_DATABASE")

	ErrEmailAlreadyUsed = errors.New("EMAIL_ALREADY_USED")

	ErrEmailNotFound = errors.New("EMAIL_NOT_FOUND")

	ErrHashPassword = errors.New("FAILED_HASHING_PASSWORD")

	ErrGenerateToken = errors.New("FAILED_GENERATE_TOKEN")

	ErrParsingString = errors.New("FAILED_PARSING_STRING_INTO_NUMBER")

	ErrSigningJWT = errors.New("FAILED_SIGNING_JWT")

	ErrClaimsJWT = errors.New("FAILED_GET_CLAIMS_FROM_JWT")

	ErrRequestTimeout = errors.New("REQUEST_TIMEOUT")

	ErrInvalidRequest = errors.New("INVALID_REQUEST")

	ErrInternalServer = errors.New("INTERNAL_SERVER_ERROR")

	ErrInvalidEmail = errors.New("INVALID_EMAIL")

	ErrParsingHTML = errors.New("FAILED_PARSING_HTML")

	ErrSendMail = errors.New("FAILED_SENDING_MAIL")

	ErrInvalidPassword = errors.New("INVALID_PASSWORD")

	ErrInvalidPhoneNumber = errors.New("INVALID_PHONE_NUMBER")

	ErrNameRequired = errors.New("NAME_REQUIRED")

	ErrEmailRequired = errors.New("EMAIL_REQUIRED")

	ErrPasswordRequired = errors.New("PASSWORD_REQUIRED")

	ErrPhoneNumberRequired = errors.New("PHONE_NUMBER_REQUIRED")

	ErrUsernameRequired = errors.New("USERNAME_REQUIRED")

	ErrAccountNotVerified = errors.New("ACCOUNT_NOT_VERIFIED")

	ErrBadRequest = errors.New("BAD_REQUEST")

	ErrFailedCreateUser = errors.New("FAILED_CREATE_USER")

)