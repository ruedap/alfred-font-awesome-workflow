$LOAD_PATH.unshift(File.expand_path(File.dirname('..')))

require 'rubygems' unless defined? Gem # rubygems is only needed in 1.8
require 'bundle/bundler/setup'
require 'alfred'
require 'lib/font_awesome'

require 'minitest/autorun'
