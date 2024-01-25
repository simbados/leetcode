const isValid = require("./20ValidParenthese");
test('Remove duplicate', () => {
    expect(isValid("()")).toBe(true);
    expect(isValid("(]")).toBe(false);
    expect(isValid("()[]{}")).toBe(true);
    expect(isValid("([]{})")).toBe(true);
    expect(isValid("([]{}")).toBe(false);
})