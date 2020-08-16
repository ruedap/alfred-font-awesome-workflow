import IconsJson from "@/assets/icons.json";

type TIconJson = Readonly<{
  changes: string[];
  ligatures: string[];
  search: {
    terms: string[];
  };
  styles: string[];
  unicode: string;
  label: string;
  voted?: boolean;
  svg: {
    [key: string]: {
      last_modified: number;
      raw: string;
      viewBox: string[];
      width: number;
      height: number;
      path: string;
    };
  };
  free: string[];
}>;

type TIconsJson = Readonly<{
  [key: string]: TIconJson;
}>;

export const getAllIconsJson = (): TIconsJson => {
  return IconsJson as TIconsJson;
};

export type TIconObject = Pick<
  TIconJson,
  "search" | "unicode" | "label" | "free"
> & {
  name: string;
};

export const getAllIconsObject = (): TIconObject[] => {
  return Object.entries(IconsJson as TIconsJson).map(([key, value]) => {
    return {
      name: key,
      label: value.label,
      unicode: value.unicode,
      search: value.search,
      styles: value.styles,
      free: value.free,
    } as TIconObject;
  });
};
