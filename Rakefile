require 'rubygems' unless defined? Gem # rubygems is only needed in 1.8

require 'yaml'
require 'plist'

config_file   = 'config.yml'
workflow_home = File.expand_path('~/Library/Application Support/Alfred 2/Alfred.alfredpreferences/workflows')

$config = YAML.load_file(config_file)
$config['plist'] = File.join($config['path'], 'info.plist')
$config['workflow_dropbox_home'] = File.join(File.expand_path($config['dropbox']), '/Alfred.alfredpreferences/workflows')

task :config do
  info = Plist::parse_xml($config['plist'])
  unless info['bundleid'].eql?($config['bundle_id'])
    info['bundleid'] = $config['bundle_id']
    File.open($config['plist'], 'wb') { |file| file.write(info.to_plist) }
  end
end

task :chdir => [:config] do
  chdir $config['path']
end

desc 'Install gems'
task 'bundle:install' => [:chdir] do
  sh %Q{bundle install --standalone --clean --without test} do |ok, res|
    puts "fail to install gems (status = #{res.exitstatus})" unless ok
  end
end

desc 'Install gems for test'
task 'bundle:install:test' => [:clean, :clobber, :chdir] do
  sh %Q{bundle install --standalone --clean} do |ok, res|
    puts "fail to install gems (status = #{res.exitstatus})" unless ok
  end
end

desc 'Update gems'
task 'bundle:update' => [:chdir] do
  sh %Q{bundle update && bundle install --standalone --clean --without test} do |ok, res|
    puts "fail to update gems (status = #{res.exitstatus})" unless ok
  end
end

desc 'Link to Alfred'
task :link => [:config] do
  ln_sf File.expand_path($config['path']), File.join(workflow_home, $config['bundle_id'])
end

desc 'Unlink from Alfred'
task :unlink => [:config] do
  rm File.join(workflow_home, $config['bundle_id'])
end

desc 'Link to Dropbox'
task 'dropbox:link' => [:config] do
  ln_sf File.expand_path($config['path']), File.join($config['workflow_dropbox_home'], $config['bundle_id'])
end

desc 'Unlink from Dropbox'
task 'dropbox:unlink' => [:config] do
  rm File.join($config['workflow_dropbox_home'], $config['bundle_id'])
end

desc 'Clean up all the extras'
task :clean => [:config] do
  rmtree File.join($config['path'], '.bundle')
  rmtree File.join($config['path'], 'bundle')
  rmtree File.join($config['path'], 'coverage')
end

desc 'Run rspec'
task :spec => [:chdir] do
  sh %Q{bundle exec rake spec} do |ok, res|
    puts "fail to run rspec (status = #{res.exitstatus})" unless ok
  end
end

task :default => :spec
