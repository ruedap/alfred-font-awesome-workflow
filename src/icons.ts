import { TResponse, TResponseItem } from "./alfred/response";
import { getAllIconsObject, TIconObject } from "./assets/icons_object";

export const toResponseItem = (iconObject: TIconObject): TResponseItem => {
  return {
    uid: iconObject.name,
    title: iconObject.name,
    subtitle: `Paste class name: fa-${iconObject.name}`,
    arg: iconObject.name,
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
