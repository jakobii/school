package school

// Session
type Session struct {
	ID         uuid.UUID
	Start, End time.Time
	Teachers   []Teacher
	Attendance []Student
}

// Duration gets the duration of a class session.
func (s *Session) Duration() (d time.Duration) {

}
