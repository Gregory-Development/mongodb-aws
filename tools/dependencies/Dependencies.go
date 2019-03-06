package dependencies

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

const rhelK8sRepo = `[kubernetes]\n
name=Kubernetes\n
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64\n
enabled=1\n
gpgcheck=1\n
repo_gpgcheck=1\n
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg\n`

const debK8sRepo = "deb https://apt.kubernetes.io/ kubernetes-xenial main"

type SYS struct {
	OS     string
	Arch   string
	Distro string
}

func ValidateExecutables(ex string) (err error) {
	if _, err = os.Stat(ex); os.IsNotExist(err) {
		return
	}

	err = nil
	return
}

func (s *SYS) DetectDistro() {
	var out bytes.Buffer
	var distro = regexp.MustCompile(`^NAME=(.+?)\n`)

	if s.OS == "linux" {
		cmd := exec.Command("/usr/bin/lsb_release", "-a")
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			s.Distro = ""
			return
		}

		result := distro.FindStringSubmatch(string(out.Bytes()))
		distStr := strings.Split(result[0], "=")
		s.Distro = strings.ToLower(distStr[1])
	}

}

func (s *SYS) DetectOS() {
	if runtime.GOOS == "linux" {
		s.OS = runtime.GOOS
		s.DetectDistro()
	} else {
		log.Fatalf("%v is not currently supported at this time", runtime.GOOS)
		os.Exit(1)
	}
}

func (s *SYS) WriteRhelRepo() {
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

func (s *SYS) WriteDebRepo() {
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

func (s *SYS) RhelInstallPackage(cmd string, pkg string) {
	var out bytes.Buffer

	install := exec.Command(cmd, "install", "-y", pkg)
	install.Stdout = &out
	err := install.Run()
	if err != nil {
		log.Fatalf("unable to install %v, are you root?", pkg)
		os.Exit(1)
	}
}

func (s *SYS) DebInstallAptKey() {
	var out bytes.Buffer
	var u = "https://packages.cloud.google.com/apt/doc/apt-key.gpg"

	install := exec.Command("apt-key", "adv", "--fetch-keys", u)
	install.Stdout = &out
	err := install.Run()
	if err != nil {
		log.Fatalf("unable to install %v, are you root?", pkg)
		os.Exit(1)
	}
}

func (s *SYS) DebInstallPackage(pkg string) {
	var out bytes.Buffer

	install := exec.Command("apt-get", "install", "-y", pkg)
	install.Stdout = &out
	err := install.Run()
	if err != nil {
		log.Fatalf("unable to install %v, are you root?", pkg)
		os.Exit(1)
	}
}

func (s *SYS) InstallDependencies(dep string) {
	if dep == "kubectl" {
		if s.Distro == "fedora" {
			s.WriteRhelRepo()
			s.RhelInstallPackage("dnf", "kubectl")
		} else if s.Distro == "centos" {
			s.WriteRhelRepo()
			s.RhelInstallPackage("yum", "kubectl")
		} else if s.Distro == "ubuntu" {
			s.WriteDebRepo()
			s.DebInstallAptKey()
			s.DebInstallPackage("kubectl")
		} else if s.Distro == "debian" {
			s.WriteDebRepo()
			s.DebInstallAptKey()
			s.DebInstallPackage("kubectl")
		} else {
			log.Fatalf("%v is not current supported", s.Distro)
			os.Exit(1)
		}
	} else if dep == "kops" {
		if s.Distro == "fedora" {
			s.WriteRhelRepo()
			s.RhelInstallPackage("dnf", "kubectl")
		} else if s.Distro == "centos" {
			s.WriteRhelRepo()
			s.RhelInstallPackage("yum", "kubectl")
		} else if s.Distro == "ubuntu" {
			s.WriteDebRepo()
			s.DebInstallAptKey()
			s.DebInstallPackage("kubectl")
		} else if s.Distro == "debian" {
			s.WriteDebRepo()
			s.DebInstallAptKey()
			s.DebInstallPackage("kubectl")
		} else {
			log.Fatalf("%v is not current supported", s.Distro)
			os.Exit(1)
		}
	} else if dep == "terraform" {
		if s.Distro == "fedora" {
			s.WriteRhelRepo()
			s.RhelInstallPackage("dnf", "kubectl")
		} else if s.Distro == "centos" {
			s.WriteRhelRepo()
			s.RhelInstallPackage("yum", "kubectl")
		} else if s.Distro == "ubuntu" {
			s.WriteDebRepo()
			s.DebInstallAptKey()
			s.DebInstallPackage("kubectl")
		} else if s.Distro == "debian" {
			s.WriteDebRepo()
			s.DebInstallAptKey()
			s.DebInstallPackage("kubectl")
		} else {
			log.Fatalf("%v is not current supported", s.Distro)
			os.Exit(1)
		}
	} else if dep == "pip" {
		if s.Distro == "fedora" {
			s.RhelInstallPackage("dnf", "python2-pip")
		} else if s.Distro == "centos" {
			s.RhelInstallPackage("yum", "python2-pip")
		} else if s.Distro == "ubuntu" {
			s.DebInstallPackage("kubectl")
		} else if s.Distro == "debian" {
			s.DebInstallPackage("kubectl")
		} else {
			log.Fatalf("%v is not current supported", s.Distro)
			os.Exit(1)
		}
	} else if dep == "aws" {
		if s.Distro == "fedora" {
			s.WriteRhelRepo()
			s.RhelInstallPackage("dnf", "kubectl")
		} else if s.Distro == "centos" {
			s.WriteRhelRepo()
			s.RhelInstallPackage("yum", "kubectl")
		} else if s.Distro == "ubuntu" {
			s.WriteDebRepo()
			s.DebInstallPackage("kubectl")
		} else if s.Distro == "debian" {
			s.WriteDebRepo()
			s.DebInstallPackage("kubectl")
		} else {
			log.Fatalf("%v is not current supported", s.Distro)
			os.Exit(1)
		}
	} else {
		return
	}
}
