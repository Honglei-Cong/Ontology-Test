
function asyncfunc(callback) {
        setTimeout(callback, 200);
}

color = 'blue';

(function(color) {
        asyncfunc(function() {
                console.log('the color is ' + color);
        });
})(color);

color = 'green';

