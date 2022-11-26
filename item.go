package search

type Bucket struct {
	maxStrLength int //max length start with current char
	//byte         byte           //current char in utf8
	str  rune             //the last reminder str
	ids  []int            //current match ids
	hash map[rune]*Bucket // str start with current char
}

func (i *Bucket) Insert(id int, str string) {
	i.insert(id, []rune(str))
}

func (i *Bucket) insert(id int, str []rune) {
	i.ids = append(i.ids, id)
	if len(str) == 0 {
		return
	}

	if len(str) > i.maxStrLength {
		i.maxStrLength = len(str)
	}

	var tmp = i.hash[str[0]]
	if tmp == nil {
		tmp = newItem(id, str)
	} else {
		tmp.insert(id, str[1:])
	}

	i.hash[str[0]] = tmp
}

func (i *Bucket) Search(str string) []int {
	return i.search([]rune(str))
}
func (i *Bucket) search(str []rune) []int {
	if len(str) == 0 {
		return i.ids
	}

	if len(i.hash) == 0 {
		return nil
	}

	if len(str) > i.maxStrLength {
		return nil
	}
	if len(str) == i.maxStrLength {
		if _, ok := i.hash[str[0]]; !ok {
			return nil
		}
	}

	var ids []int
	for _, tmpI := range i.hash {
		var tmpIds []int
		if tmpI.str == str[0] {
			tmpIds = tmpI.search(str[1:])
		} else {
			tmpIds = tmpI.search(str)
		}

		if tmpIds == nil {
			continue
		}

		ids = append(ids, tmpIds...)
	}

	return ids
}

func NewItem(id int, str string) *Bucket {
	return newItem(id, []rune(str))
}

func newItem(id int, str []rune) *Bucket {
	i := &Bucket{
		maxStrLength: len(str) - 1,
		//byte:         str[0],
		str:  str[0],
		ids:  []int{},
		hash: map[rune]*Bucket{},
	}

	i.insert(id, str[1:])
	return i
}
