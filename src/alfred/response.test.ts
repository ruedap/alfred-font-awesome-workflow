import { toJson, TResponse } from "./response";

const response: TResponse = {
  items: [
    {
      uid: "f26e",
      title: "500px",
      subtitle: "Paste class name: fa-500px",
      arg: "f26e",
      icon: {
        type: "fileicon",
        path: "./icons/500px.png",
      },
    },
    {
      uid: "f271",
      title: "calendar-plus",
      subtitle: "Paste class name: fa-calendar-plus",
      arg: "f217",
      icon: {
        type: "fileicon",
        path: "./icons/calendar-plus.png",
      },
    },
  ],
};

test("toJson", () => {
  const actual = toJson(response);
  const expected =
    '{"items":[{"uid":"f26e","title":"500px","subtitle":"Paste class name: fa-500px","arg":"f26e","icon":{"type":"fileicon","path":"./icons/500px.png"}},{"uid":"f271","title":"calendar-plus","subtitle":"Paste class name: fa-calendar-plus","arg":"f217","icon":{"type":"fileicon","path":"./icons/calendar-plus.png"}}]}';
  expect(actual).toBe(expected);
});
