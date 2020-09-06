import { TResponse, TResponseItem } from "./alfred/response";
import { getAllIconsObject, TIconObject } from "./assets/icons_object";
import { TQuery } from "./query";

export const getClassName = (iconObject: TIconObject): string => {
  const prefix = (style: TIconObject["style"]) => {
    switch (style) {
      case "brands":
        return "fab";
      case "regular":
        return "far";
      case "solid":
        return "fas";
      default:
        return style as never;
    }
  };

  const className = `${prefix(iconObject.style)} fa-${iconObject.name}`;

  return className;
};

export const toResponseItem = (iconObject: TIconObject): TResponseItem => {
  const argObj = { name: iconObject.name, style: iconObject.style } as TQuery;
  const arg = JSON.stringify(argObj);

  return {
    title: iconObject.name,
    subtitle: `Paste class name: ${getClassName(iconObject)}`,
    arg: arg,
    icon: {
      path: `./icons/${iconObject.style}/${iconObject.name}.png`,
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
