import { getAllIcons, toResponseItem } from "./icons";

test("toResponseItem()", () => {
  const faIconObject = {
    name: "font-awesome",
    free: ["brands"],
    label: "Font Awesome",
    search: { terms: ["meanpath"] },
    styles: ["brands"],
    unicode: "f2b4",
  };

  const actual = toResponseItem(faIconObject);
  const expected = {
    uid: "font-awesome",
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

  const faIcon = allIcons.items.find((icon) => icon.uid === "font-awesome");
  const faIconExpected = {
    arg: "font-awesome",
    icon: { path: "./icons/font-awesome.png" },
    subtitle: "Paste class name: fa-font-awesome",
    title: "font-awesome",
    uid: "font-awesome",
  };
  expect(faIcon).toStrictEqual(faIconExpected);
});
