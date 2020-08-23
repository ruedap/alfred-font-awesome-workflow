import { TIconObject } from "./assets/icons_object";
import { search, toJson } from "./search";

describe("search()", () => {
  const list: TIconObject[] = [
    {
      name: "AAABBB",
      search: {
        terms: ["111", "222", "333"],
      },
      unicode: "unicode",
      label: "label",
      free: ["free"],
    },
    {
      name: "BBBCCC",
      search: {
        terms: ["222", "444", "666"],
      },
      unicode: "unicode",
      label: "label",
      free: ["free"],
    },
    {
      name: "CCCDDD",
      search: {
        terms: ["333", "666", "999"],
      },
      unicode: "unicode",
      label: "label",
      free: ["free"],
    },
  ];

  test("empty query", () => {
    const keys = ["name", "search.terms"];
    const query = "";
    const actual = search(list, keys, query);
    const expected = [
      {
        item: {
          free: ["free"],
          label: "label",
          name: "AAABBB",
          search: { terms: ["111", "222", "333"] },
          unicode: "unicode",
        },
        refIndex: 0,
        score: 0,
      },
      {
        item: {
          free: ["free"],
          label: "label",
          name: "BBBCCC",
          search: { terms: ["222", "444", "666"] },
          unicode: "unicode",
        },
        refIndex: 1,
        score: 0,
      },
      {
        item: {
          free: ["free"],
          label: "label",
          name: "CCCDDD",
          search: { terms: ["333", "666", "999"] },
          unicode: "unicode",
        },
        refIndex: 2,
        score: 0,
      },
    ];

    expect(actual.length).toBe(3);
    expect(actual).toStrictEqual(expected);
  });

  test("name key", () => {
    const keys = ["name", "search.terms"];
    const query = "BBB";
    const actual = search(list, keys, query);
    const expected = [
      {
        item: {
          free: ["free"],
          label: "label",
          name: "BBBCCC",
          search: { terms: ["222", "444", "666"] },
          unicode: "unicode",
        },
        refIndex: 1,
        score: 0.001,
      },
      {
        item: {
          free: ["free"],
          label: "label",
          name: "AAABBB",
          search: { terms: ["111", "222", "333"] },
          unicode: "unicode",
        },
        refIndex: 0,
        score: 0.03,
      },
    ];

    expect(actual.length).toBe(2);
    expect(actual).toStrictEqual(expected);
  });

  test("search.terms key", () => {
    const keys = ["name", "search.terms"];
    const query = "666";
    const actual = search(list, keys, query);
    const expected = [
      {
        item: {
          free: ["free"],
          label: "label",
          name: "BBBCCC",
          search: { terms: ["222", "444", "666"] },
          unicode: "unicode",
        },
        refIndex: 1,
        score: 2.220446049250313e-16,
      },
      {
        item: {
          free: ["free"],
          label: "label",
          name: "CCCDDD",
          search: { terms: ["333", "666", "999"] },
          unicode: "unicode",
        },
        refIndex: 2,
        score: 2.220446049250313e-16,
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
        item: {
          free: ["free"],
          label: "label",
          name: "BBBCCC",
          search: { terms: ["222", "444", "666"] },
          unicode: "unicode",
        },
        refIndex: 1,
        score: 0.001,
      },
      {
        item: {
          free: ["free"],
          label: "label",
          name: "AAABBB",
          search: { terms: ["111", "222", "333"] },
          unicode: "unicode",
        },
        refIndex: 0,
        score: 0.03,
      },
    ];
    const actual = toJson(searchResult);
    const expected =
      '{"items":[{"uid":"BBBCCC","title":"BBBCCC","subtitle":"Paste class name: fa-BBBCCC","arg":"BBBCCC","icon":{"path":"./icons/BBBCCC.png"}},{"uid":"AAABBB","title":"AAABBB","subtitle":"Paste class name: fa-AAABBB","arg":"AAABBB","icon":{"path":"./icons/AAABBB.png"}}]}';

    expect(actual).toStrictEqual(expected);
  });
});
