import { TResponseItem } from "./alfred/response";
import { getAllIconsObject, TIconObject } from "./assets/icons_object";

type Query = TResponseItem["uid"];

const getIconObject = (query: Query): TIconObject | undefined => {
  const allIconsObject = getAllIconsObject();
  return allIconsObject.find((icon) => icon.name === query);
};

export const putName = (query: Query): string | null => {
  const icon = getIconObject(query);
  if (!icon) return null;
  return `fa-${icon.name}`;
};

export const putCode = (query: Query): string | null => {
  const icon = getIconObject(query);
  if (!icon) return null;
  return icon.unicode;
};

export const putRef = (query: Query): string | null => {
  const icon = getIconObject(query);
  if (!icon) return null;
  const ref = String.fromCodePoint(parseInt(icon.unicode, 16));
  return ref;
};

export const putUrl = (query: Query): string | null => {
  const icon = getIconObject(query);
  if (!icon) return null;
  return `https://fontawesome.com/icons/${icon.name}`;
};
