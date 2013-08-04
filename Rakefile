require 'rubygems' unless defined? Gem # rubygems is only needed in 1.8

require 'yaml'
require 'plist'

config_file = 'config.yml'


workflow_home=File.expand_path("~/Library/Application Support/Alfred 2/Alfred.alfredpreferences/workflows")

$config = YAML.load_file(config_file)
$config["bundleid"] = "#{$config["domain"]}.#{$config["id"]}"
$config["plist"] = File.join($config["path"], "info.plist")
$config["workflow_dbx"] = File.join(File.expand_path($config["dropbox"]), "/Alfred.alfredpreferences/workflows")

# import sub-rakefiles
FileList['*/Rakefile'].each { |file|
  import file
}

task :config do

  info = Plist::parse_xml($config["plist"])
  unless info['bundleid'].eql?($config["bundleid"])
    info['bundleid'] = $config["bundleid"]
    File.open($config["plist"], "wb") { |file| file.write(info.to_plist) }
  end
end

task :chdir => [:config] do
  chdir $config['path']
end

desc "Install Gems"
task "bundle:install" => [:chdir] do
  sh %Q{bundle install --standalone --clean} do |ok, res|
    if ! ok
      puts "fail to install gems (status = #{res.exitstatus})"
    end
  end
end

desc "Update Gems"
task "bundle:update" => [:chdir] do
  sh %Q{bundle update && bundle install --standalone --clean} do |ok, res|
    if ! ok
      puts "fail to update gems (status = #{res.exitstatus})"
    end
  end
end

desc "Install to Alfred"
task :install => [:config] do
  ln_sf File.expand_path($config["path"]), File.join(workflow_home, $config["bundleid"])
end

desc "Unlink from Alfred"
task :uninstall => [:config] do
  rm File.join(workflow_home, $config["bundleid"])
end

desc "Install to Dropbox"
task :dbxinstall => [:config] do
  ln_sf File.expand_path($config["path"]), File.join($config["workflow_dbx"], $config["bundleid"])
end

desc "Unlink from Dropbox"
task :dbxuninstall => [:config] do
  rm File.join($config["workflow_dbx"], $config["bundleid"])
end

desc "Clean up all the extras"
task :clean => [:config] do
end

desc "Remove any generated file"
task :clobber => [:clean] do
  rmtree File.join($config["path"], ".bundle")
  rmtree File.join($config["path"], "bundle")
end


