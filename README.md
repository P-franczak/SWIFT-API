# SWIFT API

## 📌 Opis projektu

SWIFT API to aplikacja RESTful napisana w języku Go, która umożliwia zarządzanie danymi kodów SWIFT.
Aplikacja pobiera dane z pliku CSV, przechowuje je w bazie danych PostgreSQL i udostępnia przez API.

### 🌟 Funkcjonalności

- Parsowanie i przechowywanie kodów SWIFT z pliku CSV.
- Przechowywanie danych w bazie PostgreSQL.
- Udostępnianie API REST do pobierania i modyfikowania danych.
- Obsługa zapytań GET, POST i DELETE.
- Możliwość pobrania wszystkich kodów SWIFT dla danego kraju.
- Obsługa kontenerów Docker do łatwego uruchamiania aplikacji.

---

## 🚀 Uruchomienie projektu

### 1️⃣ Wymagania wstępne

Aby uruchomić aplikację, upewnij się, że masz zainstalowane:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.24+](https://go.dev/dl/)

### 2️⃣ Klonowanie repozytorium

```sh
 git clone https://github.com/P-franczak/SWIFT-API.git
 cd swift-api
```

### 3️⃣ Konfiguracja bazy danych

Plik `docker-compose.yml` zawiera konfigurację PostgreSQL.

```yaml
services:
  db:
    image: postgres:latest
    ports:
      - '5433:5432' # Jeśli port 5432 jest zajęty, zmień pierwszy port na inny
```

### 4️⃣ Uruchomienie aplikacji za pomocą Dockera

```sh
docker-compose up --build -d
```

To polecenie:

- Pobiera niezbędne obrazy Dockera.
- Uruchamia bazę danych PostgreSQL.
- Buduje i uruchamia aplikację Go.

Sprawdź, czy kontenery działają:

```sh
docker ps
```

### 5️⃣ Sprawdzenie działania API

Po uruchomieniu aplikacja powinna działać na `http://localhost:8080`.

- Sprawdź status API:

```sh
curl http://localhost:8080
```

Odpowiedź powinna być:

```
Hello, SWIFT API is running!
```

### 6️⃣ Testowanie endpointów API

#### 📌 Pobranie danych konkretnego kodu SWIFT

```sh
curl -X GET http://localhost:8080/v1/swift-codes/ABCDEF12XXX
```

#### 📌 Pobranie wszystkich kodów SWIFT dla danego kraju

```sh
curl -X GET http://localhost:8080/v1/swift-codes/country/PL
```

#### 📌 Dodanie nowego kodu SWIFT

```sh
curl -X POST http://localhost:8080/v1/swift-codes -H "Content-Type: application/json" -d '{
    "address": "123 Main Street",
    "bankName": "Bank Polska",
    "countryISO2": "PL",
    "countryName": "Polska",
    "isHeadquarter": true,
    "swiftCode": "ABCDEF12XXX"
}'
```

#### 📌 Usunięcie kodu SWIFT

```sh
curl -X DELETE http://localhost:8080/v1/swift-codes/ABCDEF12XXX
```

---

## 🛠️ Struktura projektu

```
swift-api/
│── docker/
│   ├── docker-compose.yml   # Konfiguracja Dockera
│── db/
│   ├── init.sql             # Skrypt inicjalizujący bazę danych
│── internal/
│   ├── models.go            # Struktury danych
│   ├── database.go          # Obsługa bazy danych
│   ├── handlers.go          # Logika API
│── main.go                  # Główny plik aplikacji
│── README.md                # Dokumentacja
│── Dockerfile               # Konfiguracja Dockera dla aplikacji
```

---

## ✅ Testowanie aplikacji

Aby uruchomić testy jednostkowe:

```sh
go test ./...
```

---

## 🔄 Zatrzymanie i usunięcie kontenerów

```sh
docker-compose down
```

To polecenie:

- Zatrzymuje aplikację i bazę danych.
- Usuwa wszystkie powiązane kontenery.

---

## 📩 Kontakt

Jeśli masz pytania lub sugestie, skontaktuj się ze mną na [LinkedIn](https://www.linkedin.com/in/twoj-profil/) lub otwórz issue w repozytorium.

---

Teraz Twój projekt jest gotowy do udostępnienia na GitHubie! 🚀
