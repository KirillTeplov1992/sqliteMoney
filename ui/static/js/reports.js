const tabs = document.querySelectorAll('.tabs__title');
const contents = document.querySelectorAll('.tabs__content');
const navigation = document.querySelector('.tabs__navigation');

tabs.forEach(tab => {
    tab.addEventListener('click', () => {
        tabs.forEach(t => t.classList.remove('active'));
        contents.forEach(c => c.classList.remove('active'));

        tab.classList.add('active');
        document.getElementById(`tab-${tab.dataset.tab}`).classList.add('active');

        updateIndicator(tab);
    });
    if(tab.classList.contains('active')){
        updateIndicator(tab);
    }
});

function updateIndicator(tab){
    navigation.style.setProperty('--indicator-width', `${tab.offsetWidth}px`);
    navigation.style.setProperty('--indicator-left', `${tab.offsetLeft}px`);
}

fetch('http://192.168.0.104:5050/points')
    .then(response => response.json())
    .then(data => {
        const date = [];
        const amount = [];

        data.forEach(item => {
            date.push(item["date"])
            amount.push(item["amount"])
        });

        const ctx = document.getElementById('myChart'); 

        new Chart(ctx, {
	        type: 'line',
	        data: {
		        labels: date,
		        datasets: [{
			        label: 'График',
			        data: amount,
			        borderWidth: 1
		        }]
	        },
		        options: {
			    scales: {
				y: {
					beginAtZero: false

				}
			}
		}
});
    })
