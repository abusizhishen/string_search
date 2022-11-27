package string_search

type Engine struct {
	hash map[rune]*Bucket
}

func NewEngine() Engine {
	return Engine{hash: map[rune]*Bucket{}}
}

func (e *Engine) Insert(id int, str string) {
	e.insert(id, []rune(str))
}

func (e *Engine) insert(id int, str []rune) {
	if len(str) == 0 {
		return
	}

	var tmp = e.hash[str[0]]
	if tmp == nil {
		tmp = newItem(id, str)
	} else {
		tmp.insert(id, str[1:])
	}

	e.hash[str[0]] = tmp
}

func (e *Engine) InsertBatch(items []Item) {
	for _, item := range items {
		e.insert(item.Id, []rune(item.Str))
	}
}

func (e *Engine) Search(str []rune) []int {
	if len(str) == 0 || len(e.hash) == 0 {
		return nil
	}

	var ids []int
	for _, item := range e.hash {
		var tmpIds []int
		if item.str == str[0] {
			tmpIds = item.search(str[1:])
			if len(tmpIds) == 0 {
				tmpIds = item.search(str)
			}
		}else{
			tmpIds = item.search(str)
		}

		if len(tmpIds) != 0 {
			ids = append(ids, tmpIds...)
		}
	}

	return ids
}

type Item struct {
	Id  int
	Str string
}
