import { putCode, putName, putRef, putUrl } from "./put";
import { TQuery } from "./query";

const exampleEncodedQuery = (query: TQuery): string => {
  return Buffer.from(JSON.stringify(query)).toString("base64");
};

const EXAMPLE_ENCODED_QUERY_EMPTY = exampleEncodedQuery({
  name: "",
  style: "solid",
} as TQuery);

const EXAMPLE_ENCODED_QUERY_FONT_AWESOME = exampleEncodedQuery({
  name: "font-awesome",
  style: "brands",
} as TQuery);

const EXAMPLE_QUERY_EDGE_LEGACY = exampleEncodedQuery({
  name: "edge-legacy",
  style: "brands",
} as TQuery);

const EXAMPLE_QUERY_ANGRY_REGULAR = exampleEncodedQuery({
  name: "angry",
  style: "regular",
} as TQuery);

const EXAMPLE_QUERY_ANGRY_SOLID = exampleEncodedQuery({
  name: "angry",
  style: "solid",
} as TQuery);

describe("putName()", () => {
  test("empty", () => {
    const actual = putName(EXAMPLE_ENCODED_QUERY_EMPTY);
    expect(actual).toBe(null);
  });

  test("font-awesome", () => {
    const actual = putName(EXAMPLE_ENCODED_QUERY_FONT_AWESOME);
    expect(actual).toBe("fab fa-font-awesome");
  });

  test("angry (regular)", () => {
    const actual = putName(EXAMPLE_QUERY_ANGRY_REGULAR);
    expect(actual).toBe("far fa-angry");
  });

  test("angry (solid)", () => {
    const actual = putName(EXAMPLE_QUERY_ANGRY_SOLID);
    expect(actual).toBe("fas fa-angry");
  });
});

describe("putCode()", () => {
  test("empty", () => {
    const actual = putCode(EXAMPLE_ENCODED_QUERY_EMPTY);
    expect(actual).toBe(null);
  });

  test("font-awesome", () => {
    const actual = putCode(EXAMPLE_ENCODED_QUERY_FONT_AWESOME);
    expect(actual).toBe("f2b4");
  });

  // see: https://github.com/FortAwesome/Font-Awesome/blob/master/UPGRADING.md#512x513x-to-5140
  test("box-tissue", () => {
    const query = exampleEncodedQuery({
      name: "box-tissue",
      style: "solid",
    } as TQuery);
    const actual = putCode(query);
    expect(actual).toBe("e05b");
  });

  // see: https://github.com/FortAwesome/Font-Awesome/blob/master/UPGRADING.md#512x513x-to-5140
  test("edge-legacy", () => {
    const actual = putCode(EXAMPLE_QUERY_EDGE_LEGACY);
    expect(actual).toBe("e078");
  });
});

describe("putRef()", () => {
  test("empty", () => {
    const actual = putRef(EXAMPLE_ENCODED_QUERY_EMPTY);
    expect(actual).toBe(null);
  });

  test("font-awesome", () => {
    const actual = putRef(EXAMPLE_ENCODED_QUERY_FONT_AWESOME);
    const expected = String.fromCodePoint(parseInt("f2b4", 16));
    expect(actual).toStrictEqual(expected);
  });

  // see: https://github.com/FortAwesome/Font-Awesome/blob/master/UPGRADING.md#512x513x-to-5140
  test("edge-legacy", () => {
    const actual = putRef(EXAMPLE_QUERY_EDGE_LEGACY);
    const expected = String.fromCodePoint(parseInt("e078", 16));
    expect(actual).toStrictEqual(expected);
  });

  test("ad", () => {
    const query = exampleEncodedQuery({
      name: "ad",
      style: "solid",
    } as TQuery);
    const actual = putRef(query);
    const expected = String.fromCodePoint(parseInt("f641", 16));
    expect(actual).toStrictEqual(expected);
  });
});

describe("putUrl()", () => {
  test("empty", () => {
    const actual = putUrl(EXAMPLE_ENCODED_QUERY_EMPTY);
    expect(actual).toBe(null);
  });

  test("font-awesome", () => {
    const actual = putUrl(EXAMPLE_ENCODED_QUERY_FONT_AWESOME);
    expect(actual).toBe("https://fontawesome.com/icons/font-awesome");
  });
});
