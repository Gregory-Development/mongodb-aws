package config

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type Config struct {
	Config config `yaml:"config"`
}

type config struct {
	System system `yaml:"system"`
	Kubernetes kubernetes `yaml:"kubernetes"`
	AWS aws `yaml:"amazon-aws"`
}

type system struct {
	KubectlExecutable string `yaml:"kubectl-executable-path"`
	KopsExecutable string `yaml:"kops-executable-path"`
	TerraformExecutable string `yaml:"terraform-executable-path"`
	PipExecutable string `yaml:"pip-executable-path"`
	AwsCliExecutable string `yaml:"awscli-executable-path"`
}

type kubernetes struct {
	StateStore string `yaml:"state-store"`
	Zones []string `yaml:"zones"`
	NodeCount int `yaml:"node-count"`
	NodeSize string `yaml:"node-size"`
	MasterSize string `yaml:"master-size"`
	DNSZone string `yaml:"dns-zone"`
}

type aws struct {
	AccessKeyId string `yaml:"access-key-id"`
	SecretAccessKey string `yaml:"secret-access-key"`
	Region string `yaml:"region"`
}

const rhelK8sRepo = `[kubernetes]\n
name=Kubernetes\n
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64\n
enabled=1\n
gpgcheck=1\n
repo_gpgcheck=1\n
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg\n`

const debK8sRepo = "deb https://apt.kubernetes.io/ kubernetes-xenial main"

func validateExecutables (ex string) (err error) {
	if _, err = os.Stat(ex); os.IsNotExist(err) {
		return
	}

	err = nil
	return
}

func detectDistro () (dist string, err error) {
	var out bytes.Buffer
	var distro = regexp.MustCompile(`^NAME=(.+?)\n`)

	cmd := exec.Command("/usr/bin/lsb_release", "-a")
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		dist = ""
		return
	}

	result := distro.FindStringSubmatch(string(out.Bytes()))
	distStr := strings.Split(result[0], "=")
	dist = strings.ToLower(distStr[1])
	err = nil
	return
}

func detectOS () (dist string, err error) {
	if runtime.GOOS == "linux" {
		dist, err = detectDistro()
		if err != nil {
			return
		}
		err = nil
		return
	} else {
		log.Fatalf("%v is not currently supported at this time", runtime.GOOS)
		os.Exit(1)
	}
	return
}

func writeRhelRepo() {
	f, err := os.Create("/etc/yum.repos.d/kubernetes.repo")
	if err != nil {
		log.Fatalf("unable to create k8s repo file, are you root?")
		os.Exit(1)
	}
	_, err = f.WriteString(rhelK8sRepo)
	if err != nil {
		log.Fatalf("unable to write to the k8s repo file, are you root?")
		os.Exit(1)
	}
}

func writeDebRepo() {
	f, err := os.Create("/etc/apt/sources.list.d/kubernetes.list")
	if err != nil {
		log.Fatalf("unable to create k8s repo file, are you root?")
		os.Exit(1)
	}
	_, err = f.WriteString(debK8sRepo)
	if err != nil {
		log.Fatalf("unable to write to the k8s repo file, are you root?")
		os.Exit(1)
	}
}

func rhelInstallPackage (cmd string, pkg string) {
	var out bytes.Buffer

	install := exec.Command(cmd, "install", "-y", pkg)
	install.Stdout = &out
	err := install.Run()
	if err != nil {
		log.Fatalf("unable to install %v, are you root?", pkg)
		os.Exit(1)
	}
}

func debInstallPackage (pkg string) {
	var out bytes.Buffer

	install := exec.Command("apt-get", "install", "-y", pkg)
	install.Stdout = &out
	err := install.Run()
	if err != nil {
		log.Fatalf("unable to install %v, are you root?", pkg)
		os.Exit(1)
	}
}

