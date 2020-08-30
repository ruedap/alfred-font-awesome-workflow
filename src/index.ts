import { getAllIconsObject } from "./assets/icons_object";
import { putCode, putName, putRef, putUrl } from "./put";
import { getArgs, getFlagArgs, includesFlag } from "./query";
import { search, toJson } from "./search";

const FLAG = {
  FIND: "--find",
  PUT: {
    NAME: "--put-name",
    CODE: "--put-code",
    REF: "--put-ref",
    URL: "--put-url",
  },
};

const main = () => {
  const args = getArgs();
  const list = getAllIconsObject();

  if (includesFlag(args, FLAG.FIND)) {
    const query = getFlagArgs(args, FLAG.FIND);
    const searchResult = search(list, query);
    console.log(toJson(searchResult)); // NOTE: Output to Alfred
    return;
  }

  if (includesFlag(args, FLAG.PUT.NAME)) {
    const query = getFlagArgs(args, FLAG.PUT.NAME);
    const result = putName(query);
    result && console.log(result); // NOTE: Output to Alfred
    return;
  }

  if (includesFlag(args, FLAG.PUT.CODE)) {
    const query = getFlagArgs(args, FLAG.PUT.CODE);
    const result = putCode(query);
    result && console.log(result); // NOTE: Output to Alfred
    return;
  }

  if (includesFlag(args, FLAG.PUT.REF)) {
    const query = getFlagArgs(args, FLAG.PUT.REF);
    const result = putRef(query);
    result && console.log(result); // NOTE: Output to Alfred
    return;
  }

  if (includesFlag(args, FLAG.PUT.URL)) {
    const query = getFlagArgs(args, FLAG.PUT.URL);
    const result = putUrl(query);
    result && console.log(result); // NOTE: Output to Alfred
    return;
  }

  // NOTE: No flags
  console.log(toJson(search(list, ""))); // NOTE: Output to Alfred
};

main();
