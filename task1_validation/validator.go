// Package validation містить завдання 1 домашньої роботи: валідатор форми
// реєстрації, що повертає власний (custom) тип помилки з переліком УСІХ
// невалідних полів, а не лише першого.
package validation

// RegistrationForm — вхідні дані форми реєстрації, які потрібно перевірити.
type RegistrationForm struct {
	Email    string
	Password string
	Age      int
}

// ValidationError — власний тип помилки, що переносить структуровані дані:
// список назв усіх полів, які не пройшли валідацію.
//
// TODO(Завдання 1): визначте поля структури.
// Підказка: вам знадобиться щонайменше поле Fields []string.
type ValidationError struct {
	// TODO: додайте поля
}

// Error реалізує інтерфейс error.
//
// TODO(Завдання 1): поверніть читабельне повідомлення, що перелічує всі
// невалідні поля, наприклад:
// "registration invalid: fields email, password"
func (e *ValidationError) Error() string {
	// TODO: реалізуйте
	panic("not implemented")
}

// ValidateRegistration перевіряє форму реєстрації та повертає
// *ValidationError з переліком ВСІХ невалідних полів, якщо форма невалідна.
// Якщо форма валідна, повертає nil.
//
// Правила валідації:
//   - Email не може бути порожнім
//   - Password не може бути порожнім і має містити щонайменше 8 символів
//   - Age має бути в межах [0, 150]
//
// TODO(Завдання 1): реалізуйте функцію так, щоб перевірка НЕ зупинялась
// на першому невалідному полі — потрібно зібрати всі помилки одразу.
func ValidateRegistration(f RegistrationForm) error {
	// TODO: реалізуйте
	panic("not implemented")
}
