// https://leetcode-cn.com/problems/invert-binary-tree/comments/

import java.lang.System;

class Solution {
    public TreeNode invertTree(TreeNode root) {
        if(root==null) {
            return root;
        }

        if(root.left==null && root.right==null) {
            return root;
        }

        TreeNode tmpNode = root.left;
        root.left = root.right;
        root.right = tmpNode;

        if(root.left!=null && (root.left.left!=null || root.left.right!=null)) {
            invertTree(root.left);
        }
        if(root.right!=null && (root.right.left!=null || root.right.right!=null)) {
            invertTree(root.right);
        }

        return root;
    }

    public static void main(String[] args) {
        //构造树结构
        //[4,2,7,1,3,6,9]
        int[] TreeArr = new int[7];
        TreeArr[0] = 4;
        TreeArr[1] = 2;
        TreeArr[2] = 7;
        TreeArr[3] = 1;
        TreeArr[4] = 3;
        TreeArr[5] = 6;
        TreeArr[6] = 9;

//        TreeNode root = new TreeNode(4);//根节点
//        TreeNode leftnode = new TreeNode(2);
//        TreeNode rightnode = new TreeNode(7);
//        root.left = leftnode;
//        root.right = rightnode;
//
//        TreeNode leftnode_2 = new TreeNode(1);
//        TreeNode rightnode_2 = new TreeNode(3);
//        root.left.left = leftnode_2;
//        root.left.right = rightnode_2;
//
//        TreeNode leftnode_3 = new TreeNode(6);
//        TreeNode rightnode_3 = new TreeNode(9);
//        root.right.left = leftnode_3;
//        root.right.right = rightnode_3;

        TreeNode root = new TreeNode(1);//根节点
        TreeNode leftnode = new TreeNode(2);
        TreeNode rightnode = new TreeNode(3);
        root.left = leftnode;
        root.right = rightnode;

        TreeNode leftnode_2 = new TreeNode(4);
//        TreeNode rightnode_2 = new TreeNode(3);
        root.left.left = leftnode_2;
        root.left.right = null;
//
//        TreeNode leftnode_3 = new TreeNode(6);
//        TreeNode rightnode_3 = new TreeNode(9);
//        root.right.left = leftnode_3;
//        root.right.right = rightnode_3;

//        System.out.println(root.left.left.val);
        Solution solution = new Solution();
        TreeNode result = solution.invertTree(root);
        System.out.println(result.right.right);



//        TreeNode root = new TreeNode(0);//根节点
//        TreeNode leftnode = TreeNode();
//        TreeNode rightnode = TreeNode();
//        for (int i = 1; i < TreeArr.length-1; i+=2) {
//            if(leftnode==null) {
//                TreeNode leftnode = new TreeNode(TreeArr[i]);
//                root.left = leftnode;
//            }
//
//            if(rightnode==null) {
//                TreeNode rightnode = new TreeNode(TreeArr[i+1]);
//                root.right = rightnode;
//            }
//
//            TreeNode leftnode = new TreeNode(TreeArr[i]);
//            leftnode.left = leftnode;
//
//            TreeNode rightnode = new TreeNode(TreeArr[i+1]);
//            rightnode.right = rightnode;
//        }


//        TreeNode treenode_1 = TreeNode(1, null, null);
//        TreeNode treenode_3 = TreeNode(3, null, null);
//        TreeNode treenode_6 = TreeNode(6, null, null);
//        TreeNode treenode_9 = TreeNode(9, null, null);
//
//        TreeNode treenode_2 = TreeNode(2);
    }
}