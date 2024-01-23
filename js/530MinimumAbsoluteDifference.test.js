const {TreeNode, getMinimumDifference} = require("./530MinimumAbsoluteDifference");
test('minimum difference 1', () => {
    const node1 = new TreeNode(1);
    const node3 = new TreeNode(3);
    const node2 = new TreeNode(2, node1, node3);
    const node6 = new TreeNode(6);
    const node4 = new TreeNode(4, node2, node6);
    expect(getMinimumDifference(node4)).toEqual(1);
})

test('minimum difference 2', () => {
    const node12 = new TreeNode(12);
    const node49 = new TreeNode(49);
    const node48 = new TreeNode(48, node12, node49);
    const node0 = new TreeNode(0);
    const node1 = new TreeNode(1, node0, node48);
    expect(getMinimumDifference(node1)).toEqual(1);
})
