package main

import "fmt"

type Book struct {
	title  string
	author string
}

func (b Book) getInfo() { //为结构体添加方法
	//(b Book):被称作方法的值接收者，它表明getInfo函数是Book结构体类型中的一个方法，即该方法与Book类型关联。b是接收者的实例名称，可以在方法内部通过变量b来访问Book结构体中的字段
	fmt.Printf("Book title: %s, Book author: %s\n", b.title, b.author)
}

type Library struct {
	books []Book
}

func (l *Library) addBook(book Book) {
	//(l *Library):被称为指针接收者，它允许方法直接操作结构体本身而非结构体的副本，一般会在结构体中写入数据时用到。这样在addBook方法内部对l.books进行的修改就会反映到原结构体上
	l.books = append(l.books, book)
}

func (l *Library) getBooks() {
	//对于一个结构体，要么都使用值接收者方法，要么都使用指针接收者方法；由于上面的方法使用了指针接收者所以这里也要使用
	for _, book := range l.books {
		book.getInfo()
	}
}

func main() {
	lib := &Library{}
	book1 := Book{title: "123", author: "dj"}
	book2 := Book{title: "456", author: "djj"}
	lib.addBook(book1)
	lib.addBook(book2)
	lib.getBooks()
}
