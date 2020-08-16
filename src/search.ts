import Fuse from "fuse.js";
import { TIconObject, TIconsObject } from "./assets/icons_object";
import { toResponseItem } from "./icons";
import { TResponse, TResponseItem } from "./alfred/response";

type TSearchResult = Readonly<{
  item: TIconObject;
  refIndex: number;
  score: number;
}>;

export const search = (list: TIconsObject, keys: string[], query: string) => {
  if (!query) return list;

  const options = {
    includeScore: true,
    useExtendedSearch: true,
    keys: keys,
  } as const;

  const fuse = new Fuse(list, options, undefined);
  const sr = fuse.search(query) as TSearchResult[];

  return sr;
};

export const toJson = (searchResult: TSearchResult[]): string => {
  const items: TResponseItem[] = searchResult.map(({ item }) => {
    return toResponseItem(item);
  });

  const response: TResponse = { items: items };

  return JSON.stringify(response);
};
