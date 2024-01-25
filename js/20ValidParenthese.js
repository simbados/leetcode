var isValid = function(s) {
    const mapping = {
        "{": "}",
        "[": "]",
        "(": ")",
    }
    const stack = []
    for (let c of s.split('')) {
        if (Object.keys(mapping).includes(c)) {
            stack.push(c);
        } else {
            const opening = stack.pop();
            if (mapping[opening] !== c) {
                return false;
            }
        }
    }
    return stack.length === 0;
};

module.exports = isValid;