import { getArgs, getFlagArgs, includesFlag } from "./arg";
import { getAllIconsObject } from "./assets/icons_object";
import { putCode, putName, putRef, putUrl } from "./put";
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

const output = (str: string): void => {
  process.stdout.write(str); // NOTE: Output to Alfred
};

const main = () => {
  const args = getArgs();
  const list = getAllIconsObject();

  if (includesFlag(args, FLAG.FIND)) {
    const query = getFlagArgs(args, FLAG.FIND);
    const searchResult = search(list, query);
    output(toJson(searchResult));
    return;
  }

  if (includesFlag(args, FLAG.PUT.NAME)) {
    const query = getFlagArgs(args, FLAG.PUT.NAME);
    const result = putName(query);
    result && output(result);
    return;
  }

  if (includesFlag(args, FLAG.PUT.CODE)) {
    const query = getFlagArgs(args, FLAG.PUT.CODE);
    const result = putCode(query);
    result && output(result);
    return;
  }

  if (includesFlag(args, FLAG.PUT.REF)) {
    const query = getFlagArgs(args, FLAG.PUT.REF);
    const result = putRef(query);
    result && output(result);
    return;
  }

  if (includesFlag(args, FLAG.PUT.URL)) {
    const query = getFlagArgs(args, FLAG.PUT.URL);
    const result = putUrl(query);
    result && output(result);
    return;
  }

  // NOTE: No flags
  output(toJson(search(list, "")));
};

main();
