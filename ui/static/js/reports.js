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