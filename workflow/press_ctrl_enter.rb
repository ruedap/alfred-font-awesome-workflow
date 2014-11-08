#!/usr/bin/env ruby
# encoding: utf-8

require './require_helper'

argv = FontAwesome.argv(ARGV)
print FontAwesome.character_reference(argv.icon_unicode)
FontAwesome.save_config_of_recent_icons(argv.icon_id)
