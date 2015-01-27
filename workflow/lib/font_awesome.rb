# encoding: utf-8

require 'fileutils'
require 'htmlentities'
require 'ostruct'
require 'yaml'

# FontAwesome class
class FontAwesome
  VERSION          = '4.3.0.0'
  BUNDLE_ID        = 'com.ruedap.font-awesome'
  NONVOLATILE_DIR  = File.expand_path(BUNDLE_ID, '~/Library/Application Support/Alfred 2/Workflow Data/')
  CONFIG_FILE_PATH = File.join(NONVOLATILE_DIR, 'config.yml')

  FileUtils.mkdir_p(NONVOLATILE_DIR)
  ICONS = YAML.load_file(File.expand_path('./icons.yml'))['icons']

  # FontAwesome::Icon class
  class Icon
    attr_reader :id, :unicode

    def initialize(id)
      @id = id
      unicode = find_unicode_from_id(id)
      @unicode = unicode ? unicode : find_unicode_from_aliases(id)
    end

    def find_unicode_from_id(id)
      found_icon = ICONS.find { |icon| icon['id'] == id }
      found_icon ? found_icon['unicode'] : nil
    end

    def find_unicode_from_aliases(id)
      found_icon = ICONS.find do |icon|
        icon['aliases'].include?(id) if icon['aliases']
      end
      found_icon ? found_icon['unicode'] : nil
    end
  end

  def self.argv(argv)
    split_argv = argv[0].chomp.split('|||')
    os = OpenStruct.new
    os.icon_id = split_argv[0]
    os.icon_unicode = split_argv[1]
    os
  end

  def self.css_class_name(icon_id)
    "fa-#{icon_id}"
  end

  def self.character_reference(icon_unicode)
    HTMLEntities.new.decode("&#x#{icon_unicode};")
  end

  def self.glob_icons
    filenames = Dir.glob(File.expand_path('./icons/fa-*.png')).map do |path|
      md = /\/fa-(.+)\.png/.match(path)
      md && md[1] ? md[1] : nil
    end.compact.uniq.sort

    filenames.map { |name| Icon.new(name) }
  end

  def self.load_config
    unless File.exist? CONFIG_FILE_PATH
      open(CONFIG_FILE_PATH, 'w') { |f| YAML.dump({ 'version' => VERSION }, f) }
    end

    YAML.load_file(CONFIG_FILE_PATH)
  rescue StandardError
    nil
  end

  def self.save_config_of_recent_icons(icon_id)
    config_yaml = load_config

    if config_yaml['recent_icons'].instance_of?(Array)
      config_yaml['recent_icons'].delete(icon_id)
      config_yaml['recent_icons'].unshift(icon_id)
    else
      config_yaml['recent_icons'] = [icon_id]
    end

    dump_config(config_yaml)
  rescue StandardError
    nil
  end

  def self.url(icon_id)
    "http://fontawesome.io/icon/#{icon_id}/"
  end

  def self.check_version_number(config_yaml)
    config_yaml['version'] = VERSION unless config_yaml['version'] == VERSION
    config_yaml
  end
  private_class_method :check_version_number

  def self.dump_config(config_yaml)
    return nil unless config_yaml.instance_of?(Hash)
    check_version_number(config_yaml)
    open(CONFIG_FILE_PATH, 'w') { |f| YAML.dump(config_yaml, f) }
  end
  private_class_method :dump_config

  def initialize(queries = [])
    icons = FontAwesome.glob_icons
    @selected_icons = select!(queries, icons)
  end

  def select!(queries, icons)
    queries.each do |q|
      # use reject! for ruby 1.8 compatible
      icons.reject! { |icon| icon.id.index(q.downcase) ? false : true }
    end

    sort_by_recent_icons(icons)
  end

  def sort_by_recent_icons(icons)
    recent_icons = FontAwesome.load_config['recent_icons']
    return icons unless recent_icons.instance_of?(Array)

    recent_icons.reverse.each do |icon_id|
      found_icon = icons.find { |icon| icon.id == icon_id }
      icons.unshift(icons.delete(found_icon)) if found_icon.instance_of?(FontAwesome::Icon)
    end

    icons.compact
  end

  def item_hash(icon)
    {
      :uid => icon.id,
      :title => icon.id,
      :subtitle => "Paste class name: fa-#{icon.id}",
      :arg => "#{icon.id}|||#{icon.unicode}",
      :icon => { :type => 'default', :name => "./icons/fa-#{icon.id}.png" },
      :valid => 'yes',
    }
  end

  def item_xml(options = {})
    <<-XML
<item arg="#{options[:arg]}" uid="#{Time.new.to_i}-#{options[:uid]}">
<title>#{options[:title]}</title>
<subtitle>#{options[:subtitle]}</subtitle>
<icon>#{options[:icon][:name]}</icon>
</item>
    XML
  end

  def to_alfred
    item_xml = @selected_icons.map { |icon| item_xml(item_hash(icon)) }.join
    item_xml.gsub!(/(\r\n|\r|\n)/, '')
    puts xml = "<?xml version='1.0'?><items>#{item_xml}</items>"
    xml
  end
end
