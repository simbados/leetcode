var canConstruct = function(ransomNote, magazine) {
    const map = new Map();
    for (const s of ransomNote) {
        map.set(s, (map.get(s) || 0) + 1)
    }
    for (const a of magazine) {
        if (map.has(a)) {
            map.set(a, (map.get(a) - 1))
        }
    }
    return Array.from(map.values()).reduce((prev, curr) => prev && curr <= 0, true);
};

module.exports = canConstruct;