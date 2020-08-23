import { getArgs, getFlagArgs, includesFlag } from "./query";

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

describe("includesFlag()", () => {
  const FLAG_FIND = "--find";

  test("empty argv", () => {
    process.argv = ["node", "jest"];

    const args = getArgs();
    const actual = includesFlag(args, FLAG_FIND);
    expect(actual).toBeFalsy();
  });

  test("--findd", () => {
    process.argv = ["node", "jest", "--findd"];

    const args = getArgs();
    const actual = includesFlag(args, FLAG_FIND);
    expect(actual).toBeFalsy();
  });

  test("--find", () => {
    process.argv = ["node", "jest", "--find"];

    const args = getArgs();
    const actual = includesFlag(args, FLAG_FIND);
    expect(actual).toBeTruthy();
  });

  test("--find foo", () => {
    process.argv = ["node", "jest", "--find", "foo"];

    const args = getArgs();
    const actual = includesFlag(args, FLAG_FIND);
    expect(actual).toBeTruthy();
  });
});

describe("getFlagArgs()", () => {
  const FLAG_FIND = "--find";

  test("--find", () => {
    process.argv = ["node", "jest", "--find"];

    const args = getArgs();
    const actual = getFlagArgs(args, FLAG_FIND);
    expect(actual).toStrictEqual("");
  });

  test("--find foo", () => {
    process.argv = ["node", "jest", "--find", "foo"];

    const args = getArgs();
    const actual = getFlagArgs(args, FLAG_FIND);
    expect(actual).toStrictEqual("foo");
  });

  test("--find foo bar", () => {
    process.argv = ["node", "jest", "--find", "foo", "bar"];

    const args = getArgs();
    const actual = getFlagArgs(args, FLAG_FIND);
    expect(actual).toStrictEqual("foo bar");
  });
});
