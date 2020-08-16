export const getArgs = (): string[] => {
  const args = process.argv.slice(2);
  return args;
};

type Args = string[];
type Flag = string;

export const includesFlag = (args: Args, flag: Flag): boolean =>
  args.includes(flag);

export const getFlagArgs = (args: Args, flag: string): string => {
  const index = args.findIndex((value) => value === flag);
  return args.slice(index + 1).join(" ");
};
