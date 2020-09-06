import fg from "fast-glob";
import makeDir from "make-dir";
import path from "path";

// FIXME
// eslint-disable-next-line
const svgexport = require("svgexport");

const INPUT_DIR = "./font-awesome/svgs";
const INPUT_SVG_PATHS = fg.sync(`${INPUT_DIR}/**/*.svg`);
const OUTPUT_DIR = "./assets/icons";
const PARAMS = "pad 128:128";

const outputPath = (inputPath: string): string => {
  const outputDir = path.parse(inputPath.replace(INPUT_DIR, OUTPUT_DIR)).dir;
  const name = path.parse(inputPath).name;
  return `${outputDir}/${name}.png`;
};

type TSvg = Readonly<{
  input: string; // e.g. "./font-awesome/svgs/**/*.svg"
  output: string; // e.g. "./assets/icons/**/*.png pad 300:300"
}>;

const SVGS: TSvg[] = INPUT_SVG_PATHS.map((inputPath: string) => {
  return {
    input: inputPath,
    output: `${outputPath(inputPath)} ${PARAMS}`,
  } as TSvg;
});

const makeStylesDir = async (svgs: TSvg[]): Promise<string[]> => {
  const outputPaths = svgs.map((svg) => path.parse(svg.output).dir);
  const outputDirs = [...new Set(outputPaths)];

  return await Promise.all(outputDirs.map(async (dir) => await makeDir(dir)));
};

const convert = (svgs: TSvg[]) => {
  svgexport.render(svgs, process);
};

const main = async () => {
  console.log("start svg2png");

  await makeStylesDir(SVGS);

  // convert(SVGS.slice(0, 10)); // DEBUG
  convert(SVGS);
};

main();
