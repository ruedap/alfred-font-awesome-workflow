// ref: https://www.alfredapp.com/help/workflows/inputs/script-filter/json/

type TMods = Readonly<{
  alt: {
    valid: boolean;
    arg: string;
    subtitle: string;
  };
  cmd: {
    valid: boolean;
    arg: string;
    subtitle: string;
  };
}>;

type TText = Readonly<{
  copy: string;
  largetype: string;
}>;

export type TResponseItem = Readonly<{
  uid: string;
  title: string;
  subtitle: string;
  arg: string; // recommended
  icon: {
    type?: "fileicon" | "filetype";
    path: string;
  };
  valid?: boolean;
  match?: string;
  autocomplete?: string; // recommended
  type?: "default" | "file" | "file:skipcheck";
  mods?: TMods;
  text?: TText;
  quicklookurl?: string;
}>;

export type TResponse = Readonly<{
  items: TResponseItem[];
}>;

export const toJson = (response: TResponse): string => {
  return JSON.stringify(response);
};
