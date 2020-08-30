import { getAllIcons, toResponseItem } from "./icons";

test("toResponseItem()", () => {
  const iconObject = {
    name: "font-awesome",
    free: ["brands"],
    label: "Font Awesome",
    search: { terms: ["meanpath"] },
    styles: ["brands"],
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
  expect(allIcons.items.length).toBe(1448);

  const actual = allIcons.items.find((icon) => icon.arg === "font-awesome");
  const expected = {
    arg: "font-awesome",
    icon: { path: "./icons/font-awesome.png" },
    subtitle: "Paste class name: fa-font-awesome",
    title: "font-awesome",
  };
  expect(actual).toStrictEqual(expected);
});
