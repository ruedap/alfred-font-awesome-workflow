import { putName, putCode } from "./put";

describe("putName()", () => {
  test("empty", () => {
    const actual = putName("");
    expect(actual).toBe(null);
  });

  test("font-awesome", () => {
    const actual = putName("font-awesome");
    expect(actual).toBe("fa-font-awesome");
  });
});

describe("putCode()", () => {
  test("empty", () => {
    const actual = putCode("");
    expect(actual).toBe(null);
  });

  test("font-awesome", () => {
    const actual = putCode("font-awesome");
    expect(actual).toBe("f2b4");
  });

  // see: https://github.com/FortAwesome/Font-Awesome/blob/master/UPGRADING.md#512x513x-to-5140
  test("edge-legacy", () => {
    const actual = putCode("edge-legacy");
    expect(actual).toBe("f978");
  });
});
