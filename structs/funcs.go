package structs

import "time"

type NoteRepo interface {
	NewNote(name string, text string) int
	GetNote(id int) Note
	GetAll() map[int]Note
	Delete(id int)
}

type Note struct {
	name    string
	text    string
	Created time.Time
}

type MyNoteRepo struct {
	note   map[int]Note
	nextID int
}

var MyNoteRepo1 = MyNoteRepo{
	note:   make(map[int]Note),
	nextID: 0,
}

func (m *MyNoteRepo) NewNote(
	name string,
	text string,
) int {
	if name == "" && text == "" {
		return 0
	}
	m.note[m.nextID] = Note{
		name:    name,
		text:    text,
		Created: time.Now(),
	}
	m.nextID++
	return m.nextID - 1
}

func (m MyNoteRepo) GetNote(id int) Note {
	return m.note[id]
}

func (m MyNoteRepo) GetAll() map[int]Note {
	return m.note
}

func (m MyNoteRepo) Delete(id int) {
	delete(m.note, id)
}