func installDependencies(dep string) {
	if dep == "kubectl" {
		if d, _ := detectOS(); d == "fedora" {
			writeRhelRepo()
			rhelInstallPackage("dnf", "kubectl")
		} else if d, _ := detectOS(); d == "centos" {
			writeRhelRepo()
			rhelInstallPackage("yum", "kubectl")
		} else if d, _ := detectOS(); d == "ubuntu" {
			writeDebRepo()
			debInstallPackage("kubectl")
		} else if d, _ := detectOS(); d == "debian" {
			writeDebRepo()
			debInstallPackage("kubectl")
		} else {
			d, _ := detectOS()
			log.Fatalf("%v is not current supported", d)
			os.Exit(1)
		}
	} else if dep == "kops" {
		if d, _ := detectOS(); d == "fedora" {
			writeRhelRepo()
			rhelInstallPackage("dnf", "kubectl")
		} else if d, _ := detectOS(); d == "centos" {
			writeRhelRepo()
			rhelInstallPackage("yum", "kubectl")
		} else if d, _ := detectOS(); d == "ubuntu" {
			writeDebRepo()
			debInstallPackage("kubectl")
		} else if d, _ := detectOS(); d == "debian" {
			writeDebRepo()
			debInstallPackage("kubectl")
		} else {
			d, _ := detectOS()
			log.Fatalf("%v is not current supported", d)
			os.Exit(1)
		}
	} else if dep == "terraform" {
		if d, _ := detectOS(); d == "fedora" {
			writeRhelRepo()
			rhelInstallPackage("dnf", "kubectl")
		} else if d, _ := detectOS(); d == "centos" {
			writeRhelRepo()
			rhelInstallPackage("yum", "kubectl")
		} else if d, _ := detectOS(); d == "ubuntu" {
			writeDebRepo()
			debInstallPackage("kubectl")
		} else if d, _ := detectOS(); d == "debian" {
			writeDebRepo()
			debInstallPackage("kubectl")
		} else {
			d, _ := detectOS()
			log.Fatalf("%v is not current supported", d)
			os.Exit(1)
		}
	} else if dep == "pip" {
		if d, _ := detectOS(); d == "fedora" {
			rhelInstallPackage("dnf", "python2-pip")
		} else if d, _ := detectOS(); d == "centos" {
			rhelInstallPackage("yum", "python2-pip")
		} else if d, _ := detectOS(); d == "ubuntu" {
			debInstallPackage("kubectl")
		} else if d, _ := detectOS(); d == "debian" {
			debInstallPackage("kubectl")
		} else {
			d, _ := detectOS()
			log.Fatalf("%v is not current supported", d)
			os.Exit(1)
		}
	} else if dep == "aws" {
		if d, _ := detectOS(); d == "fedora" {
			writeRhelRepo()
			rhelInstallPackage("dnf", "kubectl")
		} else if d, _ := detectOS(); d == "centos" {
			writeRhelRepo()
			rhelInstallPackage("yum", "kubectl")
		} else if d, _ := detectOS(); d == "ubuntu" {
			writeDebRepo()
			debInstallPackage("kubectl")
		} else if d, _ := detectOS(); d == "debian" {
			writeDebRepo()
			debInstallPackage("kubectl")
		} else {
			d, _ := detectOS()
			log.Fatalf("%v is not current supported", d)
			os.Exit(1)
		}
	} else {
		return
	}
}

func (c *Config) ParseConfigFile (filename string) (err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return
	}

	err = validateExecutables(c.Config.System.KubectlExecutable)
	if err != nil {
		installDependencies("kubectl")
	}

	err = validateExecutables(c.Config.System.KopsExecutable)
	if err != nil {
		installDependencies("kops")
	}

	err = validateExecutables(c.Config.System.TerraformExecutable)
	if err != nil {
		installDependencies("terraform")
	}

	err = validateExecutables(c.Config.System.PipExecutable)
	if err != nil {
		installDependencies("pip")
	}

	err = validateExecutables(c.Config.System.AwsCliExecutable)
	if err != nil {
		installDependencies("aws")
	}

	err = nil
	return
}
