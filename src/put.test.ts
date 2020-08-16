import { putName } from "./put";

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
