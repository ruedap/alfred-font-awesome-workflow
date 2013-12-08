# Fixtures module
module Fixtures
  path = File.expand_path('./spec/fixtures/font-awesome-4-0-2-icons.yml')
  ICONS = YAML.load_file(path)['icons']

  def self.icon_ids
    ids = ICONS.map { |icon| icon['id'] }
    aliases = ICONS.map { |icon| icon['aliases'] }.compact.flatten
    (ids + aliases).sort
  end
end
