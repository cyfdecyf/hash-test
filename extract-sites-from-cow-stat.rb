#!/usr/bin/ruby

require 'json'

stat = JSON.parse(ARGF.read())
stat['site_info'].each_key { |k| puts k }

