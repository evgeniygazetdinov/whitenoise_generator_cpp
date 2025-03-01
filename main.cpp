#include <iostream>
#include <random>
#include <vector>
#include <chrono>
#include <portaudio.h>
#include <cstring>
#include <atomic>
#include <termios.h>
#include <unistd.h>

// Функция для неблокирующего чтения клавиши
char kbhit() {
    struct termios oldt, newt;
    char ch = 0;
    tcgetattr(STDIN_FILENO, &oldt);
    newt = oldt;
    newt.c_lflag &= ~(ICANON | ECHO);
    tcsetattr(STDIN_FILENO, TCSANOW, &newt);
    
    struct timeval tv = { 0L, 0L };
    fd_set fds;
    FD_ZERO(&fds);
    FD_SET(STDIN_FILENO, &fds);
    if (select(STDIN_FILENO + 1, &fds, NULL, NULL, &tv) > 0) {
        read(STDIN_FILENO, &ch, 1);
    }
    
    tcsetattr(STDIN_FILENO, TCSANOW, &oldt);
    return ch;
}

class WhiteNoiseGenerator {
private:
    std::mt19937 generator;
    std::uniform_real_distribution<double> distribution;
    static constexpr double SAMPLE_RATE = 44100.0;
    std::atomic<double> amplitude{0.01}; // Атомарная переменная для громкости
    std::atomic<bool> shouldStop{false};

public:
    WhiteNoiseGenerator() : 
        generator(std::chrono::system_clock::now().time_since_epoch().count()),
        distribution(-1.0, 1.0) {}

    double generateSample() {
        return distribution(generator) * amplitude.load();
    }

    // Callback функция для PortAudio
    static int paCallback(const void *inputBuffer, void *outputBuffer,
                         unsigned long framesPerBuffer,
                         const PaStreamCallbackTimeInfo* timeInfo,
                         PaStreamCallbackFlags statusFlags,
                         void *userData) {
        WhiteNoiseGenerator *generator = (WhiteNoiseGenerator*)userData;
        float *out = (float*)outputBuffer;
        (void) inputBuffer;

        if (generator->shouldStop.load()) {
            return paComplete;
        }

        for (unsigned long i = 0; i < framesPerBuffer; i++) {
            float sample = generator->generateSample();
            *out++ = sample; // Левый канал
            *out++ = sample; // Правый канал
        }
        return paContinue;
    }

    void adjustVolume(bool increase) {
        double current = amplitude.load();
        if (increase && current < 1.0) {
            amplitude.store(std::min(1.0, current + 0.001));
        } else if (!increase && current > 0.01) {
            amplitude.store(std::max(0.01, current - 0.001));
        }
        std::cout << "\rГромкость: " << static_cast<int>(amplitude.load() * 100) << "% " << std::flush;
    }

    // Воспроизведение шума
    void play() {
        PaError err = Pa_Initialize();
        if (err != paNoError) {
            std::cerr << "PortAudio error: " << Pa_GetErrorText(err) << std::endl;
            return;
        }

        PaStream *stream;
        err = Pa_OpenDefaultStream(&stream,
                                 0,          // нет входных каналов
                                 2,          // стерео выход
                                 paFloat32,  // формат сэмплов
                                 SAMPLE_RATE,
                                 paFramesPerBufferUnspecified,
                                 paCallback,
                                 this);

        if (err != paNoError) {
            std::cerr << "PortAudio error: " << Pa_GetErrorText(err) << std::endl;
            Pa_Terminate();
            return;
        }

        err = Pa_StartStream(stream);
        if (err != paNoError) {
            std::cerr << "PortAudio error: " << Pa_GetErrorText(err) << std::endl;
            Pa_CloseStream(stream);
            Pa_Terminate();
            return;
        }

        std::cout << "Воспроизведение белого шума...\n";
        std::cout << "Управление:\n";
        std::cout << "+ : увеличить громкость\n";
        std::cout << "- : уменьшить громкость\n";
        std::cout << "q : выход\n";
        std::cout << "Текущая громкость: " << static_cast<int>(amplitude.load() * 100) << "% " << std::flush;

        // Основной цикл обработки клавиш
        while (!shouldStop.load()) {
            char key = kbhit();
            if (key == '+' || key == '=') {
                adjustVolume(true);
            } else if (key == '-' || key == '_') {
                adjustVolume(false);
            } else if (key == 'q' || key == 'Q') {
                shouldStop.store(true);
                break;
            }
            Pa_Sleep(10); // Небольшая задержка для снижения нагрузки на CPU
        }

        std::cout << "\nЗавершение работы..." << std::endl;

        err = Pa_StopStream(stream);
        if (err != paNoError) {
            std::cerr << "PortAudio error: " << Pa_GetErrorText(err) << std::endl;
        }

        err = Pa_CloseStream(stream);
        if (err != paNoError) {
            std::cerr << "PortAudio error: " << Pa_GetErrorText(err) << std::endl;
        }

        Pa_Terminate();
    }
};

int main() {
    WhiteNoiseGenerator noise;
    noise.play();
    return 0;
}