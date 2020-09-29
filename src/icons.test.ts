import { TIconObject } from "./assets/icons_object";
import { getAllIcons, toResponseItem } from "./icons";

test("toResponseItem()", () => {
  const iconObject: TIconObject = {
    name: "font-awesome",
    style: "brands",
    label: "Font Awesome",
    search: { terms: ["meanpath"] },
    unicode: "f2b4",
  };

  const actual = toResponseItem(iconObject);
  const expected = {
    title: "font-awesome",
    subtitle: "Paste class name: fab fa-font-awesome",
    arg: "eyJuYW1lIjoiZm9udC1hd2Vzb21lIiwic3R5bGUiOiJicmFuZHMifQ==",
    icon: { path: "./icons/brands/font-awesome.png" },
  };
  expect(actual).toStrictEqual(expected);
});

test("getAllIcons()", () => {
  const allIcons = getAllIcons();
  expect(allIcons.items.length).toBe(1612);

  const actual = allIcons.items.find((icon) => icon.title === "font-awesome");
  const expected = {
    title: "font-awesome",
    subtitle: "Paste class name: fab fa-font-awesome",
    arg: "eyJuYW1lIjoiZm9udC1hd2Vzb21lIiwic3R5bGUiOiJicmFuZHMifQ==",
    icon: { path: "./icons/brands/font-awesome.png" },
  };
  expect(actual).toStrictEqual(expected);
});
