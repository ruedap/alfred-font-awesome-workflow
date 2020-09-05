import fg from "fast-glob";
import path from "path";

// FIXME
// eslint-disable-next-line
const svgexport = require("svgexport");
// eslint-disable-next-line
const mkdirp = require("mkdirp");

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

const makeDir = (svgs: TSvg[]): void => {
  const outputPath = svgs.map((svg) => path.parse(svg.output).dir);

  mkdirp.sync(outputPath);
};

const convert = async (svgs: TSvg[]) => {
  return svgexport.render(svgs, process);
};

const main = () => {
  console.log("start svg2png");

  makeDir(SVGS);

  // convert(SVGS.slice(0, 10)); // DEBUG
  convert(SVGS);
};

main();
