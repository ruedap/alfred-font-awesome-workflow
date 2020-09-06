import { TIconObject } from "./assets/icons_object";
import { search, toJson } from "./search";

const EXAMPLE_LIST_1: TIconObject = {
  name: "AAABBB",
  search: {
    terms: ["111", "222", "333"],
  },
  style: "brands",
  unicode: "unicode",
  label: "label",
};

const EXAMPLE_LIST_2: TIconObject = {
  name: "BBBCCC",
  search: {
    terms: ["222", "444", "666"],
  },
  style: "solid",
  unicode: "unicode",
  label: "label",
};

const EXAMPLE_LIST_3: TIconObject = {
  name: "CCCDDD",
  search: {
    terms: ["333", "666", "999"],
  },
  style: "regular",
  unicode: "unicode",
  label: "label",
};

const EXAMPLE_LIST: TIconObject[] = [
  EXAMPLE_LIST_1,
  EXAMPLE_LIST_2,
  EXAMPLE_LIST_3,
];

describe("search()", () => {
  test("empty query", () => {
    const query = "";
    const actual = search(EXAMPLE_LIST, query);
    const expected = [
      {
        item: EXAMPLE_LIST_1,
        refIndex: 0,
        score: 0,
      },
      {
        item: EXAMPLE_LIST_2,
        refIndex: 1,
        score: 0,
      },
      {
        item: EXAMPLE_LIST_3,
        refIndex: 2,
        score: 0,
      },
    ];

    expect(actual.length).toBe(3);
    expect(actual).toStrictEqual(expected);
  });

  test("name key", () => {
    const query = "BBB";
    const actual = search(EXAMPLE_LIST, query);
    const expected = [
      {
        item: EXAMPLE_LIST_2,
        refIndex: 1,
        score: 0.007943282347242817,
      },
      {
        item: EXAMPLE_LIST_1,
        refIndex: 0,
        score: 0.08589836120408871,
      },
    ];

    expect(actual.length).toBe(2);
    expect(actual).toStrictEqual(expected);
  });

  test("search.terms key", () => {
    const query = "666";
    const actual = search(EXAMPLE_LIST, query);
    const expected = [
      {
        item: EXAMPLE_LIST_2,
        refIndex: 1,
        score: 0.000020134092876783674,
      },
      {
        item: EXAMPLE_LIST_3,
        refIndex: 2,
        score: 0.000020134092876783674,
      },
    ];

    expect(actual.length).toBe(2);
    expect(actual).toStrictEqual(expected);
  });
});

describe("toJson()", () => {
  test("", () => {
    const searchResult = [
      {
        item: EXAMPLE_LIST_2,
        refIndex: 1,
        score: 0.001,
      },
      {
        item: EXAMPLE_LIST_1,
        refIndex: 0,
        score: 0.03,
      },
    ];
    const actual = toJson(searchResult);
    const expected =
      '{"items":[{"title":"BBBCCC","subtitle":"Paste class name: fa-BBBCCC","arg":"BBBCCC","icon":{"path":"./icons/BBBCCC.png"}},{"title":"AAABBB","subtitle":"Paste class name: fa-AAABBB","arg":"AAABBB","icon":{"path":"./icons/AAABBB.png"}}]}';

    expect(actual).toStrictEqual(expected);
  });
});
