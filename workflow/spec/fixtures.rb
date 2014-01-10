# Fixtures module
module Fixtures
  yaml_paths = Dir.glob('./spec/fixtures/font-awesome-*-*-*-icons.yml')
  yaml_path = File.expand_path(yaml_paths.sort.reverse.first)
  ICONS_YAML = YAML.load_file(yaml_path)['icons']

  def self.icon_ids
    ids = ICONS_YAML.map { |icon| icon['id'] }
    aliases = ICONS_YAML.map { |icon| icon['aliases'] }.compact.flatten
    (ids + aliases).sort
  end
end
