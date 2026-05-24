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