package main

func (s *server) routes() {
	s.router.GET("/", s.spit(s.handleIndex()))
}
