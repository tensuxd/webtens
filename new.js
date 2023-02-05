let a = 0

function count() {
    a++
    console.log(a);
}

count();
a = 5;
console.log(a);
count();
count();


function cs2() {
    let b = 0
    return function(){
        b++
        console.log(b);
    }
}

let d = cs2()
d()
b = 3
console.log(b);
d()
d()