import { getArgs, includesFlag, getFlagArgs } from "./query";
import { getAllIconsObject } from "./assets/icons_object";
import { search, toJson } from "./search";

const FLAG_FIND = "--find";

const main = () => {
  const args = getArgs();
  const list = getAllIconsObject();
  const keys = ["name", "search.terms"];

  if (includesFlag(args, FLAG_FIND)) {
    const query = getFlagArgs(args, FLAG_FIND);
    const searchResult = search(list, keys, query);
    console.log(toJson(searchResult));
    return;
  }

  // no args
  console.log(toJson(search(list, keys, "")));
};

main();
