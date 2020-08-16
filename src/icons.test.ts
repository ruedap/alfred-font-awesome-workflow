import { getAllIcons } from "./icons";

test("getAllIcons()", () => {
  const allIcons = getAllIcons();
  expect(allIcons.items.length).toBe(1448);

  const faIcon = allIcons.items.find((icon) => icon.uid === "font-awesome");
  const faIconExpected = {
    arg: "f2b4",
    icon: { path: "./icons/font-awesome.png" },
    subtitle: "font-awesome",
    title: "font-awesome",
    uid: "font-awesome",
  };
  expect(faIcon).toStrictEqual(faIconExpected);
});
