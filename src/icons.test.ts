import { getAllIcons, toResponseItem } from "./icons";
import { TIconObject } from "./assets/icons_object";

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
    subtitle: "Paste class name: fa-font-awesome",
    arg: "font-awesome",
    icon: { path: "./icons/font-awesome.png" },
  };
  expect(actual).toStrictEqual(expected);
});

test("getAllIcons()", () => {
  const allIcons = getAllIcons();
  expect(allIcons.items.length).toBe(1601);

  const actual = allIcons.items.find((icon) => icon.arg === "font-awesome");
  const expected = {
    arg: "font-awesome",
    icon: { path: "./icons/font-awesome.png" },
    subtitle: "Paste class name: fa-font-awesome",
    title: "font-awesome",
  };
  expect(actual).toStrictEqual(expected);
});
