import { getArgs } from "./query";

describe("getArg()", () => {
  test("empty argv", () => {
    process.argv = ["node", "jest"];

    const actual = getArgs();
    expect(actual.length).toBe(0);
  });

  test("--find", () => {
    process.argv = ["node", "jest", "--find"];

    const actual = getArgs();
    const expected = ["--find"];
    expect(actual.length).toBe(1);
    expect(actual).toStrictEqual(expected);
  });

  test("--find foo", () => {
    process.argv = ["node", "jest", "--find", "foo"];

    const actual = getArgs();
    const expected = ["--find", "foo"];
    expect(actual.length).toBe(2);
    expect(actual).toStrictEqual(expected);
  });
});
