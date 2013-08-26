module FontAwesome
  def self.icons
    Dir.glob(File.expand_path('./icon-*.png')).map do |path|
      md = /\/icon-(.+)\.png/.match(path)
      (md && md[1]) ? md[1] : nil
    end.compact.sort
  end

  def self.select!(icons, queries)
    queries.each do |q|
      # use reject! for ruby 1.8 compatible
      icons.reject! { |i| i.index(q.downcase) ? false : true }
    end
  end

  def self.item_hash(icon)
    {
      :uid      => '',
      :title    => icon,
      :subtitle => "Copy to clipboard: icon-#{icon}",
      :arg      => icon,
      :icon     => { :type => 'default', :name => "icon-#{icon}.png" },
      :valid    => 'yes',
    }
  end
end
