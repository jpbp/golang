# 🛠️ Configuração de Ambiente – Ubuntu 24.04

Este documento descreve o processo de instalação e configuração de um ambiente de desenvolvimento completo com ferramentas como Go, Kubernetes, VSCode, além de aliases úteis para terminal e produtividade.

---

## 📦 Instalações Realizadas

### ✅ Go (Golang)
```bash
sudo snap install go --classic
go version
```

### ✅ kubectl (cliente Kubernetes)
```bash
curl -LO "https://dl.k8s.io/release/$(curl -sL https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
kubectl version --client
```

### ✅ kubectx & kubens
```bash
sudo git clone https://github.com/ahmetb/kubectx /opt/kubectx
sudo ln -s /opt/kubectx/kubectx /usr/local/bin/kubectx
sudo ln -s /opt/kubectx/kubens /usr/local/bin/kubens
```

### ✅ kubent (Kube No Trouble)
```bash
KUBENT_VERSION=$(curl -s https://api.github.com/repos/doitintl/kube-no-trouble/releases/latest | grep tag_name | cut -d '"' -f4)
curl -Lo kubent "https://github.com/doitintl/kube-no-trouble/releases/download/${KUBENT_VERSION}/kubent-linux-amd64"
chmod +x kubent
sudo mv kubent /usr/local/bin/
```

---

## 🔄 Autocompletion para Zsh

Adicione ao `~/.zshrc`:

```bash
# kubectl
source <(kubectl completion zsh)

# kubectx / kubens
[[ -r "/opt/kubectx/completion/kubectx.zsh" ]] && source /opt/kubectx/completion/kubectx.zsh
[[ -r "/opt/kubectx/completion/kubens.zsh" ]] && source /opt/kubectx/completion/kubens.zsh
```

Depois:
```bash
source ~/.zshrc
```

---

## 🧠 VSCode

### Instalação via pacote `.deb`:

```bash
wget -O vscode.deb "https://code.visualstudio.com/sha/download?build=stable&os=linux-deb-x64"
sudo apt install ./vscode.deb
rm vscode.deb
```

### Extensões instaladas:

```bash
code --install-extension golang.Go
code --install-extension ms-kubernetes-tools.vscode-kubernetes-tools
code --install-extension ms-vscode-remote.remote-ssh
code --install-extension jeff-hykin.zsh
```

---

## ⚙️ Aliases úteis (`~/.zshrc` ou `~/.bashrc`)

```bash
# 🔧 Básicos do terminal
alias ll='ls -lha'
alias la='ls -A'
alias ..='cd ..'
alias ...='cd ../..'
alias c='clear'
alias h='history'
alias p='pwd'
"'
alias now='date +"%Y-%m-%d %H:%M:%S"'

# 🐙 Git
alias g='git'
alias gs='git status'
alias ga='git add .'
alias gc='git commit -m'
alias gca='git commit --amend'
alias gp='git push'
alias gpo='git push origin'
alias gl='git log --oneline --graph --all'
alias gb='git branch'
alias gco='git checkout'
alias gcb='git checkout -b'
alias gpull='git pull origin $(git branch --show-current)'
alias gpush='git push origin $(git branch --show-current)'

# 🐳 Docker
alias d='docker'
alias dps='docker ps'
alias dpa='docker ps -a'
alias di='docker images'
alias drm='docker rm'
alias drmi='docker rmi'
alias dstop='docker stop $(docker ps -q)'
alias dclean='docker system prune -af --volumes'

# ☸️ Kubernetes
alias k='kubectl'
alias kgp='kubectl get pods'
alias kgs='kubectl get svc'
alias kga='kubectl get all'
alias kl='kubectl logs'
alias kaf='kubectl apply -f'
alias kdf='kubectl delete -f'
alias kctx='kubectl config current-context'
alias kx='kubectl config use-context'
alias kns='kubectl config set-context --current --namespace'

# 🧰 Ferramentas & DevOps
alias tf='terraform'
alias tfa='terraform apply'
alias tfi='terraform init'
alias tfd='terraform destroy'
alias tfo='terraform output'
alias tfs='terraform show'
alias v='nvim'
alias py='python3'
alias pipup='pip install --upgrade pip'

# 🛠️ Customizados úteis
alias serve='python3 -m http.server'
alias weather='curl wttr.in'
alias ips='ip -4 addr | grep inet'
alias myip='curl ifconfig.me'
alias psg='ps aux | grep -i'
```