const today = new Date();
const year = today.getFullYear();
const month = String(today.getMonth()+1).padStart(2, '0');
const day = String(today.getDate()).padStart(2, '0');
const fDate = `${year}-${month}-${day}`;

const date = document.getElementById('date');
date.value = fDate;

// Получаю поле с суммой 
const amount = document.getElementById('amount');

amount.addEventListener('click', function(){
    amount.value = '';
});

// Получаю элемент с типом категории
const typeOfCategory = document.getElementById('type_of_category');

// Получаю элемент с категорией
const category = document.getElementById('category')

// Добавляем обработчик события 'change'
typeOfCategory.addEventListener('change', function(){
    // Получаем тип категории
    const tCategory = typeOfCategory.value;

    switch (tCategory) {
        case "Доход":
            get_incoms()
            break;
        case "Расход":
            get_expepenses()
            break;
        case "Перевод":
            get_accounts()
            break;
    }
});

function get_incoms() {
    fetch('http://127.0.0.1:5050/get_incoms')
        .then(response =>{
            if (!response.ok) {
                throw new Error('Ошибка сети: ${response.statusText}');
            }
            return response.json();
        })
        .then(data =>{
            category.length = 0;
            data.forEach(item =>{
                const option = document.createElement('option');
                option.value = item.id;
                option.text = item.name;
                category.add(option);
            });
        })
}

function get_accounts() {
    fetch('http://127.0.0.1:5050/get_accounts')
        .then(response =>{
            if (!response.ok) {
                throw new Error('Ошибка сети: ${response.statusText}');
            }
            return response.json();
        })
        .then(data =>{
            category.length = 0;
            data.forEach(item =>{
                const option = document.createElement('option');
                option.value = item.id;
                option.text = item.name;
                category.add(option);
            });
        })
}

function get_expepenses() {
    fetch('http://127.0.0.1:5050/get_expenses')
        .then(response =>{
            if (!response.ok) {
                throw new Error('Ошибка сети: ${response.statusText}');
            }
            return response.json();
        })
        .then(data =>{
            category.length = 0;
            data.forEach(item =>{
                const option = document.createElement('option');
                option.value = item.id;
                option.text = item.name;
                category.add(option);
            });
        })
}
    
