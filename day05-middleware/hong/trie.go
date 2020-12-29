package hong

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由， 例如： /p/：lang
	part     string  //路由中的一部分，例如： ：lang
	children []*node // 子节点 ，例如  [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有'：' 或'*' 时为true
}


func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

//  第一个匹配成功的节点，用于插入
// param  part  路由规则匹配部分
// 如果当前节点的子节点 的part 与传入的part相同，则返回这个子节点，否则返回nil
func (n *node)matchChild(part string) *node  {

	for _, child := range n.children {
		if child.part==part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {

	nodes :=make([]*node,0)
	for _, child := range n.children {
		if child.part==part ||child.isWild {
			nodes=append(nodes,child)
		}
	}
	return nodes
}


// 递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个
func (n *node)insert(pattern string,parts []string,height int)  {
	if len(parts)==height {
		n.pattern=pattern
		return
	}

	part :=parts[height]

	child :=n.matchChild(part)// 找到符合匹配的子节点
	if child==nil { //如果当前子节点没有找到，则追加一个节点到子节点，即新增一层路径
		child=&node{part:     part, isWild:   part[0]==':' || part[0]=='*'}
		n.children=append(n.children,child)

	}
	child.insert(pattern,parts,height+1)

}

func (n *node) search(parts []string, height int) *node {

	if len(parts) == height || strings.HasPrefix(n.part,"*") {
		if n.pattern=="" {
			return nil
		}
		return n
	}

	part :=parts[height]

	children :=n.matchChildren(part)

	for _, child := range children {
		result :=child.search(parts,height+1)
		if result!=nil {
			return  result

		}
	}
	return nil
}

func (n *node) travel(list *([]*node))  {
	if n.pattern!="" {
		*list =append(*list,n)

	}
	for _, child := range n.children {
		child.travel(list)
	}
}