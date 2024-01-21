var isPalindrome = function(s) {
    const initial = s.trim().toLowerCase().replaceAll(/[^a-zA-Z0-9]/g, '');
    return initial === initial.split('').reverse().join('')
};

module.exports = isPalindrome;