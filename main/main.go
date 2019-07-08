package main

import (
	"flag"
	"fmt"
	. "goes/linkedList"
	. "goes/print"
	"goes/sorts"
	"math/rand"
	"strings"
	"time"
)

func generateRandNums(max, number int) (re []int) {
	if max <= 0 {
		max = 100
	}

	for i := 0; i < number; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		t := r.Intn(max)
		re = append(re, t)
	}
	return

}

var (
	which int
)

func init() {
	flag.IntVar(&which, "which", 1, "which program")
	flag.Parse()
}

func main() {
	switch which {
	case 1: //排序
		oldNums := generateRandNums(1000, 200000)
		oldNumsCopy := make([]int, len(oldNums))
		oldNumsCopy2 := make([]int, len(oldNums))

		copy(oldNumsCopy, oldNums)
		copy(oldNumsCopy2, oldNums)

		start := time.Now()
		sorts.Bubble(&oldNums)
		fmt.Printf("冒泡排序运行时间%v秒\n", time.Since(start))

		start = time.Now()
		sorts.QuickSort2(oldNumsCopy)
		fmt.Printf("快速排序运算运行时间%v秒\n", time.Since(start))

		start = time.Now()
		sorts.QuickSort(&oldNumsCopy2)
		fmt.Printf("快速排序2运算运行时间%v\n", time.Since(start))
	case 2: //channel求和
		total := make(chan int)
		go sum(1, 3, total)
		go sum(4, 5, total)
		P(<-total+<-total, "相加的和为")
	case 3: //单项链表
		//1.创建一条元素为5个元素的单项链表
		head := &SingleNode{"张三", nil}
		head.Push(head, "李四")
		head.Push(head, "王五")
		head.Push(head, "赵六")
		head.Push(head, "李七")
		//2.遍历链表
		head.Traverse(head)
		fmt.Println("=====================================================================")
		//3.在链表头部，第3个元素，尾部各插入一个节点
		head = head.Unshift(head, "head")
		head.Rand(head, 3, "third node")
		head.Push(head, "tail node")
		head.Traverse(head)

		fmt.Println("=====================================================================")
		//4.删除node value为张三的节（如果存在），并遍历
		head = head.Delete(head, 2)
		head = head.Delete(head, 100)
		head.Traverse(head)
		//5.删除整个链表
		head = nil

	case 4: //单项循环链表
		//1.创建一条元素个数为5个单项循环链表
		head := &CircleSingleList{"张三", nil}
		head.Push(head, "李四")
		head.Push(head, "王五")
		head.Push(head, "赵六")
		head.Push(head, "谢七")
		//2.遍历链表
		head.Traverse(head)
		//3.链表的头部和链表的尾部各插入一个元素
		printEquals("遍历")
		head = head.Unshift(head, "头插")
		head.Traverse(head)
		//4.随机删除一个节点
		printEquals("随机删除")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		ri := r.Intn(8)
		P(ri, "随机数是")
		head = head.Delete(head, ri)
		head.Traverse(head)
		printEquals("随机插入")
		//5.随机插入
		head = head.Insert(head, ri, "随机")
		head.Traverse(head)
	case 5: //双向链表
		//1.创建一个元素个数为5个的双向链表
		printEquals("双向链表")
		head := &DuLinkedList{"张三", nil, nil}
		head.Push(head, "李四")
		head.Push(head, "王五")
		head.Push(head, "赵六")
		head.Push(head, "谢七")
		//2.traverse 遍历
		head.Traverse(head)
		//3.unshift and push
		printEquals("双向链表unshift")
		head = head.Unshift(head, "何二")
		head.Traverse(head)
		//4.insert
		printEquals("双向链表insert")
		head = head.Insert(head, 3, "3333")
		head.Traverse(head)
		//5.pop
		printEquals("双向链表pop")
		head.Pop(head)
		head.Traverse(head)
		//6.shift
		printEquals("双向链表shift")
		head = head.Shift(head)
		head.Traverse(head)
	case 6: //双向循环链表
		//1.创建一个元素为5个的双向链表
		printEquals("question1")
		a := "张三"
		head := &CircleDuLinkedList{&a, nil, nil}
		head.Next = head
		head.Pre = head
		b := "李四"
		head.Push(head, &b)
		c := "王五"
		head.Push(head, &c)
		d := "赵六"
		head.Push(head, &d)
		e := "谢七"
		head.Push(head, &e)
		//2.traverse
		head.Traverse(head)
		//3.push unshift
		printEquals("unshift")
		f := "UNSHIFT"
		head = head.Unshift(head, &f)
		head.Traverse(head)

		//4.pop shift
		printEquals("pop")
		head = head.Pop(head)
		head.Traverse(head)

		printEquals("shift")
		head = head.Shift(head)
		head.Traverse(head)

		//5.insert delete
		printEquals("insert")
		g := "INSERT3"
		head = head.Insert(head, 3, &g)
		head.Traverse(head)

		printEquals("delete3")
		head = head.Delete(head,3)
		head.Traverse(head)


	}
}
func sum(a, b int, total chan int) {
	time.Sleep(1e9)
	total <- a + b
}
func printEquals(tips string) {
	var index int
	if len(tips) > 4 {
		index = 100 - len(tips) - 4
	} else if len(tips) < 4 {
		index = 100 + 4 - len(tips)
	} else {
		index = 104
	}
	fmt.Println(tips + strings.Repeat("=", index))
}
