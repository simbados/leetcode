/**
 * Definition for singly-linked list.
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
    let turbo = head;
    while (turbo != null && turbo.next != null) {
        head = head.next;
        turbo = turbo.next.next;
        if (head === turbo) {
            return true
        }
    }
    return false
};

function ListNode(val) {
 this.val = val;
 this.next = null;
}

module.exports = { ListNode, hasCycle };