class FontAwesome
  attr_reader :icons

  class Icon
    attr_reader :name

    def initialize(name)
      @name = name
    end
  end

  def initialize(query = '')
    icon_names = glob_icon_names
    @icons = icon_names.map { |name| Icon.new(name) }
    select!(query.split)
  end

  def select!(queries, icons = @icons)
    queries.each do |q|
      # use reject! for ruby 1.8 compatible
      icons.reject! { |icon| icon.name.index(q.downcase) ? false : true }
    end
    icons
  end

  def item_hash(icon)
    {
      :uid      => '',
      :title    => icon,
      :subtitle => "Copy to clipboard: fa-#{icon}",
      :arg      => icon,
      :icon     => { :type => 'default', :name => "./icons/fa-#{icon}.png" },
      :valid    => 'yes',
    }
  end

  def add_items(feedback, icons = @icons)
    icons.each { |icon| feedback.add_item(item_hash(icon.name)) }
    feedback
  end

  def to_alfred(alfred)
    add_items(alfred.feedback).to_alfred
  end

  private

  def glob_icon_names
    Dir.glob(File.expand_path('./icons/fa-*.png')).map do |path|
      md = /\/fa-(.+)\.png/.match(path)
      (md && md[1]) ? md[1] : nil
    end.compact.uniq.sort
  end
end
