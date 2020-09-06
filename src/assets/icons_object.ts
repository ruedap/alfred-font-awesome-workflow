import IconsJson from "@/assets/icons.json";

type TIconStyle = "brands" | "regular" | "solid";

type TIconJson = Readonly<{
  changes: string[];
  ligatures: string[];
  search: {
    terms: string[];
  };
  styles: TIconStyle[];
  free: TIconStyle[];
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
}>;

type TIconsJson = Readonly<{
  [key: string]: TIconJson;
}>;

export const getAllIconsJson = (): TIconsJson => {
  return IconsJson as TIconsJson;
};

export type TIconObject = Pick<TIconJson, "search" | "unicode" | "label"> & {
  name: string;
  style: TIconStyle;
};

export const getIconFreeObject = (
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
    } as TIconObject;
  });
};

export const getAllIconsObject = (): TIconObject[] => {
  return Object.entries(IconsJson as TIconsJson).flatMap(([key, iconJson]) => {
    return getIconFreeObject(key, iconJson);
  });
};
