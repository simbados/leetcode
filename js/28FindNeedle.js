var strStr = function(haystack, needle) {
    let count = 0
    while (!haystack.startsWith(needle) && haystack.length > 0) {
        haystack = haystack.substring(1)
        count++
    }
    return haystack.length === 0 ? -1 : count
};

module.exports = strStr
