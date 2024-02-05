export class ListNode {
  val: number
  next: ListNode | null
 constructor(val?: number, next?: ListNode | null) {
      this.val = (val===undefined ? 0 : val)
      this.next = (next===undefined ? null : next)
  }
}

export function sortList(head: ListNode | null): ListNode | null {
    if (head === null || head.next === null) return head;
    let firstPointer = head
    let secondPointer = head
    let middle = null
    while (firstPointer && firstPointer.next) {
        firstPointer = firstPointer.next.next;
        middle = secondPointer;
        secondPointer = secondPointer.next;
    }
    middle.next = null;
    const left = sortList(head);
    const right = sortList(secondPointer);
    return mergeList(left, right)
}

function mergeList(left: ListNode | null, right: ListNode | null): ListNode {
    let head = new ListNode();
    let current = head;
    let i = left;
    let j = right;
    while (i != null && j != null) {
        if (i.val < j.val) {
            current.next = i;
            i = i.next;
        } else {
            current.next = j;
            j = j.next;
        }
        current = current.next;
    }
    if (i) current.next = i;
    if (j) current.next = j;
    return head.next;
}
