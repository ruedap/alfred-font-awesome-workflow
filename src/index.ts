import { getArgs, isFind, getFindArgsString } from "./query";
import { getAllIconsObject } from "./assets/icons_object";
import { search, toJson } from "./search";

const main = () => {
  const args = getArgs();
  const list = getAllIconsObject();
  const keys = ["name", "search.terms"];

  if (isFind(args)) {
    const query = getFindArgsString(args);
    const searchResult = search(list, keys, query);
    console.log(toJson(searchResult));
    return;
  }

  // no args
  console.log(toJson(search(list, keys, "")));
};

main();
