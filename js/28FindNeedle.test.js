const strStr = require("./28FindNeedle");
describe('Find needle', () => {
    it('should return index', () => {
        expect(strStr("sadbutsad", "sad")).toBe(0)
        expect(strStr("leetcode", "leeto")).toBe(-1)
        expect(strStr("", "leeto")).toBe(-1)
    })
})