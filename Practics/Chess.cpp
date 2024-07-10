#include <iostream>
#include <vector>
#include <string>

// Класс Фигура
class Figure {
protected:
    std::string color;
    std::string name;

public:
    Figure(std::string color, std::string name) : color(color), name(name) {}

    virtual bool isValidMove(int startX, int startY, int endX, int endY, std::vector<std::vector<Figure*>>& board) = 0;

    std::string getColor() {
        return color;
    }

    std::string getName() {
        return name;
    }
};

// Класс Доска
class Board {
private:
    std::vector<std::vector<Figure*>> board;

public:
    Board() {
        board.resize(8, std::vector<Figure*>(8, nullptr));
    }

    void printBoard() {
        for (int i = 0; i < 8; i++) {
            for (int j = 0; j < 8; j++) {
                if (board[i][j] == nullptr) {
                    std::cout << "-";
                } else {
                    std::cout << board[i][j]->getName();
                }
                std::cout << " ";
            }
            std::cout << std::endl;
        }
    }

    void setFigure(int x, int y, Figure* figure) {
        board[x][y] = figure;
    }

    Figure* getFigure(int x, int y) {
        return board[x][y];
    }
};

// Класс Игра
class Game {
private:
    Board board;
    std::string player1Name;
    std::string player2Name;
    std::string currentPlayer;
    std::string winner;

public:
    Game() {
        player1Name = "Игрок 1";
        player2Name = "Игрок 2";
        currentPlayer = player1Name;
        winner = "";
    }

    void startGame() {
        std::cout << "Игра в шахматы" << std::endl;
        std::cout << "Введите имена игроков (по умолчанию Игрок 1 и Игрок 2):" << std::endl;
        std::cout << "Игрок 1: ";
        std::getline(std::cin, player1Name);
        std::cout << "Игрок 2: ";
        std::getline(std::cin, player2Name);

        std::cout << "Первый игрок выбирает цвет фигур (белый или черный)." << std::endl;
        std::cout << "Игрок 1: ";
        std::string color;
        std::getline(std::cin, color);
        while (color != "белый" && color != "черный") {
            std::cout << "Неверный цвет. Введите белый или черный: ";
            std::getline(std::cin, color);
        }

        std::cout << "Игра началась!" << std::endl;

        if (color == "белый") {
            currentPlayer = player1Name;
        } else {
            currentPlayer = player2Name;
        }

        while (winner.empty()) {
            std::cout << "Ход игрока " << currentPlayer << std::endl;
            board.printBoard();

            std::cout << "Введите позицию фигуры (например, a2): ";
            std::string startPos;
            std::getline(std::cin, startPos);
            int startX = startPos[1] - '1';
            int startY = startPos[0] - 'a';

            std::cout << "Введите позицию для перемещения фигуры (например, a4): ";
            std::string endPos;
            std::getline(std::cin, endPos);
            int endX = endPos[1] - '1';
            int endY = endPos[0] - 'a';

            Figure* figure = board.getFigure(startX, startY);
            if (figure == nullptr || figure->getColor() != currentPlayer) {
                std::cout << "Ошибка: неверная фигура!" << std::endl;
                continue;
            }

            if (figure->isValidMove(startX, startY, endX, endY, board.getBoard())) {
                board.setFigure(endX, endY, figure);
                board.setFigure(startX, startY, nullptr);

                if (currentPlayer == player1Name) {
                    currentPlayer = player2Name;
                } else {
                    currentPlayer = player1Name;
                }
            } else {
                std::cout << "Ошибка: неверный ход!" << std::endl;
            }

            // Проверка на победу или ничью
            // ...

            // Проверка на пат
            // ...
        }

        std::cout << "Победил игрок " << winner << "!" << std::endl;
    }
};

int main() {
    Game game;
    game.startGame();

    return 0;
}