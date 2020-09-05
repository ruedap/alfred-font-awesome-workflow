import fg from "fast-glob";
import path from "path";

// FIXME
// eslint-disable-next-line
const sharp = require("sharp");
// eslint-disable-next-line
const imagemin = require("imagemin");
// eslint-disable-next-line
const imageminPngquant = require("imagemin-pngquant");

const INPUT_PNG_DIR = "./assets/icons";
const INPUT_PNG_PATHS = fg.sync(`${INPUT_PNG_DIR}/**/*.png`);

const composite = async (paths: string[]) => {
  return await Promise.all(
    paths.map(async (p) => {
      const buffer = await sharp(p)
        .flatten({ background: "#ffffff" })
        .toBuffer();

      await sharp(buffer).toFile(p, (err: Error, info: string) => {
        if (!err) return;
        throw new Error(`${err}: ${p}`);
      });
    })
  );
};

const minify = async (paths: string[]) => {
  const destination = path.parse(paths[0]).dir;
  return await imagemin(paths, {
    destination,
    plugins: [
      imageminPngquant({
        quality: [0.2, 0.4],
      }),
    ],
  });
};

const main = async () => {
  const paths = INPUT_PNG_PATHS;

  console.log("start composite");
  await composite(paths);
  console.log("end composite");

  console.log("start minify");
  await minify(paths);
  console.log("end minify");

  console.log("composite is complete");
};

main();
