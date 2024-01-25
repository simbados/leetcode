function simplifyPath(path: string): string {
    if (path.length <= 1) {
        return path;
    }
    const stack = [];
    for (let c of path.split('/')) {
        if (c !== '' && c !== '.' && c !== '..') {
            stack.push(c);
        }
        if (c === '..') {
            stack.pop();
        }
    }
    return '/' + stack.join('/');
}

export default simplifyPath