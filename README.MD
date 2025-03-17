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


## Возможные проблемы

1. **Ошибка компиляции с PortAudio:**
   ```
   fatal error: portaudio.h: No such file or directory
   ```
   Решение: Установите заголовочные файлы PortAudio (см. раздел "Установка зависимостей")

2. **Ошибка линковки с PortAudio:**
   ```
   undefined reference to `Pa_Initialize'
   ```
   Решение: Убедитесь, что библиотека PortAudio установлена и флаг `-lportaudio` указан при компиляции

3. **Ошибка при запуске:**
   ```
   PortAudio error: Unanticipated host error
   ```
   Решение: Проверьте, что звуковая система работает и у вас есть права на использование аудиоустройств

## Лицензия

MIT
