package list

import "fmt"

type List struct {
	len       int64
	nextIndex int64
	firstNode *node
}

// NewList создает новый список
func NewList() *List {
	return &List{}
}

// Len возвращает длину списка
func (l *List) Len() int64 {
	return l.len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) int64 {
	newNode := &node{value: value}
	if l.firstNode == nil {
		l.firstNode = newNode
	} else {
		current := l.firstNode
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	l.len++
	return l.len - 1
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if id < 0 || id >= l.len {
		return
	}

	if id == 0 {
		l.firstNode = l.firstNode.next
	} else {
		current := l.firstNode
		for i := int64(0); i < id-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	l.len--
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	current := l.firstNode
	prev := l.firstNode

	for current != nil {
		if current.value == value {
			if current == l.firstNode {
				l.firstNode = current.next
			} else {
				prev.next = current.next
			}
			l.len--
		}
		prev = current
		current = current.next
	}
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	for l.firstNode != nil && l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
	}

	current := l.firstNode
	prev := l.firstNode

	for current != nil {
		if current.value == value {
			prev.next = current.next
			l.len--
		} else {
			prev = current
		}
		current = current.next
	}
}

// GetByIndex возвращает значение элемента по индексу.
func (l *List) GetByIndex(id int64) (int64, bool) {
	if id < 0 || id >= l.len {
		return 0, false
	}

	current := l.firstNode
	for i := int64(0); i < id; i++ {
		current = current.next
	}
	return current.value, true
}

// GetByValue возвращает индекс первого найденного элемента по значению.
func (l *List) GetByValue(value int64) (int64, bool) {
	current := l.firstNode
	index := int64(0)

	for current != nil {
		if current.value == value {
			return index, true
		}
		current = current.next
		index++
	}

	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению.
func (l *List) GetAllByValue(value int64) ([]int64, bool) {
	indices := make([]int64, 0)
	current := l.firstNode
	index := int64(0)

	for current != nil {
		if current.value == value {
			indices = append(indices, index)
		}
		current = current.next
		index++
	}

	if len(indices) == 0 {
		return nil, false
	}

	return indices, true
}

// GetAll возвращает все элементы списка.
func (l *List) GetAll() ([]int64, bool) {
	if l.len == 0 {
		return nil, false
	}

	values := make([]int64, l.len)
	current := l.firstNode
	index := int64(0)

	for current != nil {
		values[index] = current.value
		current = current.next
		index++
	}

	return values, true
}

// Clear очищает список
func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
}

// Print выводит список в консоль
func (l *List) Print() {
	current := l.firstNode

	fmt.Print("Список: \n")
	index := 0

	for current != nil {
		fmt.Print("id =")
		fmt.Printf("%v  ", index)
		fmt.Print("значение =")
		fmt.Printf("%v  ", current.value)
		fmt.Print("\n")
		index++
		current = current.next
	}

	fmt.Println()
}
