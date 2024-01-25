// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(
        val: i32,
        left: Option<Rc<RefCell<TreeNode>>>,
        right: Option<Rc<RefCell<TreeNode>>>,
    ) -> Self {
        TreeNode { val, left, right }
    }
}
use std::cell::RefCell;
use std::cmp::max;
use std::rc::Rc;
pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if root.is_none() {
        return 0;
    }
    let binding = root.unwrap();
    let val = binding.borrow();
    let count = 0;
    return max(
        max_depth_rec(val.left.clone(), count + 1),
        max_depth_rec(val.right.clone(), count + 1),
    );
}

pub fn max_depth_rec(node: Option<Rc<RefCell<TreeNode>>>, count: i32) -> i32 {
    if node.is_none() {
        return count;
    }
    let binding = node.unwrap();
    let val = binding.borrow();
    return max(
        max_depth_rec(val.left.clone(), count + 1),
        max_depth_rec(val.right.clone(), count + 1),
    );
}

mod tests {
    use super::*;
    use std::cell::RefCell;
    use std::rc::Rc;

    #[test]
    fn test_max_depth() {
        let rightleft = Rc::new(RefCell::new(TreeNode::new(15, None, None)));
        let rightright = Rc::new(RefCell::new(TreeNode::new(7, None, None)));
        let right = Rc::new(RefCell::new(TreeNode::new(
            20,
            Option::from(rightleft),
            Option::from(rightright),
        )));
        let left = Rc::new(RefCell::new(TreeNode::new(9, None, None)));
        let root = Rc::new(RefCell::new(TreeNode::new(
            3,
            Option::from(left),
            Option::from(right),
        )));
        assert_eq!(max_depth(Option::from(root)), 3);
    }
}
