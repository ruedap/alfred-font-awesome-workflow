#!/usr/bin/env ruby
# encoding: utf-8

$LOAD_PATH.unshift(File.expand_path(File.dirname(__FILE__)))

require 'rubygems' unless defined? Gem # rubygems is only needed in 1.8
require 'logger'
require 'bundle/bundler/setup'
require 'lib/font_awesome'

# Main class
class Main
  def initialize(queries, debug = false)
    xml = FontAwesome.new(queries).to_alfred
    logging(xml) if debug
  end

  def logging(string)
    log_file = File.expand_path('~/Library/Logs/font-awesome-workflow.log')
    logger = Logger.new(log_file, 'daily')
    logger.progname = 'Font Awesome Workflow'
    logger.debug(string)
  end
end

# entry point
Main.new(ARGV)
