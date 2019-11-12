package main

import (
	"flag"
	"fmt"
	"github.com/crusj/goes/leetcode"
	. "github.com/crusj/goes/linkedList"
	. "github.com/crusj/goes/print"
	"github.com/crusj/goes/search"
	"github.com/crusj/goes/sorts"
	"github.com/crusj/goes/tree"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

func generateRandNums(max int, wg *sync.WaitGroup, oldNums *Num) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := r.Intn(max)
	oldNums.mutex.Lock()
	oldNums.value = append(oldNums.value, t)
	oldNums.mutex.Unlock()

	wg.Done()
}

var (
	which int
	num   int
)

func init() {
	flag.IntVar(&which, "which", 1, "which program")
	flag.IntVar(&num, "num", 1, "sort nums")
	flag.Parse()
}

type Num struct {
	value []int
	mutex sync.Mutex
}

func main() {
	P(runtime.NumCPU(), "cpu核心数为")
	runtime.GOMAXPROCS(runtime.NumCPU())
	switch which {
	case 1: //排序
		wg := &sync.WaitGroup{}
		P("生成随机数中.....")
		oldNums := &Num{}
		for i := 0; i < num; i++ {
			wg.Add(1)
			go generateRandNums(100000, wg, oldNums)
		}
		wg.Wait()
		//P(oldNums.value,"rand nums")
		wg.Add(1)
		start := time.Now()
		go sorts.QuickSort1(oldNums.value, wg)
		wg.Wait()
		fmt.Printf("消耗时间为%v\n", time.Since(start))
		//P(oldNums.value,"after sort")

	case 9:

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
		head = head.Delete(head, 3)
		head.Traverse(head)
	case 7: //test
		a := make(chan int, 1)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			for x := range a {
				fmt.Println(x)
			}
			wg.Done()
			fmt.Println("THE CHANNEL IS CLOSED")
		}()
		a <- 1
		close(a)
		wg.Wait()
	case 8:
		t := make([][]int, 1)
		t[0] = append(t[0], 3)
		P(t)
	case 10: //插入排序
		wg := &sync.WaitGroup{}
		P("生成随机数中.....")
		oldNums := &Num{}
		for i := 0; i < num; i++ {
			wg.Add(1)
			go generateRandNums(100000, wg, oldNums)
		}
		wg.Wait()
		start := time.Now()
		sorts.InsertSort(oldNums.value)
		P(time.Since(start), "花费时间")

	case 11: //二分查找
		wg := &sync.WaitGroup{}
		P("随机数生成中")
		oldNums := &Num{}
		for i := 0; i < num; i++ {
			wg.Add(1)
			go generateRandNums(100000, wg, oldNums)
		}
		wg.Wait()
		wg.Add(1)
		start := time.Now()
		go sorts.QuickSort1(oldNums.value, wg)
		wg.Wait()
		fmt.Printf("消耗时间为%v\n", time.Since(start))

		P(search.BinarySearch(oldNums.value, 10))
		fmt.Printf("消耗时间为%v\n", time.Since(start))
	case 12: //两数之和
		twoSum := leetcode.TwoSum{[]int{55, 24, 76, 45}, [][]int{}, 100}
		twoSum.Deal(1)
		twoSum.Empty()
		twoSum.Deal(2)
		twoSum.Empty()
		twoSum.Deal(3)
		twoSum.Empty()
	case 13: //罗马数字转证书
		romain2Integer := &leetcode.RomainToInteger{}
		romain2Integer.Romain = "MDIV"
		romain2Integer.Deal()
	case 14: //最长公共前缀
		longestCommonPrefix := &leetcode.LongestCommonPrefix{
			Input:  []string{"abc", "abcde", "abcdef"},
			Output: "",
		}
		longestCommonPrefix.Answer1()
		longestCommonPrefix.Output = ""
		longestCommonPrefix.Answer2()
		fmt.Println(longestCommonPrefix)

		input := []string{
			"zxsfas", "zxs", "zxc", "zxd",
		}
		fmt.Printf("输入字符串数组为%v,最长公共前缀为:%s\n", input, longestCommonPrefix.Answer3(input))
		fmt.Println(longestCommonPrefix.CommonPrefix("zxs", "zxc"))

		//二分法
		input = []string{
			"tzxcf", "tzxbb", "tzxc", "tzx", "tzxcftgsdf",
		}
		inputClone := make([]string, len(input))
		copy(inputClone, input)
		fmt.Printf("输入字符串数组为%v,最长公共前缀为:%s\n", inputClone, longestCommonPrefix.Answer4(input, ""))
	case 15: //括号成对
		bracket := leetcode.Bracket{
			Input:   "{}[]{[}]}",
			IsValid: false,
		}
		bracket.Deal()
		fmt.Println(bracket.IsValid)
	case 16: //测试
		sortChain := new(leetcode.SortChain)
		sortChain.ChainOne = sortChain.CreateChain([]int{
			1, 2, 4, 6,
		})
		sortChain.ChainTwo = sortChain.CreateChain([]int{
			2, 3, 7, 8,
		})
		sortChain.Deal()
		for point := sortChain.ChainOne; point != nil; {
			fmt.Println(point.Value)
			point = point.Next
		}
	case 17: //删除排序数组中的重复项目
		dsd := new(leetcode.DeleteSortArrDuplicate)
		dsd.Input = []int{
			0, 0, 1, 2, 3, 3, 4, 5, 6, 7, 8, 9,
		}
		dsd.Deal()
		fmt.Println(dsd)
	case 18: //删除数组中给定的值
		dva := new(leetcode.DeleteValueFromArray)
		dva.Input = []int{
			1, 2, 3, 4, 5, 6, 7,
		}
		fmt.Println(dva.Input)
		dva.Delete = 7
		dva.Way2()
		fmt.Println(dva)
	case 19: //实现Substr
		ss := new(leetcode.Substr)
		ss.Input = "god is satan"
		ss.Search = "god"
		ss.Deal()
		fmt.Println(ss)
		ss.Search = "sata"
		ss.Deal()
		fmt.Println(ss)
		ss.Search = "gaiyin"
		ss.Deal()
		fmt.Println(ss)
	case 20: //利用KMP算法实现Substr
		kmp := new(leetcode.Substr)
		kmp.Input = "EDaBAEDABABEDFEDEDC"
		kmp.Search = "EDED"
		kmp.Kmp()
		fmt.Println(kmp)
	case 21: //查找元素在数组中的位置，如果不存在则返回插入的位置
		SearchSortArray := new(leetcode.SearchSortArray)
		SearchSortArray.Input = []int{
			1, 3, 5, 77, 88, 111, 234, 456, 989, 1121,
		}
		SearchSortArray.Search = 4
		SearchSortArray.Deal()
		fmt.Println(SearchSortArray)
	case 22: //报数
		numberOff := new(leetcode.NumberOff)
		println(numberOff.Deal(0))
		println(numberOff.Deal(8))
	case 23: //回溯组合
		leetcode39 := new(leetcode.Leetcode39)
		leetcode39.Input = []int{2, 3, 4, 5, 6, 7, 8, 9}
		leetcode39.Target = 9
		leetcode39.Deal()
		fmt.Println(leetcode39)
	case 24: //全排列
		arrange := new(leetcode.Arrage)
		arrange.Input = []int{2, 3, 4, 5}
		arrange.Deal()
		fmt.Println(arrange)
	case 25:
		combine := new(leetcode.Combine)
		combine.Input = []int{2, 3, 4, 5}
		combine.Deal()
		fmt.Println(combine)
	case 26: //最大连续子序和
		leetcode53 := new(leetcode.Question53)
		leetcode53.Input = []int{-1, 2, 3, -2, -3, 0, -3, 5, -1, 4}
		leetcode53.Deal()
		fmt.Println(leetcode53)
	case 27: //字符串最后一个单词长度
		leetCode58 := new(leetcode.Question58)
		leetCode58.Input = "x "
		leetCode58.Deal()
		fmt.Println(leetCode58)
	case 28: //加一
		leetcode66 := new(leetcode.Question66)
		leetcode66.Input = []int{
			1, 2, 3, 4,
		}
		leetcode66.Deal()
		fmt.Println(leetcode66)
		leetcode66.Input = []int{
			9, 9, 9, 9,
		}
		leetcode66.Deal()
		fmt.Println(leetcode66)
		leetcode66.Input = []int{
			9, 9, 8, 9,
		}
		leetcode66.Deal()
		fmt.Println(leetcode66)
	case 29: //二进制求和
		leetcode67 := new(leetcode.Question67)
		leetcode67.Input = []string{"1101010", "111000"}
		leetcode67.Deal()
		fmt.Println(leetcode67)
	case 30: //求平方根
		leetCode69 := new(leetcode.Question69)
		leetCode69.Input = 8
		leetCode69.Deal()
		fmt.Println(leetCode69)
	case 31: //爬楼梯
		leetCode70 := new(leetcode.Question70)
		leetCode70.Input = 4
		leetCode70.Deal()
		fmt.Println(leetCode70)
	case 32: //删除链表中重复元素
		head := &SingleNode{"1", nil}
		head.Push(head, "2")
		head.Push(head, "2")
		head.Push(head, "3")
		head.Push(head, "3")
		leetcode83 := new(leetcode.Question83)
		leetcode83.Input = head
		leetcode83.Deal()
	case 33: //二叉树
		root := &tree.Node{Value: 1}
		node2 := &tree.Node{Value: 2}
		node3 := &tree.Node{Value: 3}
		node4 := &tree.Node{Value: 4}
		node5 := &tree.Node{Value: 5}
		node6 := &tree.Node{Value: 6}
		node7 := &tree.Node{Value: 7}
		root.SetLChild(node2)
		root.SetRChild(node3)
		node2.SetLChild(node4)
		node2.SetRChild(node5)
		node3.SetLChild(node6)
		node3.SetRChild(node7)
		//root为一棵满二叉
		tree := &tree.BinaryTree{}
		println("前序遍历")
		tree.PreOrder(root)
		println()
		println("中序遍历")
		tree.InOrder(root)
		println()
		println("后序遍历")
		tree.PostOrder(root)
	case 34: //判断两棵二叉树是否相同
		treeOne := &tree.Node{Value: 1}
		node2 := &tree.Node{Value: 2}
		node3 := &tree.Node{Value: 3}
		node4 := &tree.Node{Value: 4}
		node5 := &tree.Node{Value: 5}
		node6 := &tree.Node{Value: 6}
		node7 := &tree.Node{Value: 7}
		treeOne.SetLChild(node2)
		treeOne.SetRChild(node3)
		node2.SetLChild(node4)
		node2.SetRChild(node5)
		node3.SetLChild(node6)
		node3.SetRChild(node7)

		treeTwo := &tree.Node{Value: 1}
		node22 := &tree.Node{Value: 2}
		node23 := &tree.Node{Value: 3}
		node24 := &tree.Node{Value: 4}
		node25 := &tree.Node{Value: 5}
		node26 := &tree.Node{Value: 6}
		node27 := &tree.Node{Value: 7}
		treeTwo.SetLChild(node22)
		treeTwo.SetRChild(node23)
		node22.SetLChild(node24)
		node22.SetRChild(node25)
		node23.SetLChild(node26)
		node23.SetRChild(node27)
		leetcode100 := leetcode.Question100{}
		fmt.Println(leetcode100.Deal(treeOne, treeTwo))
	case 35: //合并有序数组
		leetcode88 := new(leetcode.Question88)
		leetcode88.ArrOne = []int{2, 3, 5, 7,}
		leetcode88.ArrTwo = []int{1, 4, 6, 7}
		leetcode88.Deal()
	case 36: //对称树的判断
		treeOne := &tree.Node{Value: 1}
		node2 := &tree.Node{Value: 2}
		node3 := &tree.Node{Value: 2}
		node4 := &tree.Node{Value: 3}
		node5 := &tree.Node{Value: 4}
		node6 := &tree.Node{Value: 4}
		node7 := &tree.Node{Value: 3}
		treeOne.SetLChild(node2)
		treeOne.SetRChild(node3)
		node2.SetLChild(node4)
		node2.SetRChild(node5)
		node3.SetLChild(node6)
		node3.SetRChild(node7)
		leetcode101 := new(leetcode.Question101)
		leetcode101.Root = treeOne
		leetcode101.Deal()
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
