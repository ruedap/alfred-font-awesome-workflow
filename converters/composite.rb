png = Dir.glob("./assets/icons/*.png")
temporary_png = './converters/__temp.png'
puts `convert -size 128x128 xc:#ffffff #{temporary_png}`

png.each do |p|
  puts "composite -gravity center -compose over #{p} #{temporary_png} #{p}"
  puts `composite -gravity center -compose over #{p} #{temporary_png} #{p}`
end

puts `rm #{temporary_png}`
