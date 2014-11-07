require 'rubygems' unless defined? Gem # rubygems is only needed in 1.8

require 'yaml'

config_file   = 'config.yml'
workflow_home = File.expand_path('~/Library/Application Support/Alfred 2/Alfred.alfredpreferences/workflows')

$config = YAML.load_file(config_file)
$config['workflow_dropbox_home'] = File.join(File.expand_path($config['dropbox']), '/Alfred.alfredpreferences/workflows')

task :chdir do
  chdir $config['path']
end

desc 'Install gems'
task 'bundle:install' => [:chdir] do
  sh %Q{bundle install --standalone --clean --without test} do |ok, res|
    puts "fail to install gems (status = #{res.exitstatus})" unless ok
  end
end

desc 'Install gems for test'
task 'bundle:install:test' => [:chdir] do
  sh %Q{bundle install --standalone --clean} do |ok, res|
    puts "fail to install gems (status = #{res.exitstatus})" unless ok
  end
end

desc 'Link to Alfred'
task :link do
  ln_sf File.expand_path($config['path']), File.join(workflow_home, $config['bundle_id'])
end

desc 'Unlink from Alfred'
task :unlink do
  rm File.join(workflow_home, $config['bundle_id'])
end

desc 'Link to Dropbox'
task 'dropbox:link' do
  ln_sf File.expand_path($config['path']), File.join($config['workflow_dropbox_home'], $config['bundle_id'])
end

desc 'Unlink from Dropbox'
task 'dropbox:unlink' do
  rm File.join($config['workflow_dropbox_home'], $config['bundle_id'])
end

desc 'Clean up all the extras'
task :clean do
  rmtree File.join($config['path'], '.bundle')
  rmtree File.join($config['path'], 'bundle')
  rmtree File.join($config['path'], 'coverage')
end

desc 'Clean up for release'
task 'clean:release' do
  ruby_version_path = File.expand_path('./.ruby-version')
  rm ruby_version_path if File.exist? ruby_version_path
  rmtree File.join($config['path'], 'coverage')
end

desc 'Run rspec'
task :spec => [:chdir] do
  sh %Q{bundle exec rake spec} do |ok, res|
    puts "fail to run rspec (status = #{res.exitstatus})" unless ok
  end
end

task :default => :spec
