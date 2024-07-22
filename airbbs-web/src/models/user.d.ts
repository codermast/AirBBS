export type User = {
	id: string
	username: string
	nickname: string
	photo: string
	mail: string
	github: string
	tel: string
	admin: string
	introduce: string
	sex: string
	registerTime: string
	loginTime: string
}

export type UserRegisterInfo = {
	id: string
	username: string
	password: string
	confirmPassword: string
	nickname: string
	mail: string
	tel: string
	code: string
}

export type UserResetPasswordInfo = {
	account: string
	password: string
	code: string
}

export type UserLoginInfo = {
	username: string
	password: string
}

export type UserQueryInfo = {
	id: string
	username: string
	nickname: string
	mail: string
	tel: string
	github: string
	introduce: string
	sex: string
}