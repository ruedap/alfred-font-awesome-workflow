module Fixtures
  path = File.expand_path('./test/fixtures/font-awesome-4-0-0-icons.yml')
  ICONS = YAML.load_file(path).compact.uniq.sort

  def self.icons
    ICONS
  end
end
