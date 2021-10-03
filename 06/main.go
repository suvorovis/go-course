package main

func main() {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	item := list.First()
	for item != nil {
		println(item.Value().(int))
		item = item.Next()
	}

	println(list.Len())

	list.Remove(list.Last())
	list.Remove(list.First())

	item = list.First()
	for item != nil {
		println(item.Value().(int))
		item = item.Next()
	}
}

type Item struct {
	prev  *Item
	next  *Item
	value interface{}
}

func (item *Item) Next() *Item {
	return item.next
}

func (item *Item) Prev() *Item {
	return item.prev
}

func (item *Item) Value() interface{} {
	return item.value
}

type List struct {
	length int
	first  *Item
	last   *Item
}

func (list *List) Len() int {
	return list.length
}

func (list *List) First() *Item {
	return list.first
}

func (list *List) Last() *Item {
	return list.last
}

func (list *List) PushFront(v interface{}) {
	list.length++
	item := &Item{
		next:  list.first,
		value: v,
	}
	if list.first != nil {
		list.first.prev = item
	}
	if list.last == nil {
		list.last = item
	}
	list.first = item
}

func (list *List) PushBack(v interface{}) {
	list.length++
	item := &Item{
		prev:  list.last,
		value: v,
	}
	if list.last != nil {
		list.last.next = item
	}
	if list.first == nil {
		list.first = item
	}
	list.last = item
}

func (list *List) Remove(i *Item) {
	list.length--
	if i.prev != nil {
		i.prev.next = i.next
	} else {
		list.first = i.next
	}
	if i.next != nil {
		i.next.prev = i.prev
	} else {
		list.last = i.prev
	}
}
