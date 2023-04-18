# Проект “Система рекомендаций”

Это онлайн-платформа, которая рекомендует пользователям новый контент, который может заинтересовать их, и побудить пользователей потреблять больше контента на этой платформе.

Рекомендации тут идут по аниме

- **Основной функционал**
    - **Регистрация Пользователя:** (не успел реализовать полноценно. К тому же, нет авторизации gmail. Костыль, одним словом)
        - Почта
        - Пароль
        - Имя
        - Пол
        - Возраст
        
    - **Авторизация Пользователя**
 
        - Почта
        - Пароль
    
    
    
    
    - **CRUD для Админа**
        
        Поля таблицы со списком всех аниме
         - id
        - author
        - year
        - views
        - image
        - genre
        - title        
            
            JSON пример:
              {
                "id": 20,
                "title": "Clannad",
                "author": "Key",
                "genre": "Romance, Drama, Supernatural",
                "year": 2007,
                "image": "https://example.com/images/clannad.jpg",
                "views": 13000
              }
        
    - **Поиск**
        - Живой поиск по названию товара
        
    - **Фильтрация**
        - По автору
        - По жанру
        - По названию
        
    - **Избранное(Любимое)**
        - Контент, который пользователю понравился
        
    - **Рекомендации**
        - популярное
        - новинка
        - по вкусу пользователя 
              - Этот пункт реализован довольно просто. Сначала, берём все жанры, какие только есть и считаем их количество в таблице favorites. Находим ТОП - 3.
              После, возвращаем JSON файлом те аниме, которых нету в favorites, но в которых встречаются жанры, которые популярны у пользователя. 
              Сначала идёт 100% попадания жанрам, после 2 из 3, потом только 1
     
     Поскольку фронта тут нет, тестить это всё можно только в постмане. Ниже вставлю все запросы:
        
        
        Create Anime: POST http://localhost:8080/content
          {
            "id": 20,
            "title": "Clannad",
            "author": "Key",
            "genre": "Romance, Drama, Supernatural",
            "year": 2007,
            "image": "https://example.com/images/clannad.jpg",
            "views": 13000
          }
          
        Read All Anime: GET http://localhost:8080/content
        
        Read by ID: GET http://localhost:8080/content/your_id
        
        Update Anime: POST http://localhost:8080/content JSON File
        
        Delete Anime: DELETE http://localhost:8080/content/your_id
        
        Search: GET http://localhost:8080/contentSearch
                Query: name - query, value - your_value
                Headers: header - X-Requested-With, value -XMLHttpRequest
                
        Filter by genre: GET http://localhost:8080/filterGenre
                Query: name - query, value - your_value(ex: drama)
        
        Filter by author: GET http://localhost:8080/filterAuthor
                Query: name - query, value - your_value(ex: Masashi Kishimoto)
    
        Filter by ABC: GET http://localhost:8080/filterABC
        
        Create favorites: POST http://localhost:8080/favorites
                                        {
                                          "userID": 12,
                                          "animeID": 79
                                        }
                                        
        Read Favorites: GET http://localhost:8080/favorites
                Query: name - query, value - user_id
                
        Delete Favorites: DELETE http://localhost:8080/favorites/your_id
        
        Recommends New: GET http://localhost:8080/new
        
        Recommends Popular: GET http://localhost:8080/popular
        
        Indivisual Recommendations: GET Indivisual Recommendations
                      Query: name - query, value - user_id
                      
        Register User: http://localhost:8080/registration
                  {
                    "gender": "male",
                    "email": "example@mail.com",
                    "age": 30,
                    "name": "John Smith",
                    "password": "password123"
                  }
                  
        Login User: GET http://localhost:8080/login/example@example.com/password
        
        Reset password: PUT http://localhost:8080/reset
            {
              "email": "example@mail.com",
              "password": "passBoss123"
            }
        
