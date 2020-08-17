require 'yaml'

path = "./assets/icons.yml"
yaml = YAML.load_file(path)
icons = []

yaml.each do |y|
  icon = {}
  icon["name"] = y.last["label"]
  icon["id"] = y.first
  icon["unicode"] = y.last["unicode"]
  icon["created"] = y.last["changes"].first
  icon["filter"] = y.last["search"]["terms"]
  icon["categories"] = "unknown"

  icons.push icon
end

YAML.dump({"icons" => icons}, File.open(path, 'w'))
