import { foo } from "./index";

test("foo()", () => {
  expect(foo("yay")).toBe("foo yay");
});
