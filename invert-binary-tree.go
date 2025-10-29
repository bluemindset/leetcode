/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

    queue := []*TreeNode{root}
    visited:= make(map[*TreeNode]bool)

    for len(queue) > 0{
        node := queue[0] 
        queue = queue[1:]
    
        if visited[node] {
			continue
		}

        if  node != nil && (node.Left != nil || node.Right != nil) {
            tmp := node.Left
            node.Left = node.Right
            node.Right = tmp
        }

		visited[node] = true
        if  node != nil &&  (node.Left != nil || node.Right != nil){
		    queue = append(queue, node.Left, node.Right)
        }
    }
    return root
}
