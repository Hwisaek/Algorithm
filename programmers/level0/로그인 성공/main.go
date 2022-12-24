package 로그인_성공

func solution(id_pw []string, db [][]string) string {
	for _, strings := range db {
		if strings[0] == id_pw[0] {
			if strings[1] == id_pw[1] {
				return "login"
			}
			return "wrong pw"
		}
	}
	return "fail"
}
