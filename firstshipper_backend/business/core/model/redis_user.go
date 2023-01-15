package model

type RedisUser struct {
	LoginSessions       []string `redis:"login_sessions"`
	ForgotPasswordToken string   `redis:"forgot_password_token"`
	UserId              string   `redis:"user_id"`
	HashId              string   `redis:"hash_id"`
	BasicUserStr        string   `redis:"basic_user_str"`
}
