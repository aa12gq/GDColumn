package validators

func ValidatePasswordConfirm(password,passwordConfirm string,errs map[string][]string)map[string][]string{
	if password != passwordConfirm{
		errs["password_confirm"] = append(errs["password_confirm"],"两次输入密码不正确！")
	}
	return errs
}


