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

export type UserLoginInfo = {
	username : string
	password : string
}