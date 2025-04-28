# Whitenoise Generator

Генератор белого шума на C++ с использованием PortAudio.

## Установка зависимостей

### Ubuntu/Debian:
```bash
sudo apt-get update
sudo apt-get install libportaudio2 libportaudio-dev
```

## Сборка

Соберите проект с помощью g++:
```bash
g++ -o whitenoise main.cpp -lportaudio -std=c++17
```

## Запуск

После успешной компиляции запустите программу:
```bash
./whitenoise
```

## Управление

- `+` : увеличить громкость
- `-` : уменьшить громкость
- `q` : выход


## Лицензия

MIT
