#!/usr/bin/env ruby
# encoding: utf-8

$LOAD_PATH.unshift(File.expand_path(File.dirname(__FILE__)))

require 'rubygems' unless defined? Gem # rubygems is only needed in 1.8
require 'bundle/bundler/setup'
require 'lib/font_awesome'

argv = FontAwesome.argv(ARGV)
print FontAwesome.character_reference(argv.icon_unicode)
