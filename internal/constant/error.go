package constant

import "errors"

var (
	ParseTokenMsg               = errors.New("parse token failed")
	TokenExpiredMsg             = errors.New("token is timed out, please log in again")
	TokenInvalidMsg             = errors.New("token has been invalidated")
	TokenNotValidYetMsg         = errors.New("token not active yet")
	TokenMalformedMsg           = errors.New("that's not even a token")
	TokenUnknownMsg             = errors.New("couldn't handle this token")
	TokenUserKickedMsg          = errors.New("user has been kicked")
	TokenDifferentPlatformIDMsg = errors.New("different platformID")
	TokenDifferentUserIDMsg     = errors.New("different userID")
	AccessMsg                   = errors.New("no permission")
	StatusMsg                   = errors.New("status is abnormal")
	DBMsg                       = errors.New("db failed")
	ArgsMsg                     = errors.New("args failed")
	CallBackMsg                 = errors.New("callback failed")
	InvitationMsg               = errors.New("invitationCode error")

	ThirdPartyMsg = errors.New("third party error")
)
