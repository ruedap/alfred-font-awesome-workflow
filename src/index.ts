import { getAllIconsObject } from "./assets/icons_object";
import { putCode, putName, putRef, putUrl } from "./put";
import { getArgs, getFlagArgs, includesFlag } from "./query";
import { search, toJson } from "./search";

const FLAG_FIND = "--find";
const FLAG_PUT_NAME = "--put-name";
const FLAG_PUT_CODE = "--put-code";
const FLAG_PUT_REF = "--put-ref";
const FLAG_PUT_URL = "--put-url";

const main = () => {
  const args = getArgs();
  const list = getAllIconsObject();
  const keys = ["name", "search.terms"];

  if (includesFlag(args, FLAG_FIND)) {
    const query = getFlagArgs(args, FLAG_FIND);
    const searchResult = search(list, keys, query);
    console.log(toJson(searchResult)); // NOTE: Output to Alfred
    return;
  }

  if (includesFlag(args, FLAG_PUT_NAME)) {
    const query = getFlagArgs(args, FLAG_PUT_NAME);
    const result = putName(query);
    result && console.log(result); // NOTE: Output to Alfred
    return;
  }

  if (includesFlag(args, FLAG_PUT_CODE)) {
    const query = getFlagArgs(args, FLAG_PUT_CODE);
    const result = putCode(query);
    result && console.log(result); // NOTE: Output to Alfred
    return;
  }

  if (includesFlag(args, FLAG_PUT_REF)) {
    const query = getFlagArgs(args, FLAG_PUT_REF);
    const result = putRef(query);
    result && console.log(result); // NOTE: Output to Alfred
    return;
  }

  if (includesFlag(args, FLAG_PUT_URL)) {
    const query = getFlagArgs(args, FLAG_PUT_URL);
    const result = putUrl(query);
    result && console.log(result); // NOTE: Output to Alfred
    return;
  }

  // NOTE: No flags
  console.log(toJson(search(list, keys, ""))); // NOTE: Output to Alfred
};

main();
