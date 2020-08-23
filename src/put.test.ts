import { putCode, putName, putRef, putUrl } from "./put";

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
  test("box-tissue", () => {
    const actual = putCode("box-tissue");
    expect(actual).toBe("e05b");
  });

  // see: https://github.com/FortAwesome/Font-Awesome/blob/master/UPGRADING.md#512x513x-to-5140
  test("edge-legacy", () => {
    const actual = putCode("edge-legacy");
    expect(actual).toBe("e078");
  });
});

describe("putRef()", () => {
  test("empty", () => {
    const actual = putRef("");
    expect(actual).toBe(null);
  });

  test("font-awesome", () => {
    const actual = putRef("font-awesome");
    const expected = String.fromCodePoint(parseInt("f2b4", 16));
    expect(actual).toStrictEqual(expected);
  });

  // see: https://github.com/FortAwesome/Font-Awesome/blob/master/UPGRADING.md#512x513x-to-5140
  xtest("edge-legacy", () => {
    const actual = putCode("edge-legacy");
    const expected = String.fromCodePoint(parseInt("f978", 16));
    expect(actual).toStrictEqual(expected);
  });

  test("ad", () => {
    const actual = putRef("ad");
    const expected = String.fromCodePoint(parseInt("f641", 16));
    expect(actual).toStrictEqual(expected);
  });
});

describe("putUrl()", () => {
  test("empty", () => {
    const actual = putUrl("");
    expect(actual).toBe(null);
  });

  test("font-awesome", () => {
    const actual = putUrl("font-awesome");
    expect(actual).toBe("https://fontawesome.com/icons/font-awesome");
  });
});
