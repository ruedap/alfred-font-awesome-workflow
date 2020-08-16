export const getArgs = (): string[] => {
  const args = process.argv.slice(2);
  return args;
};
