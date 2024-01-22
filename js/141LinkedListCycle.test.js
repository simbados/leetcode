const { ListNode, hasCycle } = require('./141LinkedListCycle');

test('Linked list cycle', () => {
    const listNode1 = new ListNode(3)
    const listNode2 = new ListNode(2)
    const listNode3 = new ListNode(0)
    const listNode4 = new ListNode(-4)
    listNode1.next = listNode2;
    listNode2.next = listNode3;
    listNode3.next = listNode4;
    listNode4.next = listNode2;
    expect(hasCycle(listNode1)).toEqual(true);
    const listNode5 = new ListNode(1)
    const listNode6 = new ListNode(2)
    listNode5.next = listNode6;
    listNode6.next = listNode5;
    expect(hasCycle(listNode5)).toEqual(true);
    const listNode7 = new ListNode(1)
    expect(hasCycle(listNode7)).toEqual(false);
})
