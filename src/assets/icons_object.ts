import IconsJson from "@/assets/icons.json";

type TIconStyle = "brands" | "regular" | "solid";

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
  style: TIconStyle;
};

export const getIconStyleObject = (
  key: string,
  iconJson: TIconJson
): TIconObject[] => {
  return iconJson.styles.map((style) => {
    return {
      name: key,
      label: iconJson.label,
      unicode: iconJson.unicode,
      search: iconJson.search,
      style: style as TIconStyle,
      free: iconJson.free,
    } as TIconObject;
  });
};

export const getAllIconsObject = (): TIconObject[] => {
  return Object.entries(IconsJson as TIconsJson).flatMap(([key, iconJson]) => {
    return getIconStyleObject(key, iconJson);
  });
};
