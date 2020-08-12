package main

import (
	distro "github.com/sipsma/bincastle-distro"
	"github.com/sipsma/bincastle-distro/src"
	"github.com/sipsma/bincastle/cmd"
	. "github.com/sipsma/bincastle/graph"
)

func main() {
	cmd.SystemDef(
		// TODO this is a bit unweidly, find way to break it down
		distro.BuildDistro(
			// build tools
			distro.Autoconf{},
			distro.Automake{},
			distro.GCC{},
			distro.GMP{},
			distro.Libtool{},
			distro.LinuxHeaders{},
			distro.M4{},
			distro.MPC{},
			distro.MPFR{},
			distro.Make{},
			distro.OpenSSL{},
			distro.PkgConfig{},
			distro.Readline{},

			// common cmdline tools (also their .so's)
			distro.Acl{},
			distro.Attr{},
			distro.Awk{},
			distro.Bzip2{},
			distro.Coreutils{},
			distro.Diffutils{},
			distro.File{},
			distro.Findutils{},
			distro.Git{},
			distro.Grep{},
			distro.Gzip{},
			distro.Inetutils{},
			distro.Iproute2{},
			distro.Less{},
			distro.Libcap{},
			distro.Patch{},
			distro.Procps{},
			distro.Psmisc{},
			distro.Sed{},
			distro.Tar{},
			distro.UtilLinux{},
			distro.Which{},
			distro.Xz{},

			// misc
			distro.CACerts{},
			distro.Ianaetc{},
			distro.Mandb{},
			distro.Manpages{},

			// langs
			distro.Golang{},
			distro.Perl5{},
			distro.Python3{},
			distro.NodeJS{},

			// user
			LayerSpec(
				Dep(distro.User{
					Name:    "sipsma",
					Shell:   "/bin/bash",
					Homedir: "/home/sipsma",
				}),
				Dep(distro.Bash{}),
				Dep(distro.OpenSSH{}),
				Dep(distro.Emacs{}),
				Dep(distro.Tmux{}),
				Dep(Wrap(src.ViaGit{
					URL:  "https://github.com/sipsma/bincastle.git",
					Ref:  "master",
					Name: "bincastle-src",
				}, MountDir("/home/sipsma/.repo/github.com/sipsma/bincastle"))),
				Dep(Wrap(src.ViaGit{
					URL:  "https://github.com/syl20bnr/spacemacs.git",
					Ref:  "develop",
					Name: "spacemacs-src",
				}, MountDir("/home/sipsma/.emacs.d"))),
				BuildDep(distro.Coreutils{}),
				BuildDep(distro.Bash{}),
				BuildDep(distro.OpenSSH{}),
				BuildDep(distro.Git{}),
				BuildDep(distro.Golang{}),
				BuildDep(distro.Ncurses{}),
				distro.BuildOpts(),
				ScratchMount(`/build`),
				Env("SSH_AUTH_SOCK", "/run/ssh-agent.sock"), // TODO this should be a helper, WithSSHSock
				Shell(
					`mkdir -p /home/sipsma`,
					`cd /build`,

					//  TODO this seems leaky
					`ln -s /inner /home/sipsma/.bincastle`,

					// TODO need a better way of updating known_hosts,
					// this is very insecure and doesn't integrate w/ the normal
					// way of adding a layer sourced from git
					`mkdir -p /home/sipsma/.ssh`,
					`ssh-keyscan github.com >> /home/sipsma/.ssh/known_hosts`,
					`git clone -b spacemacs git@github.com:sipsma/home.git /home/sipsma/.spacemacs.d`,
					`git clone git@github.com:sipsma/sipsma.dev.git /home/sipsma/.repo/github.com/sipsma/sipsma.dev`,
					`git clone git@github.com:sipsma/buildkit.git /home/sipsma/.repo/github.com/sipsma/buildkit`,

					// TODO this should be its own package
					`echo 'xterm-24bit|xterm with 24-bit direct color mode,' > terminfo`,
					`echo '   use=xterm-256color,' >> terminfo`,
					`echo '   sitm=\E[3m,' >> terminfo`,
					`echo '   ritm=\E[23m,' >> terminfo`,
					`echo '   setb24=\E[48;2;%p1%{65536}%/%d;%p1%{256}%/%{255}%&%d;%p1%{255}%&%dm,' >> terminfo`,
					`echo '   setf24=\E[38;2;%p1%{65536}%/%d;%p1%{256}%/%{255}%&%d;%p1%{255}%&%dm,' >> terminfo`,
					`echo '' >> terminfo`,
					`tic -x -o /home/sipsma/.terminfo terminfo`,

					// TODO this should be its own package
					`echo 'if [ -f "$HOME/.bashrc" ]; then . "$HOME/.bashrc"; fi' >> /home/sipsma/.profile`,
					`echo 'HISTCONTROL=ignoreboth' >> /home/sipsma/.bashrc`,
					`echo 'shopt -s histappend' >> /home/sipsma/.bashrc`,
					`echo 'HISTSIZE=1000' >> /home/sipsma/.bashrc`,
					`echo 'HISTFILESIZE=2000' >> /home/sipsma/.bashrc`,
					`echo 'shopt -s checkwinsize' >> /home/sipsma/.bashrc`,
					`echo 'set -o vi' >> /home/sipsma/.bashrc`,
					`echo 'set +o posix' >> /home/sipsma/.bashrc`,
					`echo 'export TERM=xterm-24bit' >> /home/sipsma/.bashrc`,
					`echo 'export LANG=en_US.UTF-8' >> /home/sipsma/.bashrc`,
					`echo 'export GO111MODULE=on' >> /home/sipsma/.bashrc`,
					`echo 'export PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:/usr/lib/go/bin:/home/sipsma/go/bin' >> /home/sipsma/.bashrc`,

					// TODO this should be its own package
					`echo 'set -g default-terminal "xterm-24bit"' >> /home/sipsma/.tmux.conf`,
					`echo 'set -g terminal-overrides ",xterm-24bit:Tc"' >> /home/sipsma/.tmux.conf`,
					`echo 'set -s escape-time 0' >> /home/sipsma/.tmux.conf`,

					// TODO this should be its own package
					`export GO111MODULE=on`,
					`go get golang.org/x/tools/gopls@latest`,

					`git config --global user.name "Erik Sipsma"`,
					`git config --global user.email "erik@sipsma.dev"`,
				),
			),
		),
	)
}
