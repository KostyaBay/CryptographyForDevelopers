# CryptographyForDevelopers
#### Distributed Lab

Алгоритм Ель-Гамаля. Цифровий підпис. Спрямоване шифрування

🔹 Вимоги:

Напишіть власну програмну реалізацію цифрового підпису за алгоритмом Ель-Гамаля та власну програмну реалізацію спрямованого шифрування за алгоритмом Ель-Гамаля. Зверніть увагу на розбиття повідомлення на блоки і шифрування кожного блока окремо. <br>
Зробити програмну перевірку коректності вашої реалізації шляхом накладання підпису та його перевірки, а також співпадіння оригінального повідомлення з тим що було зашифровано та розшифровано. <br>

🔹 Етапи виконання практичного завдання:

1. Власна реалізація цифрового підпису
2. Власна реалізація спрямованого шифрування
3. Перевірка коректності реалізації

🔹 Результат:

Алгоритм ЦП за Ель-Гамаля міститься у папці Signature (файл `Signature.go`):
| Функція                      | Метод                       |  Реалізовано  |
|------------------------------|-----------------------------|:-------------:|
| ● Генерація ключів | GenKey (keyLength int) |:heavy_check_mark:|
| ● Гешування повідомлення SHA-256 | Hashing (mess string) |:heavy_check_mark:|
| ● Підписування повідомлення | DigitSign (mess string, p, g, x *big.Int) |:heavy_check_mark:|
| ● Верифікація підпису | Verification (mess string, p, g, y, r, s *big.Int) |:heavy_check_mark:|

Алгоритм шифрування за Ель-Гамаля міститься у папці Encryption (файл `Encryption.go`):
| Функція                      | Метод                       |  Реалізовано  |
|------------------------------|-----------------------------|:-------------:|
| ● Шифрування повідомлення | Encrypt (mess string, p, g, y *big.Int) |:heavy_check_mark:|
| ● Розшифрування повідомлення | Decrypt (a, b, p, x *big.Int) |:heavy_check_mark:|

Тестування та демонстрація відбуваються у файлі `main.go`:
| Тестування         | Метод                      |   Реалізовано    |
|--------------------|----------------------------|:----------------:|
| ● Тест генерації ключів | Test_KeyGen(length int) |:heavy_check_mark:|
| ● Тест використання ЦП | Test_DigitSign(p, g, y, x *big.Int) |:heavy_check_mark:|
| ● Тест шифрування/розшифрування | Test_Encrypt(p, g, y, x *big.Int) |:heavy_check_mark:|

Приклад генерації ключів (довжиною 4096 бітів):
![4096_test_KeyGen](https://github.com/KostyaBay/CryptographyForDevelopers/assets/54154093/a8ef28e3-b7d0-470c-85ed-6ec4f253ad4f)

Приклад підпису та верифікації повідомлення *"Message for test!"*:
![4096_test_DS](https://github.com/KostyaBay/CryptographyForDevelopers/assets/54154093/b9b77772-0aca-4921-b914-9e378659731c)

Приклад спрямованого шифрування повідомлення:
![4096_test_encrypt](https://github.com/KostyaBay/CryptographyForDevelopers/assets/54154093/42a2887b-1695-4f3f-a18c-86219cad948b)

Приклад верифікації пошкодженого повідомлення:
![4096_Ftest_DS](https://github.com/KostyaBay/CryptographyForDevelopers/assets/54154093/41089ffb-ef1c-4145-8d7c-11990fbe04d2)

Частини даних ключів, підпису та шифротексту розділені " ; ".

Для запуску коду потрібно:
-
1. Завантажити файли **main.go** (файл для перевірки та демонстрації генерації ключів, використання ЦП, спрямоване шифрування) та **go.mod** (файл конфігурацій).
2. Також треба завантажити файл **Signature.go** з самою реалізацією ЦП Ель-Гамаля у відповідній папці та файл **Encryption.go** з реалізацією алгоритму спрямованого шифрування за Ель-Гамалем у відповідній папці.
3. Відкрити папку з завантаженими файлами (рекомендовано використовувати IDE для мови Go, наприклад [GoLand від JetBrains](https://www.jetbrains.com/go/)).
4. Cкомпілювати та запустити код (в більшості IDE виконується командою **go build**).
5. Результат буде виведено в консоль за допомогою функції fmt.Println() вбудованої бібліотеки fmt, яку перед використанням імпортуємо - import "fmt".
