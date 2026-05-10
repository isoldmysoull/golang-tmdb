package tmdb

// StatusCode type is a custom type for TMDb API response status codes.
// See https://developer.themoviedb.org/docs/errors for more details.
type StatusCode int

const (
	StatusCodeSuccess                   StatusCode = iota + 1 // 1
	StatusCodeInvalidService                                  // 2
	StatusCodeAuthenticationFailed                            // 3
	StatusCodeInvalidFormat                                   // 4
	StatusCodeInvalidParameters                               // 5
	StatusCodeInvalidID                                       // 6
	StatusCodeInvalidAPIKey                                   // 7
	StatusCodeDuplicateEntry                                  // 8
	StatusCodeServiceOffline                                  // 9
	StatusCodeSuspendedAPIKey                                 // 10
	StatusCodeInternalError                                   // 11
	StatusCodeItemUpdatedSuccessfully                         // 12
	StatusCodeItemDeletedSuccessfully                         // 13
	StatusCodeAuthenticationFailed2                           // 14
	StatusCodeFailed                                          // 15
	StatusCodeDeviceDenied                                    // 16
	StatusCodeSessionDenied                                   // 17
	StatusCodeValidationFailed                                // 18
	StatusCodeInvalidAcceptHeader                             // 19
	StatusCodeInvalidDateRange                                // 20
	StatusCodeEntryNotFound                                   // 21
	StatusCodeInvalidPage                                     // 22
	StatusCodeInvalidDate                                     // 23
	StatusCodeRequestTimedOut                                 // 24
	StatusCodeRequestCountOverLimit                           // 25
	StatusCodeMissingUsernamePassword                         // 26
	StatusCodeTooManyAppendToResponse                         // 27
	StatusCodeInvalidTimezone                                 // 28
	StatusCodeConfirmationRequired                            // 29
	StatusCodeInvalidUsernameOrPassword                       // 30
	StatusCodeAccountDisabled                                 // 31
	StatusCodeEmailNotVerified                                // 32
	StatusCodeInvalidRequestToken                             // 33
	StatusCodeResourceNotFound                                // 34
	StatusCodeInvalidToken                                    // 35
	StatusCodeTokenNoWritePermission                          // 36
	StatusCodeSessionNotFound                                 // 37
	StatusCodeNoPermissionToEdit                              // 38
	StatusCodeResourcePrivate                                 // 39
	StatusCodeNothingToUpdate                                 // 40
	StatusCodeRequestTokenNotApproved                         // 41
	StatusCodeRequestMethodNotSupported                       // 42
	StatusCodeCouldNotConnectToBackend                        // 43
	StatusCodeInvalidIDBackend                                // 44
	StatusCodeUserSuspended                                   // 45
	StatusCodeAPIUnderMaintenance                             // 46
	StatusCodeInputNotValid                                   // 47
)
