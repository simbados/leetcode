const isPalindrome = require('./125ValidPalindrome');

test('Is valid palindrome', () => {
    const palindrome1 = "A man, a plan, a canal: Panama";
    expect(isPalindrome(palindrome1)).toBeTruthy();
    const palindrome2 = "race a car";
    expect(isPalindrome(palindrome2)).toBeFalsy();
    const palindrome3 = " ";
    expect(isPalindrome(palindrome3)).toBeTruthy();
})