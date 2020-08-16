import { getAllIconsObject, TIconObject } from "./assets/icons_object";
import { TResponseItem, TResponse } from "./alfred/response";

export const toResponseItem = (iconObject: TIconObject): TResponseItem => {
  return {
    uid: iconObject.name,
    title: iconObject.name,
    subtitle: iconObject.name,
    arg: iconObject.unicode,
    icon: {
      path: `./icons/${iconObject.name}.png`,
    },
  } as TResponseItem;
};

export const getAllIcons = (): TResponse => {
  const iconObject = getAllIconsObject();

  const allIcons = iconObject.map((value) => {
    return toResponseItem(value);
  });

  const response: TResponse = {
    items: allIcons,
  } as TResponse;

  return response;
};
