import { toJson } from "./alfred/response";
import { getAllIcons } from "./icons";

export const foo = (arg: string) => `foo ${arg}`;

console.log(toJson(getAllIcons()));
