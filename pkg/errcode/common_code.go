package errcode

var (
	Success                 = NewError(20000, "success")
	RepeatError             = NewError(30000, "record repeats error")
	RepeatUsernameError     = NewError(30001, "username has been registered")
	AuthenticationFailError = NewError(30002, "authentication failed")
	EmptyRolesError         = NewError(30010, "user roles can not be empty")
	IllegalRolesError       = NewError(30011, "user roles must be in \"guest, operator, admin, root\" ")
	ServerError             = NewError(40000, "server inside error")

	IllegalTokenError = NewError(50008, "token is illegal ")
	LoggedTokenError  = NewError(50012, "other clients logged in")
	ExpiredTokenError = NewError(50014, "token is expired")
)
