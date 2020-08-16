import { getAllIconsObject } from "./assets/icons_object";
import { TResponse } from "./alfred/response";

export const getAllIcons = (): TResponse => {
  const iconObject = getAllIconsObject();

  const allIcons = iconObject.map((value) => {
    return {
      uid: value.name,
      title: value.name,
      subtitle: value.name,
      arg: value.unicode,
      icon: {
        path: `./icons/${value.name}.png`,
      },
    };
  });

  const response: TResponse = {
    items: allIcons,
  } as TResponse;

  return response;
};
