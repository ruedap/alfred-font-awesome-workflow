module FontAwesome
  def self.icons
    Dir.glob(File.expand_path('./icon-*.png')).map { |path|
      md = /\/icon-(.+)\.png/.match(path)
      (md && md[1]) ? md[1] : nil
    }.compact.sort
  end
end
