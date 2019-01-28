
/*
Given an n-ary tree, return the preorder traversal of its nodes' values.

For example, given a 3-ary tree:

      1
  3   2   4 
6   5
 

Return its preorder traversal as: [1,3,5,6,2,4].
*/

/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val,List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/

public class NaryTreePreOrderTraversal {
    public List<Integer> preorder(Node root) {
        // collection to store traversal result
        List<Integer> traversalResult = new ArrayList<Integer>();

        // perform preorder traversal recursively - passed by reference, the traversalResult list stores the full result as it is build through the recursion tree 
        preOrderTraversal(root, traversalResult);

        // return result
        return traversalResult;
    }

    public static void preOrderTraversal(Node node, List<Integer> traversalResult) {
        if(node != null) {
            // visit this node
            traversalResult.add(node.val);

            // visit the children, from left to right
            for(Node childNode: node.children) {
                preOrderTraversal(childNode, traversalResult);
            }
        }
    }
}

class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val,List<Node> _children) {
        val = _val;
        children = _children;
    }
}
