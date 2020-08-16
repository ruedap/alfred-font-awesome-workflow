import { search } from "./search";
import { TIconsObject } from "./assets/icons_object";

describe("search()", () => {
  const list: TIconsObject = [
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

    expect(actual.length).toBe(3);
    expect(actual).toStrictEqual(list);
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