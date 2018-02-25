svg = Dir.glob("./workflow/icons/*.svg")
svg.each do |s|
  p = File.basename(s, ".svg")
  puts `svgexport #{s} workflow/icons/#{p}.png pad 128:128`
  puts `rm #{s}`
end
