const searchMatrix = require("./74Search2d");
it('should test search 2d', () => {
    const matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]]
    expect(searchMatrix(matrix, 3)).toBeTruthy();
});