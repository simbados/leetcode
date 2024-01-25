import simplifyPath from "./71SimplifyPath";

it('Simplify Path', () => {
    expect(simplifyPath("/../")).toEqual('/');
    expect(simplifyPath("/home/")).toEqual('/home');
    expect(simplifyPath("/home//foo/")).toEqual('/home/foo');
    expect(simplifyPath("/a/./b/../../c/")).toEqual('/c');
    expect(simplifyPath("/...")).toEqual('/...');
})