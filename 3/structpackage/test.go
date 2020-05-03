package structpackage

type Student struct {
	Id   int
	Age  int
	Sex  byte
	addr string
}

func Test(s *Student) {
	s.addr = "bj"
	s.Id = 666
}
