export const foo = (arg: string) => `foo ${arg}`;

const expected = `{"items":[{"uid":"pos-s","title":"pos-s { position: static; }","subtitle":"Paste class name: pos-s","arg":"pos-s","icon":{"type":"fileicon","path":"icon.png"}},{"uid":"pos-r","title":"pos-r { position: relative; }","subtitle":"Paste class name: pos-r","arg":"pos-r","icon":{"type":"fileicon","path":"icon.png"}}]}`;

console.log(expected);
