import fs from "fs";

const INPUT_ICONS_JSON_PATH = "./font-awesome/metadata/icons.json";
const OUTPUT_ICONS_JSON_PATH = "./assets/icons.json";

const main = () => {
  fs.copyFileSync(INPUT_ICONS_JSON_PATH, OUTPUT_ICONS_JSON_PATH);
  const json = JSON.parse(String(fs.readFileSync(OUTPUT_ICONS_JSON_PATH)));
  console.log(
    `Outputted ${OUTPUT_ICONS_JSON_PATH} of ${
      Object.keys(json).length
    } icons successfully.`
  );
};

main();
