require 'simplecov'
require 'coveralls'

SimpleCov.formatter = SimpleCov::Formatter::MultiFormatter[
  SimpleCov::Formatter::HTMLFormatter,
  Coveralls::SimpleCov::Formatter
]
SimpleCov.start do
  add_filter '.bundle/'
end

$LOAD_PATH.unshift(File.expand_path(File.dirname('..')))

require 'rubygems' unless defined? Gem # rubygems is only needed in 1.8
require 'bundle/bundler/setup'
require 'lib/font_awesome'

require 'rspec'
require 'rexml/document'
require 'spec/fixtures'

RSpec.configure do |config|
  config.expect_with :rspec do |c|
    c.syntax = :expect  # disables `should`
  end
end
