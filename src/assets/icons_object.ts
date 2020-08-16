import IconsJson from "@/assets/icons.json";

type TIcon = Readonly<{
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

type TIcons = Readonly<{
  [key: string]: TIcon;
}>;

export const getIconsObject = (): TIcons => {
  return IconsJson as TIcons;
};
