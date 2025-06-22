package seed

import (
	"context"
	"log"

	"main/db/ent"
	"main/db/ent/product"
	"main/db/ent/productcategory"
	"main/db/ent/typeofpackaging"
	"main/db/ent/user"
	"main/db/ent/usercategory"
	"main/db/entity"
	"main/domain"
)

// SeedDatabase засеивает базу данных начальными данными, если они еще не существуют.
func SeedDatabase(ctx context.Context, client *ent.Client) {
	log.Println("Начинаем засеивание базы данных...")

	productCateroriesToCreate := []entity.ProductCategoryModel{
		{
			Name: "Тушка",
		},
		{
			Name: "Разделка",
		},
		{
			Name: "Субпродукты",
		},
		{
			Name: "Рубленные полуфабрикаты",
		},
		{
			Name: "Полуфабрикаты в маринаде",
		},
		{
			Name: "Продукция для барбекю",
		},
	}

	createdCategory := make(map[string]*ent.ProductCategory)

	for _, productCategory := range productCateroriesToCreate {
		_, err := client.ProductCategory.Query().
			Where(productcategory.NameEQ(productCategory.Name)).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				t, createErr := client.ProductCategory.
					Create().
					SetName(productCategory.Name).
					Save(ctx)
				if createErr != nil {
					log.Printf(
						"Ошибка при создании категории товара'%s': %v",
						productCategory.Name,
						createErr,
					)
				} else {
					log.Printf("Создана категория товара: %s", t.Name)
					createdCategory[productCategory.Name] = t
				}
			} else {
				log.Printf("Ошибка при поиске категории товара '%s': %v", productCategory.Name, err)
			}
		} else {
			log.Printf("Категория товара '%s' уже существует.", productCategory.Name)
			existingType, _ := client.ProductCategory.Query().Where(productcategory.NameEQ(productCategory.Name)).Only(ctx)
			createdCategory[productCategory.Name] = existingType
		}
	}

	// Убедимся, что у нас есть хотя бы одина категория для связи с товарами
	var butcheringProductCategory *ent.ProductCategory

	if t, ok := createdCategory["Разделка"]; ok {
		butcheringProductCategory = t
	} else {
		log.Fatal("Не удалось получить или создать категорию 'Разделка', невозможно продолжить засеивание проектов.")
	}

	// Создаем типы упаковок товаров

	typesOfPackagingToCreate := []entity.TypeOfPackagingModel{
		{
			Name: "Пакет",
		},
		{
			Name: "Контейнер",
		},
	}

	createdTypeOfPackaging := make(map[string]*ent.TypeOfPackaging)

	for _, typeOfPackaging := range typesOfPackagingToCreate {
		_, err := client.TypeOfPackaging.Query().
			Where(typeofpackaging.NameEQ(typeOfPackaging.Name)).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				t, createErr := client.TypeOfPackaging.
					Create().
					SetName(typeOfPackaging.Name).
					Save(ctx)
				if createErr != nil {
					log.Printf(
						"Ошибка при создании типа упаковки'%s': %v",
						typeOfPackaging.Name,
						createErr,
					)
				} else {
					log.Printf("Создан тип упаковки : %s", t.Name)
					createdTypeOfPackaging[typeOfPackaging.Name] = t
				}
			} else {
				log.Printf("Ошибка при поиске типа упаковки'%s': %v", typeOfPackaging.Name, err)
			}
		} else {
			log.Printf("Тип упаковки '%s' уже существует.", typeOfPackaging.Name)
			existingType, _ := client.TypeOfPackaging.Query().Where(typeofpackaging.NameEQ(typeOfPackaging.Name)).Only(ctx)
			createdTypeOfPackaging[typeOfPackaging.Name] = existingType
		}
	}

	// Убедимся, что у нас есть хотя бы один тип упаковки для связи с товарами
	var bagPackagingType *ent.TypeOfPackaging

	if t, ok := createdTypeOfPackaging["Пакет"]; ok {
		bagPackagingType = t
	} else {
		log.Fatal("Не удалось получить или создать тип упаковки 'Пакет', невозможно продолжить засеивание товаров.")
	}

	// 2. Засеивание товаров(Products)
	productsToCreate := []struct {
		Name               string
		Weight             float64
		ProductComposition string
		MinStorageTemp     int
		MaxStorageTemp     int
		ShelfLife          int
		Picture            string
		ProductCategoryID  int
		TypeOfPackagingID  int
	}{
		{
			Name:               `Азу`,
			Weight:             1,
			ProductComposition: "",
			MinStorageTemp:     0,
			MaxStorageTemp:     5,
			ShelfLife:          5,
			Picture:            "azu_photo.png",
			ProductCategoryID:  butcheringProductCategory.ID,
			TypeOfPackagingID:  bagPackagingType.ID,
		},
		{
			Name:               `Печень индейки`,
			Weight:             1,
			ProductComposition: "",
			MinStorageTemp:     0,
			MaxStorageTemp:     5,
			ShelfLife:          5,
			Picture:            "fork_photo.png",
			ProductCategoryID:  butcheringProductCategory.ID,
			TypeOfPackagingID:  bagPackagingType.ID,
		},
	}

	for _, productToCreate := range productsToCreate {
		_, err := client.Product.Query().Where(product.NameEQ(productToCreate.Name)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				builder := client.Product.Create().
					SetName(productToCreate.Name).
					SetWeight(productToCreate.Weight).
					SetPicture(productToCreate.Picture).
					SetProductComposition(productToCreate.ProductComposition).
					SetMinStorageTemp(productToCreate.MinStorageTemp).
					SetMaxStorageTemp(productToCreate.MaxStorageTemp)

				_, createErr := builder.Save(ctx)
				if createErr != nil {
					log.Printf(
						"Ошибка при создании товара'%s': %v",
						product.Name,
						createErr,
					)
				} else {
					log.Printf("Создан товар: %s", product.Name)
				}
			} else {
				log.Printf("Ошибка при поиске товара '%s': %v", product.Name, err)
			}
		} else {
			log.Printf("Товар '%s' уже существует.", product.Name)
		}
	}

	log.Println("Создание ролей пользователей")

	userCategoriesToCreate := []entity.UserCategoryModel{
		{
			Name: "admin",
		},
		{
			Name: "user",
		},
		{
			Name: "storekeeper",
		},
		{
			Name: "manager",
		},
	}

	for _, userCategoryToCreate := range userCategoriesToCreate {
		_, err := client.UserCategory.Query().
			Where(usercategory.NameEQ(userCategoryToCreate.Name)).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				builder := client.UserCategory.Create().
					SetName(userCategoryToCreate.Name)

				_, createErr := builder.Save(ctx)
				if createErr != nil {
					log.Printf(
						"Ошибка при создании роли пользователя: '%s': %v",
						userCategoryToCreate.Name,
						err,
					)
				} else {
					log.Printf("Создана роль пользователя: %s", userCategoryToCreate.Name)
				}
			} else {
				log.Printf("Ошибка при поиске роли '%s' : %v", userCategoryToCreate.Name, err)
			}
		}
	}

	log.Println("Создание пользователей")

	const (
		userCategoryAdmin       = 1
		userCategoryUser        = 2
		userCategoryStorekeeper = 3
		userCategoryManager     = 4
	)

	usersToCreate := []struct {
		FullName     string
		Email        string
		Password     string // Plain text password for seeding
		Username     string
		UserCategory int
	}{
		{
			FullName:     "Иван",
			Email:        "admin@example.com",
			Password:     "admin", // Use a strong password in production!
			Username:     "admin",
			UserCategory: userCategoryAdmin,
		},
		{
			FullName:     "Дмитрий",
			Email:        "admin@example.com",
			Password:     "user", // Use a strong password in production!
			Username:     "user",
			UserCategory: userCategoryUser,
		},
		{
			FullName:     "Мария",
			Email:        "storekeeper@example.com",
			Password:     "storekeeper", // Use a strong password in production!
			Username:     "storekeeper",
			UserCategory: userCategoryStorekeeper,
		},
		{
			FullName:     "Игорь",
			Email:        "manager@example.com",
			Password:     "manager", // Use a strong password in production!
			Username:     "manager",
			UserCategory: userCategoryManager,
		},
	}

	for _, userData := range usersToCreate {
		_, err := client.User.Query().Where(user.UsernameEQ(userData.Username)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// User not found, create it
				hashedPassword, hashErr := domain.HashPassword(userData.Password)
				if hashErr != nil {
					log.Printf(
						"Ошибка при хешировании пароля для пользователя '%s': %v",
						userData.Username,
						hashErr,
					)
					continue
				}

				builder := client.User.Create().
					SetFullName(userData.FullName).
					SetUsername(userData.Username).
					SetEmail(userData.Email).
					SetPasswordHash(hashedPassword).SetUserCategoryID(userData.UserCategory)

				_, createErr := builder.Save(ctx)
				if createErr != nil {
					log.Printf(
						"Ошибка при создании пользователя '%s': %v",
						userData.Username,
						createErr,
					)
				} else {
					log.Printf("Создан пользователь: %s (Роль: %s)", userData.Username, userData.UserCategory)
				}
			} else {
				log.Printf("Ошибка при поиске пользователя '%s': %v", userData.Username, err)
			}
		} else {
			log.Printf("Пользователь '%s' уже существует.", userData.Username)
		}
	}

	log.Println("Засеивание базы данных завершено.")
}

func intPtr(i int) *int {
	return &i
}
