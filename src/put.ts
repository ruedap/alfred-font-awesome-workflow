import { getAllIconsObject, TIconObject } from "./assets/icons_object";
import { getClassName } from "./icons";
import { TQuery } from "./query";

type TQueryString = string;

const getQuery = (queryString: TQueryString) =>
  JSON.parse(queryString) as TQuery;

const getIconObject = (query: TQuery): TIconObject | undefined => {
  const allIconsObject = getAllIconsObject();
  return allIconsObject.find(
    (icon) => icon.name === query.name && icon.style === query.style
  );
};

export const putName = (queryString: TQueryString): string | null => {
  const query = getQuery(queryString);
  const iconObject = getIconObject(query);
  if (!iconObject) return null;
  return getClassName(iconObject);
};

export const putCode = (queryString: TQueryString): string | null => {
  const query = getQuery(queryString);
  const icon = getIconObject(query);
  if (!icon) return null;
  return icon.unicode;
};

export const putRef = (queryString: TQueryString): string | null => {
  const query = getQuery(queryString);
  const icon = getIconObject(query);
  if (!icon) return null;
  const ref = String.fromCodePoint(parseInt(icon.unicode, 16));
  return ref;
};

export const putUrl = (queryString: TQueryString): string | null => {
  const query = getQuery(queryString);
  const icon = getIconObject(query);
  if (!icon) return null;
  return `https://fontawesome.com/icons/${icon.name}`;
};
