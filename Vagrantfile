# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|

  config.vm.box = "darrelchia/wercker-meetup-go"
  config.vm.box_version = "1.0.0"
  config.vm.box_url = "file://../wercker-meetup-go.box"

  config.vm.network "forwarded_port", guest: 9000, host: 9000, "id": "portainer-port"
  config.vm.network "forwarded_port", guest: 5000, host: 5000, "id": "golang-port"
  config.vm.network "forwarded_port", guest: 22, host: 2222, "id": "ssh"

  # You may want to install the vagrant proxyconf plugin.
  # At the command line, type : vagrant plugin install vagrant-proxyconf
  if Vagrant.has_plugin?("vagrant-proxyconf")    
    config.proxy.http = ""
    config.proxy.https = ""
    config.proxy.no_proxy = ""
  end

  config.vm.synced_folder "D:\\Meetups\\Wercker\\oracledeveloper-apac-wercker", "/home/vagrant/workspace/oracledeveloper-apac-wercker"
  config.vm.box_check_update = false

  # Example for VirtualBox:
  #
  config.vm.provider "virtualbox" do |vbox|
    vbox.memory = 2048
    vbox.cpus = 2
    vbox.name = "Wercker-Meetup-Golang"
  end
end