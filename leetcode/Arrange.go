/*
 * 实现数学中的排列组合问题
 */
 package leetcode
 import "fmt"
type Arrage struct {
  Input []int
  Output [][]int
}
func (it *Arrage) Deal(){
  t := make([]*CombineNode,0)
  it.do(t,0) 
}
type CombineNode struct {
  Index int
  Value int
}
//使用数组加node的形式是因为最先使用map值的无序性，再次需要一定顺序
func (it *Arrage) do(index []*CombineNode, length int){
  inputLen := len(it.Input)
  if len(index) == inputLen { //最大长度
    return 
  }
  for i:= 0; i<len(it.Input); i++{
    //取所有已存在的索引
    indexIndex := make(map[int]int,0)
    for _,v := range index {
      indexIndex[v.Index] = v.Value
    }
    //判断索引是否存在
    if _,ok := indexIndex[i];!ok {//不在索引里
      //取出索引对应的值
      indexValue := make([]int,0)
      for _,v := range index {
        indexValue = append(indexValue, v.Value)
      }
    
      //添加值
      indexValue = append(indexValue, it.Input[i])
      it.Output = append(it.Output, indexValue)//添加

      //copy index,map 为传引用
      indexCopy := make([]*CombineNode,0)
      for _,v := range index {
        indexCopy = append(indexCopy, &CombineNode{v.Index,v.Value}) 
      }
      indexCopy = append(indexCopy, &CombineNode{i,it.Input[i]})
      it.do(indexCopy, length + 1)
    }
  }
}
func (it *Arrage) String() string{
  return fmt.Sprintf("需要组合的值为:%v,组合的结果为:%v\n,数量为:%d", it.Input, it.Output,len(it.Output))
}
