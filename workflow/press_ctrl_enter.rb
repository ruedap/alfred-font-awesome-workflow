#!/usr/bin/env ruby
# encoding: utf-8

require './require_helper'

argv = FontAwesome.argv(ARGV)
print FontAwesome.character_reference(argv.icon_unicode)
