import { getAllIconsObject, TIconObject } from "./assets/icons_object";
import { getClassName } from "./icons";
import { TQuery } from "./query";

export type TEncodedQuery = string; // NOTE: TQuery encoded by Base64

const decodeQuery = (encodedQuery: TEncodedQuery): TQuery => {
  const queryJsonString = Buffer.from(encodedQuery, "base64").toString();

  return JSON.parse(queryJsonString) as TQuery;
};

const getIconObject = (query: TQuery): TIconObject | undefined => {
  const allIconsObject = getAllIconsObject();
  return allIconsObject.find(
    (icon) => icon.name === query.name && icon.style === query.style
  );
};

export const putName = (encodedQuery: TEncodedQuery): string | null => {
  const query = decodeQuery(encodedQuery);
  const iconObject = getIconObject(query);
  if (!iconObject) return null;
  return getClassName(iconObject);
};

export const putCode = (encodedQuery: TEncodedQuery): string | null => {
  const query = decodeQuery(encodedQuery);
  const iconObject = getIconObject(query);
  if (!iconObject) return null;
  return iconObject.unicode;
};

export const putRef = (encodedQuery: TEncodedQuery): string | null => {
  const query = decodeQuery(encodedQuery);
  const iconObject = getIconObject(query);
  if (!iconObject) return null;
  const ref = String.fromCodePoint(parseInt(iconObject.unicode, 16));
  return ref;
};

export const putUrl = (encodedQuery: TEncodedQuery): string | null => {
  const query = decodeQuery(encodedQuery);
  const iconObject = getIconObject(query);
  if (!iconObject) return null;
  return `https://fontawesome.com/icons/${iconObject.name}`;
};
