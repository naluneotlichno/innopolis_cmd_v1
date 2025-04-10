#!/bin/bash

# Функция для установки golangci-lint
install_linter() {
    echo "Устанавливаем golangci-lint..."
    if ! go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; then
        echo "Ошибка установки golangci-lint!"
        exit 1
    fi
    
    # Добавляем путь к GOPATH/bin, если его нет в PATH
    if ! grep -q "$(go env GOPATH)/bin" ~/.bashrc; then
        echo "export PATH=\$PATH:$(go env GOPATH)/bin" >> ~/.bashrc
        echo "Добавлен путь к GOPATH/bin в .bashrc. Выполните: source ~/.bashrc"
    fi
}

CONFIG_FILE=".golangci.yml"
if [ ! -f "$CONFIG_FILE" ]; then
    echo "Создаём $CONFIG_FILE..."
    cat > "$CONFIG_FILE" << 'EOL'
run:
  timeout: 5m
  modules-download-mode: readonly

linters:
  enable:
    - gofmt
    - goimports
    - govet
    - staticcheck
    - typecheck
    - errcheck
    - bodyclose
    - noctx
    - unparam
    - wastedassign
    - misspell
    - gosec
    - mnd
    - gocyclo
    - copyloopvar
    - durationcheck

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0
EOL

    # Проверка, что файл создан
    if [ -f "$CONFIG_FILE" ]; then
        echo "Файл $CONFIG_FILE успешно создан"
    else
        echo "Ошибка: не удалось создать $CONFIG_FILE!"
        exit 1
    fi
fi