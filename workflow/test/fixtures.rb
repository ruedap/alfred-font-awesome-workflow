module Fixtures
  path = File.expand_path('./test/fixtures/font-awesome-4-0-2-icons.yml')
  ICONS = YAML.load_file(path)

  def self.icon_ids
    ids = ICONS['icons'].map { |icon| icon['id'] }
    aliases = ICONS['icons'].map { |icon| icon['aliases'] }.compact.flatten
    (ids + aliases).sort
  end
end
