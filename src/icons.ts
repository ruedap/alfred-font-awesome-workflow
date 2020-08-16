import { getIconsObject } from "./assets/icons_object";
import { TResponse } from "./alfred/response";

export const getAllIcons = (): TResponse => {
  const iconObject = getIconsObject();

  const allIcons = Object.entries(iconObject).map(([key, value]) => {
    return {
      uid: key,
      title: key,
      subtitle: key,
      arg: value.unicode,
      icon: {
        path: `./icons/${key}.png`,
      },
    };
  });

  const response: TResponse = {
    items: allIcons,
  } as TResponse;

  return response;
};
