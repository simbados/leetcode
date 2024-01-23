function TreeNode(val, left, right) {
     this.val = (val===undefined ? 0 : val)
     this.left = (left===undefined ? null : left);
     this.right = (right===undefined ? null : right);
}
var getMinimumDifference = function(root) {
     let minimum = Infinity;
     const arr = [];
     function sum(node) {
          if (!node) return
          sum(node.left)
          arr.push(node.val)
          sum(node.right)
     }
     sum(root)
     for (let i = 1; i < arr.length; i++) {
          const curr = arr[i] - arr[i-1];
          if (curr < minimum) {
               minimum = curr;
          }
     }
     return minimum;
};

module.exports = { getMinimumDifference, TreeNode };