package music

type Queue struct {
	Songs []string
}

func (q *Queue) Add(song string) {
	q.Songs = append(q.Songs, song)
}
