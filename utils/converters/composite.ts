import fg from "fast-glob";
import fs from "fs";
import imagemin from "imagemin";
import imageminPngquant from "imagemin-pngquant";
import makeDir from "make-dir";
import path from "path";
import sharp from "sharp";
import util from "util";
const writeFile = util.promisify(fs.writeFile);

const INPUT_PNG_DIR = "./assets/icons";
const INPUT_PNG_PATHS = fg.sync(`${INPUT_PNG_DIR}/**/*.png`);

const composite = (paths: string[]) => {
  return Promise.all(
    paths.map(async (p) => {
      const buffer = await sharp(p)
        .flatten({ background: "#ffffff" })
        .toBuffer();

      return await sharp(buffer).toFile(p, (err: Error) => {
        if (!err) return;
        throw new Error(`${err}: ${p}`);
      });
    })
  );
};

const minify = async (paths: string[]) => {
  // NOTE: https://github.com/imagemin/imagemin/issues/191#issuecomment-611523201
  const images = await imagemin(paths, {
    plugins: [
      imageminPngquant({
        quality: [0.2, 0.4],
      }),
    ],
  });

  return await Promise.all(
    images.map(async (v) => {
      v.destinationPath = v.sourcePath;
      await makeDir(path.dirname(v.destinationPath));
      await writeFile(v.destinationPath, v.data);
    })
  );
};

const main = async () => {
  const paths = INPUT_PNG_PATHS;

  console.log("Compositing is progress...");
  await composite(paths);
  console.log("Composited successfully.");

  console.log("Minifying is progress...");
  await minify(paths);
  console.log("Minified successfully.");
};

main();
