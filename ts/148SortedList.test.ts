import {invertTree, TreeNode} from "./226InvertBinaryTree";

it('Simplify Path', () => {
    const rootNode = new TreeNode(
        4,
        new TreeNode(
            2,
            new TreeNode(1),
            new TreeNode(3)
        ),
        new TreeNode(
            7,
            new TreeNode(6),
            new TreeNode(9)
        )
    )
    const expected = new TreeNode(
        4,
        new TreeNode(
            7,
            new TreeNode(9),
            new TreeNode(6)
        ),
        new TreeNode(
            2,
            new TreeNode(3),
            new TreeNode(1)
        )
    )
    expect(invertTree(rootNode)).toEqual(expected)
});