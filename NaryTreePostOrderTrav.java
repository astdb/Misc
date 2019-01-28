
/*
Given an n-ary tree, return the postorder traversal of its nodes' values.

For example, given a 3-ary tree:

      1
  3   2   4 
6   5


 

Return its postorder traversal as: [5,6,3,2,4,1].


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

import java.util.*;

public class NaryTreePostOrderTrav {
    public List<Integer> postorder(Node root) {
        List<Integer> traversalResult = new ArrayList<Integer>();

        // post-order traversal includes visiting this node's children from left to right, then visiting the node value
        postOrderTraversal(root, traversalResult);

        return traversalResult;
    }

    public static void postOrderTraversal(Node root, List<Integer> traversalResult) {
        if (root != null) {
            // visit child nodes from  left to right
            for(Node node: root.children) {
                postOrderTraversal(node, traversalResult);
            }

            // visit this node
            traversalResult.add(root.val);            
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
