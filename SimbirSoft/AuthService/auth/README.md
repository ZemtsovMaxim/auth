# Сервис авторизации 

Это самописный gRPC сервис регистрации и аутентификации пользователей.

## Описание функций 

В настоящее время сервис включает в себя следующие функции

1. Первичная регистрация пользователя

    ```func (s *serverAPI) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {...some go code...}```

    Описание: Клиент вводит логин и пароль, тем самым регистрируясь в нашей системе (Сохраняется в БД) 

2. Вход пользователя в систему

    ```func (s *serverAPI) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) { ...some go code...} ```

    Описание: Зарагистрированный ранее клиент вводит свой логин и пароль, получая JWT токен (Генерируется с учетом secret)

3. Выход пользователя из системы

    ```func (s *serverAPI) Logout(ctx context.Context, req *api.LogoutRequest) (*api.LogoutResponse, error) { ...some go code...}```

    Описание: Токен пользователя становится невалидным, тем самым пользователь выходит из системы

4. ```func (s *serverAPI) ValidateToken(ctx context.Context, req *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {...some go code...}```

    Описание: Проверка на валидность токена (Возвращаем payload)



## Описание Makefile

### Тестирование

    ```make test```

### Запуск линтера

    ```make lint```

### Генерация кода из proto файла

    ```make generate```

### Приминение миграций к базе данных

    ```make migrate```
