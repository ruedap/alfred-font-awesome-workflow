export const getArgs = (): string[] => {
  const args = process.argv.slice(2);
  return args;
};

const findFlag = "--find";

export const isFind = (args: string[]): boolean => args.includes(findFlag);

export const getFindArgsString = (args: string[]): string => {
  const index = args.findIndex((value) => value === findFlag);
  return args.slice(index + 1).join(" ");
};
