class FontAwesome
  attr_reader :icons

  def initialize
    @icons = load_icons
  end

  def select!(queries, icons = @icons)
    queries.each do |q|
      # use reject! for ruby 1.8 compatible
      icons.reject! { |i| i.index(q.downcase) ? false : true }
    end
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

  private

  def load_icons
    Dir.glob(File.expand_path('./icons/fa-*.png')).map do |path|
      md = /\/fa-(.+)\.png/.match(path)
      (md && md[1]) ? md[1] : nil
    end.compact.uniq.sort
  end
end
