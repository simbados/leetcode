import {ListNode, sortList} from "./148SortedList";

it('Simplify Path', () => {
    const head = new ListNode(4, new ListNode(2, new ListNode(1, new ListNode(3))))
    const expected = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4))))
    expect(sortList(head)).toEqual(expected);
});