png = Dir.glob("./workflow/icons/*.png")
base_png = './converters/__base.png'
puts `convert -size 128x128 xc:#ffffff #{base_png}`

png.each do |p|
  puts "composite -gravity center -compose over #{p} #{base_png} #{p}"
  puts `composite -gravity center -compose over #{p} #{base_png} #{p}`
end

puts `rm #{base_png}`
