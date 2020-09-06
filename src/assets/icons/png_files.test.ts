import fg from "fast-glob";

const brandsDir = (path: string) => `${path}/brands`;
const regularDir = (path: string) => `${path}/regular`;
const solidDir = (path: string) => `${path}/solid`;
const INPUT_DIR = "./font-awesome/svgs";
const INPUT_DIR_BRANDS = brandsDir(INPUT_DIR);
const INPUT_DIR_REGULAR = regularDir(INPUT_DIR);
const INPUT_DIR_SOLID = solidDir(INPUT_DIR);

const OUTPUT_DIR = "./assets/icons";
const OUTPUT_DIR_BRANDS = brandsDir(OUTPUT_DIR);
const OUTPUT_DIR_REGULAR = regularDir(OUTPUT_DIR);
const OUTPUT_DIR_SOLID = solidDir(OUTPUT_DIR);

describe(OUTPUT_DIR_BRANDS, () => {
  test("length", () => {
    const actual: string[] = fg.sync(`${OUTPUT_DIR_BRANDS}/*.png`);
    const expected: string[] = fg.sync(`${INPUT_DIR_BRANDS}/*.svg`);
    expect(actual.length).toBe(expected.length);
  });
});

describe(OUTPUT_DIR_REGULAR, () => {
  test("length", () => {
    const actual: string[] = fg.sync(`${OUTPUT_DIR_REGULAR}/*.png`);
    const expected: string[] = fg.sync(`${INPUT_DIR_REGULAR}/*.svg`);
    expect(actual.length).toBe(expected.length);
  });
});

describe(OUTPUT_DIR_SOLID, () => {
  test("length", () => {
    const actual: string[] = fg.sync(`${OUTPUT_DIR_SOLID}/*.png`);
    const expected: string[] = fg.sync(`${INPUT_DIR_SOLID}/*.svg`);
    expect(actual.length).toBe(expected.length);
  });
});
